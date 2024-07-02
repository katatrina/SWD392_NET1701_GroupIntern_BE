package api

import (
	"strconv"
	"time"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392_NET1701_GroupIntern/db/sqlc"
	"github.com/katatrina/SWD392_NET1701_GroupIntern/internal/token"
	
	_ "github.com/katatrina/SWD392_NET1701_GroupIntern/docs"
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
			userGroup.POST("/login", server.loginUser)
		}
	}
	
	patientGroup := v1.Group("/patients")
	patientGroup.POST("", server.createPatient)
	patientGroup.Use(authMiddleware(server.tokenMaker))
	{
		patientGroup.POST("/appointments/examination", server.createExaminationAppointmentByPatient)
		patientGroup.GET("/appointments/examination", server.listExaminationAppointmentsByPatient)
		patientGroup.GET("/profile", server.getPatientProfile)
		patientGroup.GET("/appointments/examination/:id", server.getExaminationAppointmentByPatient)
	}
	
	serviceCategoryGroup := v1.Group("/service-categories")
	{
		serviceCategoryGroup.POST("", server.createServiceCategory)
		serviceCategoryGroup.GET("", server.listServiceCategories)
		// serviceCategoryGroup.GET("/:slug/services", server.listServicesByCategory)
		serviceCategoryGroup.GET("/:slug", server.getServiceCategoryBySlug)
		serviceCategoryGroup.PATCH("/:id", server.updateServiceCategory)
		serviceCategoryGroup.DELETE("/:id", server.deleteServiceCategory)
	}
	
	serviceGroup := v1.Group("/services")
	{
		serviceGroup.GET("", server.listServices)
		serviceGroup.POST("", server.createService)
		serviceGroup.GET("/:id", server.getService)
		serviceGroup.PATCH("/:id", server.updateService)
		serviceGroup.DELETE("/:id", server.deleteService)
	}
	
	dentistGroup := v1.Group("/dentists")
	{
		dentistGroup.POST("", server.createDentist)
		dentistGroup.GET("", server.listDentists)
		dentistGroup.GET("/:id", server.getDentist)
		dentistGroup.PATCH("/:id", server.updateDentist)
		dentistGroup.Use(authMiddleware(server.tokenMaker)).GET("/profile", server.getDentistProfile)
		dentistGroup.Use(authMiddleware(server.tokenMaker)).PATCH("/profile", server.updateDentistProfile)
	}
	
	roomGroup := v1.Group("/rooms")
	{
		roomGroup.POST("", server.createRoom)
		roomGroup.GET("", server.listRooms)
	}
	
	scheduleGroup := v1.Group("/schedules")
	
	{
		scheduleGroup.POST("/examination", server.createExaminationSchedule)
		scheduleGroup.GET("/examination", server.listExaminationSchedules)
		scheduleGroup.GET("/examination/available", server.listAvailableExaminationSchedulesByDateForPatient)
	}
	
	v1.GET("specialties", server.listSpecialties)
	
	v1.GET("/payment-methods", server.listPaymentMethods)
	
	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(":8080")
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) getIDParam(ctx *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return id, err
	}
	
	return id, nil
}

func (server *Server) getAuthorizedUserID(ctx *gin.Context) (int64, error) {
	payload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	userID, err := strconv.ParseInt(payload.Subject, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return userID, nil
}
