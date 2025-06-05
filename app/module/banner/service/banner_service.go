package service

import (
	"context"
	"log"
	"motionserver/app/module/banner/repository"
	"motionserver/app/module/banner/request"
	"motionserver/app/module/banner/response"
	"motionserver/internal/bootstrap/minio"
)

type bannerService struct {
	Repo  repository.BannerRepository
	Minio *minio.Minio
}

type BannerService interface {
	Show() (banner *response.Banner, err error)
	Store(req request.BannerRequest) (err error)
}

func NewBannerService(repo repository.BannerRepository, Minio *minio.Minio) BannerService {
	return &bannerService{
		Repo:  repo,
		Minio: Minio,
	}

}

func (_i *bannerService) Show() (banner *response.Banner, err error) {
	result, err := _i.Repo.GetBanner()
	if err != nil {
		return
	}
	if result == nil {
		return nil, err
	}
	ctx := context.Background()
	img := _i.Minio.GenerateLink(ctx, result.Image)
	banner = response.FromDomain(result, img)
	return
}

func (_i *bannerService) Store(req request.BannerRequest) (err error) {
	result, err := _i.Repo.GetBanner()
	if err != nil {
		return
	}
	ctx := context.Background()
	if result != nil {

		err = _i.Minio.DeleteFile(ctx, result.Image)
		if err != nil {
			log.Println("Error ", err.Error())
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
