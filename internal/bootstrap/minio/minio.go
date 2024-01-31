package minio

import (
	"context"
	"mime/multipart"
	"motionserver/utils/config"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog"
)

type Minio struct {
	Log   zerolog.Logger
	Cfg   *config.Config
	Minio *minio.Client
}

func NewMinio(
	Log zerolog.Logger,
	Cfg *config.Config,
) *Minio {
	return &Minio{
		Log: Log,
		Cfg: Cfg,
	}
}

func (_min *Minio) ConnectMinio(ctx context.Context) {
	cfg := _min.Cfg.Minio

	minioClient, err := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Username, cfg.Password, ""),
		Secure: true,
	})

	if err != nil {
		_min.Log.Error().Err(err).Msg("An unknown error occurred when to connect the minio-server!")
	} else {
		_min.Log.Info().Msg("Connected the minio-server succesfully!")
	}

	err = minioClient.MakeBucket(ctx, cfg.Bucket, minio.MakeBucketOptions{})
	if err != nil {
		exist, err := minioClient.BucketExists(ctx, cfg.Bucket)
		if err == nil && exist {
			_min.Minio = minioClient
			return
		}
		_min.Log.Error().Err(err).Msg("Unknown error occured")
	}
	_min.Minio = minioClient
}

func (_i *Minio) UploadFile(ctx context.Context, file multipart.FileHeader) (*string, error) {
	bucket := _i.Cfg.Minio.Bucket

	buffer, err := file.Open()
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(file.Filename)
	defer buffer.Close()
	contentType := file.Header["Content-Type"][0]
	fileName := strings.ReplaceAll(uuid.New().String(), "-", "") + ext
	fileSize := file.Size
	info, err := _i.Minio.PutObject(ctx, bucket, fileName, buffer, fileSize, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		return nil, err
	}

	return &info.Key, nil
}

func (_i *Minio) GenerateLink(ctx context.Context, fileName string) string {
	img := ""
	extension := filepath.Ext(fileName)
	extension = strings.Replace(extension, ".", "", 1)

	reqParams := make(url.Values)

	reqParams.Set("response-content-type", "image/"+extension)
	bucketName := _i.Cfg.Minio.Bucket
	presignUrl, err := _i.Minio.PresignedGetObject(ctx, bucketName, fileName, time.Duration(3600)*time.Second, reqParams)
	if err != nil {
		return img
	}
	val := presignUrl.String()
	return val
}

func (_i *Minio) DeleteFile(ctx context.Context, objectName string) error {
	bucketName := _i.Cfg.Minio.Bucket
	err := _i.Minio.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
