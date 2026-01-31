# CMS 内容管理模块

## 概述

CMS内容管理模块是模仿usercenter微服务模块创建的，用于管理文章内容。

## 功能特性

### 文章管理
- **创建文章**: 用户可以创建新文章（草稿状态）
- **更新文章**: 作者可以更新自己的文章
- **发布文章**: 作者可以将草稿发布为正式文章
- **删除文章**: 作者可以删除自己的文章
- **文章列表**: 支持分页查询、分类筛选
- **文章详情**: 查看文章完整信息
- **点赞功能**: 用户可以给文章点赞

## 数据库表结构

### article表
```sql
CREATE TABLE `article` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `delete_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
  `del_state` tinyint NOT NULL DEFAULT '0' COMMENT '删除状态 0:未删除 1:已删除',
  `version` bigint NOT NULL DEFAULT '0' COMMENT '版本号',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '文章标题',
  `content` text NOT NULL COMMENT '文章内容',
  `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '文章发布时间',
  `category` varchar(64) NOT NULL DEFAULT '' COMMENT '文章分类',
  `like_count` bigint NOT NULL DEFAULT '0' COMMENT '文章点赞数量',
  `author_id` bigint NOT NULL DEFAULT '0' COMMENT '作者id',
  `cover_image` varchar(255) NOT NULL DEFAULT '' COMMENT '封面图片',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '状态 0:草稿 1:已发布 2:已下架',
  PRIMARY KEY (`id`),
  KEY `idx_category` (`category`),
  KEY `idx_author_id` (`author_id`),
  KEY `idx_publish_time` (`publish_time`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章表';
```

## 目录结构

```
app/cms/
├── cmd/
│   ├── api/                    # API服务
│   │   ├── desc/              # API定义文件
│   │   │   ├── cms.api        # 主API定义
│   │   │   └── article/       # 文章模块API定义
│   │   │       └── article.api
│   │   ├── etc/               # 配置文件
│   │   │   └── cms.yaml
│   │   └── internal/
│   │       ├── config/        # 配置结构
│   │       ├── handler/       # HTTP处理器
│   │       ├── logic/         # 业务逻辑
│   │       │   └── article/
│   │       ├── svc/           # 服务上下文
│   │       └── types/         # 类型定义
│   └── rpc/                   # RPC服务
│       ├── pb/                # Protobuf定义
│       │   └── cms.proto
│       ├── etc/               # 配置文件
│       │   └── cms.yaml
│       └── internal/
│           ├── config/        # 配置结构
│           ├── logic/         # 业务逻辑
│           ├── server/        # gRPC服务器
│           └── svc/           # 服务上下文
└── model/                     # 数据模型
    ├── articleModel.go        # 文章模型（自定义方法）
    ├── articleModel_gen.go    # 文章模型（自动生成）
    └── vars.go                # 变量定义
```

## API接口

### 公开接口（无需登录）

#### 1. 获取文章列表
- **接口**: `POST /cms/v1/article/list`
- **请求体**:
```json
{
  "category": "技术",
  "page": 1,
  "pageSize": 10
}
```
- **响应**:
```json
{
  "list": [
    {
      "id": 1,
      "title": "文章标题",
      "content": "文章内容",
      "publishTime": "2026-01-31 10:00:00",
      "category": "技术",
      "likeCount": 100,
      "authorId": 1,
      "coverImage": "http://example.com/cover.jpg",
      "status": 1
    }
  ],
  "total": 100
}
```

#### 2. 获取文章详情
- **接口**: `POST /cms/v1/article/detail`
- **请求体**:
```json
{
  "id": 1
}
```

### 需要登录的接口

#### 3. 创建文章
- **接口**: `POST /cms/v1/article/create`
- **请求头**: `Authorization: Bearer {token}`
- **请求体**:
```json
{
  "title": "文章标题",
  "content": "文章内容",
  "category": "技术",
  "coverImage": "http://example.com/cover.jpg"
}
```
- **响应**:
```json
{
  "id": 1
}
```

#### 4. 更新文章
- **接口**: `POST /cms/v1/article/update`
- **请求头**: `Authorization: Bearer {token}`
- **请求体**:
```json
{
  "id": 1,
  "title": "新标题",
  "content": "新内容",
  "category": "新分类"
}
```

