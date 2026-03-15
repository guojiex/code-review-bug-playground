# Code Review Bug Playground - 项目说明

> 一个用于测试和展示AI代码审查能力的示例项目

---

## 📋 项目简介

**项目名称**: 电商推客联盟管理系统 (Affiliate Marketing Management System)

**项目背景**: 模拟一个真实的东南亚电商推客联盟管理后台系统，用于：
1. 展示AI代码审查的能力（包含各种典型bug）
2. 提供代码审查的测试用例
3. 作为营销案例（展示能发现哪些问题）
4. 用于训练和优化审查提示词

**技术栈**: 
- 后端：Go + Gin + GORM + MySQL
- 前端：JavaScript + React
- 其他：Java、Python示例

---

## 🎯 项目目标

### 1. 模拟真实业务场景

实现一个简化版的推客联盟系统，包含：
- 用户管理（推客、商家）
- 商品管理
- 订单管理
- 佣金计算
- 数据统计

### 2. 包含各维度的典型bug

按照五维度框架，每个维度包含多个典型bug：
- **Bug**: 空指针、数组越界、逻辑错误等
- **Security**: SQL注入、XSS、敏感信息泄露等
- **Maintainability**: 魔法数字、重复代码、命名不清等
- **Performance**: N+1查询、低效算法、内存泄漏等
- **Error Handling**: 空catch、未检查错误、错误信息不明确等

### 3. 提供对照组

每个bug文件都配对一个正确实现，方便对比：
- `xxx_buggy.go` - 有问题的代码
- `xxx_fixed.go` - 修复后的代码
- `xxx_explanation.md` - 问题说明和预期审查结果

---

## 📁 项目结构

```
code-review-bug-playground/
├── README.md                           # 本文件
├── ACTION_PLAN.md                      # 详细行动步骤
├── go.mod                              # Go模块定义
├── go.sum
│
├── backend-go/                         # Go语言示例（主要）
│   ├── cmd/
│   │   └── server/
│   │       └── main.go                # 程序入口
│   │
│   ├── internal/
│   │   ├── handler/                   # HTTP处理器
│   │   │   ├── user_handler.go
│   │   │   ├── user_handler_buggy.go   # 包含bug的版本
│   │   │   ├── user_handler_fixed.go   # 修复后的版本
│   │   │   └── user_handler.md         # 问题说明
│   │   │
│   │   ├── service/                   # 业务逻辑层
│   │   │   ├── commission_service.go
│   │   │   ├── commission_service_buggy.go
│   │   │   ├── commission_service_fixed.go
│   │   │   └── commission_service.md
│   │   │
│   │   ├── repository/                # 数据访问层
│   │   │   ├── order_repository.go
│   │   │   ├── order_repository_buggy.go
│   │   │   ├── order_repository_fixed.go
│   │   │   └── order_repository.md
│   │   │
│   │   └── model/                     # 数据模型
│   │       ├── user.go
│   │       ├── order.go
│   │       └── commission.go
│   │
│   └── test-cases/                    # 独立的测试用例
│       ├── 01-bug-null-pointer/
│       │   ├── buggy.go
│       │   ├── fixed.go
│       │   └── explanation.md
│       ├── 02-security-sql-injection/
│       │   ├── buggy.go
│       │   ├── fixed.go
│       │   └── explanation.md
│       └── ...
│
├── backend-java/                       # Java语言示例
│   ├── src/
│   │   └── main/
│   │       └── java/
│   │           └── com/
│   │               └── affiliate/
│   │                   ├── controller/
│   │                   ├── service/
│   │                   └── repository/
│   └── test-cases/
│
├── backend-python/                     # Python语言示例
│   ├── app/
│   │   ├── routes/
│   │   ├── services/
│   │   └── models/
│   └── test-cases/
│
├── frontend-js/                        # JavaScript/React示例
│   ├── src/
│   │   ├── components/
│   │   ├── services/
│   │   └── utils/
│   └── test-cases/
│
├── docs/                               # 文档
│   ├── bug-catalog.md                 # Bug目录索引
│   ├── expected-results.md            # 预期审查结果
│   └── usage-guide.md                 # 使用指南
│
└── scripts/                            # 辅助脚本
    ├── generate_bug_report.sh         # 生成bug报告
    └── test_review_quality.py         # 测试审查质量
```

---

## 🐛 Bug分类与分布

### 按维度分类

| 维度 | Bug数量 | 严重性分布 | 文件数 |
|------|---------|------------|--------|
| Bug | 15 | Critical:2, High:5, Medium:8 | 15 |
| Security | 10 | Critical:3, High:5, Medium:2 | 10 |
| Maintainability | 8 | Medium:3, Low:5 | 8 |
| Performance | 7 | Medium:5, Low:2 | 7 |
| Error Handling | 10 | High:4, Medium:6 | 10 |
| **总计** | **50** | | **50** |

