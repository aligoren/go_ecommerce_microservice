package repository

import (
	"github.com/aligoren/go_ecommerce_microservice/auth-service/database"
	"github.com/aligoren/go_ecommerce_microservice/auth-service/models"
)

func GetOne(ID int) (models.User, error) {

	user := models.User{}

	if result := database.DB.Db.Find(&user).Where("id = ?", ID); result.Error != nil {
		return user, result.Error
	}

	return user, nil

}
