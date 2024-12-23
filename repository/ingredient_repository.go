package repository

import "go_recipes/models"

type IngRepository struct {
	Mysql MySQLRepository
}

func (repo IngRepository) Update(ing *models.Ingredient) error {
	if err := repo.Mysql.db.Model(&ing).Updates(ing).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo IngRepository) Save(ing models.Ingredient) error {
	if err := repo.Mysql.db.Create(&ing).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo IngRepository) FindById(id uint) (*models.Ingredient, error) {
	var ing models.Ingredient
	if err := repo.Mysql.db.Where("ID = ?", id).Find(&ing).Error; err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return &ing, nil
}
