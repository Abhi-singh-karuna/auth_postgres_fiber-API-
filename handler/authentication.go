package handler

import (
	"os"
	"time"

	"github.com/abhishek_singh/database"
	"github.com/abhishek_singh/model"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

var UN int

func CheckPassword(psd, hsh string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hsh), []byte(psd))
	return err == nil
}

func UserbyEmail(e string) (*model.User, error) {
	db := database.DB
	var user model.User
	err := db.Where(&model.User{Email: e}).Find(&user).Error

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func UserbyUsername(u string) (*model.User, error) {
	db := database.DB
	var user model.User

	err := db.Where(&model.User{Username: u}).Find(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func Login(c *fiber.Ctx) error {

	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput

	var ud UserData

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	identity := input.Identity
	pass := input.Password

	email, err := UserbyEmail(identity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "data": err})
	}

	user, err := UserbyUsername(identity)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
	}

	if email == nil && user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	}

	if email == nil {
		ud = UserData{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}
	} else {
		ud = UserData{
			ID:       email.ID,
			Username: email.Username,
			Email:    email.Email,
			Password: email.Password,
		}
	}

	if !CheckPassword(pass, ud.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	/*  */
	db := database.DB

	tk := new(model.Token)
	c.BodyParser(tk)
	tk.Token = t
	tk.Username = user.Username
	UN = int(tk.ID)
	db.Create(&tk)

	/* 	var utk model.Token
	   	c.BodyParser(utk)
	   	id := utk.ID

	   	var uui model.Token
	   	c.BodyParser(&uui)

	   	tk.Token = uui. */

	/*  */

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success login",
		"token":   t,
	})
}

func LogOut(c *fiber.Ctx) error {

	db := database.DB
	var user model.Token
	db.Find(&user, UN)
	if user.Username == "" {
		db.First(&user, UN)
		db.Delete(&user)
	} else {
		return c.JSON(fiber.Map{
			"status":  "false",
			"message": "user is not  login",
		})

	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Success logout",
	})
}
