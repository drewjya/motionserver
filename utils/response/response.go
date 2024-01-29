package response

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Alias for any slice

func RootMessage(message string) interface{} {
	return fiber.Map{
		"@root": message,
	}
}

type Messages = interface{}

type Error struct {
	Code    int `json:"code"`
	Message any `json:"message"`
}

// error makes it compatible with the error interface
func (e *Error) Error() string {
	return fmt.Sprint(e.Message)
}

// A struct to return normal response
type Response struct {
	Code     int      `json:"code"`
	Messages Messages `json:"messages"`
	Data     any      `json:"data"`
	Meta     any      `json:"meta"`
}

// nothiing to describe this fucking variable
var IsProduction bool

// Default error handler
var ErrorHandler = func(c *fiber.Ctx, err error) error {
	resp := Response{
		Code: fiber.StatusInternalServerError,
	}

	// handle errors
	fmt.Println(reflect.TypeOf(err))
	if c, ok := err.(validator.ValidationErrors); ok {
		resp.Code = fiber.StatusUnprocessableEntity
		resp.Messages = removeTopStruct(c.Translate(trans))
	} else if c, ok := err.(*fiber.Error); ok {
		resp.Code = c.Code
		resp.Messages = RootMessage(c.Message)
	} else if c, ok := err.(*Error); ok {
		resp.Code = c.Code
		resp.Messages = c.Message

		// for ent and other errors
		if resp.Messages == nil {
			resp.Messages = RootMessage(err.Error())
		}
	} else {
		resp.Messages =
			RootMessage(err.Error())

	}

	if !IsProduction {
		log.Error().Err(err).Msg("From: Fiber's error handler")
	}

	return Resp(c, resp)
}

// function to return pretty json response
func Resp(c *fiber.Ctx, resp Response) error {
	// set status
	if resp.Code == 0 {
		resp.Code = fiber.StatusOK
	}
	c.Status(resp.Code)
	// return response
	return c.JSON(resp)
}

// remove unecessary fields from validator message
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}

	for field, msg := range fields {
		stripStruct := field[strings.Index(field, ".")+1:]
		res[toCamelCase(stripStruct)] = msg
	}

	return res
}

func toCamelCase(s string) string {
	parts := strings.Split(s, "")
	for i, part := range parts {
		parts[i] = strings.ToLower(part)
		if i > 0 {
			parts[i] = part
		}
	}
	return strings.Join(parts, "")
}
