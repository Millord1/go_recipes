package repository

import "go_recipes/models"

type QuantityRepository struct {
	Mysql MySQLRepository
}

type QuantityInterface interface {
	Save(quant models.Quantity) error
	Update(quant *models.Quantity) error
	FindById(id uint) (*models.Quantity, error)
}

func (repo QuantityRepository) Update(quant *models.Quantity) error {
	if err := repo.Mysql.db.Model(&quant).Updates(quant).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo QuantityRepository) Save(quant models.Quantity) error {
	if err := repo.Mysql.db.Create(&quant).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo QuantityRepository) FindById(id uint) (*models.Quantity, error) {
	var quant models.Quantity
	if err := repo.Mysql.db.Where("ID = ?", id).Find(&quant).Error; err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return &quant, nil
}
