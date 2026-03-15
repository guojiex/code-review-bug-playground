# Bug GO-001: 空指针异常

## Bug信息
- **ID**: GO-001
- **维度**: Bug
- **类型**: null_pointer_exception
- **严重性**: High
- **语言**: Go
- **文件**: backend-go/test-cases/01-bug-null-pointer/buggy.go

## 问题描述

函数 `GetUserProfile` 在用户不存在时会触发空指针panic。代码存在两个问题：
1. 忽略了 `FindByID` 返回的error
2. 没有检查user是否为nil就直接访问其字段

## 代码示例（有问题的代码）

```go
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
```

## 触发条件

1. 当传入的userID不存在时
2. 当FindByID返回(nil, nil)时
3. 当数据库查询出错时

**测试方法**：
```bash
# 访问一个不存在的用户ID
curl http://localhost:8080/api/users/nonexistent-id/profile
```

## 影响

### 运行时影响
- **生产环境API崩溃**: 整个goroutine会panic
- **返回500错误**: 用户看到内部服务器错误
- **服务不可用**: 如果没有recover机制，可能导致服务崩溃

### 影响范围
- **所有调用该接口的用户**: `/api/users/:id/profile` 端点完全不可用
- **系统稳定性**: panic可能导致goroutine泄漏
- **用户体验**: 用户无法查看任何用户资料

### 严重程度评估
- **CVSS评分**: 7.5 (High)
- **可触发性**: 容易触发（任何无效ID即可）
- **影响广度**: 影响所有用户

## 证据

### Go语言行为
在Go语言中，对nil指针调用方法或访问字段会导致panic：
```go
var user *User = nil
name := user.Name  // panic: runtime error: invalid memory address
```

### 函数签名分析
```go
FindByID(id string) (*User, error)
```
- 返回类型是 `*User`，可能为nil
- 当用户不存在时，某些实现会返回 `(nil, nil)`
- 当数据库错误时，某些实现会返回 `(nil, error)`

### 错误处理缺失
代码使用 `_` 忽略error，违反了Go错误处理最佳实践。

## 修复方案

### 方案1：添加完整的错误检查（推荐）

```go
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
```

### 方案2：使用辅助函数

```go
func (h *UserHandler) getUserOrFail(c *gin.Context, userID string) (*User, bool) {
    user, err := h.repo.FindByID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return nil, false
    }
    if user == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return nil, false
    }
    return user, true
}

func (h *UserHandler) GetUserProfile(c *gin.Context) {
    userID := c.Param("id")
    
    user, ok := h.getUserOrFail(c, userID)
    if !ok {
        return
    }
    
    profile := &UserProfile{
        Name:  user.Name,
        Email: user.Email,
        Bio:   "User bio",
    }
    
    c.JSON(http.StatusOK, profile)
}
```

## 预期审查结果

当AI代码审查工具审查buggy.go时，应该识别并输出：

### 必须识别的问题

```json
{
  "type": "bug",
  "subcategory": "null_pointer_exception",
  "severity": "high",
  "title": "潜在的空指针引用",
  "description": "函数GetUserProfile在user为nil时直接访问user.Name和user.Email会导致panic。FindByID在用户不存在时可能返回(nil, nil)或(nil, error)，代码忽略了error且未检查user是否为nil。",
  "line": 34,
  "evidence": "1) FindByID的返回类型是(*User, error)，user可能为nil；2) Go语言对nil指针访问字段会panic；3) 代码使用'_'忽略了error",
  "impact": "生产环境API崩溃，影响所有调用/api/users/:id/profile接口的用户。当userID无效或不存在时，会导致panic和500错误，可能造成服务不可用。",
  "suggestion": "添加错误检查和空值判断：\n1. 检查FindByID返回的error\n2. 检查user是否为nil\n3. 返回适当的HTTP状态码（404或500）"
}
```

### 评分标准
- **必须发现**: 空指针问题（否则为漏报）
- **严重性必须正确**: High级别（Critical/High都可接受）
- **必须有证据**: 说明为什么会panic
- **必须有影响**: 说明会导致什么后果
- **必须有建议**: 提供可执行的修复方案

### 不应该出现的误报
- ❌ 不应该报告"user未定义"（user是通过FindByID返回的）
- ❌ 不应该报告"类型错误"（类型是正确的）
- ❌ 不应该报告"函数名不清晰"（这是Low级别，不应该在此报告）

## 相关资源

- [Go错误处理最佳实践](https://go.dev/blog/error-handling-and-go)
- [Effective Go - Errors](https://golang.org/doc/effective_go#errors)
- [Go Code Review Comments - Error Handling](https://github.com/golang/go/wiki/CodeReviewComments#error-handling)

## 测试用例

### 正常情况
```bash
# 应该成功返回用户资料
curl http://localhost:8080/api/users/valid-user-id/profile
# 预期: 200 OK
```

### Bug触发情况
```bash
# buggy版本会panic
curl http://localhost:8080/api/users/invalid-id/profile
# buggy版本: panic
# fixed版本: 404 Not Found
```

---

*这是一个典型的Go语言空指针问题，在生产环境中非常常见且危险。*
