# TAPD SDK API 实现指南

本文档记录了如何在 go-tapd SDK 中实现一个新的 TAPD API 接口的完整流程。

## 目录

- [前置准备](#前置准备)
- [实现步骤](#实现步骤)
- [代码规范](#代码规范)
- [测试规范](#测试规范)
- [完整示例](#完整示例)

---

## 前置准备

### 1. 查看 API 文档

访问 TAPD 官方 API 文档，了解：
- API 端点路径
- HTTP 方法（GET/POST/PUT/DELETE）
- 请求参数及类型
- 响应格式及字段
- 使用示例

**示例**: https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_fields_lable.html

### 2. 查看参考实现

在 GitHub 仓库中查找类似 API 的实现作为参考，特别关注：
- 结构体定义方式
- 方法签名格式
- 测试用例写法

**参考**: https://github.com/go-tapd/tapd/pull/277/files

---

## 实现步骤

### 步骤 1: 定义数据结构

在对应的 `api_*.go` 文件中添加请求和响应结构体。

#### 文件位置规则
- 需求相关: `api_story.go`
- 缺陷相关: `api_bug.go`
- 任务相关: `api_task.go`
- 迭代相关: `api_iteration.go`
- 其他类推...

#### 请求结构体规范

```go
// 命名格式: Get{Resource}{Action}Request
type GetStoryFieldsLabelRequest struct {
    WorkspaceID *int `url:"workspace_id,omitempty"` // 项目ID
    // 其他参数...
}
```

**规范要点**:
- ✅ 所有字段使用指针类型 (`*int`, `*string`, `*int64`)
- ✅ 使用 `url` tag 标记参数名
- ✅ 添加 `omitempty` 使参数可选
- ✅ 添加中文注释说明字段用途

#### 响应结构体规范

```go
// 命名格式: {Resource}{Description}
type StoryFieldLabel struct {
    EN string `json:"en,omitempty"` // 字段英文名
    CN string `json:"cn,omitempty"` // 字段中文标签
}
```

**规范要点**:
- ✅ 字段默认使用非指针类型
- ✅ 可为空的字段使用指针类型
- ✅ 使用 `json` tag 标记 JSON 字段名
- ✅ 添加中文注释

#### 特殊类型使用

```go
// 逗号分隔值: "1,2,3"
ID: tapd.NewMulti[int64](1, 2, 3)
Fields: tapd.NewMulti[string]("id", "name", "status")

// 管道分隔枚举值: "1|2|3"
Status: tapd.NewEnum[string]("open", "resolved")
```

### 步骤 2: 在接口中添加方法定义

在对应的 Service 接口中添加方法签名。

**位置**: 在 `api_*.go` 文件中找到 `type {Resource}Service interface` 定义

```go
type StoryService interface {
    // ... 其他方法

    // GetStoryFieldsLabel 获取需求所有字段的中英文
    //
    // https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_fields_lable.html
    GetStoryFieldsLabel(ctx context.Context, request *GetStoryFieldsLabelRequest, opts ...RequestOption) ([]*StoryFieldLabel, *Response, error)
}
```

**规范要点**:
- ✅ 添加中文注释说明功能
- ✅ 添加 API 文档链接
- ✅ 第一个参数必须是 `context.Context`
- ✅ 第二个参数是请求结构体指针
- ✅ 第三个参数是可选的 `opts ...RequestOption`
- ✅ 返回值格式: `(结果, *Response, error)`

### 步骤 3: 实现具体方法

在 `api_*.go` 文件中实现具体的方法逻辑。

```go
func (s *storyService) GetStoryFieldsLabel(
    ctx context.Context, request *GetStoryFieldsLabelRequest, opts ...RequestOption,
) ([]*StoryFieldLabel, *Response, error) {
    // 1. 创建 HTTP 请求
    req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/get_fields_lable", request, opts)
    if err != nil {
        return nil, nil, err
    }

    // 2. 定义响应变量
    var labelsMap map[string]string
    resp, err := s.client.Do(req, &labelsMap)
    if err != nil {
        return nil, resp, err
    }

    // 3. 数据转换（如需要）
    labels := make([]*StoryFieldLabel, 0, len(labelsMap))
    for en, cn := range labelsMap {
        labels = append(labels, &StoryFieldLabel{
            EN: en,
            CN: cn,
        })
    }

    return labels, resp, nil
}
```

**实现模式总结**:

#### GET 请求模式
```go
func (s *service) GetResource(ctx context.Context, request *GetRequest, opts ...RequestOption) (*Resource, *Response, error) {
    req, err := s.client.NewRequest(ctx, http.MethodGet, "endpoint/path", request, opts)
    if err != nil {
        return nil, nil, err
    }

    var response struct {
        Resource *Resource `json:"resource_key"`
    }
    resp, err := s.client.Do(req, &response)
    if err != nil {
        return nil, resp, err
    }

    return response.Resource, resp, nil
}
```

#### POST 请求模式
```go
func (s *service) CreateResource(ctx context.Context, request *CreateRequest, opts ...RequestOption) (*Resource, *Response, error) {
    req, err := s.client.NewRequest(ctx, http.MethodPost, "endpoint/path", request, opts)
    if err != nil {
        return nil, nil, err
    }

    var response struct {
        Resource *Resource `json:"resource_key"`
    }
    resp, err := s.client.Do(req, &response)
    if err != nil {
        return nil, resp, err
    }

    return response.Resource, resp, nil
}
```

#### 返回列表的模式
```go
func (s *service) GetResources(ctx context.Context, request *GetRequest, opts ...RequestOption) ([]*Resource, *Response, error) {
    req, err := s.client.NewRequest(ctx, http.MethodGet, "endpoint/path", request, opts)
    if err != nil {
        return nil, nil, err
    }

    var items []*struct {
        Resource *Resource `json:"Resource"` // 注意首字母大写
    }
    resp, err := s.client.Do(req, &items)
    if err != nil {
        return nil, resp, err
    }

    resources := make([]*Resource, 0, len(items))
    for _, item := range items {
        resources = append(resources, item.Resource)
    }

    return resources, resp, nil
}
```

#### 返回计数的模式
```go
func (s *service) GetResourceCount(ctx context.Context, request *GetRequest, opts ...RequestOption) (int, *Response, error) {
    req, err := s.client.NewRequest(ctx, http.MethodGet, "endpoint/count", request, opts)
    if err != nil {
        return 0, nil, err
    }

    var response struct {
        Count int `json:"count"`
    }
    resp, err := s.client.Do(req, &response)
    if err != nil {
        return 0, resp, err
    }

    return response.Count, resp, nil
}
```

### 步骤 4: 添加单元测试

在对应的 `api_*_test.go` 文件中添加单元测试。

#### 4.1 创建测试数据文件

**路径**: `internal/testdata/api/{resource}/{endpoint_name}.json`

```json
{
  "status": 1,
  "data": {
    "id": "ID",
    "name": "标题",
    "description": "详细描述"
  },
  "info": "success"
}
```

**规范要点**:
- ✅ 必须包含 `status`, `data`, `info` 三个顶层字段
- ✅ `status` 为整数类型（通常为 1）
- ✅ `data` 包含实际响应数据
- ✅ 文件名与 API 端点对应

#### 4.2 编写测试用例

```go
func TestStoryService_GetStoryFieldsLabel(t *testing.T) {
    // 1. 创建 mock 服务器和客户端
    _, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // 2. 验证请求方法
        assert.Equal(t, http.MethodGet, r.Method)
        // 3. 验证请求路径
        assert.Equal(t, "/stories/get_fields_lable", r.URL.Path)
        // 4. 验证请求参数
        assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

        // 5. 返回测试数据
        _, _ = w.Write(loadData(t, "internal/testdata/api/story/get_fields_lable.json"))
    }))

    // 6. 调用方法
    labels, _, err := client.StoryService.GetStoryFieldsLabel(ctx, &GetStoryFieldsLabelRequest{
        WorkspaceID: Ptr(11112222),
    })

    // 7. 验证返回结果
    assert.NoError(t, err)
    assert.True(t, len(labels) > 0)

    // 8. 验证具体字段值
    labelMap := make(map[string]string)
    for _, label := range labels {
        labelMap[label.EN] = label.CN
    }
    assert.Equal(t, "ID", labelMap["id"])
    assert.Equal(t, "标题", labelMap["name"])
}
```

**测试规范**:
- ✅ 函数名格式: `Test{Service}_{MethodName}`
- ✅ 验证 HTTP 方法、路径、参数
- ✅ 使用 `assert` 进行断言
- ✅ 测试至少验证几个关键字段

#### 4.3 运行单元测试

```bash
# 运行所有测试
go test . -race

# 运行特定测试
go test . -run TestStoryService_GetStoryFieldsLabel -v
```

### 步骤 5: 添加集成测试

在 `tests/api_*_prod_test.go` 文件中添加生产环境测试。

```go
func TestStoryService_Prod_GetStoryFieldsLabel(t *testing.T) {
    // 1. 调用真实 API
    labels, _, err := createClient(t).StoryService.GetStoryFieldsLabel(ctx, &tapd.GetStoryFieldsLabelRequest{
        WorkspaceID: tapd.Ptr(workspace8591ID),
    })

    // 2. 验证响应
    assert.NoError(t, err)
    assert.NotEmpty(t, labels)

    // 3. 输出结果供人工验证
    spew.Dump(labels)
}
```

**集成测试规范**:
- ✅ 函数名格式: `Test{Service}_Prod_{MethodName}`
- ✅ 使用真实的 workspace ID
- ✅ 使用 `spew.Dump()` 输出完整响应
- ✅ 基本的错误和空值检查

#### 5.1 运行集成测试

```bash
# ⚠️ 注意：必须运行单个测试，不能运行所有 tests
go test -v -run ^TestStoryService_Prod_GetStoryFieldsLabel$ ./tests/
```

### 步骤 6: 更新功能清单

在 `features.md` 中标记 API 为已实现。

```markdown
### 需求

- [x] 获取需求所有字段的中英文
```

---

## 代码规范

### 命名规范

| 类型 | 格式 | 示例 |
|-----|------|------|
| 请求结构体 | `{Action}{Resource}Request` | `GetStoryFieldsLabelRequest` |
| 响应结构体 | `{Resource}{Description}` | `StoryFieldLabel` |
| 接口方法 | `{Action}{Resource}` | `GetStoryFieldsLabel` |
| 单元测试 | `Test{Service}_{Method}` | `TestStoryService_GetStoryFieldsLabel` |
| 集成测试 | `Test{Service}_Prod_{Method}` | `TestStoryService_Prod_GetStoryFieldsLabel` |

### 参数类型规范

```go
// ✅ 正确：请求参数使用指针
type Request struct {
    ID          *int64  `url:"id,omitempty"`
    Name        *string `url:"name,omitempty"`
    WorkspaceID *int    `url:"workspace_id,omitempty"`
}

// ✅ 正确：响应字段使用值类型（除非可能为 null）
type Response struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Description *string `json:"description"` // 可能为 null
}

// ❌ 错误：请求参数不使用指针
type Request struct {
    ID          int64  `url:"id"`      // ❌
    Name        string `url:"name"`    // ❌
}
```

### 注释规范

```go
// ✅ 正确的注释格式

// GetStoryFieldsLabel 获取需求所有字段的中英文
//
// https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_fields_lable.html
GetStoryFieldsLabel(ctx context.Context, request *GetStoryFieldsLabelRequest, opts ...RequestOption) ([]*StoryFieldLabel, *Response, error)

// StoryFieldLabel 需求字段标签
type StoryFieldLabel struct {
    EN string `json:"en,omitempty"` // 字段英文名
    CN string `json:"cn,omitempty"` // 字段中文标签
}
```

---

## 测试规范

### 单元测试要点

1. **完整性**: 覆盖所有请求参数和响应字段
2. **独立性**: 每个测试用例独立运行，不依赖其他测试
3. **可读性**: 测试代码清晰，易于理解测试意图
4. **速度**: 使用 mock 数据，测试快速执行

### 集成测试要点

1. **真实性**: 调用真实 API，使用真实数据
2. **稳定性**: 使用稳定存在的测试数据
3. **隔离性**: 单个运行，不批量执行
4. **可验证性**: 输出完整响应供人工检查

### 测试数据准备

```bash
# 单元测试数据
internal/testdata/api/
├── story/
│   ├── get_stories.json
│   ├── get_story_fields_lable.json
│   └── get_story_tcase.json
├── bug/
│   └── get_bugs.json
└── ...
```

**数据格式要求**:
```json
{
  "status": 1,        // 必须：状态码
  "data": {},         // 必须：数据内容（可以是对象或数组）
  "info": "success"   // 必须：信息提示
}
```

---

## 完整示例

以下是实现 `GetStoryFieldsLabel` API 的完整示例代码：

### 1. 数据结构定义 (api_story.go)

```go
// GetStoryFieldsLabelRequest 获取需求字段标签请求
type GetStoryFieldsLabelRequest struct {
    WorkspaceID *int `url:"workspace_id,omitempty"` // 项目ID
}

// StoryFieldLabel 需求字段标签
type StoryFieldLabel struct {
    EN string `json:"en,omitempty"` // 字段英文名
    CN string `json:"cn,omitempty"` // 字段中文标签
}
```

### 2. 接口方法定义 (api_story.go)

```go
type StoryService interface {
    // ... 其他方法

    // GetStoryFieldsLabel 获取需求所有字段的中英文
    //
    // https://open.tapd.cn/document/api-doc/API文档/api_reference/story/get_story_fields_lable.html
    GetStoryFieldsLabel(ctx context.Context, request *GetStoryFieldsLabelRequest, opts ...RequestOption) ([]*StoryFieldLabel, *Response, error)
}
```

### 3. 方法实现 (api_story.go)

```go
func (s *storyService) GetStoryFieldsLabel(
    ctx context.Context, request *GetStoryFieldsLabelRequest, opts ...RequestOption,
) ([]*StoryFieldLabel, *Response, error) {
    req, err := s.client.NewRequest(ctx, http.MethodGet, "stories/get_fields_lable", request, opts)
    if err != nil {
        return nil, nil, err
    }

    var labelsMap map[string]string
    resp, err := s.client.Do(req, &labelsMap)
    if err != nil {
        return nil, resp, err
    }

    labels := make([]*StoryFieldLabel, 0, len(labelsMap))
    for en, cn := range labelsMap {
        labels = append(labels, &StoryFieldLabel{
            EN: en,
            CN: cn,
        })
    }

    return labels, resp, nil
}
```

### 4. 测试数据 (internal/testdata/api/story/get_fields_lable.json)

```json
{
  "status": 1,
  "data": {
    "id": "ID",
    "name": "标题",
    "description": "详细描述",
    "workspace_id": "项目ID",
    "creator": "创建人",
    "created": "创建时间",
    "status": "状态"
  },
  "info": "success"
}
```

### 5. 单元测试 (api_story_test.go)

```go
func TestStoryService_GetStoryFieldsLabel(t *testing.T) {
    _, client := createServerClient(t, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        assert.Equal(t, http.MethodGet, r.Method)
        assert.Equal(t, "/stories/get_fields_lable", r.URL.Path)
        assert.Equal(t, "11112222", r.URL.Query().Get("workspace_id"))

        _, _ = w.Write(loadData(t, "internal/testdata/api/story/get_fields_lable.json"))
    }))

    labels, _, err := client.StoryService.GetStoryFieldsLabel(ctx, &GetStoryFieldsLabelRequest{
        WorkspaceID: Ptr(11112222),
    })
    assert.NoError(t, err)
    assert.True(t, len(labels) > 0)

    labelMap := make(map[string]string)
    for _, label := range labels {
        labelMap[label.EN] = label.CN
    }

    assert.Equal(t, "ID", labelMap["id"])
    assert.Equal(t, "标题", labelMap["name"])
    assert.Equal(t, "详细描述", labelMap["description"])
}
```

### 6. 集成测试 (tests/api_story_prod_test.go)

```go
func TestStoryService_Prod_GetStoryFieldsLabel(t *testing.T) {
    labels, _, err := createClient(t).StoryService.GetStoryFieldsLabel(ctx, &tapd.GetStoryFieldsLabelRequest{
        WorkspaceID: tapd.Ptr(workspace8591ID),
    })
    assert.NoError(t, err)
    assert.NotEmpty(t, labels)

    spew.Dump(labels)
}
```

### 7. 运行测试

```bash
# 单元测试
go test . -run TestStoryService_GetStoryFieldsLabel -v

# 集成测试
go test -v -run ^TestStoryService_Prod_GetStoryFieldsLabel$ ./tests/

# 所有单元测试（带竞态检测）
go test . -race
```

### 8. 更新文档 (features.md)

```markdown
### 需求

- [x] 获取需求所有字段的中英文
```

---

## 常见问题

### Q1: 如何判断使用 GET 还是 POST？

**A**: 查看 API 文档说明：
- GET: 查询、获取数据
- POST: 创建、更新数据

### Q2: 响应数据是对象还是数组？

**A**: 根据 API 文档和测试响应确定：

```go
// 单个对象
var response struct {
    Story *Story `json:"Story"`
}

// 对象数组
var items []*struct {
    Story *Story `json:"Story"`
}
```

### Q3: 字段首字母大小写如何确定？

**A**:
- JSON tag 中的字段名严格按照 API 文档
- Go 结构体字段名首字母大写（导出）
- 响应包装对象的 JSON key 通常首字母大写（如 `Story`, `Bug`）

### Q4: 什么时候需要数据转换？

**A**: 当 API 返回的格式与期望的格式不同时：
```go
// API 返回 map，但我们希望返回结构体数组
var labelsMap map[string]string  // API 返回格式
labels := make([]*StoryFieldLabel, 0)  // 转换为期望格式
```

### Q5: 测试失败如何调试？

**A**:
1. 检查测试数据格式是否正确（必须包含 status, data, info）
2. 验证 JSON tag 与 API 返回字段名匹配
3. 使用 `spew.Dump()` 输出实际响应查看
4. 对比参考实现查找差异

---

## 最佳实践

### ✅ DO

1. **先看文档，后写代码** - 理解 API 行为再开始实现
2. **参考已有实现** - 保持代码风格一致
3. **完整测试覆盖** - 单元测试 + 集成测试
4. **清晰的注释** - 中文说明 + API 链接
5. **运行所有测试** - 确保没有破坏现有功能
6. **更新文档** - 标记 API 为已实现

### ❌ DON'T

1. **不要跳过测试** - 测试是代码质量保证
2. **不要批量运行集成测试** - 可能导致 API 限流
3. **不要省略注释** - 帮助他人理解代码
4. **不要修改测试数据格式** - 保持标准格式
5. **不要忽略错误处理** - 所有错误都应正确传递

---

## 检查清单

实现完成后，使用此清单验证：

- [ ] 数据结构定义完整（请求 + 响应）
- [ ] 接口方法签名正确
- [ ] 方法实现符合模式
- [ ] 添加了测试数据文件
- [ ] 单元测试通过
- [ ] 集成测试通过（单个运行）
- [ ] 所有测试通过（`go test . -race`）
- [ ] 添加了完整注释（中文 + 链接）
- [ ] 更新了 features.md
- [ ] 代码格式化（gofumpt）

---

## 总结

实现一个 TAPD API 接口的核心步骤：

1. **定义** - 结构体（请求 + 响应）
2. **声明** - 接口方法签名
3. **实现** - 具体方法逻辑
4. **测试** - 单元测试 + 集成测试
5. **文档** - 更新功能清单

遵循此指南，保持代码风格一致，确保测试覆盖，你就能高效地为 go-tapd SDK 添加新功能！

---

**最后更新**: 2025-11-04
**维护者**: Claude Code & go-tapd Contributors