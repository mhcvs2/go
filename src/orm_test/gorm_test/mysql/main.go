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
	db, err := gorm.Open("mysql", "root:123@tcp(ali:3306)/gorm_test?charset=utf8&parseTime=true&loc=Local")
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
	fmt.Println("product.Code: ", product.Code)
	db.First(&product, "code = ?", "L1212") // find product with code l1212
	fmt.Println("product.Code: ", product.Code)
	// Update - update product's price to 2000
	fmt.Println("update")
	db.Model(&product).Update("Price", 2000)
	fmt.Println("delete")
	// Delete - delete product
	db.Delete(&product)
}

func t2() {
	db, err := gorm.Open("mysql", "root:123@tcp(ali:3306)/gorm_test?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connect success")
	}
	defer db.Close()

	// Read
	var product Product
	db.First(&product) // find product with id 1
	fmt.Println("product.Code: ", product.Code)
}

func t3() {
	db, err := gorm.Open("mysql", "root:123@tcp(ali:3306)/gorm_test?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connect success")
	}
	defer db.Close()

	// Read
	var product Product
	db.Where("code = ?", "L1212").First(&product)
	fmt.Println("product.Price: ", product.Price)
	fmt.Println("product.Code: ", product.Code)
}

func main() {
	t3()
}