package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func t1() {
	db, err := gorm.Open("mysql", "root:root.123@tcp(services.gcloud.srcb.com:30306)/test?charset=utf8&parseTime&loc=Local")
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connect success")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	fmt.Println(product.Code)
	db.First(&product, "code = ?", "L1212") // find product with code l1212
	fmt.Println(product.Code)
	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)
}

func t2() {
	db, err := gorm.Open("mysql", "root:root.123@tcp(services.gcloud.srcb.com:30306)/test?charset=utf8&parseTime&loc=Local")
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connect success")
	}
	defer db.Close()

	// Read
	var product Product
	db.First(&product, 1) // find product with id 1
	fmt.Println(product.Code)
}

func main() {
	t2()
}