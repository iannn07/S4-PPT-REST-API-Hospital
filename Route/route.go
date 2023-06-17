package Route

import (
	"github.com/gin-gonic/gin"
	"HospitalFinpro/Middleware"
	"HospitalFinpro/Controllers/AuthenticationController"
	"HospitalFinpro/Controllers/DiagnoseController"
	"HospitalFinpro/Controllers/DoctorController"
	"HospitalFinpro/Controllers/PatientController"
	"HospitalFinpro/Controllers/PaymentController"
	"HospitalFinpro/Controllers/RoomController"
)

func MakeRoute() {
	router := gin.Default()

	// Route no auth
	router.POST("login", AuthenticationController.Login)
	router.POST("register", AuthenticationController.Register)
	
	// Apply JWT middleware
	authRoutes := router.Group("/")
	authRoutes.Use(Middleware.JWTMiddleware())
	{
		authRoutes.POST("logout", AuthenticationController.Logout)

		// Doctor routes
		authRoutes.GET("api/hospital/doctors", DoctorController.SelectAll)
		authRoutes.POST("api/hospital/doctors", DoctorController.Create)
		authRoutes.GET("api/hospital/doctors/:id", DoctorController.Read)
		authRoutes.PUT("api/hospital/doctors/:id", DoctorController.Update)
		authRoutes.DELETE("api/hospital/doctors/:id", DoctorController.Delete)

		// Patient routes
		authRoutes.GET("api/hospital/patients", PatientController.SelectAll)
		authRoutes.POST("api/hospital/patients", PatientController.Create)
		authRoutes.GET("api/hospital/patients/:id", PatientController.Read)
		authRoutes.PUT("api/hospital/patients/:id", PatientController.Update)
		authRoutes.DELETE("api/hospital/patients/:id", PatientController.Delete)

		// Room routes
		authRoutes.GET("api/hospital/rooms", RoomController.SelectAll)
		authRoutes.POST("api/hospital/rooms", RoomController.Create)
		authRoutes.GET("api/hospital/rooms/:id", RoomController.Read)
		authRoutes.PUT("api/hospital/rooms/:id", RoomController.Update)
		authRoutes.DELETE("api/hospital/rooms/:id", RoomController.Delete)

		// Diagnose routes
		authRoutes.GET("api/hospital/diagnoses", DiagnoseController.SelectAll)
		authRoutes.POST("api/hospital/diagnoses", DiagnoseController.Create)
		authRoutes.GET("api/hospital/diagnoses/:id", DiagnoseController.Read)
		authRoutes.PUT("api/hospital/diagnoses/:id", DiagnoseController.Update)
		authRoutes.DELETE("api/hospital/diagnoses/:id", DiagnoseController.Delete)

		// Payment routes
		authRoutes.GET("api/hospital/payments", PaymentController.SelectAll)
		authRoutes.POST("api/hospital/payments", PaymentController.Create)
		authRoutes.GET("api/hospital/payments/:id", PaymentController.Read)
		authRoutes.PUT("api/hospital/payments/:id", PaymentController.Update)
		authRoutes.DELETE("api/hospital/payments/:id", PaymentController.Delete)
	}

	// Route Prefix Addresss
	router.Run("localhost:8080")
}