package handler

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"strings"
	_ "user-crud-service/docs"
	"user-crud-service/handler/dto"
	"user-crud-service/handler/middleware"
	"user-crud-service/model"
	"user-crud-service/repository"
	"user-crud-service/service"
)

var secret = []byte("secret")

type UserHandler struct {
	s UserService
}

func SetupUserHandler(router *gin.Engine, service UserService) {

	h := &UserHandler{
		s: service,
	}
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))

	// /swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/add/users", h.AddUsers)
	router.GET("/users/list", h.GetUsersList)
	router.PATCH("/user/:id", h.UpdateUser)
	router.DELETE("/delete/user/:id", h.DeleteUser)
	router.POST("/login", h.LoginUser)

	private := router.Group("")
	private.Use(middleware.AuthRequired)
	{
		private.GET("/user/:id", h.GetUser)
	}
}

func (h *UserHandler) AddUsers(ctx *gin.Context) {
	var req []*model.User

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Info().Msgf("UserHandler - AddUsers: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.s.AddUsers(ctx, req)
	if err != nil {
		log.Error().Msgf("UserHandler - AddUsers: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "users added successfully"})
}

func (h *UserHandler) GetUsersList(ctx *gin.Context) {
	users, err := h.s.GetUsersList(ctx.Request.Context())
	if err != nil {
		log.Info().Msgf("UserHandler - GetUsersList: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(users) == 0 {
		ctx.JSON(http.StatusOK, []model.User{})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		log.Info().Msg("UserHandler - GetUser: invalid user id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	user, err := h.s.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			log.Info().Msgf("UserHandler - GetUser: %v", err)
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		log.Error().Msgf("UserHandler - GetUser: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		log.Info().Msg("UserHandler - UpdateUser: invalid user id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req *model.UpdateUser
	if err := ctx.BindJSON(&req); err != nil {
		log.Info().Msgf("UserHandler - UpdateUser: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.s.UpdateUser(ctx, id, req)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			log.Info().Msgf("UserHandler - UpdateUser: %v", err)
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		log.Info().Msgf("UserHandler - UpdateUser: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		log.Info().Msg("UserHandler - DeleteUser: invalid user id")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	err := h.s.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			log.Info().Msgf("UserHandler - DeleteUser: %v", err)
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		log.Info().Msgf("UserHandler - DeleteUser: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *UserHandler) LoginUser(ctx *gin.Context) {
	session := sessions.Default(ctx)

	var req dto.LoginJSON

	if err := ctx.BindJSON(&req); err != nil {
		log.Info().Msgf("UserHandler - LoginUser: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if strings.Trim(req.Id, " ") == "" || strings.Trim(req.Password, " ") == "" {
		log.Info().Msg("UserHandler - LoginUser: invalid id or password")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id or password"})
		return
	}

	// Check for id and password match, from a database
	user, err := h.s.GetUser(ctx, req.Id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			log.Info().Msgf("UserHandler - LoginUser: %v", err)
			ctx.JSON(http.StatusNotFound, gin.H{"error": err})
			return
		}
		log.Error().Msgf("UserHandler - LoginUser: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	if req.Id != user.Id || !service.ComparePasswords(req.Password, []byte(user.Password)) {
		log.Info().Msg("UserHandler - LoginUser: invalid id or password")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authentication failed"})
		return
	}

	// Save the id in the session
	session.Set(middleware.Userkey, req.Id)

	if err := session.Save(); err != nil {
		log.Error().Msgf("UserHandler - LoginUser: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save session"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successfully authenticated user"})
}
