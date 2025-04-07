package routes

import (
	"guild-management/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Guild Management API",
		})
	})

	// Group untuk Guild Management
	guilds := r.Group("/guilds")
	{
		guilds.POST("/", handlers.CreateGuild)      // Tambah Guild
		guilds.GET("/", handlers.GetAllGuilds)      // Lihat Semua Guild
		guilds.GET("/:id", handlers.GetGuildByID)   // Lihat Detail Guild
		guilds.PUT("/:id", handlers.UpdateGuild)    // Update Guild
		guilds.DELETE("/:id", handlers.DeleteGuild) // Hapus Guild

		// Group untuk Member Management dalam Guild
		members := guilds.Group("/:id/members")
		{
			members.POST("/", handlers.AddMember)                // Tambah Anggota
			members.GET("/", handlers.GetAllMembers)             // Lihat Semua Anggota
			members.GET("/:member_id", handlers.GetMemberByID)   // Lihat Detail Anggota
			members.PUT("/:member_id", handlers.UpdateMember)    // Update Anggota
			members.DELETE("/:member_id", handlers.DeleteMember) // Hapus Anggota
		}
	}

	return r
}
