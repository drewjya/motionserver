package service

import (
	"context"
	"encoding/json"
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
	Minio minio.Minio
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
	GetYoutube() (youtube *schema.Youtube, err error)
	SetYoutube() (err error)
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{
		Repo: repo,
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
	return _i.Repo.Create(req.ToDomain())
}
