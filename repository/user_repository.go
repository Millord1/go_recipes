package repository

import "go_recipes/models"

type UserRepository struct {
	Mysql MySQLRepository
}

func (repo UserRepository) Update(user *models.User) error {
	if err := repo.Mysql.db.Model(&user).Updates(user).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo UserRepository) Save(user models.User) error {
	if err := repo.Mysql.db.Create(&user).Error; err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

func (repo UserRepository) FindById(id uint) (*models.User, error) {
	var user models.User
	if err := repo.Mysql.db.Where("ID = ?", id).Find(&user).Error; err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	return &user, nil
}
