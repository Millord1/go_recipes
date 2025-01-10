package repository

import (
	"go_recipes/models"
)

type DishRepository struct {
	Mysql MySQLRepository
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

/* func (repo DishRepository) GetWithIngredients(id uint) (*models.Dish, error) {
var dish models.Dish

res := repo.Mysql.db.Where("dishes.id = ?", id).Joins(
	"INNER JOIN quantities ON quantities.dish_id = dishes.id").Joins(
	"INNER JOIN ingredients ON ingredients.id = quantities.ingredient_id").Find(dish) */

/* res := repo.Mysql.db.Where("ID = ?", id).Preload("Quantities.Ingredients").Find(&dish) */
// dishes.name, quantities.num, quantities.unit, ingredients.name
/* 	res := repo.Mysql.db.Raw(`SELECT * FROM dishes
LEFT JOIN quantities ON quantities.dish_id = dishes.id
LEFT JOIN ingredients ON ingredients.id = quantities.ingredient_id
WHERE dishes.id = ?`, id).Scan(dish) */
/* 	if res.Error != nil {
		logger.Sugar.Error(res.Error)
		return nil, res.Error
	}
	testJson, _ := json.Marshal(dish)
	fmt.Println(string(testJson))

	return &dish, nil
} */

func (repo DishRepository) GetOrCreate(name string) (*models.Dish, error) {
	var dish models.Dish
	err := repo.Mysql.db.FirstOrCreate(&dish, models.Dish{Name: name}).Error
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	return &dish, nil
}
