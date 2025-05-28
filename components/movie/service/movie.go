package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/movie-rental-app/errors"
	model "github.com/movie-rental-app/models"
	repo "github.com/movie-rental-app/repository"
	uuid "github.com/satori/go.uuid"
)

type MovieService struct {
	repo *repo.Repository
	db   *gorm.DB
}

func NewMovieService(repo *repo.Repository, db *gorm.DB) *MovieService {
	return &MovieService{
		repo: repo,
		db:   db,
	}
}

func (s *MovieService) GetMovies(ctx *gin.Context, movies *[]model.Movie, limit, offset *int) error {

	//create unit of work
	uow := repo.NewUnitOfWork(s.db, true)

	defer uow.RollBack()

	processors := s.AddSearchQueries(ctx)

	err := s.repo.GetAll(&movies, model.Movie{}, limit, offset, processors)
	if err != nil {
		return err
	}

	uow.Commit()

	return nil
}

func (s *MovieService) GetMovieByID(ctx *gin.Context, movie *model.Movie, movieID uuid.UUID) error {
	//create unit of work
	uow := repo.NewUnitOfWork(s.db, true)

	defer uow.RollBack()

	if movieID == uuid.Nil {
		return errors.NewHTTPError(400, "Invalid uuid provided.")
	}

	err := s.repo.GetByID(&movie, movieID)
	if err != nil {
		return err
	}
	uow.Commit()

	return nil
}

func (s *MovieService) AddSearchQueries(ctx *gin.Context) []repo.QueryProcessor {

	var queries []repo.QueryProcessor

	if l := ctx.Query("genre"); l != "" {
		queries = append(queries, func(db *gorm.DB) *gorm.DB {
			return db.Where(`"movie"."genre" = ?`, l)
		})
	}

	if l := ctx.Query("year"); l != "" {
		queries = append(queries, func(db *gorm.DB) *gorm.DB {
			return db.Where(`EXTRACT(YEAR FROM "movie"."release_date") = ?`, l)
		})
	}

	if l := ctx.Query("title"); l != "" {
		queries = append(queries, func(db *gorm.DB) *gorm.DB {
			return db.Where(`"movie"."title" ILIKE ?`, l+"%")
		})
	}

	return queries

}
