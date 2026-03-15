# Code Review Bug Playground - 详细行动步骤

> 分阶段构建一个完整的、有bug的示例项目

---

## 📅 总体时间线

| 阶段 | 时间 | 主要任务 | 产出 |
|------|------|----------|------|
| 第1周 | Day 1-3 | 项目初始化、Go基础结构 | 可运行的基础框架 |
| 第2周 | Day 4-7 | Go核心功能+Bug | 30个Go示例 |
| 第3周 | Day 8-10 | Java/Python/JS示例 | 20个其他语言示例 |
| 第4周 | Day 11-14 | 文档完善、测试脚本 | 完整的文档和工具 |

---

## 📋 第一阶段：项目初始化（Day 1-3）

### Day 1: 创建项目结构

#### 任务清单

- [ ] 创建项目根目录和基础结构
  ```bash
  cd code-review-bug-playground
  
  # 创建Go项目目录
  mkdir -p backend-go/cmd/server
  mkdir -p backend-go/internal/{handler,service,repository,model}
  mkdir -p backend-go/test-cases
  
  # 创建其他语言目录
  mkdir -p backend-java/src/main/java/com/affiliate/{controller,service,repository}
  mkdir -p backend-python/app/{routes,services,models}
  mkdir -p frontend-js/src/{components,services,utils}
  
  # 创建文档和脚本目录
  mkdir -p docs scripts
  ```

- [ ] 初始化Go项目
  ```bash
  cd backend-go
  go mod init github.com/yourusername/code-review-bug-playground
  
  # 安装依赖
  go get -u github.com/gin-gonic/gin
  go get -u gorm.io/gorm
  go get -u gorm.io/driver/mysql
  go get -u github.com/google/uuid
  ```

- [ ] 创建基础配置文件
  ```bash
  # backend-go/.gitignore
  # backend-go/Makefile
  # backend-go/config.yaml
  ```

**交付物**:
- ✅ 完整的目录结构
- ✅ Go模块初始化
- ✅ 基础配置文件

---

### Day 2: 创建数据模型和基础框架

#### 任务清单

- [ ] 定义数据模型
  
  **backend-go/internal/model/user.go**
  ```go
  package model
  
  import "time"
  
  // User 推客用户
  type User struct {
      ID        string    `json:"id"`
      Name      string    `json:"name"`
      Email     string    `json:"email"`
      Role      string    `json:"role"` // affiliate, merchant
      Status    string    `json:"status"`
      CreatedAt time.Time `json:"created_at"`
  }
  ```

  **backend-go/internal/model/product.go**
  ```go
  package model
  
  // Product 商品
  type Product struct {
      ID             string  `json:"id"`
      Name           string  `json:"name"`
      Price          float64 `json:"price"`
      CommissionRate float64 `json:"commission_rate"`
      MerchantID     string  `json:"merchant_id"`
  }
  ```

  **backend-go/internal/model/order.go**
  ```go
  package model
  
  import "time"
  
  // Order 订单
  type Order struct {
      ID          string    `json:"id"`
      ProductID   string    `json:"product_id"`
      AffiliateID string    `json:"affiliate_id"`
      Amount      float64   `json:"amount"`
      Status      string    `json:"status"` // pending, completed, cancelled
      CreatedAt   time.Time `json:"created_at"`
  }
  ```

  **backend-go/internal/model/commission.go**
  ```go
  package model
  
  // Commission 佣金
  type Commission struct {
      ID          string  `json:"id"`
      OrderID     string  `json:"order_id"`
      AffiliateID string  `json:"affiliate_id"`
      Amount      float64 `json:"amount"`
      Status      string  `json:"status"` // pending, paid
  }
  ```

- [ ] 创建基础的HTTP服务器
  
  **backend-go/cmd/server/main.go**
  ```go
  package main
  
  import (
      "log"
      "github.com/gin-gonic/gin"
  )
  
  func main() {
      r := gin.Default()
      
      // 健康检查
      r.GET("/health", func(c *gin.Context) {
          c.JSON(200, gin.H{"status": "ok"})
      })
      
      // API路由组
      api := r.Group("/api/v1")
      {
          api.GET("/users", func(c *gin.Context) {
              c.JSON(200, gin.H{"users": []string{}})
          })
      }
      
      log.Println("Server starting on :8080")
      if err := r.Run(":8080"); err != nil {
          log.Fatal("Failed to start server:", err)
      }
  }
  ```

