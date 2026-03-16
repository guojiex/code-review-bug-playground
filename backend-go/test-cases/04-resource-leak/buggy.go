// Bug ID: GO-004
// 维度: Bug
// 类型: resource_leak
// 严重性: High
// 描述: 资源泄漏 - HTTP响应体、文件句柄等没有正确关闭

package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ExternalAPIHandler struct {
	apiURL string
}

func NewExternalAPIHandler(apiURL string) *ExternalAPIHandler {
	return &ExternalAPIHandler{apiURL: apiURL}
}

// FetchExternalData 从外部API获取数据
// BUG：HTTP响应体没有关闭，导致资源泄漏
func (h *ExternalAPIHandler) FetchExternalData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// BUG：没有defer resp.Body.Close()
	resp, err := http.Get(h.apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	// 忘记关闭响应体！每次请求都会泄漏一个连接

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// WriteLogFile 写入日志文件
// BUG：文件句柄没有关闭
func WriteLogFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var logData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&logData); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// BUG：打开文件后没有关闭
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		http.Error(w, "Failed to open log file", http.StatusInternalServerError)
		return
	}
	// 忘记 defer file.Close()

	logLine := fmt.Sprintf("[%s] %v\n", time.Now().Format(time.RFC3339), logData)
	file.WriteString(logLine)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Log written"))
}

// ProxyRequest 代理请求到另一个服务
// BUG：多个资源泄漏点
func ProxyRequest(w http.ResponseWriter, r *http.Request) {
	targetURL := r.URL.Query().Get("target")
	if targetURL == "" {
		http.Error(w, "Missing target URL", http.StatusBadRequest)
		return
	}

	// 创建新请求
	proxyReq, err := http.NewRequest(r.Method, targetURL, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// 复制请求头
	for key, values := range r.Header {
		for _, value := range values {
			proxyReq.Header.Add(key, value)
		}
	}

	client := &http.Client{Timeout: 10 * time.Second}
	// BUG：响应体没有关闭
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, "Failed to proxy request", http.StatusBadGateway)
		return
	}
	// 忘记 defer resp.Body.Close()

	// 复制响应头
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

// SetupLeakyRoutes 设置路由
func SetupLeakyRoutes(mux *http.ServeMux, handler *ExternalAPIHandler) {
	mux.HandleFunc("GET /api/external", handler.FetchExternalData)
	mux.HandleFunc("POST /api/log", WriteLogFile)
	mux.HandleFunc("/api/proxy", ProxyRequest)
}
