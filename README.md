Connect RPC 
1. https://connectrpc.com/docs/go/getting-started/ 

Twilio Docs - 
1. Send OTP - https://www.twilio.com/docs/verify/sms
2. Verify OTP - https://www.twilio.com/docs/verify/api/verification-check 


Sample .env config file

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

Database Queries for setup
1. query.sql

Steps to run
1. go mod tidy
2. buf lint
3. buf generate 
4. go run main.go