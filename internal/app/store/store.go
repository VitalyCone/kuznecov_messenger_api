package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config                *Config
	db                    *sql.DB
	chatRepository        *ChatRepository
	userRepository        *UserRepository
	chatMessageRepository *ChatMessageRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

func NewDB() *sql.DB {
	return &sql.DB{}
}

func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.databaseURL)

	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) Chat() *ChatRepository {
	if s.chatRepository != nil {
		return s.chatRepository
	}

	s.chatRepository = &ChatRepository{
		store: s,
	}

	return s.chatRepository
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

func (s *Store) ChatMessage() *ChatMessageRepository {
	if s.chatMessageRepository != nil {
		return s.chatMessageRepository
	}

	s.chatMessageRepository = &ChatMessageRepository{
		store: s,
	}

	return s.chatMessageRepository
}
