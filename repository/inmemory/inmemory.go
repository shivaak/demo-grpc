package inmemory

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shivaak/demo-grpc/entity"
	"github.com/shivaak/demo-grpc/mapper"
	"github.com/shivaak/demo-grpc/model"
	"go.uber.org/zap"
)

type Repository interface {
	CreateUser(user *entity.User) (uuid.UUID, error)
	ListUser() ([]*entity.User, error)
	GetUser(uuid uuid.UUID) (*entity.User, error)
	CreateBook(book *entity.Book) (uuid.UUID, error)
	ListBook() ([]*entity.Book, error)
}

type store struct {
	memUserStore map[uuid.UUID]*model.User
	memBookStore map[uuid.UUID]*model.Book
	logger       *zap.SugaredLogger
}

func NewInMemoryRepo(l *zap.SugaredLogger) Repository {
	return &store{
		memUserStore: make(map[uuid.UUID]*model.User),
		memBookStore: make(map[uuid.UUID]*model.Book),
		logger:       l,
	}
}

func (s *store) GetUser(id uuid.UUID) (*entity.User, error) {
	if s.memUserStore[id] == nil {
		return nil, errors.New("invalid user id")
	}
	return mapper.ModelToUser(s.memUserStore[id]), nil
}

func (s *store) ListUser() ([]*entity.User, error) {
	var valueList []*entity.User

	for _, u := range s.memUserStore {
		user := mapper.ModelToUser(u)
		valueList = append(valueList, user)
	}
	s.logger.Info("Total number of users : ", len(valueList))
	return valueList, nil
}

func (s *store) CreateUser(user *entity.User) (uuid.UUID, error) {
	if user == nil {
		return uuid.Nil, errors.New("cannot save nil user")
	}
	user.Id = uuid.New()
	s.memUserStore[user.Id] = mapper.UserToModel(user)
	return user.Id, nil
}

func (s *store) CreateBook(book *entity.Book) (uuid.UUID, error) {
	if book == nil {
		return uuid.Nil, errors.New("cannot save nil book")
	}
	book.Id = uuid.New()
	s.memBookStore[book.Id] = mapper.BookToModel(book)
	return book.Id, nil
}

func (s *store) ListBook() ([]*entity.Book, error) {
	var valueList []*entity.Book

	for _, b := range s.memBookStore {
		valueList = append(valueList, mapper.ModelToBook(b))
	}
	s.logger.Info("Total number of books : ", len(valueList))
	return valueList, nil
}
