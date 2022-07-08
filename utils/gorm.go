package utils

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	callBackAfterName = "after:logrus"
)

type GormTrace struct {
}

func NewGormTrace() *GormTrace {
	return &GormTrace{}
}

func (plugin *GormTrace) Name() string {
	return "tracePlugin"
}

func (plugin *GormTrace) Initialize(db *gorm.DB) (err error) {
	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

func after(db *gorm.DB) {
	ctx := db.Statement.Context
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	logrus.WithContext(ctx).Info(sql)
	return
}
