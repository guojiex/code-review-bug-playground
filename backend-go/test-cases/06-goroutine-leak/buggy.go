// Bug ID: GO-006
// 维度: Bug
// 类型: goroutine_leak
// 严重性: High
// 描述: goroutine泄漏 - 启动的goroutine没有正确退出机制

package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type Task struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Result    string    `json:"result,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

var tasks = make(map[string]*Task)

// ProcessTask 处理长时间运行的任务
// BUG：启动goroutine但没有取消机制，可能永久运行
func ProcessTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input map[string]string
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	taskID := input["task_id"]
	task := &Task{
		ID:        taskID,
		Status:    "processing",
		CreatedAt: time.Now(),
	}
	tasks[taskID] = task

	// BUG：启动goroutine没有context控制，无法取消
	go func() {
		// 模拟长时间运行的任务
		for i := 0; i < 100; i++ {
			time.Sleep(1 * time.Second)
			// 如果客户端断开连接，这个goroutine仍会继续运行
		}
		task.Status = "completed"
		task.Result = "Task finished"
	}()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// StreamEvents 流式事件推送
// BUG：没有检测客户端断开连接，goroutine永久运行
func StreamEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	// BUG：没有使用context监听客户端断开
	// 即使客户端关闭连接，这个循环也会继续运行
	for i := 0; ; i++ {
		event := map[string]interface{}{
			"id":        i,
			"timestamp": time.Now().Unix(),
			"message":   "Event data",
		}
		
		data, _ := json.Marshal(event)
		w.Write([]byte("data: "))
		w.Write(data)
		w.Write([]byte("\n\n"))
		flusher.Flush()
		
		time.Sleep(1 * time.Second)
	}
}

// PollStatus 轮询任务状态
// BUG：channel没有缓冲且没有select超时，可能导致goroutine阻塞
func PollStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	taskID := r.URL.Query().Get("task_id")
	
	// BUG：创建无缓冲channel，如果没有接收方会永久阻塞
	resultChan := make(chan *Task)
	
	// BUG：启动goroutine但如果下面的接收超时，发送会永久阻塞
	go func() {
		time.Sleep(2 * time.Second)
		task := tasks[taskID]
		resultChan <- task // 如果没人接收，这里会永久阻塞
	}()

	// 等待结果，但没有超时控制
	task := <-resultChan
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// WatchResource 监控资源变化
// BUG：多层goroutine嵌套，没有统一的退出机制
func WatchResource(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resourceID := r.PathValue("id")
	
	// BUG：启动监控goroutine，但没有停止机制
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		// BUG：ticker没有stop，会一直运行
		
		for range ticker.C {
			// BUG：内层又启动新的goroutine
			go func() {
				// 执行某些检查
				time.Sleep(2 * time.Second)
				// 这个goroutine也会一直存在
			}()
			
			// 检查资源状态
			checkResource(resourceID)
		}
	}()

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Watching started"))
}

func checkResource(id string) {
	// 模拟资源检查
	time.Sleep(100 * time.Millisecond)
}

// SetupTaskRoutes 设置路由
func SetupTaskRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/tasks", ProcessTask)
	mux.HandleFunc("GET /api/events", StreamEvents)
	mux.HandleFunc("GET /api/tasks/status", PollStatus)
	mux.HandleFunc("GET /api/watch/{id}", WatchResource)
}
