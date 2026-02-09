# CMS内容管理模块创建总结

## 📋 项目概述

本CMS内容管理模块完全模仿`usercenter`微服务模块创建，实现了文章的增删改查、发布、点赞等功能。

## 🎯 功能特性

### 核心功能
- ✅ 创建文章（草稿状态）
- ✅ 更新文章
- ✅ 发布文章
- ✅ 删除文章
- ✅ 文章列表（支持分页、分类筛选）
- ✅ 文章详情
- ✅ 文章点赞

### 技术特性
- ✅ JWT身份验证
- ✅ 权限控制（作者验证）
- ✅ 缓存机制（Redis）
- ✅ 事务支持
- ✅ 链路追踪（Jaeger）
- ✅ 服务监控（Prometheus）

## 📊 数据库设计

### Article表字段
| 字段 | 类型 | 说明 |
|------|------|------|
| id | bigint | 文章ID（主键）|
| title | varchar(255) | 文章标题 |
| content | text | 文章内容 |
| publish_time | datetime | 发布时间 |
| category | varchar(64) | 文章分类 |
| like_count | bigint | 点赞数量 |
| author_id | bigint | 作者ID |
| cover_image | varchar(255) | 封面图片 |
| status | tinyint | 状态（0:草稿 1:已发布 2:已下架）|
| create_time | datetime | 创建时间 |
| update_time | datetime | 更新时间 |
| delete_time | datetime | 删除时间 |
| del_state | tinyint | 删除状态 |
| version | bigint | 版本号 |

### 索引设计
- PRIMARY KEY: `id`
- INDEX: `idx_category` (分类查询)
- INDEX: `idx_author_id` (作者查询)
- INDEX: `idx_publish_time` (时间排序)
- INDEX: `idx_status` (状态筛选)

## 🏗️ 项目结构

```
app/cms/
├── cmd/
│   ├── api/                        # API服务
│   │   ├── cms.go                  # 入口文件
│   │   ├── desc/                   # API定义
│   │   │   ├── cms.api
│   │   │   └── article/article.api
│   │   ├── etc/cms.yaml            # 配置文件
│   │   └── internal/
│   │       ├── config/             # 配置结构
│   │       ├── handler/            # HTTP处理器
│   │       ├── logic/article/      # 业务逻辑
│   │       ├── svc/                # 服务上下文
│   │       └── types/              # 类型定义
│   └── rpc/                        # RPC服务
│       ├── cms.go                  # 入口文件
│       ├── pb/cms.proto            # Proto定义
│       ├── etc/cms.yaml            # 配置文件
│       └── internal/
│           ├── config/             # 配置结构
│           ├── logic/              # 业务逻辑
│           ├── server/             # gRPC服务器
│           └── svc/                # 服务上下文
├── model/                          # 数据模型
│   ├── articleModel.go             # 自定义方法
│   ├── articleModel_gen.go         # 自动生成
│   └── vars.go
├── README.md                       # 模块文档
├── API_TEST.md                     # API测试文档
└── start.sh                        # 启动脚本
```

## 🔧 创建步骤

### 1. 数据库设计
创建SQL文件 `deploy/sql/looklook_cms.sql`
- 定义article表结构
- 添加必要的索引
- 设置字段注释

### 2. API定义
在 `app/cms/cmd/api/desc/` 目录下：
- 创建 `cms.api` - 主API文件
- 创建 `article/article.api` - 文章模块API定义
- 定义请求/响应结构
- 区分需要登录和公开的接口

### 3. RPC定义
在 `app/cms/cmd/rpc/pb/` 目录下：
- 创建 `cms.proto` 文件
- 定义RPC服务方法
- 定义消息结构

### 4. 代码生成

#### 生成Model
```bash
cd app/cms/model
goctl model mysql ddl -src ../../../deploy/sql/looklook_cms.sql -dir . -c -style=goZero
```

#### 生成API代码
```bash
cd app/cms/cmd/api/desc
goctl api go -api *.api -dir ../ -style=goZero
```

#### 生成RPC代码
```bash
cd app/cms/cmd/rpc/pb
goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../
sed -i "" 's/,omitempty//g' *.pb.go
```

### 5. Model层实现

#### articleModel.go 自定义方法
- `Trans()` - 事务支持
- `FindPageListByPage()` - 分页查询
- `UpdateLikeCount()` - 更新点赞数
- `FindAllByCategory()` - 按分类查询

#### articleModel_gen.go 修改
- `Insert()` - 添加session参数支持事务
- `Update()` - 添加session参数支持事务

### 6. RPC层实现

