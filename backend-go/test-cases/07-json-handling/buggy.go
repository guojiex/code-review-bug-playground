// Bug ID: GO-007
// 维度: Bug
// 类型: json_handling_issues
// 严重性: Medium
// 描述: JSON处理问题 - 类型不匹配、缺少验证、序列化错误

package handler

import (
	"encoding/json"
	"net/http"
	"time"
)

// UserRequest 用户请求结构
type UserRequest struct {
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Email    string  `json:"email"`
	Balance  float64 `json:"balance"`
	IsActive bool    `json:"is_active"`
}

// CreateUser 创建用户
// BUG：没有验证JSON字段的有效性
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user UserRequest
	
	// BUG：解码时没有检查额外的未知字段
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// BUG：没有验证必填字段
	// name可能为空，age可能为负数，email格式可能错误
	
	// BUG：没有验证数值范围
	// age可能是负数或超大值，balance可能是负数

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created",
		"user":    user,
	})
}

// GetUserData 获取用户数据
// BUG：返回的JSON包含循环引用
func GetUserData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// BUG：创建循环引用的结构
	type Node struct {
		Value int   `json:"value"`
		Next  *Node `json:"next,omitempty"`
	}

	node1 := &Node{Value: 1}
	node2 := &Node{Value: 2}
	node1.Next = node2
	node2.Next = node1 // 循环引用！

	w.Header().Set("Content-Type", "application/json")
	// BUG：这会导致序列化错误或无限循环
	json.NewEncoder(w).Encode(node1)
}

// UpdateSettings 更新设置
// BUG：使用interface{}导致类型不确定
func UpdateSettings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// BUG：使用map[string]interface{}无法保证类型安全
	var settings map[string]interface{}
	
	err := json.NewDecoder(r.Body).Decode(&settings)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// BUG：直接使用interface{}值，没有类型断言
	// 如果客户端发送错误类型，运行时会panic
	maxConnections := settings["max_connections"].(int) // 可能panic
	timeout := settings["timeout"].(int)                // 可能panic
	
	// BUG：没有检查必需的字段是否存在
	enableCache := settings["enable_cache"].(bool) // 如果字段不存在会panic

	response := map[string]interface{}{
		"max_connections": maxConnections,
		"timeout":         timeout,
		"enable_cache":    enableCache,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetReport 生成报告
// BUG：序列化特殊类型时出错
func GetReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type Report struct {
		Timestamp time.Time          `json:"timestamp"`
		Data      map[string]float64 `json:"data"`
		Channel   chan int           `json:"channel"` // BUG：channel不能序列化
		Function  func()             `json:"function"` // BUG：function不能序列化
	}

	report := Report{
		Timestamp: time.Now(),
		Data: map[string]float64{
			"revenue": 10000.50,
			"cost":    5000.25,
		},
		Channel:  make(chan int),
		Function: func() {},
	}

	w.Header().Set("Content-Type", "application/json")
	// BUG：channel和function字段无法序列化
	json.NewEncoder(w).Encode(report)
}

// ParseNumbers 解析数字
// BUG：数字精度问题
func ParseNumbers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// BUG：使用float64解析大整数会丢失精度
	var data struct {
		LargeNumber float64 `json:"large_number"` // 对于很大的整数会丢失精度
		Precision   float64 `json:"precision"`    // 浮点数精度问题
	}

	json.NewDecoder(r.Body).Decode(&data)

	// BUG：直接比较浮点数
	if data.Precision == 0.1+0.2 { // 可能不相等！
		w.Write([]byte("Equal"))
	} else {
		w.Write([]byte("Not equal"))
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// SetupJSONRoutes 设置路由
func SetupJSONRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/users", CreateUser)
	mux.HandleFunc("GET /api/users/data", GetUserData)
	mux.HandleFunc("PUT /api/settings", UpdateSettings)
	mux.HandleFunc("GET /api/report", GetReport)
	mux.HandleFunc("POST /api/numbers", ParseNumbers)
}
