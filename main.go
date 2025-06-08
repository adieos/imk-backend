package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/adieos/imk-backend/cmd"
	"github.com/adieos/imk-backend/config"
	"github.com/adieos/imk-backend/controller"
	"github.com/adieos/imk-backend/dto"
	"github.com/adieos/imk-backend/middleware"
	"github.com/adieos/imk-backend/repository"
	"github.com/adieos/imk-backend/routes"
	"github.com/adieos/imk-backend/service"
	"github.com/adieos/imk-backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() error {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %v", err)
	}

	// Get the directory of the executable
	ex, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error getting executable path: %v", err)
	}
	exPath := filepath.Dir(ex)

	// List of possible locations for .env file
	envLocations := []string{
		filepath.Join(cwd, ".env"),
		filepath.Join(exPath, ".env"),
		// Add the expected location when run as a service
	}

	// Try to load .env from each location
	for _, loc := range envLocations {
		err := godotenv.Load(loc)
		if err == nil {
			fmt.Printf("Loaded .env from: %s\n", loc)
			return nil
		}
	}

	return fmt.Errorf("no .env file found in any of the expected locations")
}

func handlePanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var err error
				if e, ok := r.(error); ok {
					err = e
				} else {
					err = fmt.Errorf("%v", r)
				}
				fmt.Printf("\n[recovery] panic occurred: %v\n", err)
				stack := debug.Stack()
				fmt.Fprintln(os.Stderr, string(stack))

				res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PANIC_OCCURED, err.Error(), nil)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, res)
			}
		}()

		ctx.Next()
	}
}

func main() {
	if err := loadEnv(); err != nil {
		log.Printf("Warning: %v", err)
		log.Println("Continuing without .env file. Ensure all necessary environment variables are set.")
	}

	fmt.Println("Setting up database connection...")
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)
	fmt.Println("Database connection established")

	// Check if there are any command-line arguments
	if len(os.Args) > 1 {
		fmt.Println("Running commands...")
		cmd.Commands(db)
		return
	}

	fmt.Println("Initializing services...")

	// Initilization package

	var (
		jwtService service.JWTService = service.NewJWTService()

		// Implementation Dependency Injection
		// Repository
		userRepository repository.UserRepository = repository.NewUserRepository(db)
		bsRepository   repository.BSRepository   = repository.NewBSRepository(db)

		// Service
		userService service.UserService = service.NewUserService(userRepository, jwtService)
		bsService   service.BSService   = service.NewBSService(bsRepository, jwtService)

		// Controller
		userController controller.UserController = controller.NewUserController(userService)
		bsController   controller.BSController   = controller.NewBSController(bsService)
	)
	fmt.Println("Services initialized")

	fmt.Println("Setting up server...")
	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	server.Use(handlePanic())
	server.MaxMultipartMemory = 30 * 1024 * 1024
	server.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Route Not Found",
		})
	})

	server.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/api/rickroll", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.youtube.com/watch?v=dQw4w9WgXcQ")
	})

	fmt.Println("Setting up routes...")
	// routes
	routes.User(server, userController, jwtService)
	routes.BS(server, bsController, jwtService)

	server.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port == "" {
		port = "7777"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	fmt.Printf("Starting server on %s\n", serve)
	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
