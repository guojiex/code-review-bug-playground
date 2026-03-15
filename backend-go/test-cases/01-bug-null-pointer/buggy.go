// Bug ID: GO-001
// 维度: Bug
// 类型: null_pointer_exception
// 严重性: High
// 描述: 用户查询未检查空值，会导致panic

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

// GetUserProfile 获取用户资料（有BUG的版本）
// 问题：当用户不存在时，FindByID返回nil，直接访问user.Name会panic
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userID := c.Param("id")
	
	// BUG：没有检查error
	user, _ := h.repo.FindByID(userID)
	
	// BUG：user可能为nil，直接访问会panic
	profile := &UserProfile{
		Name:  user.Name,   // 空指针！
		Email: user.Email,  // 空指针！
		Bio:   "User bio",
	}
	
	c.JSON(http.StatusOK, profile)
}
