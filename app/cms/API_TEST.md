# CMS API 测试用例

以下是使用curl命令测试CMS API的示例。

## 环境变量设置
```bash
# API服务地址
export API_HOST="http://127.0.0.1:1005"

# 登录获取token（需要先有usercenter服务）
# 假设已获取token
export TOKEN="your_jwt_token_here"
```

## 1. 获取文章列表（无需登录）

### 获取所有文章
```bash
curl -X POST "${API_HOST}/cms/v1/article/list" \
  -H "Content-Type: application/json" \
  -d '{
    "page": 1,
    "pageSize": 10
  }'
```

### 按分类获取文章
```bash
curl -X POST "${API_HOST}/cms/v1/article/list" \
  -H "Content-Type: application/json" \
  -d '{
    "category": "技术",
    "page": 1,
    "pageSize": 10
  }'
```

## 2. 获取文章详情（无需登录）

```bash
curl -X POST "${API_HOST}/cms/v1/article/detail" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1
  }'
```

## 3. 创建文章（需要登录）

```bash
curl -X POST "${API_HOST}/cms/v1/article/create" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${TOKEN}" \
  -d '{
    "title": "Go-Zero微服务实践",
    "content": "本文介绍如何使用go-zero框架开发微服务...",
    "category": "技术",
    "coverImage": "http://example.com/cover.jpg"
  }'
```

## 4. 更新文章（需要登录）

```bash
curl -X POST "${API_HOST}/cms/v1/article/update" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${TOKEN}" \
  -d '{
    "id": 1,
    "title": "Go-Zero微服务最佳实践",
    "content": "更新后的文章内容...",
    "category": "技术"
  }'
```

## 5. 发布文章（需要登录）

```bash
curl -X POST "${API_HOST}/cms/v1/article/publish" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${TOKEN}" \
  -d '{
    "id": 1
  }'
```

## 6. 点赞文章（需要登录）

```bash
curl -X POST "${API_HOST}/cms/v1/article/like" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${TOKEN}" \
  -d '{
    "id": 1
  }'
```

## 7. 删除文章（需要登录）

```bash
curl -X POST "${API_HOST}/cms/v1/article/delete" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${TOKEN}" \
  -d '{
    "id": 1
  }'
```

## 完整测试流程示例

```bash
#!/bin/bash

API_HOST="http://127.0.0.1:1005"

# 假设已经登录获取到token
TOKEN="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."

echo "1. 创建文章"
ARTICLE_ID=$(curl -s -X POST "${API_HOST}/cms/v1/article/create" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${TOKEN}" \
  -d '{
    "title": "测试文章",
    "content": "这是一篇测试文章的内容",
    "category": "测试"
  }' | jq -r '.id')

echo "文章ID: ${ARTICLE_ID}"

echo ""
echo "2. 发布文章"
curl -s -X POST "${API_HOST}/cms/v1/article/publish" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${TOKEN}" \
  -d "{
    \"id\": ${ARTICLE_ID}
  }" | jq

echo ""
echo "3. 查看文章详情"
curl -s -X POST "${API_HOST}/cms/v1/article/detail" \
  -H "Content-Type: application/json" \
  -d "{
    \"id\": ${ARTICLE_ID}
  }" | jq

echo ""
echo "4. 点赞文章"
curl -s -X POST "${API_HOST}/cms/v1/article/like" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer ${TOKEN}" \
  -d "{
    \"id\": ${ARTICLE_ID}
  }" | jq

echo ""
echo "5. 获取文章列表"
curl -s -X POST "${API_HOST}/cms/v1/article/list" \
  -H "Content-Type: application/json" \
  -d '{
    "page": 1,
    "pageSize": 10
  }' | jq
```

## 响应示例

### 成功响应
```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "id": 1
  }
}
```

### 错误响应
```json
{
  "code": 10001,
  "msg": "参数错误"
}
```

## 常见错误码

- `401`: 未授权，token无效或过期
- `403`: 无权限操作
- `404`: 资源不存在
- `500`: 服务器内部错误

## 注意事项

1. 所有需要登录的接口都需要在请求头中携带 `Authorization: Bearer {token}`
2. token需要先通过usercenter服务的登录/注册接口获取
3. 只有文章作者才能更新、发布、删除自己的文章
4. 文章状态: 0-草稿, 1-已发布, 2-已下架
