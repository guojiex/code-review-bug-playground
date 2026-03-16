// Bug ID: GO-008
// 维度: Bug
// 类型: authentication_security
// 严重性: Critical
// 描述: 认证和安全问题 - 弱密码处理、JWT漏洞、会话管理问题

package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Session struct {
	Username  string
	Token     string
	CreatedAt time.Time
}

var sessions = make(map[string]*Session)

// Login 用户登录
// BUG：多个安全问题
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var creds Credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// BUG：明文存储和比较密码
	if creds.Username == "admin" && creds.Password == "admin123" {
		// BUG：使用简单的字符串作为token，容易猜测
		token := creds.Username + "_" + time.Now().Format("20060102")
		
		session := &Session{
			Username:  creds.Username,
			Token:     token,
			CreatedAt: time.Now(),
		}
		sessions[token] = session

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"token": token,
		})
		return
	}

	// BUG：登录失败时返回详细信息，可能被用于枚举用户
	http.Error(w, "Username or password incorrect", http.StatusUnauthorized)
}

// GetProfile 获取用户资料
// BUG：token验证不安全
func GetProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// BUG：从查询参数获取token，容易泄露
	token := r.URL.Query().Get("token")
	
	// BUG：没有检查token是否为空
	session, exists := sessions[token]
	if !exists {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// BUG：没有检查session是否过期
	// session可能是几天前的，但仍然有效

	profile := map[string]interface{}{
		"username": session.Username,
		"role":     "admin", // BUG：硬编码角色
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// ChangePassword 修改密码
// BUG：多个安全漏洞
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Username    string `json:"username"`
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// BUG：没有验证用户身份（没有检查token）
	// BUG：没有验证旧密码
	// BUG：没有密码强度要求
	// BUG：没有防止暴力破解的机制

	// BUG：密码更新后，旧的session应该失效但没有处理
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Password changed successfully"))
}

// AdminAction 管理员操作
// BUG：权限检查不足
func AdminAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// BUG：从请求头获取用户角色，客户端可以伪造
	role := r.Header.Get("X-User-Role")
	
	// BUG：简单的字符串比较，容易绕过
	if role != "admin" {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// 执行管理员操作
	action := r.URL.Query().Get("action")
	
	// BUG：没有验证action的合法性
	// BUG：没有审计日志记录谁执行了什么操作
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Action executed: " + action))
}

// ResetPassword 重置密码
// BUG：重置链接不安全
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// BUG：使用可预测的重置token
	resetToken := data.Email + "_reset"
	
	// BUG：重置链接没有过期时间
	// BUG：没有限制重置尝试次数
	
	resetURL := "http://example.com/reset?token=" + resetToken
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Reset link sent",
		"link":    resetURL, // BUG：在响应中暴露重置链接
	})
}

// ValidateToken 验证Token（辅助函数）
// BUG：token验证逻辑有问题
func ValidateToken(token string) bool {
	// BUG：简单的字符串包含检查
	if strings.Contains(token, "admin") {
		return true
	}
	
	// BUG：没有验证token格式
	// BUG：没有验证签名
	
	return len(token) > 10 // 只检查长度
}

// SetupAuthRoutes 设置路由
func SetupAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/login", Login)
	mux.HandleFunc("GET /api/profile", GetProfile)
	mux.HandleFunc("POST /api/password/change", ChangePassword)
	mux.HandleFunc("POST /api/admin/action", AdminAction)
	mux.HandleFunc("POST /api/password/reset", ResetPassword)
}
