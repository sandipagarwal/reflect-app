package v1

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	feedbackServices "github.com/iReflect/reflect-app/apps/feedback/services"
)

//FeedbackController ...
type FeedbackController struct {
	FeedbackService feedbackServices.FeedbackService
}

// Routes for Feedback
func (ctrl FeedbackController) Routes(r *gin.RouterGroup) {
	r.GET("/", ctrl.List)
	r.GET("/:id/", ctrl.Get)
}

// Get feedback
func (ctrl FeedbackController) Get(c *gin.Context) {
	id := c.Param("id")
	userID, _ := c.Get("userID")
	feedbackResponse, err := ctrl.FeedbackService.Get(id, userID.(uint))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Feedback not found", "error": err})
		return
	}

	c.JSON(http.StatusOK, feedbackResponse)
}

// List Feedbacks
func (ctrl FeedbackController) List(c *gin.Context) {
	statuses := c.QueryArray("status")
	userID, _ := c.Get("userID")
	perPage, _ := c.GetQuery("perPage")
	parsedPerPage, err := strconv.Atoi(perPage)
	if err != nil {
		parsedPerPage = -1

	}
	response, err := ctrl.FeedbackService.TeamList(userID.(uint), statuses, parsedPerPage)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Feedbacks not found", "error": err})
		return
	}
	c.JSON(http.StatusOK, response)
}
