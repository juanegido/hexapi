package courses

import (
	"github.com/gin-gonic/gin"
	mooc "github.com/juanegido/hexapi/internal"
	"net/http"
)

type CreateRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request CreateRequest
		if err := ctx.BindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course := mooc.NewCourse(request.ID, request.Name, request.Duration)
		if err := courseRepository.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
