package service

import (
	"context"
	"hello/dao"

	"github.com/sirupsen/logrus"
)

func List(ctx context.Context) {
	logrus.WithContext(ctx).Info("Hello")
	dao.Get(ctx)
}
