package repository

import (
	"go_recipes/models"
)

type QuantityRepository struct {
	Mysql MySQLRepository
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

func (repo QuantityRepository) GetOrCreate(quantity models.Quantity) (*models.Quantity, error) {
	var qtt models.Quantity
	err := repo.Mysql.db.FirstOrCreate(&qtt, quantity).Error
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	return &qtt, nil
}
