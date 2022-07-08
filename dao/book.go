package dao

import (
	"context"
	"hello/model"
	"hello/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Use(utils.NewGormTrace())
	// db.AutoMigrate(&model.Book{})
	// db.Create(&model.Book{Id: 1, Name: "things", Auther: "Lucifer"})
}

func Get(ctx context.Context) {
	logrus.WithContext(ctx).Info("Hello")
	var book model.Book
	db.WithContext(ctx).First(&book, 1)
	logrus.WithContext(ctx).Info(book)

	db.WithContext(ctx).First(&book, 2)
	logrus.WithContext(ctx).Info(book)
}
