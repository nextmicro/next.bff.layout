package data

import (
	"context"

	"github.com/nextmicro/logger"
	shortUrlV1 "github.com/nextmicro/next.data.layout/api/shorturl/v1"
	"next.bff.layout/internal/biz"
)

type ShortUrlRepo struct {
	data *Data
	log  logger.Logger
}

func NewShortUrlRepo(data *Data, logger logger.Logger) biz.ShortUrlRepo {
	return &ShortUrlRepo{
		data: data,
		log:  logger,
	}
}

func (s *ShortUrlRepo) Expand(ctx context.Context, shorten string) (string, error) {
	reply, err := s.data.shortUrl.Expand(ctx, &shortUrlV1.ExpandRequest{
		Shorten: shorten,
	})
	if err != nil {
		return "", err
	}

	return reply.Url, nil
}
func (s *ShortUrlRepo) Shorten(ctx context.Context, url string) (string, error) {
	reply, err := s.data.shortUrl.Shorten(ctx, &shortUrlV1.ShortenRequest{
		Url: url,
	})
	if err != nil {
		return "", err
	}

	return reply.Shorten, nil
}
