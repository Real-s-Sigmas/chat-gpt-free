package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	usersTable    = "users"
	chatsTable    = "chats"
	messagesTable = "messages"
)

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	if err := AddTables(db); err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AddTables(db *sqlx.DB) error {
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s("+
		"id uuid, "+
		"username text, "+
		"email text, "+
		"password_hash text, "+
		")", usersTable)
	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s("+
		"id uuid, "+
		"type text, "+
		"title text, "+
		"user_id uuid, "+
		")", chatsTable)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s("+
		"id uuid, "+
		"text, "+
		"isAI bool, "+
		"id_chat uuid"+
		")", messagesTable)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
