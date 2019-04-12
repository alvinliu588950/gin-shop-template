package model

import (
	"teashop/db"
	"time"
)

type Product struct {
	Id uint `gorm:"primary_key"`
	Name string
	Price float32
	Img string // url to image storage
	CreatedAt time.Time  
	UpdatedAt time.Time  
	DeletedAt *time.Time 
}

func init() {
	err := db.Orm.AutoMigrate(&Product{}).Error
	if err != nil {
		panic(err)
	}
}

func(p *Product) Create() (err error) {
	return db.Orm.Create(p).Error
}

func(p *Product) GetAll() (products []Product, err error) {
	err = db.Orm.Find(&products).Error
	if err != nil {
		panic(err)
	}
	return products, err
}

func(p *Product) Update() (err error) {
	// Save will include all fields when perform the Updating SQL, even it is not changed
	return db.Orm.Save(p).Error
}

func(p *Product) Delete() (err error) {
	// Soft delete
	return db.Orm.Delete(p).Error
}