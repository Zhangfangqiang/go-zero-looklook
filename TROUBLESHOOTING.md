# Troubleshooting 502 Bad Gateway Error

## Problem
```
curl -X POST "http://127.0.0.1:8888/cms/v1/article/list" \
  -H "Content-Type: application/json" \
  -d '{"page": 1, "pageSize": 10}'
```
Returns: **502 Bad Gateway**

## Root Cause
The error occurs because **no services are running**:
- ✗ Nginx gateway is not listening on port 8888
- ✗ CMS API service is not running on port 1005
- ✗ CMS RPC service is not running on port 2005

## Quick Fix

### Option A: Using Docker Compose (Recommended)

1. **Start environment dependencies:**
   ```bash
   cd /Users/a111/Documents/code/go/go-zero-looklook
   docker-compose -f docker-compose-env.yml up -d
   ```

2. **Start application services:**
   ```bash
   docker-compose up -d
   ```

3. **Verify services:**
   ```bash
   docker ps
   ```

4. **Test API:**
   ```bash
   curl -X POST "http://127.0.0.1:8888/cms/v1/article/list" \
     -H "Content-Type: application/json" \
     -d '{"page": 1, "pageSize": 10}'
   ```

### Option B: Using Local Startup Script

1. **Run the startup script:**
   ```bash
   cd /Users/a111/Documents/code/go/go-zero-looklook
   ./start-local.sh
   ```

2. **Choose option 4 to check service status**

3. **If MySQL/Redis are not running, you need to start them first**

### Option C: Manual Service Startup

1. **Start MySQL and Redis** (if not already running via Docker):
   ```bash
   # Using Docker for dependencies only
   docker-compose -f docker-compose-env.yml up -d mysql redis
   ```

2. **Update CMS RPC configuration** to use localhost:
   ```bash
   # Edit app/cms/cmd/rpc/etc/cms.yaml
   # Change:
   #   redis:6379 → 127.0.0.1:6379
   #   mysql:3306 → 127.0.0.1:3306 (or 33069)
   ```

3. **Update CMS API configuration**:
   ```bash
   # Edit app/cms/cmd/api/etc/cms.yaml
   # Ensure CmsRpcConf endpoints point to 127.0.0.1:2005
   ```

4. **Start CMS services:**
   ```bash
   bash app/cms/start.sh
   ```

5. **Start nginx gateway locally:**
   ```bash
   # Install nginx if needed
   brew install nginx
   
   # Copy and modify config
   cp deploy/nginx/conf.d/looklook-gateway.conf /tmp/looklook-gateway.conf
   
   # Edit /tmp/looklook-gateway.conf:
   # Change: listen 8081 → listen 8888
   # Change: proxy_pass http://looklook:1005 → proxy_pass http://127.0.0.1:1005
   
   # Start nginx with custom config
   nginx -c /tmp/nginx.conf
   ```

### Option D: Direct API Testing (Bypass Nginx)

Test the CMS API directly without nginx:

```bash
# Make sure CMS services are running first
bash app/cms/start.sh

# Test directly on port 1005
curl -X POST "http://127.0.0.1:1005/cms/v1/article/list" \
  -H "Content-Type: application/json" \
  -d '{"page": 1, "pageSize": 10}'
```

## Checking Service Status

### Check what's running:
```bash
# Check if nginx is running
lsof -i :8888

# Check if CMS API is running
lsof -i :1005

# Check if CMS RPC is running
lsof -i :2005

# Check all services
./start-local.sh  # Then choose option 4
```

### Check logs:
```bash
# CMS API logs
tail -f app/cms/cmd/api/cms-api.log

# CMS RPC logs
tail -f app/cms/cmd/rpc/cms-rpc.log

# Nginx logs (if using Docker)
docker logs nginx-gateway
```

## Common Issues

### Issue 1: "connect: connection refused" in logs
**Cause:** Backend service (CMS API) is not running
**Solution:** Start the CMS services using `bash app/cms/start.sh`

### Issue 2: Services start but immediately exit
**Cause:** MySQL or Redis not accessible
**Solution:** 
- Check if MySQL is running: `mysql -h127.0.0.1 -P33069 -uroot -pPXDN93VRKUm8TeE7 -e "SELECT 1"`
- Check if Redis is running: `redis-cli -h 127.0.0.1 -p 6379 -a G62m50oigInC30sf ping`
- Start dependencies: `docker-compose -f docker-compose-env.yml up -d mysql redis`

### Issue 3: "no such host" errors in logs
**Cause:** Configuration files use Docker hostnames (mysql, redis) but running locally
**Solution:** Update config files to use 127.0.0.1 instead of Docker hostnames

### Issue 4: Port already in use
**Cause:** Another service is using the port
**Solution:** 
```bash
# Find what's using the port
lsof -i :1005

# Kill the process
kill <PID>
```

## Architecture Overview

```
Client (curl)
    ↓
Port 8888 (nginx-gateway)
    ↓
Port 1005 (cms-api)
    ↓
Port 2005 (cms-rpc)
    ↓
Port 3306 (MySQL) + Port 6379 (Redis)
```

All parts of this chain must be running for the API to work.

## Next Steps

1. **First time setup?** Follow: `doc/chinese/01-开发环境搭建.md`
2. **Understanding the gateway?** Read: `doc/chinese/02-nginx网关.md`
3. **Full documentation:** Check the `doc/` directory

## Quick Commands Reference

```bash
# Start everything with Docker
docker-compose -f docker-compose-env.yml up -d && docker-compose up -d

# Check Docker services
docker ps

# Stop everything
docker-compose down && docker-compose -f docker-compose-env.yml down

# Start CMS only (local)
bash app/cms/start.sh

# Check service status
./start-local.sh  # option 4

# Stop all local services
./start-local.sh  # option 5

# View logs
tail -f app/cms/cmd/api/cms-api.log
tail -f app/cms/cmd/rpc/cms-rpc.log
```

## Still Having Issues?

1. Check if Docker is installed: `docker --version`
2. Check if Go is installed: `go version`
3. Check project README: `README-cn.md`
4. Review the service-specific README: `app/cms/README.md`
