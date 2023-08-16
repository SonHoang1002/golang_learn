package server_side

import (
	"awesomeProject/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *sql.DB
)

func InitMySQLConnection() *sql.DB {
	once.Do(func() {
		connectionMySQL, err := sql.Open("mysql", "root:son12345@(127.0.0.1:3306)/golang_db_1")
		connectionMySQL.Ping()
		if err != nil {
			fmt.Println("error from connection ", err)
		}
		instance = connectionMySQL
	})
	return instance
}

func CreateUser(name string, age int) error {
	db := InitMySQLConnection()
	var _username = name
	var _age = age
	var _createdAt = time.Now()

	query := `
	INSERT INTO users (username, age, created_at) VALUES (?,?,?)
	`
	_, err := db.Exec(query, _username, _age, _createdAt)
	return err
}
func GetAllUser() []models.UserModel {

	var results []models.UserModel
	db := InitMySQLConnection()
	query := `
     SELECT username, age FROM users
`
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error executing query:", err)
		return results
	}
	defer rows.Close()

	for rows.Next() {
		var user models.UserModel
		err := rows.Scan(&user.Username, &user.Age)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		results = append(results, user)
	}
	return results

}

func UpdateUser(id int, username string, age int) error {
	var _username = username
	var _age = age
	var _createdAt = time.Now()
	db := InitMySQLConnection()
	query := `
       UPDATE users SET username = ?, age = ?,created_at = ? WHERE id = ? `
	_, err := db.Exec(query, _username, _age, _createdAt, id)
	return err
}

func DeleteUser(id int) error {
	db := InitMySQLConnection()
	query := `
       DELETE FROM users WHERE id = ? `
	_, err := db.Exec(query, id)
	return err
}
