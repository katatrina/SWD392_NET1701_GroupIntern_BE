package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392/db/sqlc"
	"github.com/katatrina/SWD392/internal/token"

	_ "github.com/katatrina/SWD392/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router     *gin.Engine
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(store db.Store) *Server {
	tokenMaker := token.NewJWTMaker("a3d1b2c3e4f5678910a1b2c3d4e5f67890abcdef1234567890abcdef12345678")

	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			userGroup.POST("", server.createCustomer)
			userGroup.POST("/login", server.loginUser)
		}
	}

	authorized := v1.Group("/")
	authorized.Use(authMiddleware(server.tokenMaker))
	authorized.POST("appointments/examination", server.createExaminationAppointment)

	v1.GET("/service-categories", server.listAllServiceCategories)

	v1.GET("/schedules/examination", server.listExaminationSchedulesByDateAndServiceCategory)

	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(":8080")
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
