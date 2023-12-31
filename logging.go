package main

import (
	"context"
	"micro-grcp/types"
	"time"

	"github.com/sirupsen/logrus"
)

type loggingService struct {
	next PriceFetcher
}

func NewLoggingService(next PriceFetcher) PriceFetcher {
	return &loggingService{
		next: next,
	}
}

func (s *loggingService) FetchPrice(ctx context.Context, ticker string) (price float64, err error) {
	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"requestID": ctx.Value(types.RequestID),
			"took":      time.Since(begin),
			"err":       err,
			"price":     price,
		}).Info("fechPrice")
	}(time.Now())

	return s.next.FetchPrice(ctx, ticker)
}
