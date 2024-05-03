package service

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestTwilioSendOTP(t *testing.T) {
	// *** Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	type args struct {
		mobile string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "TEST 1",
			args:    args{mobile: os.Getenv("MOBILE")},
			want:    "OTP Generated",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TwilioSendOTP(tt.args.mobile)
			if (err != nil) != tt.wantErr {
				t.Errorf("TwilioSendOTP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("TwilioSendOTP() = %v, want %v", got, tt.want)
			}
		})
	}
}
