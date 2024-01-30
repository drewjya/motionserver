package controller

import (
	"log"
	"motionserver/app/module/gallery/request"
	"motionserver/app/module/gallery/service"
	"motionserver/utils/paginator"
	"motionserver/utils/response"

	"github.com/gofiber/fiber/v2"
)

type galleryController struct {
	galleryService service.GalleryService
}

type GalleryController interface {
	Index(c *fiber.Ctx) error
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

func (_i *galleryController) Store(c *fiber.Ctx) error {
	req := new(request.GalleryRequest)
	if err := response.ParseAndValidate(c, req); err != nil {
		return err
	}
	file, err := c.FormFile("image")
	if err.Error() != "" {
		return err
	}
	log.Println(file)

	err = _i.galleryService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("Gallery Successfully created"),
	})
}
