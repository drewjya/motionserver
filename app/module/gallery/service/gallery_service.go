package service

import (
	"context"
	"motionserver/app/module/gallery/repository"
	"motionserver/app/module/gallery/request"
	"motionserver/app/module/gallery/response"
	"motionserver/internal/bootstrap/minio"
	"motionserver/utils/paginator"
)

type galleryService struct {
	Repo  repository.GalleryRepository
	Minio *minio.Minio
}

type GalleryService interface {
	All(req request.GalleriesRequest) (categories []*response.Gallery, paging paginator.Pagination, err error)
	Store(req request.GalleryRequest) (err error)
	Update(id uint64, req request.GalleryRequest) (err error)
}

func NewGalleryService(repo repository.GalleryRepository, Minio *minio.Minio) GalleryService {
	return &galleryService{
		Repo:  repo,
		Minio: Minio,
	}

}

func (_i *galleryService) All(req request.GalleriesRequest) (galleries []*response.Gallery, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.GetProducts(req)
	if err != nil {
		return
	}
	ctx := context.Background()
	for _, v := range results {
		img := _i.Minio.GenerateLink(ctx, v.Image)
		galleries = append(galleries, response.FromDomain(v, img))
	}
	return

}

func (_i *galleryService) Store(req request.GalleryRequest) (err error) {
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

func (_i *galleryService) Update(id uint64, req request.GalleryRequest) (err error) {
	gallery, err := _i.Repo.FindOne(id)
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
	return _i.Repo.Update(id, req.ToDomain())

}