实现以下Logic：
- `createArticleLogic.go` - 创建文章
- `updateArticleLogic.go` - 更新文章
- `publishArticleLogic.go` - 发布文章
- `deleteArticleLogic.go` - 删除文章
- `getArticleListLogic.go` - 获取列表
- `getArticleDetailLogic.go` - 获取详情
- `likeArticleLogic.go` - 点赞文章

### 7. API层实现

实现以下Logic：
- `createArticleLogic.go` - 调用RPC创建
- `updateArticleLogic.go` - 调用RPC更新
- `publishArticleLogic.go` - 调用RPC发布
- `deleteArticleLogic.go` - 调用RPC删除
- `articleListLogic.go` - 调用RPC获取列表
- `articleDetailLogic.go` - 调用RPC获取详情
- `likeArticleLogic.go` - 调用RPC点赞

### 8. 配置文件

#### API配置 (cms.yaml)
- 端口: 1005
- JWT配置
- RPC客户端配置
- 监控配置
- 日志配置

#### RPC配置 (cms.yaml)
- 端口: 2005
- 数据库配置
- Redis配置
- 监控配置
- 日志配置

## 💡 关键技术点

### 1. 从Context获取用户ID
```go
userId := ctxdata.GetUidFromCtx(l.ctx)
```

### 2. 对象拷贝
```go
_ = copier.Copy(&article, item)
```

### 3. 错误包装
```go
return nil, errors.Wrapf(err, "create article failed, req: %+v", req)
```

### 4. 事务使用
```go
err := l.svcCtx.ArticleModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
    // 事务操作
    return nil
})
```

### 5. 缓存机制
- 自动缓存文章详情查询
- 更新/删除时自动失效缓存
- 使用Redis作为缓存存储

### 6. 分页查询
```go
articles, total, err := l.svcCtx.ArticleModel.FindPageListByPage(
    l.ctx, page, pageSize, category, status
)
```

## 🚀 启动方式

### 方式1: 使用启动脚本
```bash
cd app/cms
./start.sh
```

### 方式2: 手动启动

启动RPC服务：
```bash
cd app/cms/cmd/rpc
go run cms.go -f etc/cms.yaml
```

启动API服务：
```bash
cd app/cms/cmd/api
go run cms.go -f etc/cms.yaml
```

## 🧪 测试

详见 `API_TEST.md` 文件，包含：
- 完整的curl测试用例
- 各接口的请求响应示例
- 完整的测试流程脚本

## 📝 与usercenter的对比

| 特性 | usercenter | cms |
|------|-----------|-----|
| 表数量 | 2张(user, user_auth) | 1张(article) |
| 事务使用 | 注册时插入两张表 | 单表操作 |
| 缓存 | 用户信息缓存 | 文章详情缓存 |
| 权限 | 基于JWT | 基于JWT+作者验证 |
| 分页 | 无 | 支持分页查询 |
| 特殊功能 | 微信授权 | 点赞计数 |

## ✅ 已实现功能清单

- [x] SQL表设计
- [x] API定义文件
- [x] RPC Proto定义
- [x] Model代码生成及自定义
- [x] RPC逻辑实现
- [x] API逻辑实现
- [x] 配置文件
- [x] 事务支持
- [x] 缓存机制
- [x] 权限控制
- [x] 分页查询
- [x] 启动脚本
- [x] 测试文档
- [x] README文档

## 🎓 学习要点

1. **模块化设计**: API、RPC、Model三层分离
2. **代码生成**: 使用goctl自动生成基础代码
3. **事务处理**: Model层提供Trans方法
4. **错误处理**: 统一使用errors.Wrapf包装
5. **缓存策略**: 查询缓存、更新失效
6. **权限控制**: JWT + 业务权限验证
7. **配置管理**: yaml配置文件
8. **微服务通信**: API调用RPC

## 📚 参考文档

- `doc/chinese/04-用户服务.md` - usercenter服务文档
- `app/usercenter/` - usercenter源码
- go-zero官方文档

## 🔗 服务端口

- CMS API: `1005`
- CMS RPC: `2005`
- Prometheus (API): `4012`
- Prometheus (RPC): `4013`

## 📌 注意事项

1. 需要先启动MySQL和Redis服务
2. 需要先执行SQL脚本创建数据库
3. 需要登录的接口必须携带JWT token
4. 只有作者可以操作自己的文章
5. 文章状态: 0-草稿, 1-已发布, 2-已下架

---

**创建时间**: 2026年1月31日  
**作者**: 模仿usercenter模块创建  
**技术栈**: go-zero + MySQL + Redis + Jaeger + Prometheus
