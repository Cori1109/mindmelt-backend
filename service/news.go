package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Cori1109/mindmelt-backend/models"
	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type News struct {
	Content string `json:"content"`
	Createdat string `json:"createdAt"`
}

func (r *Repository) CreateNews(context *fiber.Ctx) error {
	news := News{}

	err := context.BodyParser(&news)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	news.Createdat = fmt.Sprintf("%s", time.Now())

	validator := validator.New()
	err = validator.Struct(News{})

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": err},
		)
		return err
	}

	err = r.DB.Create(&news).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not create news"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "news has been successfully added"})

	return nil
}

func (r *Repository) DeleteNews(context *fiber.Ctx) error {
	newsModel := &models.News{}

	id := context.Params("news_id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "NEWS_ID cannot be empty"})
		return nil
	}

	err := r.DB.Delete(newsModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not delete news"})
		return err.Error
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "news has been successfully deleted"})

	return nil
}

func (r *Repository) UpdateNews(context *fiber.Ctx) error {
	id := context.Params("news_id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{"message": "NEWS_ID cannot be empty"})
		return nil
	}
	fmt.Printf("news_id: %s, %T", id, id)

	newsModel := &models.News{}

	news := News{}

	err := context.BodyParser(&news)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Model(newsModel).Where("news_id = ?", id).Updates(news).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not update news"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "news has been successfully updated"})

	return nil
}

func (r *Repository) GetNews(context *fiber.Ctx) error {
	newsModels := &[]models.News{}

	err := r.DB.Find(newsModels).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{"message": "could not get news"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "news gotten successfully",
		"data": newsModels,
	})

	return nil
}
