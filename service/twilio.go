package service

import (
	"os"

	"github.com/saurabh-sde/otp-authentication-go/utility"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

func TwilioSendOTP(mobile string) (string, error) {
	client := twilio.NewRestClient()

	params := &verify.CreateVerificationParams{}
	params.SetTo(mobile)
	params.SetChannel("sms")
	utility.Print(nil, "Sending OTP to mobile:", mobile)
	resp, err := client.VerifyV2.CreateVerification(os.Getenv("TWILIO_SERVICE_SID"), params)
	if err != nil {
		utility.Print(&err, err.Error())
		return "Error in sending OTP", err
	} else {
		if resp.Status != nil {
			utility.Print(&err, "SendOTP status", *resp.Status)
		} else {
			utility.Print(&err, "SendOTP status", resp.Status)
		}
	}
	utility.Print(&err, "SendOTP Success")
	return "OTP Generated", err
}

func TwilioVerifyOTP(mobile string, code string) (bool, string, error) {
	client := twilio.NewRestClient()

	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(mobile)
	params.SetCode(code)
	utility.Print(nil, "Verifying OTP for mobile:", mobile)

	resp, err := client.VerifyV2.CreateVerificationCheck(os.Getenv("TWILIO_SERVICE_SID"), params)
	if err != nil {
		utility.Print(&err, err.Error())
		return false, "pending", err
	} else {
		if resp.Status != nil {
			utility.Print(&err, "OTPVerification status: ", *resp.Status)
		} else {
			utility.Print(&err, "OTPVerification status: ", resp.Status)
		}
	}
	// `pending`, `approved`, or `canceled`
	return *resp.Status == "approved", *resp.Status, err
}
