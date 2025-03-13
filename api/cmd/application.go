package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"stonks-api/cmd/database"
	"stonks-api/internal/recommendations"
	"stonks-api/internal/stocks"
	"stonks-api/internal/stocks/services"
	"syscall"
	"time"

	authMiddleware "stonks-api/cmd/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type application struct {
	config          *Config
	db              database.Database
	server          *echo.Echo
	env             string
	stocks          *stocks.Module
	recommendations *recommendations.Module
}

func (app *application) setup(env string) error {
	app.env = env

	var err error
	app.config, err = LoadConfig(env)
	if err != nil {
		return fmt.Errorf("can't read config: %v", err)
	}

	app.db, err = database.NewPostgresDatabase(app.config.GetConnectionString())
	if err != nil {
		return fmt.Errorf("can't connect to database: %v", err)
	}
	fmt.Println("CockroachDB connection established")

	if err := database.RunMigrations(app.config.GetConnectionString()); err != nil {
		return err
	}

	// Initialize modules
	app.stocks = stocks.NewModule(app.db)
	apiConfig := services.ExternalAPIConfig{
		URL:        app.config.ExternalStocksAPI.URL,
		AuthHeader: app.config.ExternalStocksAPI.AuthHeader,
		AuthToken:  app.config.ExternalStocksAPI.AuthToken,
	}
	app.stocks.StockService.SetExternalAPIConfig(apiConfig)
	app.recommendations = recommendations.NewModule(app.db)

	// Setup HTTP server
	app.server = echo.New()
	app.server.HideBanner = true

	app.server.Use(middleware.Logger())
	app.server.Use(middleware.Recover())
	app.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{app.config.Server.AllowedOrigin},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
		AllowHeaders: []string{echo.HeaderContentType, "X-API-Key"},
	}))

	app.setupRoutes()

	return nil
}

func (app *application) setupRoutes() {
	// Health check route
	app.server.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status":  "ok",
			"service": app.config.Service.Name,
			"env":     app.env,
		})
	})
	apiV1 := app.server.Group("/api/v1/stonks-api")
	apiV1.Use(authMiddleware.APIKeyAuth(app.config.Server.APIKey))

	// Register module routes
	app.stocks.RegisterRoutes(apiV1)
	app.recommendations.RegisterRoutes(apiV1)
}

func (app *application) startServer() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		addr := app.config.GetServerAddress()
		fmt.Printf("Started %s server on %s (%s)\n",
			app.config.Service.Name, addr, app.env)

		if err := app.server.Start(addr); err != nil {
			fmt.Printf("Server stopped: %v\n", err)
		}
	}()

	<-quit
	fmt.Println("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("error during server shutdown: %v", err)
	}

	// Close database connection
	if app.db != nil {
		if err := app.db.Close(); err != nil {
			return fmt.Errorf("error closing database connection: %v", err)
		}
	}

	fmt.Println("Server gracefully stopped")
	return nil
}
