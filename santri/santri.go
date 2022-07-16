package santri

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lshinkuro/go-fiber-tutorial/database"
)

type Person struct {
	gorm.Model
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	Origin    string `json:"origin"`
}

type Job struct {
	gorm.Model
	Type   string `json:"type" validate:"required,min=3,max=32"`
	Salary int    `json:"salary" validate:"required,number"`
}

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required,min=3,max=32"`
	IsActive *bool  `json:"isActive" validate:"required"`
	Email    string `json:"email" validate:"required,email,min=6,max=32"`
	Job      Job    `json:"job" validate:"dive"`
}
type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(user User) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func (ctx *User) checkValidateRequest(c *fiber.Ctx) (*User, error) {
	if err := c.BodyParser(ctx); err != nil {
		return ctx, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := ValidateStruct(*ctx)
	if errors != nil {
		return ctx, c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return ctx, nil
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users []User
	db.Find(&users)
	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var user User
	db.Find(&user, id)
	return c.JSON(user)
}

func NewUser(c *fiber.Ctx) error {
	db := database.DBConn
	var user User
	requestBody := new(User)

	requestBody.checkValidateRequest(c)

	user.Name = requestBody.Name
	user.Email = requestBody.Email
	user.IsActive = requestBody.IsActive
	user.Job = requestBody.Job
	db.Create(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	db := database.DBConn

	id := c.Params("id")
	var user User
	db.Find(&user, id)

	requestBody := new(User)

	requestBody.checkValidateRequest(c)

	user.Name = requestBody.Name
	user.Email = requestBody.Email
	user.IsActive = requestBody.IsActive
	user.Job = requestBody.Job

	db.Save(&user)
	return c.JSON(&user)
}