### 按语言分类

| 语言 | 文件数 | 说明 |
|------|--------|------|
| Go | 30 | 主要示例，完整的系统 |
| Java | 8 | 常见企业级场景 |
| Python | 7 | Web和数据处理场景 |
| JavaScript | 5 | 前端常见问题 |

---

## 📝 Bug示例说明

每个bug文件都包含：

### 1. 代码文件（`xxx_buggy.go`）
```go
// Bug ID: GO-001
// 维度: Bug
// 类型: null_pointer_exception
// 严重性: High
// 描述: 用户查询未检查空值

func GetUserProfile(userID string) *UserProfile {
    user := userRepo.FindByID(userID)  // 可能返回nil
    return &UserProfile{
        Name:  user.Name,   // 空指针！
        Email: user.Email,
    }
}
```

### 2. 修复文件（`xxx_fixed.go`）
```go
// Fixed version of GO-001

func GetUserProfile(userID string) (*UserProfile, error) {
    user, err := userRepo.FindByID(userID)
    if err != nil {
        return nil, fmt.Errorf("failed to find user: %w", err)
    }
    if user == nil {
        return nil, ErrUserNotFound
    }
    
    return &UserProfile{
        Name:  user.Name,
        Email: user.Email,
    }, nil
}
```

### 3. 说明文档（`xxx.md`）
```markdown
# Bug GO-001: 空指针异常

## 问题描述
函数GetUserProfile在用户不存在时会触发空指针panic。

## 触发条件
当userID不存在或无效时，FindByID返回nil。

## 影响
- 生产环境API崩溃
- 影响所有调用该接口的用户
- 返回500错误

## 预期审查结果
- 维度: Bug
- 类型: null_pointer_exception
- 严重性: High
- 证据: user可能为nil，直接访问会panic
- 建议: 添加空值检查和错误处理
```

---

## 🎓 使用场景

### 1. 测试代码审查能力
```bash
# 运行审查测试
./scripts/test_review_quality.py \
  --input backend-go/test-cases/01-bug-null-pointer/buggy.go \
  --expected backend-go/test-cases/01-bug-null-pointer/explanation.md
```

### 2. 训练提示词
- 使用bug文件训练AI识别问题
- 对比预期结果优化提示词
- 评估不同模型效果

### 3. 营销展示
```markdown
# 展示文案
"我们的AI代码审查系统在测试中：
- 发现了50个典型bug中的47个（94%检出率）
- 误报率仅为8%
- Medium级别准确度达到89%"
```

### 4. 开发者教育
- 展示常见错误模式
- 提供最佳实践对比
- 学习如何修复问题

---

## 🚀 快速开始

### 1. 克隆项目
```bash
git clone <repo-url>
cd code-review-bug-playground
```

### 2. 查看bug目录
```bash
cat docs/bug-catalog.md
```

### 3. 运行示例
```bash
# Go示例
cd backend-go
go run cmd/server/main.go

# 查看某个bug
cat internal/handler/user_handler_buggy.go
cat internal/handler/user_handler.md
```

### 4. 测试审查
```bash
# 使用你的AI审查工具审查buggy文件
your-review-tool internal/handler/user_handler_buggy.go

# 对比预期结果
diff <actual-result> <expected-result>
```

---

## 📊 质量评估指标

用这个项目评估代码审查系统的质量：

### 1. 准确性指标
```
准确率 = 正确识别的bug数 / 总bug数
误报率 = 误报的问题数 / 总报告问题数
漏报率 = 未识别的bug数 / 总bug数
```

### 2. 严重性准确度
```
Critical级别准确度 = 正确评级的Critical问题数 / 总Critical问题数
High级别准确度 = ...
Medium级别准确度 = ...
```

### 3. 证据完整性
```
证据完整率 = 有完整证据链的问题数 / 总报告问题数
```

---

## 🔄 持续更新

### 添加新的bug用例

1. **选择场景**: 确定业务场景和bug类型
2. **创建目录**: 在相应语言目录下创建
3. **编写代码**: 创建buggy、fixed、explanation三个文件
4. **更新索引**: 在`docs/bug-catalog.md`中添加
5. **测试验证**: 用AI工具测试审查效果

### 维护建议

- 定期review bug库的覆盖度
- 根据实际发现的新bug类型补充
- 更新预期审查结果
- 收集AI审查的实际效果

---

## 📚 相关文档

- [详细行动步骤](./ACTION_PLAN.md)
- [Bug目录索引](./docs/bug-catalog.md) - 待创建
- [预期审查结果](./docs/expected-results.md) - 待创建
- [使用指南](./docs/usage-guide.md) - 待创建

---

## 📞 联系方式

如有问题或建议，请联系项目维护者。

---

## 📄 许可证

MIT License

---

*本项目用于展示和测试AI代码审查能力，包含的bug均为教学目的*
