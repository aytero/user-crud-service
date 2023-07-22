package handler

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
    "user-crud-service/internal/handler/dto"
)

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
    router.GET("/users/list", h.GetUsersList)
    //router.POST("/add/users/", h.AddUsers)
    //router.GET("/user/:id", h.GetUser)
    //router.PATCH("/user/:id", h.UpdateUser)
    //router.POST("/login", h.LoginUser)
    //router.DELETE("/delete/user/:id", h.DeleteUser)
}

// GetUsersList -
func (h *UserHandler) GetUsersList(ctx *gin.Context) {

    offset, err := strconv.ParseInt(ctx.Query("offset"), 10, 32)
    if err != nil {
        offset = 0
    }
    limit, err := strconv.ParseInt(ctx.Query("limit"), 10, 32)
    if err != nil {
        limit = 1000 // todo limit = max value
    }

    users, err := h.s.GetUsersList(ctx.Request.Context(), int32(limit), int32(offset))
    // todo
    //usersResponse := dto.ParseFromEntitySlice(users)
    usersResponse := users
    if err != nil {
        //log.Infof("CourierHandler - GetCouriers: %w", err)
        ctx.JSON(http.StatusBadRequest, dto.BadRequestResponse{})
        return
    }

    ctx.JSON(http.StatusOK, usersResponse)
}
