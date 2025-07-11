# TODO

API 文档：https://open.tapd.cn/document/api-doc/API%E6%96%87%E6%A1%A3/

> [!NOTE]  
> API 数量较多，因个人精力有限，暂未完全实现所有 API。如有希望实现的 API 可创建 issue 说明，后续的开发工作将优先考虑。
>
> 当然，也欢迎大家一起参与开发，共同完善。期待您的 PR。


```
开发说明：

1、所有请求参数的字段均为指针类型，响应字段默认非指针类型，若存在null值可使用指针类型
2、尽可能以精简的请求参数或结构体、响应参数或结构体
3、支持逗号分隔的列表，如：1,2,3，请使用 *Multi[T] 结构体，如 ID 则为 *Multi[int]，如 Fields 则为 *Multi[string]。使用时可使用 `NewMulti` 函数创建
4、支持枚举的列表，如：1|2|3，请使用 *Enum[T] 结构体，如 ID 则为 *Enum[int]，如 Fields 则为 *Enum[string]。使用时可使用 `NewEnum` 函数创建
```

## 研发协作API

### 需求

- [x] 创建需求: 待二审或重构
- [ ] 创建需求分类
- [ ] 复制需求
- [ ] 获取需求与其它需求的所有关联关系
- [x] 获取需求: 待二审或重构
- [x] 获取需求数量: 待二审或重构
- [ ] 获取保密需求
- [ ] 获取保密需求数量
- [x] 获取需求分类
- [x] 获取需求分类数量
- [x] 获取指定分类需求数量
- [x] 获取需求变更历史
- [ ] 获取需求变更次数
- [x] 获取需求自定义字段配置
- [ ] 获取需求与测试用例关联关系
- [ ] 获取需求前后置关系
- [ ] 批量新增或修改需求前后置关系
- [ ] 批量删除需求前后置关系
- [ ] 获取需求保密信息
- [ ] 批量修改保密信息
- [ ] 获取需求类别
- [x] 更新需求: 待二审或重构
- [ ] 更新需求的需求类别
- [ ] 获取需求所有字段及候选值
- [ ] 获取需求所有字段的中英文
- [x] 获取需求模板列表
- [x] 获取需求模板字段
- [ ] 更新需求分类
- [ ] 获取回收站下的需求
- [x] 获取需求关联的缺陷
- [ ] 解除需求缺陷关联关系
- [ ] 更新父需求
- [ ] 创建需求与缺陷关联关系
- [ ] 创建需求与测试用例关联关系
- [ ] 获取视图对应的需求列表
- [x] 转换需求ID成列表queryToken
- [ ] 创建需求关联关系

### 缺陷

- [ ] 创建缺陷
- [ ] 复制缺陷
- [ ] 获取缺陷变更历史
- [ ] 获取缺陷变更次数
- [ ] 获取缺陷自定义字段配置
- [x] 获取缺陷
- [ ] 获取缺陷数量
- [ ] 获取缺陷与其它缺陷的所有关联关系
- [ ] 获取缺陷模板列表
- [ ] 获取缺陷模板字段
- [ ] 获取视图对应的缺陷列表
- [ ] 获取缺陷所有字段及候选值
- [ ] 获取缺陷所有字段的中英文
- [x] 更新缺陷
- [ ] 获取回收站下的缺陷
- [ ] 获取缺陷关联的需求ID
- [ ] 转换缺陷ID成列表queryToken
- [ ] 缺陷说明

### 迭代

- [x] 创建迭代
- [ ] 获取迭代自定义字段配置
- [x] 获取迭代
- [x] 获取迭代数量
- [x] 更新迭代
- [ ] 获取迭代变更历史
- [ ] 获取迭代仪表盘自定义卡片内容
- [ ] 修改迭代仪表盘自定义卡片内容
- [ ] 锁定迭代
- [ ] 解锁迭代
- [x] 获取迭代类别列表
- [x] 获取迭代模板列表
- [ ] 获取迭代模板字段配置
- [ ] 获取迭代类别默认模板字段配置

### 任务

