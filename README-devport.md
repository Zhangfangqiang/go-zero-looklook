### 开发端口

#### 服务端口

| 服务名称       | API 服务端口(1xxx) | RPC 服务端口(2xxx) | 其他服务端口(3xxx)           |
|------------|----------------|----------------|------------------------|
| order      | 1001           | 2001           | mq-3001                |
| payment    | 1002           | 2002           |                        |
| travel     | 1003           | 2003           |                        |
| usercenter | 1004           | 2004           |                        |
| cms        | 1005           | 2005           |                        |
| mqueue     | -              | -              | job-3002、schedule-3003 |

#### Prometheus 端口

⚠️ 线上容器是相互隔离的，所以线上可以全部设置为相同的端口。本地因为是在同一个容器中进行开发，为了防止端口冲突，需要分别设置端口。

| 服务名称             | Prometheus 端口 |
|------------------|---------------|
| order-api        | 4001          |
| order-rpc        | 4002          |
| order-mq         | 4003          |
| payment-api      | 4004          |
| payment-rpc      | 4005          |
| travel-api       | 4006          |
| travel-rpc       | 4007          |
| usercenter-api   | 4008          |
| usercenter-rpc   | 4009          |
| mqueue-job       | 4010          |
| mqueue-scheduler | 4011          |
| cms-api          | 4012          |
| cms-rpc          | 4013          |
