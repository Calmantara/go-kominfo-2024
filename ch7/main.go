package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID        uint64    `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	DoB       time.Time `json:"dob"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `json:"deleted_at,omitempty"`

	// relations
	UserMediaSocials []UserMediaSocial `json:"user_media_socials" gorm:"foreignKey:UserID;references:ID"`
}

type UserMediaSocial struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

// example
type UserSocialMedia struct {
	ID       uint64 `json:"id" gorm:"column:id"`
	Username string `json:"username" gorm:"column:username"`
	Title    string `json:"title" gorm:"column:title"`
	Url      string `json:"url" gorm:"column:url"`
}

func main() {
	host := "127.0.0.1"
	port := "15432"
	user := "postgres"
	password := "mysecretpassword"
	dbname := "go_kominfo"

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// select users
	db.
		Table("users").
		Where("id = ?", 1).
		Update("last_name", "gorm")

	users := []User{}
	db.
		Table("users").
		Model(&User{}).
		Preload("UserMediaSocials", func(db *gorm.DB) *gorm.DB {
			// nested preload:
			// db.Preload("OtherData")

			// select in nested preload
			// we can use where clause also
			return db.Select("user_id, title, url")
		}).
		Where("id = $1", 1).
		Find(&users)
	fmt.Println(users)

	// join user and user social media
	ums := []UserSocialMedia{}
	db.
		Table("users u").
		Select("u.id, u.username, ums.title, ums.url").
		Joins("JOIN user_media_socials ums on u.id = ums.user_id").
		Find(&ums)
	fmt.Println(ums)

}

func sqlway() {
	host := "127.0.0.1"
	port := "15432"
	user := "postgres"
	password := "mysecretpassword"
	dbname := "go_kominfo"

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// connect ke database postgres
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	// menutup koneksi database ketika
	// program go sudah selesai dieksekusi
	defer conn.Close()

	// get data
	rows, err := conn.Query(
		`SELECT 
			id, 
			username, 
			first_name, 
			last_name, dob, 
			deleted_at 
		FROM users WHERE deleted_at is null 
		AND id = $1`, 1)
	if err != nil {
		panic(err)
	}
	// mapping data
	users := []User{}
	for rows.Next() {
		user := User{}
		rows.Scan(
			&user.ID,
			&user.Username,
			&user.FirstName,
			&user.LastName,
			&user.DoB,
			&user.DeletedAt)

		users = append(users, user)
	}
	fmt.Println(users)

	// $1 dan $2 adalah argumen
	res, err := conn.Exec(
		"UPDATE users SET last_name = $1 WHERE id = $2",
		"golang kominfo", 1)
	if err != nil {
		panic(err)
	}

	val, _ := res.RowsAffected()
	if val <= 0 {
		fmt.Println("no rows updated")
	}
}
