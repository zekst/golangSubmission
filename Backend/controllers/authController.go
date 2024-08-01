package controllers

import (
	"fmt"
	"someName/database"
	"someName/models"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

var roles = []string{"sales", "accountant", "hr", "admin"}

func assignPermissions(role string) []string {
	var permissions []string

	role = strings.ToLower(strings.TrimSpace(role))

	if role == roles[0] || role == roles[2] || role == roles[3] {
		permissions = append(permissions, "read", "write")
	} else if role == roles[1] {
		permissions = append(permissions, "read")
	}

	return permissions
}

func assignDepartments(role string) []string {

	var department []string

	role = strings.ToLower(strings.TrimSpace(role))

	switch role {
	case roles[0]:
		department = append(department, "Customer Management", "Billing Management")
	case roles[1]:
		department = append(department, "Payroll Management", "Billing Management")
	case roles[2]:
		department = append(department, "Payroll Management")
	case roles[3]:
		department = append(department, "User Management")
	}

	return department
}

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["email"] == "" || data["password"] == "" || data["role"] == "" {
		return c.JSON("required fields can not be empty")
	}

	password, error := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	if error != nil {
		return c.JSON(fmt.Sprintf("something wrong while encrypting the passowrd %s", error))
	}

	user := models.User{
		Name:        data["name"],
		Email:       data["email"],
		Password:    password,
		Role:        data["role"],
		Permissions:  assignPermissions(data["role"]),
		Departments: assignDepartments(data["role"]),
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": user,
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
		"info": user,
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
