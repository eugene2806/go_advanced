package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	config            *Config
	db                *sql.DB
	userRepository    *UserRepository
	articleRepository *ArticleRepository
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (s *Storage) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseURI)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {

		return err
	}

	s.db = db

	log.Println("Database connection...")

	return nil
}

func (s *Storage) Close() {
	s.db.Close()
}

func (s *Storage) User() *UserRepository {
	if s.userRepository != nil {

		return s.userRepository
	}

	s.userRepository = &UserRepository{
		storage: s,
	}

	return s.userRepository
}

func (s *Storage) Article() *ArticleRepository {
	if s.articleRepository != nil {

		return s.articleRepository
	}

	s.articleRepository = &ArticleRepository{
		storage: s,
	}

	return s.articleRepository
}
