#!/bin/bash

# Local startup script for go-zero-looklook project
# This script helps start services without Docker

set -e

# Color definitions
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}=====================================${NC}"
echo -e "${BLUE}  Go-Zero-Looklook Local Startup${NC}"
echo -e "${BLUE}=====================================${NC}"
echo ""

# Step 1: Check prerequisites
echo -e "${YELLOW}Step 1: Checking prerequisites...${NC}"

# Check Go
if ! command -v go &> /dev/null; then
    echo -e "${RED}✗ Go is not installed${NC}"
    exit 1
fi
echo -e "${GREEN}✓ Go is installed: $(go version)${NC}"

# Check MySQL
MYSQL_HOST="127.0.0.1"
MYSQL_PORT="33069"
MYSQL_USER="root"
MYSQL_PASS="PXDN93VRKUm8TeE7"

echo -e "${YELLOW}Checking MySQL connection (${MYSQL_HOST}:${MYSQL_PORT})...${NC}"
if mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USER} -p${MYSQL_PASS} -e "SELECT 1" > /dev/null 2>&1; then
    echo -e "${GREEN}✓ MySQL is running${NC}"
    MYSQL_AVAILABLE=true
else
    echo -e "${RED}✗ MySQL is not accessible on ${MYSQL_HOST}:${MYSQL_PORT}${NC}"
    echo -e "${YELLOW}  Trying default port 3306...${NC}"
    MYSQL_PORT="3306"
    if mysql -h${MYSQL_HOST} -P${MYSQL_PORT} -u${MYSQL_USER} -p${MYSQL_PASS} -e "SELECT 1" > /dev/null 2>&1; then
        echo -e "${GREEN}✓ MySQL is running on port 3306${NC}"
        MYSQL_AVAILABLE=true
    else
        echo -e "${RED}✗ MySQL is not running${NC}"
        echo -e "${YELLOW}  Please start MySQL first or use Docker Compose${NC}"
        MYSQL_AVAILABLE=false
    fi
fi

# Check Redis
REDIS_HOST="127.0.0.1"
REDIS_PORT="6379"
REDIS_PASS="G62m50oigInC30sf"

echo -e "${YELLOW}Checking Redis connection (${REDIS_HOST}:${REDIS_PORT})...${NC}"
if redis-cli -h ${REDIS_HOST} -p ${REDIS_PORT} -a ${REDIS_PASS} ping > /dev/null 2>&1; then
    echo -e "${GREEN}✓ Redis is running${NC}"
    REDIS_AVAILABLE=true
else
    echo -e "${RED}✗ Redis is not running${NC}"
    echo -e "${YELLOW}  Please start Redis first or use Docker Compose${NC}"
    REDIS_AVAILABLE=false
fi

echo ""

# Step 2: Show available options
echo -e "${BLUE}=====================================${NC}"
echo -e "${BLUE}Available Startup Options:${NC}"
echo -e "${BLUE}=====================================${NC}"
echo ""
echo "1. Start ALL services (requires MySQL + Redis)"
echo "2. Start CMS services only"
echo "3. Start specific service"
echo "4. Show service status"
echo "5. Stop all services"
echo "6. Exit"
echo ""

read -p "Choose an option (1-6): " choice

case $choice in
    1)
        if [ "$MYSQL_AVAILABLE" = false ] || [ "$REDIS_AVAILABLE" = false ]; then
            echo -e "${RED}Cannot start services - MySQL or Redis not available${NC}"
            echo -e "${YELLOW}Please start dependencies first using Docker Compose:${NC}"
            echo -e "  docker-compose -f docker-compose-env.yml up -d"
            exit 1
        fi

        echo -e "${GREEN}Starting all services...${NC}"

        # Update configs if needed
        echo -e "${YELLOW}Updating configuration files...${NC}"

        # Start services using modd (if available)
        if command -v modd &> /dev/null; then
            echo -e "${GREEN}Starting services with modd...${NC}"
            modd
        else
            echo -e "${YELLOW}modd not found. Starting services manually...${NC}"
            bash app/cms/start.sh &
            # Add other service start scripts here
        fi
        ;;

    2)
        echo -e "${GREEN}Starting CMS services...${NC}"
        bash app/cms/start.sh
        ;;

    3)
        echo "Available services:"
        echo "  - usercenter"
        echo "  - travel"
        echo "  - order"
        echo "  - payment"
        echo "  - cms"
        read -p "Enter service name: " service_name

        if [ -f "app/${service_name}/start.sh" ]; then
            bash "app/${service_name}/start.sh"
        else
            echo -e "${RED}Service start script not found${NC}"
        fi
        ;;

    4)
        echo -e "${BLUE}Checking service status...${NC}"
        echo ""

        ports=(1001 1002 1003 1004 1005 2001 2002 2003 2004 2005)
        services=("order-api" "payment-api" "travel-api" "usercenter-api" "cms-api" "order-rpc" "payment-rpc" "travel-rpc" "usercenter-rpc" "cms-rpc")

        for i in "${!ports[@]}"; do
            port="${ports[$i]}"
            service="${services[$i]}"
            if lsof -i :$port > /dev/null 2>&1; then
                echo -e "${GREEN}✓ ${service} (port ${port}) - RUNNING${NC}"
            else
                echo -e "${RED}✗ ${service} (port ${port}) - NOT RUNNING${NC}"
            fi
        done

        echo ""
        echo -e "${YELLOW}Nginx gateway (port 8888):${NC}"
        if lsof -i :8888 > /dev/null 2>&1; then
            echo -e "${GREEN}✓ RUNNING${NC}"
        else
            echo -e "${RED}✗ NOT RUNNING${NC}"
        fi
        ;;

    5)
        echo -e "${YELLOW}Stopping all services...${NC}"

        # Kill processes by port
        for port in 1001 1002 1003 1004 1005 2001 2002 2003 2004 2005; do
            pid=$(lsof -t -i:$port 2>/dev/null || echo "")
            if [ ! -z "$pid" ]; then
                kill $pid 2>/dev/null && echo -e "${GREEN}✓ Stopped service on port ${port}${NC}"
            fi
        done

        # Kill nginx if running locally
        pid=$(lsof -t -i:8888 2>/dev/null || echo "")
        if [ ! -z "$pid" ]; then
            kill $pid 2>/dev/null && echo -e "${GREEN}✓ Stopped nginx on port 8888${NC}"
        fi

        echo -e "${GREEN}All services stopped${NC}"
        ;;

    6)
        echo "Exiting..."
        exit 0
        ;;

    *)
        echo -e "${RED}Invalid option${NC}"
        exit 1
        ;;
esac

echo ""
echo -e "${BLUE}=====================================${NC}"
echo -e "${GREEN}Done!${NC}"
echo -e "${BLUE}=====================================${NC}"
