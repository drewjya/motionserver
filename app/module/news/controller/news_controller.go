package controller

import (
	"fmt"
	"motionserver/app/module/news/request"
	"motionserver/app/module/news/service"
	koderor "motionserver/utils/error"
	"motionserver/utils/paginator"
	"motionserver/utils/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type newsController struct {
	newsService service.NewsService
}

// FindOne implements NewsController.
func (_i *newsController) FindOne(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}
	news, err := _i.newsService.FindOne(id)
	if err != nil {
		return err
	}
	return response.Resp(c, response.Response{
		Data:     news,
		Messages: response.RootMessage("success retrieve news"),
		Code:     fiber.StatusOK,
	})
}

// Delete implements NewsController.
func (_i *newsController) Delete(c *fiber.Ctx) error {

	id, err := strconv.ParseUint(c.Params("id"), 10, 64)

	if err != nil {
		return err
	}
	err = _i.newsService.Delete(id)
	if err != nil {
		return err
	}
	return response.Resp(c, response.Response{
		Data:     nil,
		Messages: response.RootMessage("success delete news"),
		Code:     fiber.StatusCreated,
	})
}

// Update implements NewsController.
func (_i *newsController) Update(c *fiber.Ctx) error {
	var req request.NewsRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return err
	}

	err = _i.newsService.Update(id, req)
	if err != nil {
		return err
	}
	return response.Resp(c, response.Response{
		Data:     nil,
		Messages: response.RootMessage("success update news"),
		Code:     fiber.StatusCreated,
	})
}

type NewsController interface {
	Index(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindOne(c *fiber.Ctx) error
}

func NewNewsController(newsService service.NewsService) NewsController {
	return &newsController{
		newsService: newsService,
	}
}

func (_i *newsController) Store(c *fiber.Ctx) error {
	req := new(request.NewsRequest)
	file, errNew := c.FormFile("image")
	val := c.FormValue("berita")
	fmt.Println(val)
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
	fmt.Println(req.File)
	err = _i.newsService.Store(*req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("News Successfully created"),
	})
}

func (_i *newsController) Index(c *fiber.Ctx) error {
	paginate, err := paginator.Paginate(c)
	if err != nil {
		return err
	}

	var req request.NewssRequest
	req.Pagination = paginate

	news, paging, err := _i.newsService.All(req)
	if err != nil {
		return err
	}

	return response.Resp(c, response.Response{
		Messages: response.RootMessage("success retrieve news"),
		Data:     news,
		Meta:     paging,
	})
}
