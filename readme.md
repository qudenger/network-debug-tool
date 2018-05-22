### tcp/udp网络调试助手 web版
为解决嵌入式设备调试TCP协议无法连接内网TCP Server，此程序可部署在服务器上，通过web页面进行tcp消息的接收和发送。

#### 编译：
CGO_ENABLED=0 GOOS=linux go build -o tcpsrv ./*.go
CGO_ENABLED=0 GOOS=linux go build -o tcpweb ./*.go

#### 部署：
目录结构
tcp 
  tcpsrv
web
  tcpweb
public
  assets
  index.html

cd tcp
nohup ./tcpsrv &
cd web
nohup ./tcpweb &

#### 使用:
设备连接tcp server地址：serverip:12345
web访问地址： http://serverip:1234

#### docker:
