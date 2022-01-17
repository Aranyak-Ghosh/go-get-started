package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID          string `gorm:"primary_key;size:36"`
	Name        string
	Email       string `gorm:"unique_index"`
	ImgUrl      string
	Country     string
	Public      bool `gorm:"default:true"`
	DateOfBirth time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	Follows []User `gorm:"many2many:user_follows"`
}

func initDatabase() *gorm.DB {
	var connString = "host=localhost user=postgres password=rexTqGq*tuk7Xs&a dbname=gorm port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
	return db
}

func Seed(context *gorm.DB) {
	context.Create(&User{
		ID:          uuid.NewString(),
		Name:        "Aranyak Ghosh",
		Email:       "aranyakghosh@gmail.com",
		Country:     "India",
		Public:      true,
		DateOfBirth: time.Date(1996, time.December, 16, 0, 0, 0, 0, time.UTC),
	})

	context.Create(&User{
		ID:          uuid.NewString(),
		Name:        "Shailika Garg",
		Email:       "shailika.garg@gmail.com",
		Country:     "India",
		Public:      true,
		DateOfBirth: time.Date(1992, time.August, 21, 0, 0, 0, 0, time.UTC),
	})
}

func Migrate(context *gorm.DB) {
	context.AutoMigrate(&User{})
}

func List(context *gorm.DB) []User {
	var users []User
	context.Model(&User{}).Preload("Follows").Find(&users)
	return users
}

func Follow(context *gorm.DB, follower string, following string) {
	context.Model(&User{ID: follower}).Association("Follows").Append(&User{ID: following})
}

func QueryJoin(context *gorm.DB) ([]User, int64) {
	var users []User
	var count int64

	query := context.Joins("JOIN user_follows on user_follows.user_id = users.id").Joins("JOIN users as followed on user_follows.follow_id = followed.id").Where("followed.name like ?", "%Shailika%").Preload("Follows").Model(&User{})

	query.Count(&count)
	query.Preload("Follows").Find(&users)
	return users, count
}

func main() {
	dbContext := initDatabase()
	Migrate(dbContext)
	// Seed(dbContext)

	users, count := QueryJoin(dbContext)
	fmt.Printf("%+v\n", users)
	fmt.Printf("%v\n", count)

	// Follow(dbContext, users[0].ID, users[1].ID)
}
