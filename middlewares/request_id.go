package middlewares

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"golang.org/x/net/context"
)

type RequestIdMiddlewareConfig struct {
	HeaderName          string
	ContextKey          string
	LoggerKey           string
	IgnoredPathPrefixes []string
	IDGenerator         func() string
}

type RequestIdMiddleware struct {
	config *RequestIdMiddlewareConfig
}

func NewRequestIdMiddleware(config *RequestIdMiddlewareConfig) *RequestIdMiddleware {
	return &RequestIdMiddleware{
		config: config,
	}
}

func (m *RequestIdMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ignore if path prefix is ignored
		for _, pathPrefix := range m.config.IgnoredPathPrefixes {
			if strings.HasPrefix(c.Request.RequestURI, pathPrefix) {
				c.Next()
				return
			}
		}

		// if nothing passed, generate a default ULID.
		requestIdValue := c.GetHeader(m.config.HeaderName)
		if requestIdValue == "" {
			requestIdValue = m.config.IDGenerator()
			c.Request.Header.Add(m.config.HeaderName, requestIdValue)
		}

		// update logger and requestid in context
		logger := zerolog.Ctx(c.Request.Context())
		ctx := logger.With().Str(m.config.LoggerKey, requestIdValue).Logger().WithContext(c.Request.Context())
		ctx = context.WithValue(ctx, m.config.ContextKey, requestIdValue)
		c.Request = c.Request.WithContext(ctx)

		// update response
		c.Header(m.config.HeaderName, requestIdValue)

		c.Next()
	}
}
