package service

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/astaxie/beego/orm"
	"github.com/saurabh-sde/otp-authentication-go/auth"
	"github.com/saurabh-sde/otp-authentication-go/client"
	auth_servicev1 "github.com/saurabh-sde/otp-authentication-go/internal/gen/auth_service/v1"
	"github.com/saurabh-sde/otp-authentication-go/messagingQueue/publish"
	"github.com/saurabh-sde/otp-authentication-go/model"
	"github.com/saurabh-sde/otp-authentication-go/utility"
	"github.com/spf13/cast"
)

type AuthService struct{}

func (s *AuthService) SignupWithPhoneNumber(ctx context.Context, req *connect.Request[auth_servicev1.SignupWithPhoneNumberRequest]) (*connect.Response[auth_servicev1.SignupWithPhoneNumberResponse], error) {
	utility.Print(nil, "Request headers: ", req.Header())

	res := connect.NewResponse(&auth_servicev1.SignupWithPhoneNumberResponse{})
	if req.Msg.Mobile == "" || len(req.Msg.Mobile) < 10 {
		res.Msg.Message = "Invalid Mobile Number"
	} else {
		// check user exists or not
		user, err := model.GetUserByMobile(req.Msg.Mobile)
		if err != nil && err != orm.ErrNoRows {
			utility.Print(&err, "Error getting user")
			return res, nil
		} else if user != nil {
			// user is already registered and verified
			if user.Status == "approved" {
				res.Msg.Message = "User already registered. Please use Login"
				return res, nil
			}

			utility.Print(nil, "GENERATING OTP - Existing User: ", user.Id)
			// using OTP service client for OTP generation
			generateOTPResp, err := client.GenerateOTPClient(req.Msg.Mobile)
			if err != nil {
				res.Msg.Message = "Error in sending OTP"
				utility.Print(&err, "Error sending OTP to user")
				return res, nil
			}
			res.Msg.Message = generateOTPResp
			return res, nil
		}

		// insert user
		newUser := &model.AppUser{
			Name:   req.Msg.Name,
			Mobile: req.Msg.Mobile,
		}
		id, err := model.AddUser(newUser)
		if err != nil {
			res.Msg.Message = "Error in register User"
			utility.Print(&err, "Error in adding user")
			return res, nil
		}
		utility.Print(nil, "Added new user: ", id)
		utility.Print(nil, "GENERATING OTP - New User: ", id)

		generateOTPResp := "OTP sent to Mobile number"

		// TODO: Send OTP using Messaging Queue for new user
		// using OTP service client for OTP generation
		// generateOTPResp, err := client.GenerateOTPClient(req.Msg.Mobile)
		// if err != nil {
		// 	res.Msg.Message = "Error in sending OTP"
		// 	utility.Print(&err, "Error sending OTP to user")
		// 	return res, nil
		// }

		// Use MQ for send OTP for new user
		err = publish.PublishToMQ(newUser.Mobile)
		if err != nil {
			res.Msg.Message = "Error in sending OTP"
			utility.Print(&err, "Error sending OTP to user")
			return res, nil
		}

		res.Msg.Message = generateOTPResp
	}
	return res, nil
}

func (s *AuthService) VerifyPhoneNumber(ctx context.Context, req *connect.Request[auth_servicev1.VerifyPhoneNumberRequest]) (*connect.Response[auth_servicev1.VerifyPhoneNumberResponse], error) {
	utility.Print(nil, "VerifyPhoneNumber Request headers: ", req.Header())

	res := connect.NewResponse(&auth_servicev1.VerifyPhoneNumberResponse{})

	// verify otp code with mobile
	isValid, respStatus, err := TwilioVerifyOTP(req.Msg.Mobile, req.Msg.Code)
	if err != nil {
		res.Msg.Message = "Error in validating OTP"
		utility.Print(&err, res.Msg.Message)
		return res, nil
	}
	if !isValid {
		res.Msg.Message = "OTP Verification Failed"
		utility.Print(&err, "Verification Failed: ", respStatus)
		return res, nil
	}
	user, err := model.GetUserByMobile(req.Msg.Mobile)
	if err != nil || user == nil {
		utility.Print(&err, "Error getting user")
		return res, err
	}
	user.Status = respStatus
	err = model.UpdateUser(user)
	if err != nil {
		utility.Print(&err, "Error updating user")
		return res, err
	}
	res.Msg.Message = "Registered Successfully"
	res.Msg.Valid = true
	return res, nil
}

