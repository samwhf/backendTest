package postgres

import (
	"fmt"
	"log"

	"github.com/jackc/pgx"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/samwhf/backendTest/config/flags"
)

type Client struct {
	conn *pgx.Conn
	DB   *gorm.DB
}

var client *Client

// New is a postgress database constructor
func New(cfg flags.Configuration) *Client {
	url := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Database)

	db, err := gorm.Open("postgres", url)
	if err != nil {
		log.Fatalf("could not create postgres connection, err=%v ", err)
	}

	client = &Client{DB: db}

	return client
}

func (c *Client) Close() error {
	return c.DB.Close()
}

func GetClient() *Client {
	return client
}
