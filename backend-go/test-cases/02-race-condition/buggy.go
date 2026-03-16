// Bug ID: GO-002
// 维度: Bug
// 类型: race_condition
// 严重性: Critical
// 描述: 并发访问共享变量时没有使用互斥锁，导致数据竞态

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Counter 计数器服务（有BUG的版本）
type Counter struct {
	count int // BUG：多个goroutine同时读写，没有加锁保护
}

// 全局计数器实例
var globalCounter = &Counter{count: 0}

// IncrementHandler 增加计数器
// BUG：多个并发请求会导致数据竞态
func IncrementHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// BUG：没有加锁，多个goroutine同时访问count变量
	globalCounter.count++
	
	response := map[string]int{"count": globalCounter.count}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetCountHandler 获取当前计数
// BUG：读取操作也没有加锁保护
func GetCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// BUG：读取count时没有加锁
	response := map[string]int{"count": globalCounter.count}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ResetHandler 重置计数器
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// BUG：重置操作也没有加锁
	globalCounter.count = 0
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Counter reset"))
}

// SetupCounterRoutes 设置路由
func SetupCounterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/counter/increment", IncrementHandler)
	mux.HandleFunc("GET /api/counter", GetCountHandler)
	mux.HandleFunc("POST /api/counter/reset", ResetHandler)
}
