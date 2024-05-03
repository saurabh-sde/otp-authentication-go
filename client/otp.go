package client

import (
	"context"
	"net/http"
	"os"

	"connectrpc.com/connect"
	otp_servicev1 "github.com/saurabh-sde/otp-authentication-go/internal/gen/otp_service/v1"
	"github.com/saurabh-sde/otp-authentication-go/internal/gen/otp_service/v1/otp_servicev1connect"
	"github.com/saurabh-sde/otp-authentication-go/utility"
)

func GenerateOTPClient(mobile string) (string, error) {
	client := otp_servicev1connect.NewOTPServiceClient(
		http.DefaultClient,
		"http://"+os.Getenv("LOCAL_HOST"),
	)
	res, err := client.GenerateOTP(
		context.Background(),
		connect.NewRequest(&otp_servicev1.GenerateOTPRequest{Mobile: mobile}),
	)
	if err != nil {
		utility.Print(&err)
		return err.Error(), err
	}

	return res.Msg.Message, err
}
