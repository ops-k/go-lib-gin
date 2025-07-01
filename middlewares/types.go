package middlewares

import "github.com/gin-gonic/gin"

// Routes contains multiple routes
type Middlewares []Middleware

// Route interface
type Middleware interface {
	Handler() gin.HandlerFunc
}
