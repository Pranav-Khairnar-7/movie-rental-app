package service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/movie-rental-app/errors"
	model "github.com/movie-rental-app/models"
	repo "github.com/movie-rental-app/repository"
	uuid "github.com/satori/go.uuid"
)

type CartService struct {
	repo *repo.Repository
	db   *gorm.DB
}

func NewCartService(repo *repo.Repository, db *gorm.DB) *CartService {
	return &CartService{
		repo: repo,
		db:   db,
	}
}

func (s *CartService) CreateCart(ctx *gin.Context, cart *model.Cart) error {

	//create unit of work
	uow := repo.NewUnitOfWork(s.db, false)

	defer uow.RollBack()

	var user model.User

	//check if user exist
	err := s.repo.GetByID(&user, cart.UserID)
	if err != nil {
		return nil
	}

	if user.ID != cart.UserID {
		return errors.NewHTTPError(400, "User does not exist!")
	}

	//check if cart already exist
	var tempCart []model.Cart
	var queryProcessors []repo.QueryProcessor

	queryProcessors = append(queryProcessors, func(db *gorm.DB) *gorm.DB {
		return db.Where(`"cart"."user_id" = ?`, cart.UserID)
	})
	err = s.repo.GetAll(&tempCart, model.Cart{}, nil, nil, queryProcessors)
	if err != nil {
		return nil
	}

	if len(tempCart) > 0 {
		return errors.NewHTTPError(400, "User already has a cart created.")
	}

	//create cart
	cart.ID = uuid.NewV4()
	cart.CreatedAt = time.Now()
	cart.UpdatedAt = time.Now()
	cart.DeletedAt = nil

	for i := 0; i < len(cart.CartItem); i++ {
		cart.CartItem[i].ID = uuid.NewV4()
		cart.CartItem[i].CreatedAt = time.Now()
		cart.CartItem[i].UpdatedAt = time.Now()
		cart.CartItem[i].DeletedAt = nil
	}

	err = s.repo.Create(cart)
	if err != nil {
		return err
	}

	uow.Commit()
	return nil
}

func (s *CartService) GetCartByID(ctx *gin.Context, cart *model.Cart, cartID uuid.UUID) error {
	//create unit of work
	uow := repo.NewUnitOfWork(s.db, true)

	defer uow.RollBack()

	if cartID == uuid.Nil {
		return errors.NewHTTPError(400, "Invalid uuid provided.")
	}

	var queries []repo.QueryProcessor

	queries = append(queries, func(db *gorm.DB) *gorm.DB {
		return db.
			Joins(`JOIN cart_item ON cart_item.cart_id = cart.id`).
			Joins(`JOIN movie ON movie.id = cart_item.movie_id`).
			Preload("CartItem").
			Preload("CartItem.Movie")
	})

	queries = append(queries, func(db *gorm.DB) *gorm.DB {
		return db.Where(`"cart"."id" = ?`, cartID)
	})

	err := s.repo.GetWithFilters(&cart, queries)
	if err != nil {
		return err
	}

	uow.Commit()

	return nil
}
