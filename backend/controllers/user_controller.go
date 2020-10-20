package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Sujitnapa21/app/ent"
	"github.com/Sujitnapa21/app/ent/role"
	"github.com/Sujitnapa21/app/ent/sex"
	"github.com/Sujitnapa21/app/ent/title"
	"github.com/Sujitnapa21/app/ent/user"
	"github.com/gin-gonic/gin"
)

// UserController defines the struct for the user controller
type UserController struct {
	client *ent.Client
	router gin.IRouter
}

// User defines the struct for the user controller
type User struct {
	Email string
	Password string
	Title int
	Name string
	Sex int
	Birthday string
	Age string
	Tel string
	Role int

}

// CreateUser handles POST requests for adding user entities
// @Summary Create user
// @Description Create user
// @ID create-user
// @Accept   json
// @Produce  json
// @Param user body ent.User true "User entity"
// @Success 200 {object} ent.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [post]
func (ctl *UserController) CreateUser(c *gin.Context) {
	obj := User{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}

	t, err := ctl.client.Title.
		Query().
		Where(title.IDEQ(int(obj.Title))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "title not found",
		})
		return
	}

	s, err := ctl.client.Sex.
		Query().
		Where(sex.IDEQ(int(obj.Sex))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "sex not found",
		})
		return
	}

	r, err := ctl.client.Role.
		Query().
		Where(role.IDEQ(int(obj.Role))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "role not found",
		})
		return
	}

	birthday, err := time.Parse(time.RFC3339, obj.Birthday)
	
	u, err := ctl.client.User.
		Create().
		SetEmail(obj.Email).
		SetPassword(obj.Password).
		SetTitle(t).
		SetName(obj.Name).
		SetSex(s).
		SetAge(obj.Age).
		SetBirthday(birthday).
		SetTel(obj.Tel).
		SetRole(r).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetUser handles GET requests to retrieve a user entity
// @Summary Get a user entity by ID
// @Description get user by ID
// @ID get-user
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} ent.User
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [get]
func (ctl *UserController) GetUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListUser handles request to get a list of user entities
// @Summary List user entities
// @Description list user entities
// @ID list-user
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users [get]
func (ctl *UserController) ListUser(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	users, err := ctl.client.User.
		Query().
		WithTitle().
		WithSex().
		WithRole().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

// DeleteUser handles DELETE requests to delete a user entity
// @Summary Delete a user entity by ID
// @Description get user by ID
// @ID delete-user
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [delete]
func (ctl *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.User.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateUser handles PUT requests to update a user entity
// @Summary Update a user entity by ID
// @Description update user by ID
// @ID update-user
// @Accept   json
// @Produce  json
// @Param id path int true "User ID"
// @Param user body ent.User true "User entity"
// @Success 200 {object} ent.User
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /users/{id} [put]
func (ctl *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := User{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "user binding failed",
		})
		return
	}

	t, err := ctl.client.Title.
		Query().
		Where(title.IDEQ(int(obj.Title))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "title not found",
		})
		return
	}

	s, err := ctl.client.Sex.
		Query().
		Where(sex.IDEQ(int(obj.Sex))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "sex not found",
		})
		return
	}

	r, err := ctl.client.Role.
		Query().
		Where(role.IDEQ(int(obj.Role))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "role not found",
		})
		return
	}
	time, err := time.Parse(time.RFC3339, obj.Birthday)
	u, err := ctl.client.User.
		UpdateOneID(int(id)).
		SetEmail(obj.Email).
		SetPassword(obj.Password).
		SetName(obj.Name).
		SetAge(obj.Age).
		SetBirthday(time).
		SetTel(obj.Tel).
		SetTitle(t).
		SetSex(s).
		SetRole(r).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, u)
}

// NewUserController creates and registers handles for the User controller
func NewUserController(router gin.IRouter, client *ent.Client) *UserController {
	uc := &UserController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *UserController) register() {
	users := ctl.router.Group("/users")

	// CRUD
	users.POST("", ctl.CreateUser)
	users.GET(":id", ctl.GetUser)
	users.GET("", ctl.ListUser)
	users.DELETE(":id", ctl.DeleteUser)
	users.PUT(":id", ctl.UpdateUser)
	
}
