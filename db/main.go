package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Id struct {
	ID uint `gorm:"primarykey"`
}

func (id *Id) TableName() string {
	return "test"
}

func main() {

	dbUrl := getDbUrl()

	fmt.Println("Testing database")

	dsn := "root:password@tcp(" + dbUrl + ":3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connected to database")

	fmt.Println("Creating a dummy test table")
	createError := db.Exec("CREATE TABLE IF NOT EXISTS test (ID INT AUTO_INCREMENT PRIMARY KEY)").Error
	if createError != nil {
		panic(fmt.Errorf("error creating test table: %s", createError))
	}

	fmt.Println("Inserting a row into the test table")
	insertError := db.Create(&Id{}).Error
	if insertError != nil {
		panic(fmt.Errorf("error inserting test table: %s", insertError))
	}

	fmt.Println("Reading a row from the test table")
	newRow := Id{}
	readError := db.First(&newRow).Error
	if readError != nil {
		panic(fmt.Errorf("error reading test table: %s", readError))
	}

	fmt.Println("Updating a row from the test table")
	deleteError := db.Delete(&newRow).Error
	if deleteError != nil {
		panic(fmt.Errorf("error deleting test table: %s", deleteError))
	}

	fmt.Println("Deleting the test table")
	deleteTableError := db.Exec("DROP TABLE IF EXISTS `test`").Error
	if deleteTableError != nil {
		panic(fmt.Errorf("error deleting test table: %s", deleteTableError))
	}

}

func getDbUrl() string {
	if os.Getenv("DB_URL") == "" {
		return "localhost"
	}
	return os.Getenv("DB_URL")
}
