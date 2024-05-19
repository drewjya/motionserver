package controller

import (
	"motionserver/app/module/banner/request"
	"motionserver/app/module/banner/service"
	koderor "motionserver/utils/error"
	"motionserver/utils/response"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type bannerController struct {
	bannerService service.BannerService
}

type BannerController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
}

func NewBannerController(bannerService service.BannerService) BannerController {
	return &bannerController{
		bannerService: bannerService,
	}
}

func (_i *bannerController) Index(c *fiber.Ctx) error {

	banner, err := _i.bannerService.Show()
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("success retrieve banner"),
		Data:     banner,
	})
}

func (_i *bannerController) Store(c *fiber.Ctx) error {
	req := new(request.BannerRequest)
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

	err = _i.bannerService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("Banner Successfully created"),
	})
}
