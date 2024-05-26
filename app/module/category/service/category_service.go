package service

import (
	"context"
	"encoding/json"
	"fmt"
	"motionserver/app/database/schema"
	"motionserver/app/module/category/repository"
	"motionserver/app/module/category/request"
	"motionserver/app/module/category/response"
	"motionserver/internal/bootstrap/minio"
	"motionserver/utils/paginator"
	"net/http"
)

type categoryService struct {
	Repo  repository.CategoryRepository
	Minio *minio.Minio
}

// Delete implements CategoryService.
func (_i *categoryService) Delete(id uint) (err error) {
	category, err := _i.Repo.FindOne(id)
	if err != nil {
		return err
	}
	if len(category.Image) != 0 {
		ctx := context.Background()
		err = _i.Minio.DeleteFile(ctx, category.Image)
		if err != nil {
			return err
		}
	}
	return _i.Repo.Delete(id)
}

// Update implements CategoryService.
func (_i *categoryService) Update(req request.CategoryRequest, id uint) (err error) {
	category, err := _i.Repo.FindOne(id)
	if err != nil {
		return err
	}
	if len(category.Image) != 0 {
		ctx := context.Background()
		err = _i.Minio.DeleteFile(ctx, category.Image)
		if err != nil {
			return err
		}
	}
	if req.File != nil {
		fmt.Println(req.File.Filename)
		ctx := context.Background()
		val, err := _i.Minio.UploadFile(ctx, *req.File)
		if err != nil {
			fmt.Println("ERROR IMAGE", err)
			return err
		}
		req.Image = *val
	}
	return _i.Repo.Create(req.ToDomain())
}

// GetYoutube implements CategoryService.
func (_i *categoryService) GetYoutube() (youtube *schema.Youtube, err error) {
	return _i.Repo.GetYoutube()
}

// SetYoutube implements CategoryService.
func (_i *categoryService) SetYoutube() (err error) {

	res, err := http.Get(schema.URL)
	if err != nil {
		return
	}

	defer res.Body.Close()
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return err
	}
	err = _i.Repo.DeleteYoutube()
	if err != nil {
		return
	}
	for _, item := range result["items"].([]interface{}) {
		curr := item.(map[string]interface{})["snippet"].(map[string]interface{})
		currVid := item.(map[string]interface{})["id"].(map[string]interface{})
		thumbnailURL := curr["thumbnails"].(map[string]interface{})["default"].(map[string]interface{})["url"].(string)

		// Create a new Youtube entry
		youtube := schema.Youtube{
			PublishedAt: curr["publishedAt"].(string),
			VideoId:     currVid["videoId"].(string),
			Thumbnail:   thumbnailURL,
		}

		// Insert the entry into the database
		err = _i.Repo.SetYoutube(&youtube)

		break
	}

	return
}

type CategoryService interface {
	All(req request.CategoriesRequest) (categories []*response.Categories, paging paginator.Pagination, err error)
	Store(req request.CategoryRequest) (err error)

	Update(req request.CategoryRequest, id uint) (err error)
	Delete(id uint) (err error)
	GetYoutube() (youtube *schema.Youtube, err error)
	SetYoutube() (err error)
}

func NewCategoryService(repo repository.CategoryRepository, minio *minio.Minio) CategoryService {
	return &categoryService{
		Repo:  repo,
		Minio: minio,
	}
}

func (_i *categoryService) All(req request.CategoriesRequest) (categories []*response.Categories, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.GetCategories(req)
	if err != nil {
		return
	}
	ctx := context.Background()
	for _, v := range results {
		if len(v.Image) != 0 {

			image := _i.Minio.GenerateLink(ctx, v.Image)
			categories = append(categories, response.FromDomain(v, &image))
		} else {
			categories = append(categories, response.FromDomain(v, nil))
		}
	}
	return

}

func (_i *categoryService) Store(req request.CategoryRequest) (err error) {
	if req.File != nil {
		fmt.Println(req.File.Filename)
		ctx := context.Background()
		val, err := _i.Minio.UploadFile(ctx, *req.File)
		if err != nil {
			fmt.Println("ERROR IMAGE", err)
			return err
		}
		req.Image = *val
	}
	return _i.Repo.Create(req.ToDomain())
}
