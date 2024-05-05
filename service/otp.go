package service

import (
	"context"

	"connectrpc.com/connect"
	otp_servicev1 "github.com/saurabh-sde/otp-authentication-go/internal/gen/otp_service/v1"
	"github.com/saurabh-sde/otp-authentication-go/utility"
)

type OTPService struct{}

func (s *OTPService) GenerateOTP(ctx context.Context, req *connect.Request[otp_servicev1.GenerateOTPRequest]) (*connect.Response[otp_servicev1.GenerateOTPResponse], error) {
	utility.Print(nil, "GenerateOTP Request: headers: ", req.Header(), " req: ", req.Msg.Mobile)

	res := connect.NewResponse(&otp_servicev1.GenerateOTPResponse{})

	resp, err := TwilioSendOTP(req.Msg.Mobile)
	if err != nil {
		res.Msg.Message = "Failed to Send OTP"
		utility.Print(&err, "Failed to Send OTP")
		return res, err
	}
	utility.Print(nil, resp)
	res.Msg.Message = "OTP sent to Mobile number"
	return res, nil
}
