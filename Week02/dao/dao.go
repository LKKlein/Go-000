package dao

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct {
	ID   string
	Name string
	Age  int
}

const (
	USER_NAME = "root"
	PASS_WORD = "root"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "test"
	CHARSET   = "utf8"
)

var DB *sql.DB
var ErrorNotFound = errors.New("record not found")

func init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

	var err error
	DB, err = sql.Open("mysql", dbDSN)
	if err != nil {
		log.Panicf("mysql data source err: %+v\n", err.Error())
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(20)
	DB.SetConnMaxLifetime(100 * time.Second)

	if err := DB.Ping(); err != nil {
		log.Panicf("mysql connect err: %+v\n", err.Error())
	}
}

func FindUserByID(id string) (*User, error) {
	row := DB.QueryRow("select id,name,age from t_user where id = ?", id)
	resUser := User{}
	if err := row.Scan(&resUser.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrorNotFound
		}
		return nil, errors.Wrap(err, "QueryUserByIDErr")
	}
	return &resUser, nil
}
