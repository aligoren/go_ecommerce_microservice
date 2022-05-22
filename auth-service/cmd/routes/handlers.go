package routes

import (
	"errors"
	"github.com/aligoren/go_ecommerce_microservice/auth-service/models"
	"github.com/aligoren/go_ecommerce_microservice/auth-service/repository"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func GetUserByID(ctx *fiber.Ctx) error {

	id, _ := ctx.ParamsInt("id", 0)

	user, err := repository.GetOne(id)

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "User not found",
			Error:      true,
			Data:       nil,
		})
	}

	if user.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "User not found",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(user)
}

func GetAllUsers(ctx *fiber.Ctx) error {

	users, err := repository.GetAllUsers()

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "User not found",
			Error:      true,
			Data:       err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(users)
}

func Create(ctx *fiber.Ctx) error {

	user := new(models.User)

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Request body couldn't parse",
			Error:      true,
			Data:       nil,
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "User couldn't create",
			Error:      true,
			Data:       nil,
		})
	}

	u := models.User{
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  string(password),
		CreatedAt: time.Now(),
	}

	_, err = repository.Create(&u)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "User couldn't create",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(models.ResponseModel{
		StatusCode: 200,
		Message:    "User created successfully",
		Error:      false,
		Data:       nil,
	})
}

func Update(ctx *fiber.Ctx) error {

	userModel := new(models.User)

	if err := ctx.BodyParser(userModel); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Request body couldn't parse",
			Error:      true,
			Data:       nil,
		})
	}

	user, err := repository.GetOne(userModel.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "User couldn't find",
			Error:      true,
			Data:       nil,
		})
	}

	user.FirstName = userModel.FirstName
	user.LastName = userModel.LastName
	user.Email = userModel.Email

	if userModel.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(userModel.Password), 12)

		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
				StatusCode: 500,
				Message:    "User couldn't update",
				Error:      true,
				Data:       nil,
			})
		}

		user.Password = string(password)
	}

	_, err = repository.Update(&user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "User couldn't update",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(models.ResponseModel{
		StatusCode: 200,
		Message:    "User updated successfully",
		Error:      false,
		Data:       nil,
	})

}

func Delete(ctx *fiber.Ctx) error {

	userModel := new(models.User)

	if err := ctx.BodyParser(userModel); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Response body couldn't parse",
			Error:      true,
			Data:       nil,
		})
	}

	user, err := repository.GetOne(userModel.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "User couldn't delete",
			Error:      true,
			Data:       nil,
		})
	}

	_, err = repository.Delete(&user)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "User couldn't delete",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(models.ResponseModel{
		StatusCode: 200,
		Message:    "User deleted successfully",
		Error:      false,
		Data:       nil,
	})

}

func Login(ctx *fiber.Ctx) error {
	userModel := new(models.User)

	if err := ctx.BodyParser(userModel); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Response body couldn't parse",
			Error:      true,
			Data:       nil,
		})
	}

	user, err := repository.GetByEmail(userModel.Email)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "User couldn't delete",
			Error:      true,
			Data:       nil,
		})
	}

	isPasswordValid, err := checkPassword(user.Password, userModel.Password)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "User couldn't login",
			Error:      true,
			Data:       nil,
		})
	}

	if !isPasswordValid {
		return ctx.Status(http.StatusUnauthorized).JSON(models.ResponseModel{
			StatusCode: 401,
			Message:    "User couldn't found",
			Error:      false,
			Data:       nil,
		})
	}

	user.Password = ""

	return ctx.Status(http.StatusOK).JSON(models.ResponseModel{
		StatusCode: 200,
		Message:    "User logged in successfully",
		Error:      false,
		Data:       user,
	})
}

func checkPassword(registeredPassword string, formPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(registeredPassword), []byte(formPassword))

	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
