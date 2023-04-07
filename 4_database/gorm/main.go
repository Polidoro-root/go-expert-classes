package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Category struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	Product []Product `gorm:"many2many:products_categories";`
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID string
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:products_categories";`
	SerialNumber SerialNumber
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/expert"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	db.Create(&Product{
		Name:  "Notebook",
		Price: 1000.0,
	})

	var product []Product
	db.First(&product, 1)

	var products Product
	db.Find(&products)

	category := Category{Name: "Eletronicos"}
	db.Create(&category)

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: "1",
	})

	db.Create(&Product{
		Name:       "Notebook",
		Price:      1000.0,
		CategoryID: category.ID,
	})

	// [respectively] belongs to | has one
	db.Preload("Category").Preload("SerialNumber").Find(&products)

	var categories []Category

	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	tx := db.Begin()
	var c Category
	err = tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&c, 1).Error

	if err != nil {
		panic(err)
	}

	c.Name = "Eletronicos"

	tx.Debug().Save(&c)
	tx.Commit()
}
