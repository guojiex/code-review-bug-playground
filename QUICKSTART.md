# Bug Playground 快速启动指南

## 🎯 项目目标

这个项目用于：
1. ✅ 展示AI代码审查能力
2. ✅ 提供测试用例集
3. ✅ 训练和优化提示词
4. ✅ 营销案例展示

## 📂 当前状态

### 已完成
- ✅ 项目README和ACTION_PLAN
- ✅ 第一个完整的bug示例（GO-001: 空指针异常）
- ✅ Bug文档模板

### 待完成
- [ ] 49个其他bug示例
- [ ] Bug目录索引
- [ ] 测试脚本
- [ ] 完整的文档

## 🚀 快速使用

### 1. 查看第一个bug示例

```bash
# 查看有问题的代码
cat backend-go/test-cases/01-bug-null-pointer/buggy.go

# 查看修复后的代码
cat backend-go/test-cases/01-bug-null-pointer/fixed.go

# 查看详细说明
cat backend-go/test-cases/01-bug-null-pointer/explanation.md
```

### 2. 测试AI审查工具

```bash
# 用你的AI审查工具审查buggy.go
your-ai-tool review backend-go/test-cases/01-bug-null-pointer/buggy.go

# 对比结果
# 检查是否识别出：
# - 类型: null_pointer_exception
# - 严重性: High
# - 有完整的证据和影响说明
```

## 📋 下一步工作

### 立即可做（今天）

1. **创建更多bug示例**
   
   复制第一个示例作为模板：
   ```bash
   cp -r backend-go/test-cases/01-bug-null-pointer backend-go/test-cases/02-security-sql-injection
   ```
   
   然后修改：
   - buggy.go - 改为SQL注入的bug
   - fixed.go - 改为修复后的代码
   - explanation.md - 更新说明文档

2. **创建bug目录索引**
   
   在 `docs/bug-catalog.md` 中列出所有bug：
   ```markdown
   # Bug目录索引
   
   ## 按维度分类
   
   ### Bug维度
   - GO-001: 空指针异常 [High]
   - GO-002: 数组越界 [High]
   ...
   
   ### Security维度
   - GO-011: SQL注入 [Critical]
   ...
   ```

3. **开始实现正常功能**
   
   按照ACTION_PLAN Day 2-3的步骤，创建：
   - 数据模型（model/）
   - Repository层
   - Service层
   - Handler层

### 本周可做

1. **完成Go语言的40个bug示例**
   - Bug维度: 10个
   - Security维度: 10个
   - Performance维度: 7个
   - Maintainability维度: 8个
   - Error Handling维度: 5个

2. **创建测试脚本**
   ```python
   # scripts/test_review_quality.py
   # 自动测试所有bug示例
   ```

3. **完善文档**
   - Bug目录索引
   - 使用指南
   - 预期结果汇总

## 💡 创建Bug的技巧

### 1. 选择真实场景

基于推客联盟系统的真实场景：
- 用户管理
- 佣金计算
- 订单处理
- 数据统计

### 2. 确保bug可触发

每个bug都应该：
- 有明确的触发条件
- 能实际运行
- 真的会出问题

### 3. 提供完整说明

每个bug都要有：
- 问题描述
- 触发条件
- 影响分析
- 证据说明
- 修复方案
- 预期审查结果

### 4. 难度分级

- **简单bug**: 明显的空指针、SQL注入
- **中等bug**: N+1查询、并发问题
- **困难bug**: 微妙的逻辑错误、性能问题

## 📊 质量标准

每个bug示例必须：
- ✅ 代码可以编译
- ✅ bug真实存在
- ✅ 修复确实有效
- ✅ 说明文档完整
- ✅ 符合命名规范

## 🔗 相关文档

- [README.md](./README.md) - 项目总览
- [ACTION_PLAN.md](./ACTION_PLAN.md) - 详细行动步骤
- [第一个bug示例](./backend-go/test-cases/01-bug-null-pointer/) - 参考模板

## 🎯 成功指标

完成后应该达到：
- ✅ 50个完整的bug示例
- ✅ 覆盖5个维度
- ✅ 覆盖4种语言
- ✅ 每个bug都有buggy/fixed/explanation
- ✅ 可以用来测试审查工具的准确率

---

**开始构建吧！第一个示例已经完成，可以作为模板快速创建其他示例。** 🚀
