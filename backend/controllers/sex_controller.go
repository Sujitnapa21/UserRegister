package controllers

import (
	"context"
	"strconv"

	"github.com/Sujitnapa21/app/ent"
	"github.com/Sujitnapa21/app/ent/sex"
	"github.com/gin-gonic/gin"
)

// SexController defines the struct for the sex controller
type SexController struct {
	client *ent.Client
	router gin.IRouter
}

// Sex defines the struct for the sex controller
type Sex struct {
	Sexname string
}

// CreateSex handles POST requests for adding sex entities
// @Summary Create sex
// @Description Create sex
// @ID create-sex
// @Accept   json
// @Produce  json
// @Param sex body ent.Sex true "Sex entity"
// @Success 200 {object} ent.Sex
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /sexs [post]
func (ctl *SexController) CreateSex(c *gin.Context) {
	obj := ent.Sex{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "sex binding failed",
		})
		return
	}

	s, err := ctl.client.Sex.
		Create().
		SetSexname(obj.Sexname).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, s)
}

// GetSex handles GET requests to retrieve a sex entity
// @Summary Get a sex entity by ID
// @Description get sex by ID
// @ID get-sex
// @Produce  json
// @Param id path int true "Sex ID"
// @Success 200 {object} ent.Sex
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /sexs/{id} [get]
func (ctl *SexController) GetSex(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	s, err := ctl.client.Sex.
		Query().
		Where(sex.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListSex handles request to get a list of sex entities
// @Summary List sex entities
// @Description list sex entities
// @ID list-sex
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Sex
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /sexs [get]
func (ctl *SexController) ListSex(c *gin.Context) {
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

	sexs, err := ctl.client.Sex.
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

	c.JSON(200, sexs)
}

// NewSexController creates and registers handles for the Sex controller
func NewSexController(router gin.IRouter, client *ent.Client) *SexController {
	sc := &SexController{
		client: client,
		router: router,
	}

	sc.register()

	return sc

}

func (ctl *SexController) register() {
	sexs := ctl.router.Group("/sexs")

	sexs.POST("", ctl.CreateSex)
	sexs.GET(":id", ctl.GetSex)
	sexs.GET("", ctl.ListSex)

}
