package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanawat/app/ent"
	"github.com/thanawat/app/ent/problemstatus"
)

// ProblemStatusController defines the struct for the problemstatus controller
type ProblemStatusController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateProblemStatus handles POST requests for adding problemstatus entities
// @Summary Create problemstatus
// @Description Create problemstatus
// @ID create problemstatus
// @Accept   json
// @Produce  json
// @Param problemstatus body ent.ProblemStatus true "ProblemStatus entity"
// @Success 200 {object} ent.ProblemStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemstatus [post]
func (ctl *ProblemStatusController) CreateProblemStatus(c *gin.Context) {
	obj := ent.ProblemStatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "problemstatus binding failed",
		})
		return
	}

	u, err := ctl.client.ProblemStatus.
		Create().
		SetProblemstatus(obj.Problemstatus).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetProblemStatus handles GET requests to retrieve a problemstatus entity
// @Summary Get a problemstatus entity by ID
// @Description get problemstatus by ID
// @ID get problemstatus
// @Produce  json
// @Param id path int true "ProblemStatus ID"
// @Success 200 {object} ent.ProblemStatus
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemstatus/{id} [get]
func (ctl *ProblemStatusController) GetProblemStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.ProblemStatus.
		Query().
		Where(problemstatus.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListUser handles request to get a list of problemstatus entities
// @Summary List problemstatus entities
// @Description list problemstatus entities
// @ID list problemstatus
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.ProblemStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemstatus [get]
func (ctl *ProblemStatusController) ListProblemStatus(c *gin.Context) {
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

	problemstatus, err := ctl.client.ProblemStatus.
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

	c.JSON(200, problemstatus)
}

// DeleteProblemStatus handles DELETE requests to delete a problemstatus entity
// @Summary Delete a problemstatus entity by ID
// @Description get problemstatus by ID
// @ID delete problemstatus
// @Produce  json
// @Param id path int true "ProblemStatus ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemstatus/{id} [delete]
func (ctl *ProblemStatusController) DeleteProblemStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.ProblemStatus.
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

// UpdateProblemStatus handles PUT requests to update a problemstatus entity
// @Summary Update a problemstatus entity by ID
// @Description update problemstatus by ID
// @ID update problemstatus
// @Accept   json
// @Produce  json
// @Param id path int true "ProblemStatus ID"
// @Param problemstatus body ent.ProblemStatus true "ProblemStatus entity"
// @Success 200 {object} ent.ProblemStatus
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problemstatus/{id} [put]
func (ctl *ProblemStatusController) UpdateProblemStatus(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.ProblemStatus{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "problemstatus binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.ProblemStatus.
		UpdateOneID(int(id)).
		SetProblemstatus(obj.Problemstatus).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, u)
}

// NewProblemStatusController creates and registers handles for the problemstatus controller
func NewProblemStatusController(router gin.IRouter, client *ent.Client) *ProblemStatusController {
	uc := &ProblemStatusController{
		client: client,
		router: router,
	}

	uc.register()

	return uc

}

func (ctl *ProblemStatusController) register() {
	problemstatus := ctl.router.Group("/problemstatus")

	problemstatus.GET("", ctl.ListProblemStatus)

	// CRUD
	problemstatus.POST("", ctl.CreateProblemStatus)
	problemstatus.GET(":id", ctl.GetProblemStatus)
	problemstatus.PUT(":id", ctl.UpdateProblemStatus)
	problemstatus.DELETE(":id", ctl.DeleteProblemStatus)
}
