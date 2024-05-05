package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/saurabh-sde/otp-authentication-go/utility"
)

func CreateJWT(mobile string) (bearerToken string, err error) {
	// create jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"mobile": mobile})

	// get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		utility.Print(&err, "Err in SignedString")
		return
	}
	return tokenString, err
}
func ParseBearerToken(header http.Header) (resp string, err error) {
	// get token from request
	// assuming 'Authorization': 'Bearer <YOUR_TOKEN_HERE>' as standard
	reqToken := header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	tokenString := strings.TrimSpace(splitToken[1])

	// Check if user is authenticated by parsing token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the alg for SigningMethodHS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		utility.Print(&err, "Error  parsing token")
		return
	}

	if !token.Valid {
		err = errors.New("Invalid token")
		utility.Print(&err, "Invalid token")
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		err = errors.New("Invalid token")
		utility.Print(&err, "Invalid token claims")
		return
	}

	utility.Print(nil, "Successfully Parsed Token: ", claims)
	return claims["mobile"].(string), nil
}
