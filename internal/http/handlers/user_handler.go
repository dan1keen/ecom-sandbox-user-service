package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"user-service/internal/services"

	"github.com/bytedance/gopkg/util/logger"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uh *UserHandler) GetProfile() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Request.Header.Get("X-User-ID"))
		if err != nil {
			logger.Errorf("user id is not int %v", err)

			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": fmt.Sprintf("Error while getting user id %v", err)},
			)
			return
		}

		user, err := uh.userService.GetUserById(c.Request.Context(), userID)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusInternalServerError,
				gin.H{"error": fmt.Sprintf("Error while getting user %v", err)},
			)
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
