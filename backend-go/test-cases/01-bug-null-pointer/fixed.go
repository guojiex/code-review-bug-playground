// Fixed version of GO-001
// 修复了空指针异常问题

package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// UserRepository 模拟的用户仓储接口
type UserRepository interface {
	FindByID(id string) (*User, error)
}

// User 用户模型
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserProfile 用户资料
type UserProfile struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Bio   string `json:"bio"`
}

type UserHandler struct {
	repo UserRepository
}

// GetUserProfile 获取用户资料（修复后的版本）
// 修复：添加了错误检查和空值判断
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userID := c.Param("id")
	
	// 修复1：检查error
	user, err := h.repo.FindByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to find user",
		})
		return
	}
	
	// 修复2：检查user是否为nil
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}
	
	// 现在可以安全地访问user的字段
	profile := &UserProfile{
		Name:  user.Name,
		Email: user.Email,
		Bio:   "User bio",
	}
	
	c.JSON(http.StatusOK, profile)
}
