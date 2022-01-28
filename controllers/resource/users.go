package resources

import (
	"github.com/gofiber/fiber"
	"github.com/hoanggggg5/learn-go/config"
	"github.com/hoanggggg5/learn-go/models"
	"github.com/huuhait/go-learn/controllers/entities"
)

func userToEntity(user *models.User) entities.User {
	return entities.User{
		ID:        user.ID,
		UID:       user.UID,
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		UpdatedAt: user.UpdatedAt,
		CreatedAt: user.CreatedAt,
	}
}

func CreateUser(c *fiber.Ctx) {
	type Payload struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	payload := new(Payload)

	if err := c.BodyParser(payload); err != nil {
		c.Status(422).JSON("LỖI")
	}

	user := &models.User{
		Email:    payload.Email,
		Password: payload.Password,
	}

	config.Database.Create(&user)

	return c.JSON(userToEntity(user))
}