#### 5. 发布文章
- **接口**: `POST /cms/v1/article/publish`
- **请求头**: `Authorization: Bearer {token}`
- **请求体**:
```json
{
  "id": 1
}
```

#### 6. 删除文章
- **接口**: `POST /cms/v1/article/delete`
- **请求头**: `Authorization: Bearer {token}`
- **请求体**:
```json
{
  "id": 1
}
```

#### 7. 点赞文章
- **接口**: `POST /cms/v1/article/like`
- **请求头**: `Authorization: Bearer {token}`
- **请求体**:
```json
{
  "id": 1
}
```
- **响应**:
```json
{
  "likeCount": 101
}
```

## 部署配置

### API服务配置 (cms.yaml)
```yaml
Name: cms-api
Host: 0.0.0.0
Port: 1005
Mode: dev

JwtAuth:
  AccessSecret: ae0536f9-6450-4606-8e13-5a19ed505da0

Prometheus:
  Host: 0.0.0.0
  Port: 4010
  Path: /metrics

Telemetry:
  Name: cms-api
  Endpoint: http://jaeger:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger

Log:
  ServiceName: cms-api
  Level: error

CmsRpcConf:
  Endpoints:
    - 127.0.0.1:2005
  NonBlock: true
```

### RPC服务配置 (cms.yaml)
```yaml
Name: cms-rpc
ListenOn: 0.0.0.0:2005
Mode: dev

Prometheus:
  Host: 0.0.0.0
  Port: 4011
  Path: /metrics

Telemetry:
  Name: cms-rpc
  Endpoint: http://jaeger:14268/api/traces
  Sampler: 1.0
  Batcher: jaeger

Log:
  ServiceName: cms-rpc
  Level: error

Redis:
  Host: redis:6379
  Type: node
  Pass: G62m50oigInC30sf
  Key: cms-rpc

DB:
  DataSource: root:PXDN93VRKUm8TeE7@tcp(mysql:3306)/looklook_cms?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai

Cache:
  - Host: redis:6379
    Pass: G62m50oigInC30sf
```

## 启动服务

### 1. 初始化数据库
```bash
# 执行SQL脚本创建数据库和表
mysql -h127.0.0.1 -P33069 -uroot -p < deploy/sql/looklook_cms.sql
```

### 2. 启动RPC服务
```bash
cd app/cms/cmd/rpc
go run cms.go -f etc/cms.yaml
```

### 3. 启动API服务
```bash
cd app/cms/cmd/api
go run cms.go -f etc/cms.yaml
```

## 代码生成命令

### 生成API代码
```bash
cd app/cms/cmd/api/desc
goctl api go -api *.api -dir ../ -style=goZero
```

### 生成RPC代码
```bash
cd app/cms/cmd/rpc/pb
goctl rpc protoc *.proto --go_out=../ --go-grpc_out=../ --zrpc_out=../
sed -i "" 's/,omitempty//g' *.pb.go
```

### 生成Model代码
```bash
cd app/cms/model
goctl model mysql ddl -src ../../../deploy/sql/looklook_cms.sql -dir . -c -style=goZero
```

## 技术栈

- **框架**: go-zero
- **数据库**: MySQL 8.0
- **缓存**: Redis
- **链路追踪**: Jaeger
- **监控**: Prometheus + Grafana
- **ORM**: sqlx + cache

## 特性说明

### 1. 缓存策略
- 使用go-zero内置的cache机制
- 文章详情查询使用缓存
- 自动缓存失效机制

### 2. 事务支持
- Model层提供Trans方法支持事务操作
- Insert和Update方法支持传入session参数

### 3. 分页查询
- 支持按分类筛选
- 支持按状态筛选
- 默认按发布时间倒序排列

### 4. 权限控制
- JWT token验证
- 文章操作权限验证（仅作者可操作）

## 开发参考

本模块完全模仿usercenter模块的开发模式：
- API层调用RPC层
- RPC层处理业务逻辑
- Model层处理数据库操作
- 使用copier进行对象拷贝
- 使用errors.Wrapf包装错误信息
- 从context中获取用户ID: `ctxdata.GetUidFromCtx(ctx)`
