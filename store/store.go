package store

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Store struct {
	config            *Config
	db                *sql.DB
	userRepository    *UserRepository
	articleRepository *ArticleRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func (store *Store) Open() error {
	db, err := sql.Open("postgres", store.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	store.db = db
	log.Println("Database connection created successfully!")
	return nil
}

func (store *Store) Close() {
	store.db.Close()
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		store: s,
	}
	return nil
}

func (s *Store) Article() *ArticleRepository {
	if s.articleRepository != nil {
		return s.articleRepository
	}
	s.articleRepository = &ArticleRepository{
		store: s,
	}
	return nil
}
