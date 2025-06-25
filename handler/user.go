package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thanthtooaung-coding/user-crud-go-fiber/internal/domain/user"
	"gorm.io/gorm"
	"net/http"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	u := new(user.User)
	if err := c.BodyParser(u); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if result := h.db.Create(&u); result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.Status(http.StatusCreated).JSON(u)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	var users []user.User
	if result := h.db.Find(&users); result.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var u user.User
	if result := h.db.First(&u, id); result.Error != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(u)
}
