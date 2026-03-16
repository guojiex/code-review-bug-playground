// Bug ID: GO-005
// 维度: Bug
// 类型: improper_error_handling
// 严重性: Medium
// 描述: 错误处理不当 - 忽略错误、错误信息泄露、不恰当的错误返回

package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Order struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
}

// CreateOrder 创建订单
// BUG：错误处理不当，可能泄露内部信息
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var order Order
	// BUG：忽略了Decode可能返回的错误
	json.NewDecoder(r.Body).Decode(&order)
	
	// 模拟数据库操作
	err := saveOrderToDB(&order)
	if err != nil {
		// BUG：直接将内部错误信息暴露给用户
		// 可能泄露数据库结构、连接信息等敏感数据
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

// GetOrder 获取订单详情
// BUG：忽略多个错误
func GetOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	orderIDStr := r.PathValue("id")
	
	// BUG：忽略转换错误
	orderID, _ := strconv.Atoi(orderIDStr)
	
	// 模拟数据库查询
	order, err := getOrderFromDB(orderID)
	if err != nil {
		// BUG：没有区分"未找到"和"内部错误"
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	// BUG：忽略编码错误
	json.NewEncoder(w).Encode(order)
}

// UploadFile 上传文件
// BUG：多处错误处理缺失
func UploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// BUG：没有限制文件大小，可能导致内存耗尽
	r.ParseMultipartForm(32 << 20) // 32MB
	
	file, handler, err := r.FormFile("file")
	if err != nil {
		// BUG：错误信息过于详细，可能泄露服务器信息
		http.Error(w, fmt.Sprintf("Error retrieving file: %v", err), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// BUG：没有验证文件名，可能导致路径遍历攻击
	dst, err := os.Create("uploads/" + handler.Filename)
	if err != nil {
		// BUG：直接暴露文件系统错误
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// BUG：忽略Copy的返回值，不知道是否完全写入
	io.Copy(dst, file)

	w.Write([]byte("File uploaded successfully"))
}

// DeleteOrder 删除订单
// BUG：错误被完全忽略
func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	orderIDStr := r.PathValue("id")
	orderID, _ := strconv.Atoi(orderIDStr)
	
	// BUG：完全忽略删除操作的错误
	deleteOrderFromDB(orderID)
	
	// 无论删除是否成功都返回200
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Order deleted"))
}

// 模拟的数据库操作函数
func saveOrderToDB(order *Order) error {
	return fmt.Errorf("database connection failed: host=localhost user=admin password=secret123 dbname=orders")
}

func getOrderFromDB(id int) (*Order, error) {
	return nil, fmt.Errorf("SQL error: SELECT * FROM orders WHERE id = %d - table does not exist", id)
}

func deleteOrderFromDB(id int) error {
	return fmt.Errorf("cannot delete order: foreign key constraint violation")
}

// SetupOrderRoutes 设置路由
func SetupOrderRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/orders", CreateOrder)
	mux.HandleFunc("GET /api/orders/{id}", GetOrder)
	mux.HandleFunc("DELETE /api/orders/{id}", DeleteOrder)
	mux.HandleFunc("POST /api/upload", UploadFile)
}
