# Bug目录索引

> 所有bug示例的快速索引和查找

---

## 📊 总览

| 维度 | Bug数量 | Critical | High | Medium | Low |
|------|---------|----------|------|--------|-----|
| Bug | 1/15 | 0 | 1 | 0 | 0 |
| Security | 0/10 | 0 | 0 | 0 | 0 |
| Maintainability | 0/8 | 0 | 0 | 0 | 0 |
| Performance | 0/7 | 0 | 0 | 0 | 0 |
| Error Handling | 0/10 | 0 | 0 | 0 | 0 |
| **总计** | **1/50** | **0** | **1** | **0** | **0** |

**进度**: 2% (1/50) ■□□□□□□□□□

---

## 📋 按维度分类

### 1️⃣ Bug维度（1/15）

| ID | 名称 | 严重性 | 语言 | 文件路径 | 状态 |
|----|----|--------|------|----------|------|
| GO-001 | 空指针异常 | High | Go | test-cases/01-bug-null-pointer/ | ✅ 完成 |
| GO-002 | 数组越界 | High | Go | test-cases/02-bug-array-bounds/ | ⏳ 待创建 |
| GO-003 | 并发数据竞争 | Critical | Go | test-cases/03-bug-data-race/ | ⏳ 待创建 |
| GO-004 | 资源泄漏 | High | Go | test-cases/04-bug-resource-leak/ | ⏳ 待创建 |
| GO-005 | 逻辑错误（佣金计算） | Critical | Go | test-cases/05-bug-logic-error/ | ⏳ 待创建 |
| GO-006 | 未检查错误返回值 | Medium | Go | test-cases/06-bug-unchecked-error/ | ⏳ 待创建 |
| GO-007 | goroutine泄漏 | High | Go | test-cases/07-bug-goroutine-leak/ | ⏳ 待创建 |
| GO-008 | channel未关闭 | Medium | Go | test-cases/08-bug-channel-not-closed/ | ⏳ 待创建 |
| GO-009 | defer位置不当 | Medium | Go | test-cases/09-bug-defer-in-loop/ | ⏳ 待创建 |
| GO-010 | 类型断言未检查 | High | Go | test-cases/10-bug-type-assertion/ | ⏳ 待创建 |

### 2️⃣ Security维度（0/10）

| ID | 名称 | 严重性 | 语言 | 文件路径 | 状态 |
|----|----|--------|------|----------|------|
| GO-011 | SQL注入 | Critical | Go | test-cases/11-security-sql-injection/ | ⏳ 待创建 |
| GO-012 | XSS漏洞 | High | Go | test-cases/12-security-xss/ | ⏳ 待创建 |
| GO-013 | 路径遍历 | Critical | Go | test-cases/13-security-path-traversal/ | ⏳ 待创建 |
| GO-014 | 敏感信息泄露 | High | Go | test-cases/14-security-data-exposure/ | ⏳ 待创建 |
| GO-015 | 硬编码密钥 | Critical | Go | test-cases/15-security-hardcoded-secret/ | ⏳ 待创建 |
| GO-016 | 弱加密算法 | High | Go | test-cases/16-security-weak-crypto/ | ⏳ 待创建 |
| GO-017 | 命令注入 | Critical | Go | test-cases/17-security-command-injection/ | ⏳ 待创建 |
| GO-018 | 不安全的随机数 | Medium | Go | test-cases/18-security-weak-random/ | ⏳ 待创建 |
| GO-019 | CORS配置错误 | Medium | Go | test-cases/19-security-cors-misconfiguration/ | ⏳ 待创建 |
| GO-020 | 权限控制缺失 | High | Go | test-cases/20-security-missing-authorization/ | ⏳ 待创建 |

### 3️⃣ Performance维度（0/7）

| ID | 名称 | 严重性 | 语言 | 文件路径 | 状态 |
|----|----|--------|------|----------|------|
| GO-021 | N+1查询 | Medium | Go | test-cases/21-performance-n-plus-1/ | ⏳ 待创建 |
| GO-022 | 低效算法 | Medium | Go | test-cases/22-performance-inefficient-algorithm/ | ⏳ 待创建 |
| GO-023 | 内存泄漏 | High | Go | test-cases/23-performance-memory-leak/ | ⏳ 待创建 |
| GO-024 | 大对象拷贝 | Medium | Go | test-cases/24-performance-unnecessary-copy/ | ⏳ 待创建 |
| GO-025 | 字符串拼接 | Low | Go | test-cases/25-performance-string-concat/ | ⏳ 待创建 |
| GO-026 | 同步阻塞IO | Medium | Go | test-cases/26-performance-blocking-io/ | ⏳ 待创建 |
| GO-027 | 频繁GC压力 | Low | Go | test-cases/27-performance-gc-pressure/ | ⏳ 待创建 |

