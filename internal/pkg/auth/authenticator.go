package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

type Authenticator struct {
	jwtSigningKey []byte
}

func NewAuthenticator(jwtSigningKey string) *Authenticator {
	return &Authenticator{
		jwtSigningKey: []byte(jwtSigningKey),
	}
}

const (
	Audience      = "lucidstack"
	Issuer        = "lucidstack"
	BearerToken   = "Bearer"
	TypeUserToken = "user"
)

func (a *Authenticator) GenerateUserToken(userID string, organizationID string, admin bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":             userID,
		"organization_id": organizationID,
		"iss":             Issuer,
		"aud":             Audience,
		"type":            TypeUserToken,
		"admin":           admin,
		"iat":             int(time.Now().Unix()),
		"exp":             int(time.Now().Add(time.Hour * 24).Unix()),
	})

	tokenString, err := token.SignedString(a.jwtSigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (a *Authenticator) ValidateUserContext(c *gin.Context) (*AuthenticatedUser, error) {
	tokenType, token, err := a.extractToken(c)

	if err != nil {
		return nil, err
	}

	switch tokenType {
	case BearerToken:
		return a.validateUserToken(token)
	default:
		return nil, fmt.Errorf("invalid token type")
	}
}

func (a *Authenticator) validateUserToken(tokenString string) (*AuthenticatedUser, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return a.jwtSigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		tokenType, ok := claims["type"]
		if !ok {
			return nil, fmt.Errorf("invalid token type")
		}

		tokenTypeString, ok := tokenType.(string)
		if !ok {
			return nil, fmt.Errorf("invalid token type")
		}

		if tokenTypeString != TypeUserToken {
			return nil, fmt.Errorf("invalid token type")
		}

		id, ok := claims["sub"]

		if !ok {
			return nil, fmt.Errorf("invalid access token")
		}

		userIDString, ok := id.(string)

		if !ok {
			return nil, fmt.Errorf("invalid access token")
		}

		organizationID, ok := claims["organization_id"]

		if !ok {
			return nil, fmt.Errorf("invalid access token")
		}

		organizationIDString, ok := organizationID.(string)

		if !ok {
			return nil, fmt.Errorf("invalid access token")
		}

		admin, ok := claims["admin"]
		if !ok {
			return nil, fmt.Errorf("invalid access token")
		}

		adminBool, ok := admin.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid access token")
		}

		return &AuthenticatedUser{
			ID:             userIDString,
			OrganizationID: organizationIDString,
			Admin:          adminBool,
		}, nil
	} else {
		return nil, fmt.Errorf("invalid access token")
	}
}

func (a *Authenticator) extractToken(c *gin.Context) (string, string, error) {
	authorizationHeader := c.GetHeader("Authorization")

	if len(authorizationHeader) == 0 {
		return "", "", fmt.Errorf("authorization header is not set")
	}

	components := strings.Split(authorizationHeader, " ")

	if len(components) != 2 {
		return "", "", fmt.Errorf("invalid access token")
	}

	return components[0], components[1], nil
}
