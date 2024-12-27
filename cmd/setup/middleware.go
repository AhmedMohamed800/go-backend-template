package setup

import (
	"bytes"
	"time"

	"github.com/AhmedMohamed800/go-backend-template/config"
	"github.com/AhmedMohamed800/go-backend-template/utils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupMiddleware(e *echo.Echo, config *config.Config) {


	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: config.APIConfig.CORSOrigins,
		AllowMethods: config.APIConfig.CORSMethods,
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Start time middleware
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("start_time", time.Now())
			return next(c)
		}
	})

	// Logger middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `${time_rfc3339} | ${method} ${uri} | ${status} | ${custom} | ${remote_ip} | ${user_agent}` + "\n",
		CustomTimeFormat: time.RFC3339,
		CustomTagFunc: func(c echo.Context, buf *bytes.Buffer) (int, error) {
			start := c.Get("start_time").(time.Time)
			end := time.Now()
			latencyStr := utils.CustomLatency(start, end)
			return buf.WriteString(latencyStr)
		},
		Output: nil, // You can direct this to a file if necessary
	}))

	// Recover middleware to handle panics
	e.Use(middleware.Recover())
}
