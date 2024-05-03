package service

import (
	"context"
	"log"

	"connectrpc.com/connect"
	"github.com/astaxie/beego/orm"
	"github.com/saurabh-sde/otp-authentication-go/client"
	auth_servicev1 "github.com/saurabh-sde/otp-authentication-go/internal/gen/auth_service/v1"
	"github.com/saurabh-sde/otp-authentication-go/model"
	"github.com/saurabh-sde/otp-authentication-go/utility"
	"github.com/spf13/cast"
)

type AuthService struct{}

func (s *AuthService) SignupWithPhoneNumber(ctx context.Context, req *connect.Request[auth_servicev1.SignupWithPhoneNumberRequest]) (*connect.Response[auth_servicev1.SignupWithPhoneNumberResponse], error) {
	log.Println("Request headers: ", req.Header())

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
	return res, nil
}

func (s *AuthService) VerifyPhoneNumber(ctx context.Context, req *connect.Request[auth_servicev1.VerifyPhoneNumberRequest]) (*connect.Response[auth_servicev1.VerifyPhoneNumberResponse], error) {
	log.Println("VerifyPhoneNumber Request headers: ", req.Header())

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

	// add login event
	_, err = model.AddLogEvent(&model.LogEvent{EventName: "login", UserId: cast.ToString(user.Id)})
	if err != nil {
		// skip returning error as login was success
		utility.Print(&err, "Error Adding login event")
	}

	return res, nil
}

func (s *AuthService) LoginWithPhoneNumber(ctx context.Context, req *connect.Request[auth_servicev1.LoginWithPhoneNumberRequest]) (*connect.Response[auth_servicev1.LoginWithPhoneNumberResponse], error) {
	log.Println("LoginWithPhoneNumber Request headers: ", req.Header())

	res := connect.NewResponse(&auth_servicev1.LoginWithPhoneNumberResponse{})

	user, err := model.GetUserByMobile(req.Msg.Mobile)
	if err != nil || user == nil {
		utility.Print(&err, "Error getting user")
		return res, nil
	}

	utility.Print(nil, "GENERATING LOGIN OTP: ", res.Msg.UserId)
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

func (s *AuthService) ValidatePhoneNumberLogin(ctx context.Context, req *connect.Request[auth_servicev1.ValidatePhoneNumberLoginRequest]) (*connect.Response[auth_servicev1.ValidatePhoneNumberLoginResponse], error) {
	log.Println("ValidatePhoneNumberLogin Request headers: ", req.Header())

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

	// return response
	res.Msg.Message = "OTP Login Successful"
	res.Msg.Valid = true

	return res, nil
}

func (s *AuthService) GetProfile(ctx context.Context, req *connect.Request[auth_servicev1.GetProfileRequest]) (*connect.Response[auth_servicev1.GetProfileResponse], error) {
	log.Println("GetProfile Request headers: ", req.Header())

	res := connect.NewResponse(&auth_servicev1.GetProfileResponse{})
	// check user exists or not
	user, err := model.GetUserById(req.Msg.UserId)
	if err != nil || user == nil {
		utility.Print(&err, "Error getting user")
		return res, nil
	}
	// return resp
	res.Msg.Id = cast.ToString(user.Id)
	res.Msg.Name = user.Name
	res.Msg.Mobile = user.Mobile
	res.Msg.Status = user.Status
	return res, nil
}
