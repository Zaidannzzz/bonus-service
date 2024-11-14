package jwt

import (
	"backend-api-test/config"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ValidateToken(encodedToken, secretKey string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}

func GenerateToken(email, tokenName string, userUUID string, conf *config.Config) (string, error) {

	claim := jwt.MapClaims{}
	claim["email"] = email
	//claim["role_uuid"] = roleUUID
	//claim["role_name"] = roleName
	claim["user_uuid"] = userUUID

	/*
		tokenName is between at and rt
		at is access token
		rt is refresh token
	*/
	if tokenName == "at" {
		claim["token_name"] = "access_token"
		claim["exp"] = time.Now().Add(time.Minute * 60 * 1 * 24).Unix() // AccessToken expires after  24 Hours
	} else if tokenName == "rt" {
		claim["token_name"] = "refresh_token"
		claim["exp"] = time.Now().Add(time.Minute * 60 * 1 * 24).Unix() // RefreshToken expires after 24 Hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	secretKey := []byte(conf.JWT.SecretKey)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ExtractTokenEmail(token *jwt.Token) (*string, error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	email := fmt.Sprintf("%v", claims["email"])
	return &email, nil
}

func ExtractTokenUserUUID(token *jwt.Token) (*string, error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	userUUID := fmt.Sprintf("%v", claims["user_uuid"])
	return &userUUID, nil
}

func ExtractTokenRoleID(token *jwt.Token) (*string, error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	roleUUID := fmt.Sprintf("%v", claims["role_uuid"])
	return &roleUUID, nil
}

func ExtractTokenRoleName(token *jwt.Token) (*string, error) {
	claims, _ := token.Claims.(jwt.MapClaims)
	roleName := fmt.Sprintf("%v", claims["role_name"])
	return &roleName, nil
}
