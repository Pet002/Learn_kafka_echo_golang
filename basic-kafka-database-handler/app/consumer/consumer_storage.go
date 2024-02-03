package consumer

import (
	"database/sql"
	"fmt"
	"log"
)

type consumerStorage struct {
	db *sql.DB
}

func NewConsumerStorage(db *sql.DB) *consumerStorage {
	return &consumerStorage{
		db: db,
	}
}

func (s *consumerStorage) Insert(u User) error {
	log.Println(u.Name + " " + u.Surname)
	stmt := fmt.Sprintf("INSERT INTO users (name, surname) VALUES ('%s', '%s')", u.Name, u.Surname)
	_, err := s.db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}
