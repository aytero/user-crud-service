package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

const Userkey = "user"

func AuthRequired(ctx *gin.Context) {
	session := sessions.Default(ctx)
	user := session.Get(Userkey)
	//if user == nil || user != c.Param("id") {
	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	ctx.Next()
}
