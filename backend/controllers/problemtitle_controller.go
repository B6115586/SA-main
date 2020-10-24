package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/thanawat/app/ent/problemtitle"

	"github.com/gin-gonic/gin"
	"github.com/thanawat/app/ent"
)

// ProblemTitleController defines the struct for the problemtitle controller
type ProblemTitleController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateProblemTitle handles POST requests for adding problemtitle entities
// @Summary Create problemtitle
// @Description Create problemtitle
// @ID create problemtitle
// @Accept   json
// @Produce  json
// @Param problemtitle body ent.ProblemTitle true "ProblemTitle entity"
// @Success 200 {object} ent.ProblemTitle
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemtitles [post]
func (ctl *ProblemTitleController) CreateProblemTitle(c *gin.Context) {
	obj := ent.ProblemTitle{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "problemtitle binding failed",
		})
		return
	}

	u, err := ctl.client.ProblemTitle.
		Create().
		SetProblemtitle(obj.Problemtitle).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetProblemTitle handles GET requests to retrieve a problemtitle entity
// @Summary Get a problemtitle entity by ID
// @Description get problemtitle by ID
// @ID get problemtitle
// @Produce  json
// @Param id path int true "ProblemTitle ID"
// @Success 200 {object} ent.ProblemTitle
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemtitles/{id} [get]
func (ctl *ProblemTitleController) GetProblemTitle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.ProblemTitle.
		Query().
		Where(problemtitle.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListUser handles request to get a list of problemtitle entities
// @Summary List problemtitle entities
// @Description list problemtitle entities
// @ID list problemtitle
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.ProblemTitle
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemtitles [get]
func (ctl *ProblemTitleController) ListProblemTitle(c *gin.Context) {
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

	problemtitles, err := ctl.client.ProblemTitle.
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

	c.JSON(200, problemtitles)
}

// DeleteProblemTitle handles DELETE requests to delete a problemtitle entity
// @Summary Delete a problemtitle entity by ID
// @Description get problemtitle by ID
// @ID delete problemtitle
// @Produce  json
// @Param id path int true "ProblemTitle ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemtitles/{id} [delete]
func (ctl *ProblemTitleController) DeleteProblemTitle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ProblemTitle.
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

// UpdateProblemTitle handles PUT requests to update a problemtitle entity
// @Summary Update a problemtitle entity by ID
// @Description update problemtitle by ID
// @ID update problemtitle
// @Accept   json
// @Produce  json
// @Param id path int true "ProblemTitle ID"
// @Param problemtitle body ent.ProblemTitle true "ProblemTitle entity"
// @Success 200 {object} ent.ProblemTitle
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemtitles/{id} [put]
func (ctl *ProblemTitleController) UpdateProblemTitle(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.ProblemTitle{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "problemtitle binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.ProblemTitle.
		UpdateOneID(int(id)).
		SetProblemtitle(obj.Problemtitle).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, u)
}

// NewProblemTitleController creates and registers handles for the problemtitle controller
func NewProblemTitleController(router gin.IRouter, client *ent.Client) *ProblemTitleController {
	uc := &ProblemTitleController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *ProblemTitleController) register() {
	problemtitles := ctl.router.Group("/problemtitles")

	problemtitles.GET("", ctl.ListProblemTitle)

	// CRUD
	problemtitles.POST("", ctl.CreateProblemTitle)
	problemtitles.GET(":id", ctl.GetProblemTitle)
	problemtitles.PUT(":id", ctl.UpdateProblemTitle)
	problemtitles.DELETE(":id", ctl.DeleteProblemTitle)
}
