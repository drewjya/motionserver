package service

import (
	"context"
	"motionserver/app/module/compro/repository"
	"motionserver/app/module/compro/request"
	"motionserver/app/module/compro/response"
	"motionserver/internal/bootstrap/minio"
)

type comproService struct {
	Repo  repository.ComproRepository
	Minio *minio.Minio
}

type ComproService interface {
	Show() (compro *response.Compro, err error)
	Store(req request.ComproRequest) (err error)
}

func NewComproService(repo repository.ComproRepository, Minio *minio.Minio) ComproService {
	return &comproService{
		Repo:  repo,
		Minio: Minio,
	}

}

func (_i *comproService) Show() (compro *response.Compro, err error) {
	result, err := _i.Repo.GetCompro()
	if err != nil {
		return
	}
	if result == nil {
		return nil, err
	}
	ctx := context.Background()
	img := _i.Minio.GenerateLink(ctx, result.Image)
	compro = response.FromDomain(result, img)
	return
}

func (_i *comproService) Store(req request.ComproRequest) (err error) {
	result, err := _i.Repo.GetCompro()
	if err != nil {
		return
	}
	ctx := context.Background()
	if result != nil {

		err = _i.Minio.DeleteFile(ctx, result.Image)
		if err != nil {
			return
		}
	}
	if req.File != nil {
		ctx := context.Background()
		val, err := _i.Minio.UploadFile(ctx, *req.File)
		if err != nil {
			return err
		}
		req.Image = *val
	}
	return _i.Repo.Create(req.ToDomain())
}
