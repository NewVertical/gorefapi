package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io"
	"log"
	"os"
)

type Database struct {
	ConnectInfo ConnectInfo `json:"database"`
}
type ConnectInfo struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Pass     string `json:"password"`
	Database string `json:"database"`
}

func (c ConnectInfo) New() ConnectInfo {
	jsonFile, err := os.Open("./config/db.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println("Error Closing File")
		}
	}(jsonFile)

	byteValue, _ := io.ReadAll(jsonFile)
	var database Database

	marshall_err := json.Unmarshal(byteValue, &database)
	if marshall_err != nil {
		fmt.Println("Json File improperly formatted")
		return ConnectInfo{}
	}

	c.Host = database.ConnectInfo.Host
	c.Port = database.ConnectInfo.Port
	c.User = database.ConnectInfo.User
	c.Pass = database.ConnectInfo.Pass
	c.Database = database.ConnectInfo.Database

	return c
}
func (c ConnectInfo) ConnectionString() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Pass, c.Host, c.Port, c.Database)
}
func (c ConnectInfo) Connect() *sql.DB {
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Pass, c.Database)

	var err error
	DB, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}
	return DB
}

func (c ConnectInfo) Execute(handler func(db *sql.DB) ([]any, error)) ([]any, error) {
	db := c.Connect()
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(db)
	return handler(db)
}
func (c ConnectInfo) ExecuteSelect(query string, handler func(rows *sql.Rows) ([]any, error)) ([]any, error) {
	db := c.Connect()
	defer func(DB *sql.DB) {
		err := DB.Close()
		if err != nil {

		}
	}(db)
	rows, err := db.Query("SELECT * FROM lessons LIMIT $1 OFFSET $2", 10, 0)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	return handler(rows)
}
