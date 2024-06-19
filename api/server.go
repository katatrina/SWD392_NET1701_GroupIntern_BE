package api

import (
	"time"
	
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
	
	// Setup cors
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	config.AllowAllOrigins = true
	router.Use(cors.New(config))
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	v1 := router.Group("/api/v1")
	{
		userGroup := v1.Group("/users")
		{
			userGroup.POST("", server.createPatient)
			userGroup.POST("/login", server.loginUser)
		}
	}
	
	authorized := v1.Group("/")
	authorized.Use(authMiddleware(server.tokenMaker))
	{
		patientGroup := authorized.Group("/patients")
		{
			patientGroup.POST("/appointments/examination", server.createExaminationAppointmentByPatient)
			patientGroup.GET("/appointments/examination", server.listExaminationAppointmentsByPatient)
			patientGroup.GET("", server.getPatientInfo)
			patientGroup.GET("/appointments/examination/:id", server.getExaminationAppointmentByPatient)
		}
	}
	
	serviceCategoryGroup := v1.Group("/service-categories")
	{
		serviceCategoryGroup.GET("", server.listServiceCategories)
		serviceCategoryGroup.GET("/:slug/services", server.listServicesOfOneCategory)
		serviceCategoryGroup.GET("/:slug", server.getServiceCategoryBySlug)
		serviceCategoryGroup.PATCH("/:id", server.updateServiceCategory)
		serviceCategoryGroup.DELETE("/:id", server.deleteServiceCategory)
	}
	
	serviceGroup := v1.Group("/services")
	{
		serviceGroup.POST("", server.createService)
		// serviceGroup.PATCH("/:id", server.updateService)
	}
	
	v1.GET("/schedules/examination", server.listExaminationSchedulesByDate)
	
	v1.GET("/payment-methods", server.listAllPaymentMethods)
	
	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(":8080")
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
