# Connect RPC implementation for OTP authentication using Twilio  

# Tasks
Basic gRPC Service using Connect.
- two basic connect services: auth-service and otp-service.

## Auth Service
Authentication service that manages user accounts and profiles.

### Features
1. Signup with Phone Number: Allow users to create an account using a phone
number.
2. OTP Generation: Trigger an OTP generation via the OTP service when a new
account is created.
3. Verify Account: Allow users to verify their accounts using the OTP received.
4. Login: Enable users to log in using their phone number and OTP.
5. Get Profile: Users can retrieve their profile data.

### RPCs to Implement
- SignupWithPhoneNumber
- VerifyPhoneNumber
- LoginWithPhoneNumber
- ValidatePhoneNumberLogin
- GetProfile

## OTP Service
service dedicated to generating OTPs.

### Feature
1. Generate OTP: When invoked, generate a randomised OTP using Twilio's API and
provide it to the requesting service.

### RPCs to Implement
- GenerateOTP


# Database Integration
Integrating a PostgreSQL database to persist data
related to user activities.
## BeegGo ORM 
1. Perform CRUD operations
### Database Operations
1. Persist User Profiles: Store user profile information upon account creation.
2. Log Events: Record login and logout activities.
<br>

## Configurations

### Sample .env config file
```
SECRET=
TWILIO_AUTH_TOKEN=
TWILIO_ACCOUNT_SID=
TWILIO_SERVICE_SID=
DB_USER=""
DB_PASSWORD=""
DB_NAME=""
DB_HOST="localhost"
PORT="5432"
DB_DRIVER="postgres"
LOCAL_HOST="localhost:8080"
MOBILE=""
```

### Database Queries for setup
     query.sql

## Steps to run
- go mod tidy
- buf lint
- buf generate 
- update .env with all required values
- go run main.go
- Use postman to import proto files for testing RPC methods
> internal/proto/auth_service/v1/auth.proto

> internal/proto/otp_service/v1/otp.proto

# Package Documentations
## ConnectRPC
1. https://connectrpc.com/docs/go/getting-started/ 

## Twilio Docs 
1. Send OTP - https://www.twilio.com/docs/verify/sms
2. Verify OTP - https://www.twilio.com/docs/verify/api/verification-check 

## Beego ORM
1. https://pkg.go.dev/github.com/beego/beego/v2/client/orm#section-readme
