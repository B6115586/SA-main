package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thanawat/app/ent"
)

type ProblemController struct {
	client *ent.Client
	router gin.IRouter
}

type Problem struct {
	Problemstatus int
	Problemtitle  int
	Room          int
	User          int
	Probleminfo   string
	Added         string
}

// CreateProblem handles POST requests for adding problem entities
// @Summary Create problem
// @Description Create problem
// @ID create-problem
// @Accept   json
// @Produce  json
// @Param problem body Problem true "Problem entity"
// @Success 200 {object} Problem
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /problem [post]
func (ctl *ProblemController) CreateProblem(c *gin.Context) {
	obj := Problem{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "problem binding failed",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Added)
	pv, err := ctl.client.Problem.
		Create().
		SetUserID(obj.User).
		SetAdddate(time).
		SetRoomID(obj.Room).
		SetProblemstatusID(obj.Problemstatus).
		SetProblemtitleID(obj.Problemtitle).
		SetProbleminfo(obj.Probleminfo).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, pv)
}

// ListProblem handles request to get a list of problem entities
// @Summary List problem entities
// @Description list problem entities
// @ID list-problem
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Problem
// @Failure 500 {object} gin.H
// @Router /problem [get]
func (ctl *ProblemController) ListProblem(c *gin.Context) {
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

	problem, err := ctl.client.Problem.
		Query().
		WithUser().
		WithProblemstatus().
		WithRoom().
		WithProblemtitle().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, problem)
}

// NewProblemController creates and registers handles for the problem controller
func NewProblemController(router gin.IRouter, client *ent.Client) *ProblemController {
	pvc := &ProblemController{
		client: client,
		router: router,
	}

	pvc.register()

	return pvc

}

func (ctl *ProblemController) register() {
	problem := ctl.router.Group("/problem")

	problem.POST("", ctl.CreateProblem)
	problem.GET("", ctl.ListProblem)

}
