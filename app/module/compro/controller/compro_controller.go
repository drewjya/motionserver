package controller

import (
	"motionserver/app/module/compro/request"
	"motionserver/app/module/compro/service"
	koderor "motionserver/utils/error"
	"motionserver/utils/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type comproController struct {
	comproService service.ComproService
}

type ComproController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
}

func NewComproController(comproService service.ComproService) ComproController {
	return &comproController{
		comproService: comproService,
	}
}

func (_i *comproController) Index(c *fiber.Ctx) error {

	compro, err := _i.comproService.Show()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("success retrieve compro"),
		Data:     compro,
	})
}

func (_i *comproController) Store(c *fiber.Ctx) error {
	req := new(request.ComproRequest)
	file, errNew := c.FormFile("image")
	var vale koderor.KodeError
	var err error
	if errNew != nil {
		errVall := errNew.Error()
		if errVall == "there is no uploaded file associated with the given key" {
			errVall = "Image field is required"
		}

		vale = koderor.New("image", errVall)
	}
	err = response.ParseAndValidate(c, req)
	if err != nil || vale != nil {
		if err != nil && vale == nil {
			val := err.(validator.ValidationErrors)
			return koderor.NewErrors(&val, nil)
		}
		if err == nil && vale != nil {
			return koderor.NewErrors(nil, vale.(*koderor.ErrorKode))
		}
		val := err.(validator.ValidationErrors)
		return koderor.NewErrors(&val, vale.(*koderor.ErrorKode))
	}
	req.File = file

	err = _i.comproService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("Compro Successfully created"),
	})
}
