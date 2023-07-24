package handler

import (
	"encoding/json"
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"strings"
	"user-crud-service/internal/handler/dto"
	"user-crud-service/internal/handler/middleware"
	"user-crud-service/model"
	"user-crud-service/repository"
	//swaggerFiles "github.com/swaggo/files"
	//ginSwagger "github.com/swaggo/gin-swagger"
)

// todo
var secret = []byte("secret")

type UserHandler struct {
	s UserService
}

//func NewUserHandler(service UserService) *UserHandler {
//    h := &UserHandler{
//        s: service,
//    }
//    return h
//}

func SetupUserHandler(router *gin.Engine, service UserService) {

	h := &UserHandler{
		s: service,
	}
	//router.Use(middleware.DefaultStructuredLogger())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// Setup the cookie store for session management
	router.Use(sessions.Sessions("mysession", cookie.NewStore(secret)))

	//swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	//router.GET("/swagger/*any", swaggerHandler)

	router.GET("/users/list", h.GetUsersList)
	router.POST("/add/users", h.AddUsers)
	router.GET("/user/:id", h.GetUser)
	router.PATCH("/user/:id", h.UpdateUser)
	router.POST("/login", h.LoginUser)
	router.DELETE("/delete/user/:id", h.DeleteUser)
}

// GetUsersList -
func (h *UserHandler) GetUsersList(ctx *gin.Context) {
	// todo ctx
	users, err := h.s.GetUsersList(ctx.Request.Context())
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			ctx.JSON(http.StatusOK, []model.User{})
			return

		}
		log.Info().Msgf("UserHandler - GetUsersList: %v", err)
		ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
		return
	}

	// todo
	//usersResponse := dto.ParseFromEntitySlice(users)

	ctx.JSON(http.StatusOK, users)
	//ctx.JSON(http.StatusOK, []model.User{})
}

func (h *UserHandler) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		log.Info().Msg("UserHandler - GetUser: invalid user id")
		ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
		return
	}
	user, err := h.s.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			log.Info().Msgf("UserHandler - GetUser: %v", err)
			ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
			return
		}
		log.Error().Msgf("UserHandler - GetUser: %v", err)
		ctx.JSON(http.StatusInternalServerError, dto.InternalServerErrorResponse{})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) AddUsers(ctx *gin.Context) {

	body := ctx.Request.Body
	defer body.Close()

	data, _ := io.ReadAll(body)

	var req []*model.User

	err := json.Unmarshal(data, &req)
	//err := ctx.ShouldBindJSON(&req)
	if err != nil {
		log.Info().Msgf("UserHandler - AddUsers: %v", err)
		ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
		return
	}

	users, err := h.s.AddUsers(ctx, req)
	if err != nil {
		log.Error().Msgf("UserHandler - AddUsers: %v", err)
		ctx.JSON(http.StatusInternalServerError, dto.InternalServerErrorResponse{})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		log.Info().Msg("UserHandler - UpdateUser: invalid user id")
		ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
		return
	}
	var req *model.User

	if err := ctx.BindJSON(&req); err != nil {
		log.Info().Msgf("UserHandler - UpdateUser: %v", err)
		ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
		return
	}
	user, err := h.s.UpdateUser(ctx, id, req)
	if err != nil {
		log.Info().Msgf("UserHandler - UpdateUser: %v", err)
		ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		log.Info().Msg("UserHandler - DeleteUser: invalid user id")
		ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
		return
	}
	err := h.s.DeleteUser(ctx, id)
	if err != nil {
		log.Info().Msgf("UserHandler - DeleteUser: %v", err)
		ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
		return
	}
	// todo
	ctx.JSON(http.StatusOK, gin.H{"id": id})
}

type LoginForm struct {
	Id       string `form:"id" binding:"required"`
	Password string `form:"password" binding:"required"`
}
type LoginJSON struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *UserHandler) LoginUser(ctx *gin.Context) {
	session := sessions.Default(ctx)

	//var req struct {
	//	Id       string
	//	Password string
	//}

	var req LoginForm

	req.Id = ctx.PostForm("id")
	req.Password = ctx.PostForm("password")
	if strings.Trim(req.Id, " ") == "" || strings.Trim(req.Password, " ") == "" {

		log.Info().Msg("UserHandler - LoginUser: invalid id or password")
		ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
		return
	}

	//if err := ctx.ShouldBindJSON(&req); err != nil {
	//	log.Info().Msgf("UserHandler - LoginUser: %v", err)
	//	ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
	//	return
	//}

	// Check for id and password match, from a database
	user, err := h.s.GetUser(ctx, req.Id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			log.Info().Msgf("UserHandler - LoginUser: %v", err)
			ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
			return
		}
		log.Error().Msgf("UserHandler - LoginUser: %v", err)
		ctx.JSON(http.StatusInternalServerError, dto.InternalServerErrorResponse{})
		return
	}
	if req.Id != user.Id || !middleware.CheckPasswordHash(req.Password, user.Password) {
		log.Info().Msg("UserHandler - LoginUser: invalid id or password")
		// todo
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}
	const userkey = "user"

	// Save the id in the session
	session.Set(userkey, req.Id) // set this to the users ID

	if err := session.Save(); err != nil {
		log.Error().Msgf("UserHandler - LoginUser: %v", err)
		ctx.JSON(http.StatusInternalServerError, dto.InternalServerErrorResponse{})
		return
	}
	// todo
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
	ctx.JSON(http.StatusOK, user)
}