- [x] 创建任务
- [x] 获取任务变更历史
- [x] 获取任务变更次数
- [ ] 获取任务自定义字段配置
- [x] 获取任务
- [x] 获取任务数量
- [ ] 更新任务
- [ ] 获取回收站下的任务
- [ ] 获取视图对应的任务列表
- [x] 获取任务字段信息

### 测试

- [ ] 创建测试用例
- [ ] 批量创建测试用例
- [ ] 创建测试用例目录
- [ ] 创建测试计划
- [ ] 分配测试用例
- [ ] 创建测试计划和需求关联关系
- [ ] 创建测试计划和测试用例关联关系
- [ ] 解除测试计划和需求关联关系
- [ ] 解除测试用例关联并移出测试计划
- [ ] 执行测试用例
- [ ] 获取测试用例关联的需求
- [ ] 获取测试用例目录
- [ ] 获取测试用例目录数量
- [ ] 获取测试用例自定义字段配置
- [ ] 获取测试用例字段所有字段及候选值
- [ ] 获取测试用例执行结果
- [ ] 获取测试用例
- [ ] 获取测试用例数量
- [ ] 获取测试计划关联bug
- [ ] 获取测试计划测试结果
- [ ] 获取测试计划执行进度
- [ ] 获取测试计划与测试用例关联关系
- [ ] 获取测试计划
- [ ] 获取测试计划数量
- [ ] 测试用例移出测试计划
- [ ] 更新测试用例
- [ ] 编辑测试计划
- [ ] 获取测试计划所有字段及候选值
- [ ] 获取测试计划关联的需求

### 发布

- [ ] 创建发布计划接口
- [ ] 获取发布评审依据
- [ ] 获取发布计划接口
- [ ] 获取发布计划数量接口
- [ ] 获取发布评审
- [ ] 更新发布计划接口
- [ ] 创建发布评审
- [ ] 创建发布评审依据
- [ ] 获取发布评审数量接口
- [ ] 获取发布评审自定义字段
- [ ] 获取发布评审模板
- [ ] 获取发布评审日志

### 源码

- [ ] 保存Commit提交数据
- [ ] 获取GIT关联提交数据(GitCommit)
- [ ] 获取指定commit关联的业务对象

### Wiki

- [ ] 创建 wiki
- [ ] 获取 wiki
- [ ] 获取 Wiki 数量
- [ ] 更新 wiki
- [ ] 获取wiki drawio数据
- [ ] 获取wiki关注人数据
- [ ] 获取wiki关注人数量
- [ ] 获取wiki可访问范围人员及用户组
- [ ] 获取wiki标签信息
- [ ] 获取wiki标签信息数量
- [ ] 获取wiki附件数量

### 看板

- [ ] 新建看板工作项
- [ ] 获取看板工作项接口
- [ ] 更新看板工作项
- [ ] 获取看板板块

### 评论

- [x] 添加评论接口
- [x] 获取评论
- [x] 获取评论数量
- [x] 更新评论接口

### 报表

- [x] 获取项目报告

### 附件

- [x] 获取附件
- [x] 获取单个附件下载链接
- [x] 获取单个图片下载链接
- [x] 获取单个文档下载链接

### 度量

- [x] 获取状态流转时间

### 工时

- [x] 创建工时花费
- [x] 获取工时花费
- [x] 获取工时花费的数量
- [x] 更新工时花费

### 项目

- [ ] 获取子项目信息
- [ ] 获取项目信息
- [x] 获取指定项目成员
- [ ] 添加项目成员
- [ ] 获取公司项目列表
- [ ] 获取用户组ID对照关系
- [ ] 获取用户参与的项目列表
- [ ] 获取项目成员列表
- [ ] 获取项目自定义字段
- [ ] 更新项目信息
- [ ] 获取项目文档
- [x] 获取成员活动日志 —— ⚠️ 因无权限，暂未测试过，请谨慎使用

### 项目集

- [ ] 根据视图id获取项目集视图工作项列表
- [ ] 项目集批量关联/取消关联、修改授权范围项目
- [ ] 项目集批量关联/取消关联业务对象

### 工作流

