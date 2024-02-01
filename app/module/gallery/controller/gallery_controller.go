package controller

import (
	"motionserver/app/module/gallery/request"
	"motionserver/app/module/gallery/service"
	koderor "motionserver/utils/error"
	"motionserver/utils/paginator"
	"motionserver/utils/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type galleryController struct {
	galleryService service.GalleryService
}

type GalleryController interface {
	Index(c *fiber.Ctx) error
	Show(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
}

func NewGalleryController(galleryService service.GalleryService) GalleryController {
	return &galleryController{
		galleryService: galleryService,
	}
}

func (_i *galleryController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req request.GalleriesRequest
	req.Pagination = paginate

	galleries, paging, err := _i.galleryService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("success retrieve galleries"),
		Data:     galleries,
		Meta:     paging,
	})
}

func (_i *galleryController) Show(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}
	gallery, err := _i.galleryService.Show(id)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("success retrieve gallery"),
		Data:     gallery,
	})
}

func (_i *galleryController) Store(c *fiber.Ctx) error {
	req := new(request.GalleryRequest)
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

	err = _i.galleryService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("Gallery Successfully created"),
	})
}
