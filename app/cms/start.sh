#!/bin/bash

# CMS服务启动脚本

echo "==================================="
echo "启动CMS内容管理服务"
echo "==================================="

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# 检查MySQL是否运行
echo -e "${YELLOW}检查MySQL连接...${NC}"
if ! mysql -h127.0.0.1 -P33069 -uroot -pPXDN93VRKUm8TeE7 -e "SELECT 1" > /dev/null 2>&1; then
    echo -e "${RED}错误: MySQL连接失败，请先启动MySQL服务${NC}"
    exit 1
fi
echo -e "${GREEN}MySQL连接正常${NC}"

# 检查数据库是否存在
echo -e "${YELLOW}检查数据库looklook_cms...${NC}"
DB_EXISTS=$(mysql -h127.0.0.1 -P33069 -uroot -pPXDN93VRKUm8TeE7 -e "SHOW DATABASES LIKE 'looklook_cms';" | grep looklook_cms)

if [ -z "$DB_EXISTS" ]; then
    echo -e "${YELLOW}数据库不存在，创建数据库和表...${NC}"
    mysql -h127.0.0.1 -P33069 -uroot -pPXDN93VRKUm8TeE7 < deploy/sql/looklook_cms.sql
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}数据库创建成功${NC}"
    else
        echo -e "${RED}数据库创建失败${NC}"
        exit 1
    fi
else
    echo -e "${GREEN}数据库已存在${NC}"
fi

# 检查Redis是否运行
echo -e "${YELLOW}检查Redis连接...${NC}"
if ! redis-cli -h 127.0.0.1 -p 6379 -a G62m50oigInC30sf ping > /dev/null 2>&1; then
    echo -e "${RED}警告: Redis连接失败，服务可能无法正常运行${NC}"
else
    echo -e "${GREEN}Redis连接正常${NC}"
fi

# 进入项目根目录
cd "$(dirname "$0")/../.."

echo ""
echo -e "${YELLOW}==================================="
echo "启动CMS RPC服务"
echo "===================================${NC}"

# 启动RPC服务
cd app/cms/cmd/rpc
nohup go run cms.go -f etc/cms.yaml > cms-rpc.log 2>&1 &
RPC_PID=$!
echo -e "${GREEN}CMS RPC服务已启动 (PID: $RPC_PID)${NC}"
echo "日志文件: app/cms/cmd/rpc/cms-rpc.log"

# 等待RPC服务启动
sleep 3

echo ""
echo -e "${YELLOW}==================================="
echo "启动CMS API服务"
echo "===================================${NC}"

# 启动API服务
cd ../api
nohup go run cms.go -f etc/cms.yaml > cms-api.log 2>&1 &
API_PID=$!
echo -e "${GREEN}CMS API服务已启动 (PID: $API_PID)${NC}"
echo "日志文件: app/cms/cmd/api/cms-api.log"

echo ""
echo -e "${GREEN}==================================="
echo "服务启动完成！"
echo "===================================${NC}"
echo "CMS RPC服务: 127.0.0.1:2005 (PID: $RPC_PID)"
echo "CMS API服务: 127.0.0.1:1005 (PID: $API_PID)"
echo ""
echo "查看日志:"
echo "  RPC: tail -f app/cms/cmd/rpc/cms-rpc.log"
echo "  API: tail -f app/cms/cmd/api/cms-api.log"
echo ""
echo "停止服务:"
echo "  kill $RPC_PID $API_PID"
echo ""
echo "API文档: http://127.0.0.1:1005/cms/v1/"
