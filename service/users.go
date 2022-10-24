package service

import (
	"fmt"
	"net/http"

	"github.com/Cori1109/mindmelt-backend/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type User_data struct {
	User_id string `json:"USER_ID"`
	First_name string `json:"FIRST_NAME"`
	Last_name string `json:"LAST_NAME"`
	User_role string `json:"USER_ROLE"`
	User_pass string `json:"USER_PASS"`
	Overall_score string `json:"OVERALL_SCORE"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateUser(context *fiber.Ctx) error {
	user := User_data{}

	err := context.BodyParser(&user)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	validator := validator.New()
	err = validator.Struct(User_data{})

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": err},
		)
		return err
	}

	err = r.DB.Create(&user).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create user"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "user has been successfully added"})

	return nil
}

func (r *Repository) UpdateUser(context *fiber.Ctx) error {
	id := context.Params("vr_app_session")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "VR_APP_SESSION cannot be empty"})
		return nil
	}

	userModel := &models.User_data{}

	user := User_data{}

	err := context.BodyParser(&user)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Model(userModel).Where("vr_app_session = ?", id).Updates(user).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not update user"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "user has been successfully updated"})

	return nil
}

func (r *Repository) DeleteUser(context *fiber.Ctx) error {
	userModel := &models.User_data{}

	id := context.Params("vr_app_session")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "VR_APP_SESSION cannot be empty"})
		return nil
	}

	err := r.DB.Delete(userModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete user"})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "user has been successfully deleted"})

	return nil
}

func (r *Repository) GetUsers(context *fiber.Ctx) error {
	userModels := &[]models.User_data{}

	err := r.DB.Find(userModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get users"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "users gotten successfully",
		"data": userModels,
	})

	return nil
}

func (r *Repository) GetUserByID(context *fiber.Ctx) error {
	id := context.Params("id")
	fmt.Println("VR_APP_SESSION: ", id)

	userModel := &models.User_data{}

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "VR_APP_SESSION cannot be empty"})
		return nil
	}

	err := r.DB.Where("vr_app_session = ?", id).First(userModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get user"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "users id gotten successfully",
		"data": userModel,
	})

	return nil
}