- [ ] 创建内存数据存储（模拟数据库）
  
  **backend-go/internal/repository/memory_store.go**
  ```go
  package repository
  
  import (
      "sync"
      "code-review-bug-playground/internal/model"
  )
  
  // MemoryStore 内存数据存储（避免依赖真实数据库）
  type MemoryStore struct {
      users       map[string]*model.User
      products    map[string]*model.Product
      orders      map[string]*model.Order
      commissions map[string]*model.Commission
      mu          sync.RWMutex
  }
  
  func NewMemoryStore() *MemoryStore {
      return &MemoryStore{
          users:       make(map[string]*model.User),
          products:    make(map[string]*model.Product),
          orders:      make(map[string]*model.Order),
          commissions: make(map[string]*model.Commission),
      }
  }
  ```

**交付物**:
- ✅ 4个数据模型定义
- ✅ 可运行的HTTP服务器
- ✅ 内存数据存储

---

### Day 3: 创建第一个正常功能

#### 任务清单

- [ ] 实现用户管理Repository
  
  **backend-go/internal/repository/user_repository.go**
  ```go
  package repository
  
  import (
      "errors"
      "github.com/google/uuid"
      "code-review-bug-playground/internal/model"
  )
  
  type UserRepository struct {
      store *MemoryStore
  }
  
  func NewUserRepository(store *MemoryStore) *UserRepository {
      return &UserRepository{store: store}
  }
  
  func (r *UserRepository) Create(user *model.User) error {
      r.store.mu.Lock()
      defer r.store.mu.Unlock()
      
      if user.ID == "" {
          user.ID = uuid.New().String()
      }
      
      r.store.users[user.ID] = user
      return nil
  }
  
  func (r *UserRepository) FindByID(id string) (*model.User, error) {
      r.store.mu.RLock()
      defer r.store.mu.RUnlock()
      
      user, exists := r.store.users[id]
      if !exists {
          return nil, errors.New("user not found")
      }
      return user, nil
  }
  
  func (r *UserRepository) List() ([]*model.User, error) {
      r.store.mu.RLock()
      defer r.store.mu.RUnlock()
      
      users := make([]*model.User, 0, len(r.store.users))
      for _, user := range r.store.users {
          users = append(users, user)
      }
      return users, nil
  }
  ```

- [ ] 实现用户Service层
  
  **backend-go/internal/service/user_service.go**
  ```go
  package service
  
  import (
      "errors"
      "time"
      "code-review-bug-playground/internal/model"
      "code-review-bug-playground/internal/repository"
  )
  
  type UserService struct {
      repo *repository.UserRepository
  }
  
  func NewUserService(repo *repository.UserRepository) *UserService {
      return &UserService{repo: repo}
  }
  
  func (s *UserService) CreateUser(name, email, role string) (*model.User, error) {
      if name == "" || email == "" {
          return nil, errors.New("name and email are required")
      }
      
      user := &model.User{
          Name:      name,
          Email:     email,
          Role:      role,
          Status:    "active",
          CreatedAt: time.Now(),
      }
      
      if err := s.repo.Create(user); err != nil {
          return nil, err
      }
      
      return user, nil
  }
  
  func (s *UserService) GetUser(id string) (*model.User, error) {
      return s.repo.FindByID(id)
  }
  
  func (s *UserService) ListUsers() ([]*model.User, error) {
      return s.repo.List()
  }
  ```

- [ ] 实现用户Handler层
  
  **backend-go/internal/handler/user_handler.go**
  ```go
  package handler
  
  import (
      "net/http"
      "github.com/gin-gonic/gin"
      "code-review-bug-playground/internal/service"
  )
  
  type UserHandler struct {
      service *service.UserService
  }
  
  func NewUserHandler(service *service.UserService) *UserHandler {
      return &UserHandler{service: service}
  }
  
  type CreateUserRequest struct {
      Name  string `json:"name" binding:"required"`
      Email string `json:"email" binding:"required,email"`
      Role  string `json:"role" binding:"required"`
  }
  
  func (h *UserHandler) CreateUser(c *gin.Context) {
      var req CreateUserRequest
      if err := c.ShouldBindJSON(&req); err != nil {
          c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
          return
      }
      
      user, err := h.service.CreateUser(req.Name, req.Email, req.Role)
      if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }
      
      c.JSON(http.StatusCreated, user)
  }
  
  func (h *UserHandler) GetUser(c *gin.Context) {
      id := c.Param("id")
      
      user, err := h.service.GetUser(id)
      if err != nil {
          c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
          return
      }
      
      c.JSON(http.StatusOK, user)
  }
  
  func (h *UserHandler) ListUsers(c *gin.Context) {
      users, err := h.service.ListUsers()
      if err != nil {
          c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
          return
      }
      
      c.JSON(http.StatusOK, gin.H{"users": users})
  }
  ```

