package repository

import (
	"github.com/aligoren/go_ecommerce_microservice/auth-service/database"
	"github.com/aligoren/go_ecommerce_microservice/auth-service/models"
)

func GetOne(ID int) (models.User, error) {

	user := models.User{}

	if result := database.DB.Db.Select("id", "first_name", "last_name", "email", "created_at", "updated_at").Where("id = ?", ID).First(&user); result.Error != nil {
		return user, result.Error
	}

	return user, nil

}

func GetByEmail(email string) (models.User, error) {

	user := models.User{}

	if result := database.DB.Db.Select("id", "first_name", "last_name", "email", "created_at", "updated_at", "password").Where("email = ?", email).Find(&user); result.Error != nil {
		return user, result.Error
	}

	return user, nil

}

func GetAllUsers() ([]models.User, error) {
	var users []models.User

	if result := database.DB.Db.Select("id", "first_name", "last_name", "email", "created_at", "updated_at").Find(&users); result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func Create(user *models.User) (*models.User, error) {
	if err := database.DB.Db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func Update(user *models.User) (*models.User, error) {
	if err := database.DB.Db.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func Delete(user *models.User) (*models.User, error) {
	if err := database.DB.Db.Delete(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
