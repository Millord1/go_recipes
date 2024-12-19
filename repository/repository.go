package repository

import (
	"go_recipes/models"
	"go_recipes/utils"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLRepository struct {
	User     string
	Password string
	Protocol string
	Address  string
	Port     string
	Name     string
	db       *gorm.DB
}

var logger utils.Logger = utils.NewLogger("repository.log")

func getMySQLRepo() MySQLRepository {
	return MySQLRepository{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Protocol: os.Getenv("DB_PROTOCOL"),
		Address:  os.Getenv("DB_ADDRESS"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
}

func DbConnect(envFile string) *MySQLRepository {
	// Init connection to database from specified env file variables
	err := godotenv.Load(envFile)

	if err != nil {
		logger.Sugar.Fatal("Error loading " + envFile + " file")
	}

	repo := getMySQLRepo()

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: repo.User + ":" + repo.Password + "@" + repo.Protocol +
			"(" + repo.Address + ":" + repo.Port + ")/" + repo.Name +
			"?charset=utf8mb4&parseTime=True&loc=Local",
	}), &gorm.Config{})

	repo.db = db

	if err != nil {
		logger.Sugar.Fatal("Failed to connect database")
	}

	return &repo
}

func Migrate() error {
	repo := DbConnect(utils.GetEnvFile().Name)
	return repo.db.Transaction(func(tx *gorm.DB) error {
		// User table
		if err := checkUserTable(tx); err != nil {
			logger.Sugar.Fatal(err)
			return err
		}

		tableNames := map[string]interface{}{
			"recipe":     models.Recipe{},
			"quantity":   models.Quantity{},
			"ingredient": models.Ingredient{},
			"dish":       models.Dish{},
			"category":   models.Category{},
		}

		for name, model := range tableNames {
			if err := checkTable(tx, name, &model); err != nil {
				logger.Sugar.Fatal(err)
				return err
			}
		}
		return nil
	})
}

func checkTable(db *gorm.DB, tableName string, model interface{}) error {
	if !db.Migrator().HasTable(tableName) {
		return db.AutoMigrate(&model)
	}
	return nil
}

func checkUserTable(db *gorm.DB) error {
	var user = &models.User{}
	if db.Migrator().HasTable("users") {
		// if table exists, check potential missing columns
		if !db.Migrator().HasColumn(user, "dishes") {
			return db.Migrator().AddColumn(user, "dishes")
		}
		if !db.Migrator().HasColumn(user, "favorites") {
			return db.Migrator().AddColumn(user, "favorites")
		}
		return nil
	}

	// if table doesn't exist, create it
	return db.Migrator().AutoMigrate(user)
}
