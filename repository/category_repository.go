package repository

import "go_recipes/models"

type CategoryRepository struct {
	Mysql MySQLRepository
}

func (repo CategoryRepository) Update(cat *models.Category) error {
	if err := repo.Mysql.db.Model(&cat).Updates(cat).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo CategoryRepository) Save(cat models.Category) error {
	if err := repo.Mysql.db.Create(&cat).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo CategoryRepository) FindById(id uint) (*models.Category, error) {
	var cat models.Category
	if err := repo.Mysql.db.Where("ID = ?", id).Find(&cat).Error; err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return &cat, nil
}
