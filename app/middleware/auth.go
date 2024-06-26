package middleware

import (
	"fmt"
	"log"
	"motionserver/app/database/schema"
	"motionserver/utils/config"
	"motionserver/utils/response"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Protected(refresh bool) fiber.Handler {
	conf := config.NewConfig()

	if conf.Middleware.Jwt.Secret == "" {
		panic("JWT secret is not set")
	}

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(conf.Middleware.Jwt.Secret),
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			jwtc := new(JWTClaims)
			_, err := jwt.ParseWithClaims(user.Raw, jwtc, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}
				return []byte(conf.Middleware.Jwt.Secret), nil
			})

			if err != nil {
				return c.Status(fiber.StatusUnauthorized).
					JSON(response.Response{
						Code:     fiber.StatusUnauthorized,
						Data:     nil,
						Messages: response.RootMessage("unauthorized"),
						Meta:     nil,
					})

			} else {
				c.Locals("token", jwtc)
				if jwtc.Type == conf.Middleware.Jwt.AccessKey && !refresh {
					return c.Next()
				}
				if jwtc.Type == conf.Middleware.Jwt.RefreshKey && refresh {
					return c.Next()
				}

				return c.Status(fiber.StatusUnauthorized).
					JSON(response.Response{
						Code:     fiber.StatusUnauthorized,
						Data:     nil,
						Messages: response.RootMessage("wrong_token_type"),
						Meta:     nil,
					})
			}

		},
	})
}

func ByRole(role schema.Role) fiber.Handler {
	conf := config.NewConfig()

	if conf.Middleware.Jwt.Secret == "" {
		panic("JWT secret is not set")
	}

	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(conf.Middleware.Jwt.Secret),
		ErrorHandler: jwtError,
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			jwtc := new(JWTClaims)
			_, err := jwt.ParseWithClaims(user.Raw, jwtc, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}
				return []byte(conf.Middleware.Jwt.Secret), nil
			})

			if err != nil {
				return c.Status(fiber.StatusUnauthorized).
					JSON(response.Response{
						Code:     fiber.StatusUnauthorized,
						Data:     nil,
						Messages: response.RootMessage("unauthorized"),
						Meta:     nil,
					})

			} else {
				c.Locals("token", jwtc)
				fmt.Println(role == schema.Role(jwtc.Roles), role, jwtc.Roles)
				if jwtc.Type == conf.Middleware.Jwt.AccessKey && role == schema.Role(jwtc.Roles) {
					return c.Next()
				}

				return c.Status(fiber.StatusUnauthorized).
					JSON(response.Response{
						Code:     fiber.StatusUnauthorized,
						Data:     nil,
						Messages: response.RootMessage("wrong_token_type"),
						Meta:     nil,
					})
			}

		},
	})
}

func jwtError(c *fiber.Ctx, err error) error {

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(response.Response{
				Code:     fiber.StatusUnauthorized,
				Data:     nil,
				Messages: response.RootMessage("unauthorized"),
				Meta:     nil,
			})

	}
	return c.Next()
}

type JWTClaims struct {
	UserId uint64 `json:"user_id"`
	Type   string `json:"type"`
	Roles  string `json:"roles"`
	jwt.RegisteredClaims
}

type TokenData struct {
	UserId    uint64 `json:"user_id"`
	Roles     string `json:"roles"`
	TokenType string `json:"token_type"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenParams struct {
	data TokenData
	exp  time.Time
	conf *config.Config
}

func generateToken(params TokenParams) (string, error) {
	log.Println(params.exp)
	claims := &JWTClaims{
		UserId: params.data.UserId,
		Type:   params.data.TokenType,
		Roles:  params.data.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(params.exp),
			ID:        strconv.FormatUint(params.data.UserId, 10),
			Issuer:    params.conf.App.Name,
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(params.conf.Middleware.Jwt.Secret))
}

func GenerateTokenUser(token TokenData) (*TokenResponse, error) {

	config := config.NewConfig()
	accessType := token
	accessType.TokenType = config.Middleware.Jwt.AccessKey

	accessParams := TokenParams{
		data: accessType,
		exp:  time.Now().AddDate(0, 0, 1),
		conf: config,
	}
	access, err := generateToken(accessParams)
	if err != nil {
		return nil, err
	}
	refreshType := token
	refreshType.TokenType = config.Middleware.Jwt.RefreshKey
	refreshParams := TokenParams{
		data: refreshType,
		exp:  time.Now().AddDate(0, 1, 0),
		conf: config,
	}
	refresh, err := generateToken(refreshParams)
	if err != nil {
		return nil, err
	}
	return &TokenResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
