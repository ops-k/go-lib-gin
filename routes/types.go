package routes

import "github.com/gin-gonic/gin"

// Routes contains multiple routes
type Routes []Route

// Setup all the route
func (r Routes) BindTo(engine *gin.Engine) {
	for _, route := range r {
		route.BindTo(engine)
	}
}

// Route interface
type Route interface {
	BindTo(engine *gin.Engine)
}