func (s *AuthService) LoginWithPhoneNumber(ctx context.Context, req *connect.Request[auth_servicev1.LoginWithPhoneNumberRequest]) (*connect.Response[auth_servicev1.LoginWithPhoneNumberResponse], error) {
	utility.Print(nil, "LoginWithPhoneNumber Request headers: ", req.Header())

	res := connect.NewResponse(&auth_servicev1.LoginWithPhoneNumberResponse{})

	user, err := model.GetUserByMobile(req.Msg.Mobile)
	if err != nil || user == nil {
		utility.Print(&err, "Error getting user")
		return res, nil
	}

	utility.Print(nil, "GENERATING LOGIN OTP: ", res.Msg.UserId)
	// Login flow does not need messaging queue acc. to requirements
	// using OTP service client for OTP generation
	generateOTPResp, err := client.GenerateOTPClient(req.Msg.Mobile)
	if err != nil {
		res.Msg.Message = "Error in sending OTP"
		utility.Print(&err, "Error sending OTP to user")
		return res, nil
	}
	res.Msg.Message = generateOTPResp
	res.Msg.UserId = cast.ToString(user.Id)
	return res, nil
}

func (s *AuthService) ValidatePhoneNumberLogin(ctx context.Context, req *connect.Request[auth_servicev1.ValidatePhoneNumberLoginRequest]) (*connect.Response[auth_servicev1.ValidatePhoneNumberLoginResponse], error) {
	utility.Print(nil, "ValidatePhoneNumberLogin Request headers: ", req.Header())

	res := connect.NewResponse(&auth_servicev1.ValidatePhoneNumberLoginResponse{})

	// check user exists
	user, err := model.GetUserByMobile(req.Msg.Mobile)
	if err != nil || user == nil {
		utility.Print(&err, "Error fetching user by mobile")
		return res, nil
	}

	// verify otp code with mobile
	isValid, respStatus, err := TwilioVerifyOTP(req.Msg.Mobile, req.Msg.Code)
	if err != nil {
		res.Msg.Message = "Error in validating OTP"
		utility.Print(&err, res.Msg.Message)
		return res, nil
	}
	if !isValid {
		res.Msg.Message = "OTP Verification Failed"
		utility.Print(&err, "Verification Failed: ", respStatus)
		return res, nil
	}

	// add login event
	_, err = model.AddLogEvent(&model.LogEvent{EventName: "login", UserId: cast.ToString(user.Id)})
	if err != nil {
		// skip returning error as login was success
		utility.Print(&err, "Error Adding login event")
	}

	// send jwt auth token
	loginToken, err := auth.CreateJWT(user.Mobile)
	if err != nil {
		utility.Print(&err, "Error generating login token")
		res.Msg.Message = "Invalid Server error"
		return res, err
	}

	// return response
	res.Msg.Message = "OTP Login Successful"
	res.Msg.Valid = true
	res.Msg.Token = loginToken
	return res, nil
}

func (s *AuthService) GetProfile(ctx context.Context, req *connect.Request[auth_servicev1.GetProfileRequest]) (*connect.Response[auth_servicev1.GetProfileResponse], error) {
	utility.Print(nil, "GetProfile Request headers: ", req.Header())

	res := connect.NewResponse(&auth_servicev1.GetProfileResponse{})

	// check user exists or not
	user, err := model.GetUserById(req.Msg.UserId)
	if err != nil || user == nil {
		utility.Print(&err, "Error getting user")
		res.Msg.Status = "User not found"
		err := errors.New("User not found")
		return res, err
	}

	// validate bearer token with mobile from bearer token value and user.Mobile from database using userId
	mobileFromToken, err := auth.ParseBearerToken(req.Header())
	if err != nil {
		utility.Print(&err, "Error in parsing token")
		res.Msg.Status = "Unauthenticated request"
		return res, err
	}

	// check valid mobile
	if mobileFromToken != user.Mobile {
		utility.Print(&err, "Error in validating token with mobile")
		res.Msg.Status = "Unauthenticated request"
		err := errors.New(res.Msg.Status)
		return res, err
	}

	// return resp
	res.Msg.Id = cast.ToString(user.Id)
	res.Msg.Name = user.Name
	res.Msg.Mobile = user.Mobile
	res.Msg.Status = user.Status
	return res, nil
}
