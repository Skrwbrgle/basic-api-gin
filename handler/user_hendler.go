package handler

import (
	"net/http"
	"restfull-api/m/v2/domain"
	"restfull-api/m/v2/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUseCase *usecase.UserUseCase
}

func NewUserHandler(router *gin.Engine, uu *usecase.UserUseCase) {
	handler := &UserHandler{
		UserUseCase: uu,
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is online",
		})
	})

	router.GET("/users/:id", handler.GetUserByID)
	router.POST("/users", handler.CreateUser)
	router.PUT("/users/:id", handler.UpdateUser)
	router.DELETE("/users/:id", handler.DeleteUser)
	router.GET("/users", handler.ListUsers)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	user, err := h.UserUseCase.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.UserUseCase.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "email address has already been"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	user.ID = id

	if user.Email == "" && user.Name == "" {
		existingUser, err := h.UserUseCase.GetUserByID(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user"})
			return
		}
		user.Email = existingUser.Email
		user.Name = existingUser.Name
	} else if user.Email == "" {
		existingUser, err := h.UserUseCase.GetUserByID(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user"})
			return
		}
		user.Email = existingUser.Email
	} else if user.Name == "" {
		existingUser, err := h.UserUseCase.GetUserByID(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get user"})
			return
		}
		user.Name = existingUser.Name
	}
	res, err := h.UserUseCase.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not update user"})
		return
	}

	if numAff, err := res.RowsAffected(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to update user"})
		return
	} else if numAff == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "user id does not exist"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	res, err := h.UserUseCase.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if numAff, err := res.RowsAffected(); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "failed to delete user"})
		return
	} else if numAff == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "user id does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.UserUseCase.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not list users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
