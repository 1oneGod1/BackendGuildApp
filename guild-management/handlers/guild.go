package handlers

import (
	"guild-management/config"
	"guild-management/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateGuild - Tambah guild baru
func CreateGuild(c *gin.Context) {
	var guild models.Guild

	if err := c.ShouldBindJSON(&guild); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if err := config.DB.Create(&guild).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, guild)
}

// GetAllGuilds - Lihat semua guild
func GetAllGuilds(c *gin.Context) {
	var guilds []models.Guild
	config.DB.Find(&guilds)
	c.JSON(http.StatusOK, guilds)
}

// GetGuildByID - Lihat detail guild berdasarkan ID
func GetGuildByID(c *gin.Context) {
	var guild models.Guild
	id := c.Param("id")

	if err := config.DB.First(&guild, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guild not found"})
		return
	}

	c.JSON(http.StatusOK, guild)
}

// UpdateGuild - Mengubah data guild
func UpdateGuild(c *gin.Context) {
	var guild models.Guild
	id := c.Param("id")

	if err := config.DB.First(&guild, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guild not found"})
		return
	}

	if err := c.ShouldBindJSON(&guild); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	config.DB.Save(&guild)
	c.JSON(http.StatusOK, guild)
}

// DeleteGuild - Menghapus guild
func DeleteGuild(c *gin.Context) {
	var guild models.Guild
	id := c.Param("id")

	if err := config.DB.Delete(&guild, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Guild not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Guild deleted successfully"})
}