- [ ] 更新main.go集成所有组件

**交付物**:
- ✅ 完整的用户管理功能（增删改查）
- ✅ 可以运行和测试
- ✅ 三层架构（Handler-Service-Repository）

---

## 📋 第二阶段：创建Bug示例（Day 4-7）

### 创建Bug的通用规则

每创建一个bug示例，都要包含三个文件：
1. `xxx_buggy.go` - 有问题的代码
2. `xxx_fixed.go` - 修复后的代码  
3. `xxx.md` - 问题说明文档

### Bug文档模板

```markdown
# Bug [ID]: [标题]

## Bug信息
- **ID**: GO-001
- **维度**: Bug / Security / Maintainability / Performance / Error Handling
- **类型**: null_pointer_exception / sql_injection / magic_numbers / ...
- **严重性**: Critical / High / Medium / Low
- **语言**: Go
- **文件**: internal/handler/xxx_buggy.go

## 问题描述
[简要描述问题是什么]

## 代码示例
```go
// buggy代码片段
```

## 触发条件
[在什么条件下会触发这个bug]

## 影响
- [具体影响1]
- [具体影响2]
- [影响范围]

## 证据
[为什么这是个问题的具体证据]

## 修复方案
```go
// fixed代码片段
```

## 预期审查结果
当AI工具审查buggy代码时，应该输出：
- **维度**: [维度]
- **类型**: [类型]
- **严重性**: [严重性]
- **证据**: [应该识别的证据]
- **影响**: [应该识别的影响]
- **建议**: [应该给出的修复建议]
```

---

### Day 4: 创建10个Bug维度示例

[详细的bug创建任务，参照之前的ACTION_PLAN]

---

## 🎯 快速启动指南

### 方式1：逐步构建（推荐学习）

按照Day 1-14的步骤逐步构建项目。

### 方式2：快速生成骨架

```bash
# 运行项目初始化脚本
./scripts/init_project.sh

# 这个脚本会：
# 1. 创建所有目录
# 2. 初始化Go模块
# 3. 创建基础文件
# 4. 生成模板文件
```

### 方式3：使用AI辅助

利用AI助手快速生成bug示例：
```
提示词：
"请帮我创建一个Go语言的SQL注入bug示例，包含：
1. buggy版本（有SQL注入漏洞）
2. fixed版本（修复后）
3. 说明文档（按照模板格式）

场景：用户管理的查询接口"
```

---

## 📝 质量检查清单

### 代码质量
- [ ] 所有Go代码可以编译
- [ ] 所有buggy版本确实有bug
- [ ] 所有fixed版本确实修复了问题
- [ ] 代码注释清晰
- [ ] 符合Go代码规范

### 文档质量
- [ ] 每个bug都有完整的说明文档
- [ ] 说明文档包含所有必需字段
- [ ] Bug目录索引完整
- [ ] 使用指南清晰

### 测试质量
- [ ] 测试脚本可以运行
- [ ] 质量评估脚本工作正常
- [ ] 可以生成bug报告

---

## 🎯 成功标准

项目完成后应该达到：

1. **完整性**
   - ✅ 50个完整的bug示例
   - ✅ 每个bug都有buggy/fixed/explanation
   - ✅ 覆盖5个维度
   - ✅ 覆盖4种语言

2. **可用性**
   - ✅ 代码可以运行
   - ✅ 文档清晰完整
   - ✅ 有测试工具
   - ✅ 有使用指南

3. **质量**
   - ✅ Bug真实可触发
   - ✅ 修复确实有效
   - ✅ 说明准确详细
   - ✅ 预期结果合理

4. **实用性**
   - ✅ 可用于测试审查工具
   - ✅ 可用于训练提示词
   - ✅ 可用于营销展示
   - ✅ 可用于开发者教育

---

## 🔗 参考资源

- [Go代码规范](../.cursor/rules)
- [代码审查规则](../.cursor/code-review-rules.md)
- [提示词模板](../docs/prompt-templates/)
- [实施指南](../docs/implementation-guide.md)

---

*按照这个行动计划，可以在2-4周内完成一个高质量的bug playground项目*
