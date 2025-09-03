package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	config             *Config
	db                 *sql.DB
	userRepository     *UserRepository
	arcticleRepository *ArcticleRepository
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	log.Println("DB open succesfully")
	storage.db = db
	return nil
}

func (storage *Storage) Close() {
	storage.db.Close()
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

func (s *Storage) Article() *ArcticleRepository {
	if s.arcticleRepository != nil {
		return s.arcticleRepository
	}
	s.arcticleRepository = &ArcticleRepository{
		storage: s,
	}
	return s.arcticleRepository
}
