package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type Anomalie struct {
	Id         uint32
	Session_id string
	Frequency  float64
	Time       time.Time
}

func Base() *gorm.DB {
	dsn := "host=localhost user=delilahl password=123456 dbname=postgres port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&Anomalie{})
	return db
}

func main() {
	db := Base()
	db.Create(&Anomalie{Session_id: "qweerttytuy-trytu-tyuiubjj-dfghjughxj", Frequency: 1.1111111, Time: time.Now().UTC()})
	fmt.Println("See you base ;-)")
}
