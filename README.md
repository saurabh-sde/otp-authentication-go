# Connect RPC implementation for OTP authentication using Twilio  

# Tasks
Basic gRPC Service using Connect.
- `auth-service and otp-service`.

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
- `SignupWithPhoneNumber`
- `VerifyPhoneNumber`
- `LoginWithPhoneNumber`
- `ValidatePhoneNumberLogin`
- `GetProfile`

## OTP Service
service dedicated to generating OTPs.

### Feature
1. Generate OTP: When invoked, generate a randomised OTP using Twilio's API and
provide it to the requesting service.

### RPCs to Implement
- `GenerateOTP`

<br>

# Database Integration
Integrating a `PostgreSQL` database to persist data
related to user activities.
## BeegGo ORM 
1. Performing all CRUD operations
### Database Operations
1. `Persist User Profiles`: Store user profile information upon account creation.
2. `Log Events`: Record login and logout activities.
<br>
<br>

# Messaging with Broker

Implement a messaging broker to facilitate communication between services.

## Auth Service
- Publish Events: Send a `SendOTP` message on the `verification` topic when a
new account is created.

## OTP Service
- Consume Events: Listen to the `verification` topic, receive the `SendOTP` message, and proceed to send the OTP to the user's phone number using Twilio.
<br>
<br>

# Configurations

### Sample `.env` config file
```
SECRET=<JWT Token secret>
TWILIO_AUTH_TOKEN=<get from twilio>
TWILIO_ACCOUNT_SID=<get from twilio>
TWILIO_SERVICE_SID=<get from twilio>
DB_USER="<set user>"
DB_PASSWORD="<set password>"
DB_NAME="auth"
DB_HOST="localhost"
DB_PORT="5432"
DB_DRIVER="postgres"
LOCAL_HOST="localhost:8080"
HOSTNAME="localhost"
PORT="8080"
RABBIT_HOST="localhost"
RABBIT_USER="guest"
RABBIT_PASSWORD="guest"
RABBIT_PORT="5672"
MOBILE="<used for running Send OTP test>"
```

### Run services after installation
1. Postgres - 
```
brew install postgresql@14
# start service
brew services start postgresql
```
2. RabbitMQ - https://www.rabbitmq.com/docs/install-homebrew
```
brew install rabbitmq
# starts a local RabbitMQ node
brew services start rabbitmq

# highly recommended: enable all feature flags on the running node
/opt/homebrew/sbin/rabbitmqctl enable_feature_flag all
```

### Setup Postgres Database
     query.sql

# Steps to run
- clone repo
- `go mod tidy`
- `buf lint`
- `buf generate` 
- update `.env `with all required values
- start `postgres and rabbitMQ` service
- `go run main.go`
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

## Rabbit MQ
1. `Topics` - https://www.rabbitmq.com/tutorials/tutorial-five-go

<br>

# Next Steps 

1. Docker setup using ` docker-compose.yml `
- Setup in single docker `docker-compose up`