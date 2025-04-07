package handlers

import (
	"guild-management/config"
	"guild-management/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddMember - Menambahkan anggota baru ke guild
func AddMember(c *gin.Context) {
	var member models.Member
	guildID := c.Param("id")

	// Konversi guildID dari string ke int
	id, err := strconv.Atoi(guildID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid guild ID"})
		return
	}

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	// Konversi int ke uint dan set pada Member
	member.GuildID = uint(id)

	if err := config.DB.Create(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, member)
}

// GetAllMembers - Lihat semua anggota dalam guild
func GetAllMembers(c *gin.Context) {
	var members []models.Member
	guildID := c.Param("id")

	config.DB.Where("guild_id = ?", guildID).Find(&members)
	c.JSON(http.StatusOK, members)
}

// GetMemberByID - Lihat anggota berdasarkan ID
func GetMemberByID(c *gin.Context) {
	var member models.Member
	memberID := c.Param("member_id")

	if err := config.DB.First(&member, memberID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	c.JSON(http.StatusOK, member)
}

// UpdateMember - Mengubah data anggota
func UpdateMember(c *gin.Context) {
	var member models.Member
	memberID := c.Param("member_id")

	if err := config.DB.First(&member, memberID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	config.DB.Save(&member)
	c.JSON(http.StatusOK, member)
}

// DeleteMember - Menghapus anggota dari guild
func DeleteMember(c *gin.Context) {
	var member models.Member
	memberID := c.Param("member_id")

	if err := config.DB.Delete(&member, memberID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member deleted successfully"})
}
