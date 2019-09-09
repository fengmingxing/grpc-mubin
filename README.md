# grpc-mubin
1、yum install autoconf automake libtool  （centos 系统）

2、protocal buffer安装
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip 
unzip解压后，将可执行文件:/bin/protoc拷贝到/usr/bin/下面
protoc --version查看版本
->  libprotoc 3.9.1

3、安装 golang protobuf
执行命令：vi ~/.bashrc
新增：export PATH=$PATH:$GOPATH/bin
执行source ~/.bashrc
执行
go get -u github.com/golang/protobuf/proto // golang protobuf 库
go get -u github.com/golang/protobuf/protoc-gen-go //protoc --go_out 工具

4、安装 gRPC-go
go get google.golang.org/grpc

5、编写相关文件和代码

5.1、编写 proto 文件，helloworld/helloworld.proto

5.2、使用protoc命令生成相关文件

命令：protoc --go_out=plugins=grpc:. helloworld.proto

结果：

ls

helloworld.pb.go    helloworld.proto

5.3、编写server和client代码

6、将服务端制作成镜像

6.1、先执行go build server.go 编译成可执行文件

6.2、执行docker build -t="grpc-server:latest" . -f Dockerfile
