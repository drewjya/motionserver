package service

import (
	"context"
	auth "motionserver/app/module/auth/repository"
	"motionserver/app/module/news/repository"
	"motionserver/app/module/news/request"
	"motionserver/app/module/news/response"
	"motionserver/internal/bootstrap/minio"
	"motionserver/utils/paginator"
)

type newsService struct {
	Repo  repository.NewsRepository
	Auth  auth.AuthRepository
	Minio *minio.Minio
}

// Delete implements NewsService.
func (_i *newsService) Delete(id uint64) (err error) {
	return _i.Repo.DeleteNews(uint(id))
}

type NewsService interface {
	All(req request.NewssRequest) (newss []*response.News, paging paginator.Pagination, err error)

	Store(req request.NewsRequest) (err error)
	Update(id uint64, req request.NewsRequest) (err error)
	Delete(id uint64) (err error)
}

func NewNewsService(repo repository.NewsRepository, Minio *minio.Minio, Auth auth.AuthRepository) NewsService {
	return &newsService{
		Repo:  repo,
		Auth:  Auth,
		Minio: Minio,
	}

}

func (_i *newsService) All(req request.NewssRequest) (newss []*response.News, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.FindNews(req)
	if err != nil {
		return
	}
	ctx := context.Background()
	for _, v := range results {
		img := _i.Minio.GenerateLink(ctx, v.Image)
		newss = append(newss, response.FromDomain(v, img))
	}
	return
}

func (_i *newsService) Store(req request.NewsRequest) (err error) {
	request := req.ToDomain()
	if req.File != nil {
		ctx := context.Background()
		val, err := _i.Minio.UploadFile(ctx, *req.File)
		if err != nil {
			return err
		}
		req.Image = *val
	}

	return _i.Repo.Create(request)

}
func (_i *newsService) Update(id uint64, req request.NewsRequest) (err error) {
	gallery, err := _i.Repo.FindOne(uint(id))
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = _i.Minio.DeleteFile(ctx, gallery.Image)
	if err != nil {
		return err
	}

	val, err := _i.Minio.UploadFile(ctx, *req.File)
	if err != nil {
		return err
	}
	req.Image = *val
	return _i.Repo.Update(uint(id), req.ToDomain())
}
