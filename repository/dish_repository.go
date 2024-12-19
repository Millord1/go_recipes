package repository

import "go_recipes/models"

type DishRepository struct {
	Mysql MySQLRepository
}

type DishInterface interface {
	Save(dish models.Dish) error
	Update(dish *models.Dish) error
	FindById(id uint) (*models.Dish, error)
}

func (repo DishRepository) Update(dish *models.Dish) error {
	if err := repo.Mysql.db.Model(&dish).Updates(dish).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo DishRepository) Save(dish models.Dish) error {
	if err := repo.Mysql.db.Create(&dish).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo DishRepository) FindById(id uint) (*models.Dish, error) {
	var dish models.Dish
	if err := repo.Mysql.db.Where("ID = ?", id).Find(&dish).Error; err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return &dish, nil
}
