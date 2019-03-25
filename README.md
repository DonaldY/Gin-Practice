> 使用 `golang` 开发web应用

### 框架选择

1. web 框架：gin


### 包结构

golang 按包来分层

同一包下的方法均能调用

```
├─common     # 公共参数
├─config     # 配置文件
├─middleware # 中间件
├─model      # 数据库操作/结构体定义
├─handler    # 处理层
├─router     # 路由
├─test       # 测试
├─utils      # 工具
└─main.go    # 启动
```


