package controllers

import (
	"context"
	"strconv"

	"github.com/Sujitnapa21/app/ent"
	"github.com/Sujitnapa21/app/ent/role"
	"github.com/gin-gonic/gin"
)

// RoleController defines the struct for the role controller
type RoleController struct {
	client *ent.Client
	router gin.IRouter
}

// Role defines the struct for the role controller
type Role struct {
	Rolename string
}

// CreateRole handles POST requests for adding role entities
// @Summary Create role
// @Description Create role
// @ID create-role
// @Accept   json
// @Produce  json
// @Param role body ent.Role true "Role entity"
// @Success 200 {object} ent.Role
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles [post]
func (ctl *RoleController) CreateRole(c *gin.Context) {
	obj := ent.Role{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "role binding failed",
		})
		return
	}

	r, err := ctl.client.Role.
		Create().
		SetRolename(obj.Rolename).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, r)
}

// GetRole handles GET requests to retrieve a role entity
// @Summary Get a role entity by ID
// @Description get role by ID
// @ID get-role
// @Produce  json
// @Param id path int true "Role ID"
// @Success 200 {object} ent.Role
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles/{id} [get]
func (ctl *RoleController) GetRole(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r, err := ctl.client.Role.
		Query().
		Where(role.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// ListRole handles request to get a list of role entities
// @Summary List role entities
// @Description list role entities
// @ID list-role
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Role
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /roles [get]
func (ctl *RoleController) ListRole(c *gin.Context) {
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

	roles, err := ctl.client.Role.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, roles)
}

// NewRoleController creates and registers handles for the Role controller
func NewRoleController(router gin.IRouter, client *ent.Client) *RoleController {
	rc := &RoleController{
		client: client,
		router: router,
	}

	rc.register()

	return rc

}

func (ctl *RoleController) register() {
	roles := ctl.router.Group("/roles")

	roles.POST("", ctl.CreateRole)
	roles.GET(":id", ctl.GetRole)
	roles.GET("", ctl.ListRole)

}
