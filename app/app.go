package app

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	cartCon "github.com/movie-rental-app/components/cart/controller"
	cartSer "github.com/movie-rental-app/components/cart/service"
	movieCon "github.com/movie-rental-app/components/movie/controller"
	movieSer "github.com/movie-rental-app/components/movie/service"
	userCon "github.com/movie-rental-app/components/user/controller"
	userSer "github.com/movie-rental-app/components/user/service"

	repo "github.com/movie-rental-app/repository"
	router "github.com/movie-rental-app/route"
)

type App struct {
	Environment Environment
	Port        string
	Host        string
	User        string
	Password    string
	DBName      string
	Repo        repo.Repository
}

func NewApp() *App {
	return &App{
		Environment: LOCAL,
		Port:        "8080",
		Host:        "localhost",
		User:        "root",
		Password:    "",
		DBName:      "movie_rental",
	}
}

func (a *App) Init() error {

	a.Environment = Environment(os.Getenv("ENVIRONMENT"))

	if a.Environment == "" {
		a.Environment = LOCAL
	}
	a.Port = os.Getenv("PORT")
	a.Host = os.Getenv("DB_HOST")
	a.User = os.Getenv("DB_USER")
	a.Password = os.Getenv("DB_PASSWORD")
	a.DBName = os.Getenv("DB_NAME")

	// conect to db
	db, err := Connect()
	if err != nil {
		return err
	}

	fmt.Println("Connected to db", db)

	//create a repository
	repo := repo.GetRepository(db)

	//create service and controllers
	userService := userSer.NewUserService(repo, db)
	userController := userCon.NewUserController(userService)

	movieService := movieSer.NewMovieService(repo, db)
	movieController := movieCon.NewMovieController(movieService)

	cartService := cartSer.NewCartService(repo, db)
	cartController := cartCon.NewCartController(cartService)

	//start server
	baseRouter := router.GetRouter()
	router.RegisterUserRoutes(baseRouter, userController)
	router.RegisterMovieRoutes(baseRouter, movieController)
	router.RegisterCartRoutes(baseRouter, cartController)

	baseRouter.Run("localhost:8080")
	fmt.Println("running on port 8080")

	return nil
}

func Connect() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	environment := os.Getenv("ENVIRONMENT")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPass == "" || dbName == "" {
		log.Fatal("Missing required environment variables for database connection")
		return nil, errors.New("Missing ENV keys")
	}

	sslMode := "disable"
	if environment != "local" {
		sslMode = "require"
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, sslMode,
	)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, errors.New("Failed to connect to database.")

	}

	// Optional: Enable logging for GORM
	db.LogMode(true)

	log.Println("Connected to PostgreSQL database successfully!")
	return db, nil
}
