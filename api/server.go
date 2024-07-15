package api

import (
	"strconv"
	"strings"
	"time"
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	db "github.com/katatrina/SWD392_NET1701_GroupIntern_BE/db/sqlc"
	"github.com/katatrina/SWD392_NET1701_GroupIntern_BE/internal/token"
	"github.com/katatrina/SWD392_NET1701_GroupIntern_BE/internal/util"
	
	_ "github.com/katatrina/SWD392_NET1701_GroupIntern_BE/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	router     *gin.Engine
	store      db.Store
	config     util.Config
	tokenMaker token.Maker
}

func NewServer(store db.Store, config util.Config) *Server {
	tokenMaker := token.NewJWTMaker(config.TokenSymmetricKey)
	
	server := &Server{
		store:      store,
		tokenMaker: tokenMaker,
		config:     config,
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
			userGroup.Use(authMiddleware(server.tokenMaker)).PATCH("/password", server.changeUserPassword)
		}
	}
	
	patientGroup := v1.Group("/patients")
	patientGroup.POST("", server.createPatient)
	patientGroup.GET("", server.listPatients)
	patientGroup.GET("/:id", server.getPatient)
	patientGroup.Use(authMiddleware(server.tokenMaker))
	{
		// Examination appointment
		patientGroup.POST("/appointments/examination", server.createExaminationAppointmentByPatient)
		patientGroup.GET("/appointments/examination", server.listExaminationAppointmentsByPatient)
		patientGroup.GET("/appointments/examination/:id", server.getExaminationAppointmentByPatient)
		
		// Treatment appointment
		patientGroup.GET("/appointments/treatment", server.listTreatmentAppointmentsByPatient)
		patientGroup.GET("/appointments/treatment/:id", server.getTreatmentAppointmentByPatient)
	}
	
	serviceCategoryGroup := v1.Group("/service-categories")
	{
		serviceCategoryGroup.POST("", server.createServiceCategory)
		serviceCategoryGroup.GET("", server.listServiceCategories)
		serviceCategoryGroup.GET("/:slug", server.getServiceCategoryBySlug)
		serviceCategoryGroup.PUT("/:id", server.updateServiceCategory)
		serviceCategoryGroup.DELETE("/:id", server.deleteServiceCategory)
	}
	
	serviceGroup := v1.Group("/services")
	{
		serviceGroup.GET("", server.listServices)
		serviceGroup.POST("", server.createService)
		serviceGroup.GET("/:id", server.getService)
		serviceGroup.PUT("/:id", server.updateService)
		serviceGroup.DELETE("/:id", server.deleteService)
	}
	
	dentistGroup := v1.Group("/dentists")
	{
		dentistGroup.POST("", server.createDentist)
		dentistGroup.GET("", server.listDentists)
		dentistGroup.GET("/:id", server.getDentist)
		dentistGroup.PUT("/:id", server.updateDentist)
		dentistGroup.DELETE("/:id", server.deleteDentist)
		
		dentistGroup.GET("/:id/schedules/examination", server.listExaminationSchedulesOfDentist)
		
		dentistGroup.GET("/:id/schedules/treatment", server.listTreatmentSchedulesOfDentist)
	}
	
	roomGroup := v1.Group("/rooms")
	{
		roomGroup.POST("", server.createRoom)
		roomGroup.GET("", server.listRooms)
		roomGroup.PUT("/:id", server.updateRoom)
		roomGroup.DELETE("/:id", server.deleteRoom)
	}
	
	scheduleGroup := v1.Group("/schedules")
	{
		// Examination schedule
		scheduleGroup.POST("/examination", server.createExaminationSchedule)
		scheduleGroup.GET("/examination", server.listExaminationSchedules)
		scheduleGroup.GET("/examination/:id/patients", server.listPatientsByExaminationSchedule)
		
		// Available examination schedules for patient
		scheduleGroup.GET("/examination/available", server.listAvailableExaminationSchedulesByDateForPatient)
	}
	
	appointmentGroup := v1.Group("/appointments")
	{
		// Examination appointment
		appointmentGroup.PATCH("/examination/:id/cancel", server.cancelExaminationAppointmentByPatient)
		// appointmentGroup.PATCH("/examination/:id/complete", server.completeExaminationAppointmentByDentist)
		
		// Treatment appointment
		appointmentGroup.POST("/treatment", server.createTreatmentAppointment)
		appointmentGroup.GET("/treatment", server.listTreatmentAppointments)
		appointmentGroup.GET("/treatment/:id/patients", server.listPatientsOfTreatmentAppointment)
		
		// appointmentGroup.PATCH("/treatment/:id/complete", server.completeTreatmentAppointmentByPatient)
	}
	
	v1.GET("specialties", server.listSpecialties)
	
	v1.GET("/payment-methods", server.listPaymentMethods)
	
	server.router = router
}

func (server *Server) Start() error {
	return server.router.Run(server.config.HTTPServerAddress)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// getLastIDParam returns the ID parameter of the URL for the current request.
func (server *Server) getLastIDParam(ctx *gin.Context) (int64, error) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return id, err
	}
	
	return id, nil
}

// getMiddleIDParam returns the ID parameter of the URL for the current request.
func (server *Server) getMiddleIDParam(ctx *gin.Context) (int64, error) {
	idParam := ctx.Param("id")
	sanitizedID := strings.Trim(idParam, "/")
	id, err := strconv.ParseInt(sanitizedID, 10, 64)
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
