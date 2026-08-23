package routes

import (
	"jsdc-api/handlers"
	"jsdc-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes
	r.GET("/", handlers.HealthCheck)
	r.GET("/health", handlers.HealthCheck)

	// Auth routes (no token required)
	auth := r.Group("/api/auth")
	{
		auth.POST("/login", handlers.Login)
	}

	// Protected routes (token required)
	api := r.Group("/api")
	api.Use(middleware.AuthRequired())
	{
		// Auth
		api.POST("/auth/logout", handlers.Logout)
		api.GET("/auth/me", handlers.GetCurrentUser)
		api.POST("/auth/change-password", handlers.ChangePassword)
		api.GET("/auth/permissions", handlers.GetPermissions)

		// RBAC
		api.GET("/menus", handlers.GetAllMenus)
		api.POST("/menus", handlers.CreateMenu)
		api.PUT("/menus/:id", handlers.UpdateMenu)
		api.DELETE("/menus/:id", handlers.DeleteMenu)
		api.GET("/levels", handlers.GetAllLevels)
		api.POST("/levels", handlers.CreateLevel)
		api.PUT("/levels/:id", handlers.UpdateLevel)
		api.DELETE("/levels/:id", handlers.DeleteLevel)
		api.GET("/hak-akses", handlers.GetAllHakAkses)
		api.GET("/hak-akses/role/:level_id", handlers.GetHakAksesByRole)
		api.PUT("/hak-akses/role/:level_id", handlers.SaveHakAksesForRole)

		// Dashboard
		api.GET("/dashboard", handlers.GetDashboard)

		// Users
		api.GET("/users", handlers.GetAllUsers)
		api.GET("/users/:id", handlers.GetUser)
		api.POST("/users", handlers.CreateUser)
		api.PUT("/users/:id", handlers.UpdateUser)
		api.DELETE("/users/:id", handlers.DeleteUser)
		api.GET("/users/area/:area_id", handlers.GetUsersByArea)

		// Kelas
		api.GET("/kelas", handlers.GetAllKelas)
		api.GET("/kelas/:id", handlers.GetKelas)
		api.POST("/kelas", handlers.CreateKelas)
		api.PUT("/kelas/:id", handlers.UpdateKelas)
		api.DELETE("/kelas/:id", handlers.DeleteKelas)

		// Area
		api.GET("/area", handlers.GetAllArea)
		api.GET("/area/:id", handlers.GetArea)
		api.POST("/area", handlers.CreateArea)
		api.PUT("/area/:id", handlers.UpdateArea)
		api.DELETE("/area/:id", handlers.DeleteArea)

		// Nilai Lulus
		api.GET("/nilai-lulus", handlers.GetAllNilaiLulus)
		api.GET("/nilai-lulus/:id", handlers.GetNilaiLulus)
		api.POST("/nilai-lulus", handlers.CreateNilaiLulus)
		api.PUT("/nilai-lulus/:id", handlers.UpdateNilaiLulus)
		api.DELETE("/nilai-lulus/:id", handlers.DeleteNilaiLulus)

		// Daftar (Pendaftaran)
		api.GET("/daftar", handlers.GetAllDaftar)
		api.GET("/daftar/:id", handlers.GetDaftar)
		api.POST("/daftar", handlers.CreateDaftar)
		api.PUT("/daftar/:id", handlers.UpdateDaftar)
		api.DELETE("/daftar/:id", handlers.DeleteDaftar)
		api.GET("/daftar/count/today", handlers.GetCountDaftarToday)

		// Peserta
		api.GET("/peserta", handlers.GetAllPeserta)
		api.GET("/peserta/search", handlers.SearchPeserta)
		api.GET("/peserta/:id", handlers.GetPeserta)
		api.GET("/peserta/:id/soal", handlers.GetPesertaSoal)
		api.GET("/peserta/:id/praktek", handlers.GetPraktekByPeserta)
		api.POST("/peserta/:id/praktek", handlers.SubmitPraktek)
		api.GET("/peserta/:id/comment", handlers.GetCommentByPeserta)
		api.POST("/peserta/:id/comment", handlers.SubmitComment)

		// Soal
		api.GET("/soal", handlers.GetAllSoal)
		api.GET("/soal/:id", handlers.GetSoal)
		api.POST("/soal", handlers.CreateSoal)
		api.PUT("/soal/:id", handlers.UpdateSoal)
		api.DELETE("/soal/:id", handlers.DeleteSoal)

		// Jenis Dokumen
		api.GET("/jenis-dokumen", handlers.GetAllJenisDokumen)
		api.GET("/jenis-dokumen/:id", handlers.GetJenisDokumen)
		api.POST("/jenis-dokumen", handlers.CreateJenisDokumen)
		api.PUT("/jenis-dokumen/:id", handlers.UpdateJenisDokumen)
		api.DELETE("/jenis-dokumen/:id", handlers.DeleteJenisDokumen)

		// Daftar Dokumen
		api.GET("/daftar-dokumen", handlers.GetAllDaftarDokumen)
		api.POST("/daftar-dokumen", handlers.CreateDaftarDokumen)
		api.PUT("/daftar-dokumen/:id", handlers.UpdateDaftarDokumen)
		api.DELETE("/daftar-dokumen/:id", handlers.DeleteDaftarDokumen)
	}
}
