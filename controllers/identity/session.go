package identity

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/learn-go/config"
	"github.com/hoanggggg5/learn-go/controllers/entities"
	"github.com/hoanggggg5/learn-go/models"
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

func Login(c *fiber.Ctx) error {
	type Payload struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	payload := new(Payload)

	if err := c.BodyParser(payload); err != nil {
		return c.JSON("Lỗi")
	}

	user := &models.User{
		Email:    payload.Email,
		Password: payload.Password,
	}

	session, _ := config.Store.Get(c)

	if result := config.Database.First(&user, "Email = ?", payload.Email); result.Error != nil {
		session.Destroy()
		return c.Status(422).JSON("error")
	}

	session.Set("email", user.Email)

	if err := session.Save(); err != nil {
		return c.JSON("LỖI lưu cookie")
	}

	return c.Status(201).JSON(userToEntity(user))
}