- [ ] 获取工作流流转细则
- [ ] 获取工作流结束状态
- [x] 获取所有结束状态
- [ ] 获取工作流状态中英文名对应关系
- [ ] 获取工作流起始状态
- [ ] 获取项目下的工作流列表
- [ ] 获取并行工作节点和状态的对应关系

### 配置

- [ ] 创建自定义字段（需求及缺陷）
- [ ] 更新下拉类型自定义字段候选值
- [ ] 更新需求下拉类型自定义字段候选值
- [ ] 更新缺陷下拉类型自定义字段候选值
- [ ] 更新级联自定义字段侯选值
- [ ] 创建模块接口
- [ ] 创建版本接口
- [ ] 获取模块接口
- [ ] 获取模块数量接口
- [ ] 获取版本接口
- [ ] 获取版本数量接口
- [ ] 更新模块接口
- [ ] 创建基线接口
- [ ] 创建特性接口
- [ ] 复制需求类别接口
- [ ] 复制缺陷配置接口
- [ ] 更新基线接口
- [ ] 更新特性接口
- [ ] 获取特性接口
- [ ] 获取特性数量接口
- [ ] 获取基线接口
- [ ] 获取基线数量接口
- [ ] 更新版本接口
- [ ] 获取项目配置开关

### 标签

- [x] 获取自定义标签
- [x] 获取标签数量
- [x] 创建标签
- [x] 更新标签

### 公共存储

- [ ] 删除数据
- [ ] 查询数据
- [ ] 保存数据
- [ ] 更新数据
- [ ] 条件语法

### webhook

- [x] 需求/任务/缺陷类
  - [x] `story::create`
  - [x] `story::update`
  - [x] `story::delete`
  - [x] `task::create`
  - [x] `task::update`
  - [x] `task::delete`
  - [x] `bug::create`
  - [x] `bug::update`
  - [x] `bug::delete`
- [ ] 评论类：需求/任务/缺陷
  - [x] `story_comment::add`
  - [x] `story_comment::update`
  - [x] `story_comment::delete`
  - [x] `task_comment::add`
  - [x] `task_comment::update`
  - [x] `task_comment::delete`
  - [x] `bug_comment::add`
  - [x] `bug_comment::update`
  - [x] `bug_comment::delete`
- [x] 迭代
  - [x] `iteration::create`
  - [x] `iteration::update`
  - [x] `iteration::delete`

### 用户

- [x] 获取角色ID对照关系

## 轻协作API文档

### 工作项

- [ ] 添加工作项
- [ ] 更新工作项
- [ ] 获取工作项
- [ ] 获取工作项数量
- [ ] 添加分组
- [ ] 更新分组
- [ ] 获取分组
- [ ] 获取分组数量
- [ ] 获取工作项动态
- [ ] 获取工作项动态数量
- [ ] 获取工作项自定义字段配置
- [ ] 获取工作项所有字段的中英文
- [ ] 添加工作项与其他业务对象的关联关系
- [ ] 获取关联需求
- [ ] 获取关联缺陷
- [ ] 解除工作项与其他业务对象的关联关系
- [ ] 获取回收站内的工作项

### 空间

- [ ] 获取空间信息
- [ ] 添加空间成员
- [ ] 获取空间成员列表
- [ ] 新建空间
- [ ] 获取用户所有参与的空间

### 评论

- [x] 添加评论
- [x] 更新评论
- [x] 获取评论
- [x] 获取评论数量

### 附件

- [ ] 附件上传
- [ ] 上传base64图片
- [ ] 获取单个附件下载链接
- [ ] 获取附件

### 应用集成-工蜂 

- [ ] 创建工作项和Git分支关联关系
- [ ] 保存Commit提交数据
- [ ] 关联代码仓库与TAPD空间
- [ ] 解除工作项和Git分支关联
- [ ] 获取工作项和Git分支的关联关系
- [ ] 获取分支关联工作项
- [ ] 获取GIT关联提交数据(GitCommit)
- [ ] 获取代码仓库与TAPD关联空间列表
- [ ] 解除commit与工作项关联关系
- [ ] 获取commit关联的工作项
