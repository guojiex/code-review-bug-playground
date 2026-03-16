// Bug ID: GO-003
// 维度: Bug
// 类型: sql_injection
// 严重性: Critical
// 描述: 直接拼接SQL查询字符串，存在SQL注入风险

package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductHandler struct {
	db *sql.DB
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{db: db}
}

// SearchProducts 搜索产品（有SQL注入漏洞）
// BUG：直接拼接用户输入到SQL查询中
func (h *ProductHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从查询参数获取搜索关键词
	keyword := r.URL.Query().Get("keyword")
	
	// BUG：直接拼接SQL，存在SQL注入风险
	// 例如：keyword = "test' OR '1'='1"
	query := "SELECT id, name, description, price FROM products WHERE name LIKE '%" + keyword + "%'"
	
	rows, err := h.db.Query(query)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price); err != nil {
			continue
		}
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProductByID 根据ID获取产品
// BUG：使用字符串拼接构建SQL查询
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 从URL路径参数获取ID
	productID := r.PathValue("id")
	
	// BUG：直接拼接用户输入
	query := "SELECT id, name, description, price FROM products WHERE id = " + productID
	
	var product Product
	err := h.db.QueryRow(query).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err == sql.ErrNoRows {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// SetupProductRoutes 设置路由
func SetupProductRoutes(mux *http.ServeMux, handler *ProductHandler) {
	mux.HandleFunc("GET /api/products/search", handler.SearchProducts)
	mux.HandleFunc("GET /api/products/{id}", handler.GetProductByID)
}
