package repository

import "go_recipes/models"

type StepRepository struct {
	Mysql MySQLRepository
}

func (repo StepRepository) Update(step *models.Step) error {
	if err := repo.Mysql.db.Model(&step).Updates(step).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo StepRepository) Save(step models.Step) error {
	if err := repo.Mysql.db.Create(&step).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo StepRepository) FindById(id uint) (*models.Step, error) {
	var step models.Step
	if err := repo.Mysql.db.Where("ID = ?", id).Find(&step).Error; err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return &step, nil
}