### 4️⃣ Maintainability维度（0/8）

| ID | 名称 | 严重性 | 语言 | 文件路径 | 状态 |
|----|----|--------|------|----------|------|
| GO-028 | 魔法数字 | Low | Go | test-cases/28-maintainability-magic-numbers/ | ⏳ 待创建 |
| GO-029 | 重复代码 | Low | Go | test-cases/29-maintainability-duplication/ | ⏳ 待创建 |
| GO-030 | 过长函数 | Low | Go | test-cases/30-maintainability-long-function/ | ⏳ 待创建 |
| GO-031 | 过深嵌套 | Low | Go | test-cases/31-maintainability-deep-nesting/ | ⏳ 待创建 |
| GO-032 | 命名不清晰 | Info | Go | test-cases/32-maintainability-poor-naming/ | ⏳ 待创建 |
| GO-033 | 上帝类 | Low | Go | test-cases/33-maintainability-god-class/ | ⏳ 待创建 |
| GO-034 | 注释不清 | Info | Go | test-cases/34-maintainability-poor-comments/ | ⏳ 待创建 |
| GO-035 | 硬编码配置 | Low | Go | test-cases/35-maintainability-hardcoded-config/ | ⏳ 待创建 |

### 5️⃣ Error Handling维度（0/10）

| ID | 名称 | 严重性 | 语言 | 文件路径 | 状态 |
|----|----|--------|------|----------|------|
| GO-036 | 空catch块 | Medium | Go | test-cases/36-error-empty-catch/ | ⏳ 待创建 |
| GO-037 | 错误信息不明确 | Low | Go | test-cases/37-error-unclear-message/ | ⏳ 待创建 |
| GO-038 | panic未恢复 | High | Go | test-cases/38-error-unrecovered-panic/ | ⏳ 待创建 |
| GO-039 | 错误类型不当 | Low | Go | test-cases/39-error-type-mismatch/ | ⏳ 待创建 |
| GO-040 | 错误日志缺失 | Medium | Go | test-cases/40-error-missing-log/ | ⏳ 待创建 |

---

## 🔍 快速查找

### 按严重性查找

#### Critical (0)
- 待添加...

#### High (1)
- **GO-001**: 空指针异常 [Bug]

#### Medium (0)
- 待添加...

#### Low (0)
- 待添加...

### 按语言查找

#### Go语言 (1/40)
- GO-001 ✅

#### Java (0/8)
- 待添加...

#### Python (0/7)
- 待添加...

#### JavaScript (0/5)
- 待添加...

---

## 📖 使用方法

### 查看某个bug

```bash
# 查看GO-001的三个文件
cat backend-go/test-cases/01-bug-null-pointer/buggy.go
cat backend-go/test-cases/01-bug-null-pointer/fixed.go
cat backend-go/test-cases/01-bug-null-pointer/explanation.md
```

### 测试审查工具

```bash
# 审查所有Bug维度的示例
for dir in backend-go/test-cases/0{1..10}-*; do
    your-ai-tool review "$dir/buggy.go"
done
```

### 批量测试

```bash
# 使用测试脚本
python scripts/test_review_quality.py --all
```

---

## 📊 统计信息

### 完成进度
- **总体**: 2% (1/50)
- **Bug维度**: 6.7% (1/15)
- **Security维度**: 0% (0/10)
- **Performance维度**: 0% (0/7)
- **Maintainability维度**: 0% (0/8)
- **Error Handling维度**: 0% (0/10)

### 严重性分布（目标）
- **Critical**: 0/5 (0%)
- **High**: 1/20 (5%)
- **Medium**: 0/17 (0%)
- **Low**: 0/8 (0%)

---

## 🎯 下一步

### 优先级1: 高危bug（Critical + High）
建议优先创建这些bug示例：
1. GO-003: 并发数据竞争 [Critical]
2. GO-005: 逻辑错误（佣金计算） [Critical]
3. GO-011: SQL注入 [Critical]
4. GO-013: 路径遍历 [Critical]
5. GO-015: 硬编码密钥 [Critical]

### 优先级2: Medium bug
重要但不紧急的问题

### 优先级3: Low bug
代码质量问题

---

*本目录会随着项目进展持续更新*
