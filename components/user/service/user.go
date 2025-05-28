package service

import (
	"context"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/movie-rental-app/errors"
	model "github.com/movie-rental-app/models"
	repo "github.com/movie-rental-app/repository"
	uuid "github.com/satori/go.uuid"
)

type UserService struct {
	repo *repo.Repository
	db   *gorm.DB
}

func NewUserService(repo *repo.Repository, db *gorm.DB) *UserService {
	return &UserService{
		repo: repo,
		db:   db,
	}
}

func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {

	//create a new unit of work
	uow := repo.NewUnitOfWork(s.db, false)
	defer uow.RollBack()

	var users []model.User
	var limit int = 1
	var offset int = 0

	var queryProcessors []repo.QueryProcessor

	processor := func(db *gorm.DB) *gorm.DB {
		return db.Where(`"user"."email" = ?`, user.Email)
	}

	queryProcessors = append(queryProcessors, processor)

	//check if user already exists
	err := s.repo.GetAll(&users, model.User{}, &limit, &offset, queryProcessors)
	if err != nil {
		return err
	}

	if len(users) > 0 {
		return errors.NewHTTPError(400, "User with this email already exists.")
	}

	user.ID = uuid.NewV4()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.DeletedAt = nil

	err = s.repo.Create(user)
	if err != nil {
		return err
	}

	uow.Commit()

	return nil
}
