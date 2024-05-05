package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/astaxie/beego/orm"
	"github.com/joho/godotenv"
	"github.com/saurabh-sde/otp-authentication-go/internal/gen/auth_service/v1/auth_servicev1connect"
	"github.com/saurabh-sde/otp-authentication-go/internal/gen/otp_service/v1/otp_servicev1connect"
	"github.com/saurabh-sde/otp-authentication-go/messagingQueue/recieve"
	"github.com/saurabh-sde/otp-authentication-go/service"
	"github.com/saurabh-sde/otp-authentication-go/utility"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	_ "github.com/lib/pq"
)

func init() {
	// *** Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %+v", err)
	}

	// *** Connect to the database

	// driver
	dbDriver := os.Getenv("DB_DRIVER")
	orm.RegisterDriver(dbDriver, orm.DRPostgres)

	// Set the database connection str
	// connStr := "postgres://" + DBUser + ":" + DBPassword + "@" + DBHost + ":" + DBPort + "/" + DBName
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		dbDriver,
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// register database
	err = orm.RegisterDataBase("default", dbDriver, dsn)
	if err != nil {
		log.Fatalf("Error connecting DB: %+v", err)
	}
	// debug sql logs true
	orm.Debug = true
}

func main() {
	mux := http.NewServeMux()
	// service
	otpService := &service.OTPService{}
	autService := &service.AuthService{}

	mux.Handle(otp_servicev1connect.NewOTPServiceHandler(otpService))
	mux.Handle(auth_servicev1connect.NewAuthServiceHandler(autService))

	go recieve.InitializeMQConsumer()

	utility.Print(nil, "Initializing server: ", os.Getenv("LOCAL_HOST"))
	http.ListenAndServe(
		os.Getenv("LOCAL_HOST"),
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
