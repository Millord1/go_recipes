package repository

import "go_recipes/models"

type RecipeRepository struct {
	Mysql MySQLRepository
}

func (repo RecipeRepository) Update(recipe *models.Recipe) error {
	if err := repo.Mysql.db.Model(&recipe).Updates(recipe).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo RecipeRepository) Save(recipe models.Recipe) error {
	if err := repo.Mysql.db.Create(&recipe).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo RecipeRepository) FindById(id uint) (*models.Recipe, error) {
	var recipe models.Recipe
	if err := repo.Mysql.db.Where("ID = ?", id).Find(&recipe).Error; err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return &recipe, nil
}
