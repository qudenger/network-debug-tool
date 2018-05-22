## tcp网络调试助手 web版

此程序为方便嵌入式设备调试TCP协议，部署在linux服务器上，通过web页面即可进行tcp消息的接收和发送。

![截图](https://raw.githubusercontent.com/qudenger/network-debug-tool/master/screen-print.jpg)

#### 编译与部署：

以linux服务为例：

```
cd tcp
CGO_ENABLED=0 GOOS=linux go build -o tcpsrv ./*.go
nohup ./tcpsrv &
```

```
cd web
CGO_ENABLED=0 GOOS=linux go build -o tcpweb ./*.go
nohup ./tcpweb &
```

#### 使用:

设备连接tcp server地址：serverip:12345

web访问地址： http://serverip:1234

#### 使用docker部署:

