package handlers

import (
	"strconv"

	"github.com/catalinfl/infinite-scroll/models"
	"github.com/catalinfl/infinite-scroll/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {

	rows, err := utils.Database.Query("SELECT * FROM dfdf")

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	defer rows.Close()

	result := make([]models.User, 0)

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		result = append(result, user)
	}

	return c.JSON(result)
}

func CreateUser(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := utils.Database.QueryRow("INSERT INTO dfdf (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&user.ID)

	if result != nil {
		return c.Status(500).SendString(result.Error())
	}

	return c.JSON(&models.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user := models.User{}

	row := utils.Database.QueryRow("SELECT * FROM dfdf WHERE id = $1", id)

	err := row.Scan(&user.ID, &user.Name, &user.Email)

	if err != nil {
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}

func PaginationUser(c *fiber.Ctx) error {

	pageStr := c.Query("page")

	page, err := strconv.Atoi(pageStr)

	if err != nil {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	rows, err := utils.Database.Query("SELECT * FROM dfdf LIMIT 2 OFFSET $1", page)

	if err != nil {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	defer rows.Close()

	result := make([]models.User, 0)

	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email)

		if err != nil {
			c.Status(fiber.ErrBadRequest.Code)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		result = append(result, user)
	}

	return c.JSON(result)
}

func CursorPaginationPosts(c *fiber.Ctx) error {
	pageStr := c.Query("cursor")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	limit := 2
	offset := page * limit

	rows, err := utils.Database.Query("SELECT * FROM posts ORDER BY id LIMIT $1 OFFSET $2", limit, offset)

	if err != nil {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	defer rows.Close()

	result := make([]models.Posts, 0)

	for rows.Next() {
		post := models.Posts{}
		err := rows.Scan(&post.ID, &post.Name, &post.Description)

		if err != nil {
			c.Status(fiber.ErrBadRequest.Code)
			return c.JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		result = append(result, post)
	}

	return c.JSON(fiber.Map{
		"posts": result,
	})
}

func CreateDescription(c *fiber.Ctx) error {

	var post models.Posts

	if err := c.BodyParser(&post); err != nil {
		c.Status(fiber.ErrBadRequest.Code)
		return c.JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	result := utils.Database.QueryRow("INSERT INTO posts (name, description) VALUES ($1, $2) RETURNING id", post.Name, post.Description).Scan(&post.ID)

	if result != nil {
		return c.Status(500).SendString(result.Error())
	}

	return c.JSON(&models.Posts{
		ID:          post.ID,
		Name:        post.Name,
		Description: post.Description,
	})
}
