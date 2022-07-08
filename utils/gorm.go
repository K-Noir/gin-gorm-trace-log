package utils

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	callBackBeforeName = "core:before"
	callBackAfterName  = "core:after"
	startTime          = "_start_time"
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
	// 开始前 主要用于统计时间
	// _ = db.Callback().Create().Before("gorm:before_create").Register(callBackBeforeName, before)
	// _ = db.Callback().Query().Before("gorm:query").Register(callBackBeforeName, before)
	// _ = db.Callback().Delete().Before("gorm:before_delete").Register(callBackBeforeName, before)
	// _ = db.Callback().Update().Before("gorm:setup_reflect_value").Register(callBackBeforeName, before)
	// _ = db.Callback().Row().Before("gorm:row").Register(callBackBeforeName, before)
	// _ = db.Callback().Raw().Before("gorm:raw").Register(callBackBeforeName, before)

	// 结束后 主要用于组合日志
	_ = db.Callback().Create().After("gorm:after_create").Register(callBackAfterName, after)
	_ = db.Callback().Query().After("gorm:after_query").Register(callBackAfterName, after)
	_ = db.Callback().Delete().After("gorm:after_delete").Register(callBackAfterName, after)
	_ = db.Callback().Update().After("gorm:after_update").Register(callBackAfterName, after)
	_ = db.Callback().Row().After("gorm:row").Register(callBackAfterName, after)
	_ = db.Callback().Raw().After("gorm:raw").Register(callBackAfterName, after)
	return
}

func before(db *gorm.DB) {
	db.InstanceSet(startTime, time.Now())
	return
}

func after(db *gorm.DB) {
	ctx := db.Statement.Context
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)
	logrus.WithContext(ctx).Info(sql)
	return
}
