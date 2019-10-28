# golang-socket
提供一个在家可以使用公司内网环境的工具

### 使用方法
service端运行在公司内网

client运行在自己本机

打通这个需要公司路由器做一个端口映射到运行service上面

### todo 
- adsl拨号公网ip变动问题
    - 使用websocket由service发起一个长连接，建立一个全双工通道
    - service发起一个tcp长连接（尝试但是未能成功）
    - 最笨的就是service端实时上报自己的公网ip
- 登录验证
- 数据加密
