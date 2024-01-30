package service

import (
	"motionserver/app/module/gallery/repository"
	"motionserver/app/module/gallery/request"
	"motionserver/app/module/gallery/response"
	"motionserver/utils/paginator"
)

type galleryService struct {
	Repo repository.GalleryRepository
}

type GalleryService interface {
	All(req request.GalleriesRequest) (categories []*response.Gallery, paging paginator.Pagination, err error)
	Store(req request.GalleryRequest)(err error)
}

func NewGalleryService(repo repository.GalleryRepository) GalleryService {
	return &galleryService{
		Repo: repo,
	}

}

func (_i *galleryService) All(req request.GalleriesRequest) (galleries []*response.Gallery, paging paginator.Pagination, err error) {
	results, paging, err := _i.Repo.GetProducts(req)
	if err != nil {
		return
	}
	for _, v := range results {
		galleries = append(galleries, response.FromDomain(v))
	}
	return
	
}

func (_i *galleryService)Store(req request.GalleryRequest)(err error){
	return _i.Repo.Create(req.ToDomain())
}