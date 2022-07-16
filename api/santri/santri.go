package santri

import (
	"html"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/lshinkuro/go-fiber-tutorial/api/database"
	"golang.org/x/crypto/bcrypt"
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
	Password string `json:"password" validate:"required,min=4,max=100"`
	Job      Job    `json:"job" validate:"dive"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)

	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) PrepareToSave() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

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
		return &User{}, c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := ValidateStruct(*ctx)
	if errors != nil {
		return &User{}, c.Status(fiber.StatusBadRequest).JSON(errors)
	}
	return ctx, nil
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*User, error) {
	var err error
	users := User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &User{}, err
	}
	return &users, err
}

func GetUsers(c *fiber.Ctx) error {
	db := database.DBConn
	var users User
	users.FindAllUsers(db)
	// db.Find(&users)
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
	var user User
	requestBody := new(User)

	requestBody.checkValidateRequest(c)

	user.Name = requestBody.Name
	user.Email = requestBody.Email
	user.IsActive = requestBody.IsActive
	user.Job = requestBody.Job
	user.SaveUser(database.DBConn)
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
