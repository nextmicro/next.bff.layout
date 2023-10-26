package biz

import (
	"context"

	"github.com/nextmicro/logger"
)

// ShortUrlRepo is a Greater repo.
type ShortUrlRepo interface {
	Expand(ctx context.Context, shorten string) (string, error)
	Shorten(ctx context.Context, url string) (string, error)
}

type ShortUrlUsecase struct {
	repo ShortUrlRepo
	log  logger.Logger
}

func NewShortUrlUsecase(repo ShortUrlRepo, logger logger.Logger) *ShortUrlUsecase {
	return &ShortUrlUsecase{repo: repo, log: logger}
}

func (uc *ShortUrlUsecase) Expand(ctx context.Context, shorten string) (string, error) {
	uc.log.WithContext(ctx).Infof("Expand: %v", shorten)
	return uc.repo.Expand(ctx, shorten)
}
