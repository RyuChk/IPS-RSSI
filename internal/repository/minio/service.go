package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	wireminio "git.cie.com/ips/wire-provider/minio"
	"github.com/ZecretBone/ips-rssi-service/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=service.go -destination=mock_minio/mock_minio.go -package=mock_minio
type Service interface {
	Get(ctx context.Context, filepath string) ([]byte, error)
	Upload(ctx context.Context, filePath string, file []byte) error
	GetSignedURL(ctx context.Context, fileName, filepath string) (string, error)
	ListObject()
}

type service struct {
	client *minio.Client
	cfg    config.MinioConfig
}

func ProvideMinioService(config wireminio.Config, cfg config.MinioConfig) Service {
	conn, err := wireminio.NewConnection(&config)
	if err != nil {
		panic(err)
	}
	return &service{
		client: conn.Client(),
		cfg:    cfg,
	}
}

func (s *service) GetSignedURL(ctx context.Context, fileName, filepath string) (string, error) {

	// Set request parameters
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	ttl := time.Duration(1000) * time.Second

	// Gernerate presigned get object url.
	presignedURL, err := s.client.PresignedGetObject(ctx, s.cfg.Bucket, filepath, ttl, reqParams)
	if err != nil {
		log.Error().Str("file_path", filepath).Err(err).Msg("error creating download url")
		return "", err
	}
	log.Info().Str("PresignedURL", presignedURL.RawPath)
	return presignedURL.String(), nil
}

func (s *service) ListObject() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	objectCh := s.client.ListObjects(ctx, s.cfg.Bucket, minio.ListObjectsOptions{
		Prefix:    "",
		Recursive: true,
	})
	for object := range objectCh {
		if object.Err != nil {
			fmt.Println(object.Err)
			return
		}
		fmt.Println(object)
	}
}

func (s *service) Get(ctx context.Context, filepath string) ([]byte, error) {
	object, err := s.client.GetObject(ctx, s.cfg.Bucket, filepath, minio.GetObjectOptions{})
	if err != nil {
		log.Error().Str("filepath", filepath).Err(err).Msg("error download template cannot find document")
		return nil, err
	}

	defer object.Close()
	file, err := io.ReadAll(object)
	if err != nil {
		log.Error().Str("filepath", filepath).Err(err).Msg("error reading pdf template")
		return nil, err
	}

	return file, nil
}

func (s *service) Upload(ctx context.Context, filePath string, file []byte) error {
	media := bytes.NewReader(file)

	uploadInfo, err := s.client.PutObject(ctx, s.cfg.Bucket, filePath, media, int64(len(file)), minio.PutObjectOptions{})
	if err != nil {
		log.Error().Str("file_path", filePath).Err(err).Msg("error upload file")
		return err
	}

	log.Info().Str("file_path", filePath).Any("upload_info", uploadInfo).Msg("successfully upload file")
	return nil
}
