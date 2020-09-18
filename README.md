#### 运行服务
```
go build -o gin-learn-todo main.go

./gin-learn-todo

```
####  Gin

```
Gin框架 学习
https://www.kancloud.cn/shuangdeyu/gin_book/949445

Gin学习文档
https://learnku.com/docs/gin-gonic/2019/examples-jsonp/6162

```
#### Swagger
```
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/gin-swagger/swaggerFiles

go run --tags=doc main.go
http://127.0.0.1:8888/swagger/index.html
...
在项目的根目录：windows下执行swag.exe init ，linux下执行swag init
参考地址: https://blog.csdn.net/hjxzb/article/details/84899100
go build -tags=doc
```
#### Go-micro
```
```
#### Docker

#### K8s

#### Sentry

### Gorm

####  通过 Omit方法可以设置忽略的字段

### 性能分析工具 gops pprof

### todo list
```
1.grpc grpc-getway done
2.etcd done
3.jwt  done
4.zinkin
5.kong 鉴权,限流,降级
6.gops   done
7.swagger done
8.
```

### 单元测试
```
# grpc client 接入zipKin trace
go test -v -count=1 -run=TestOrder ./client_test.go

```