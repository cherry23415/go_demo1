# go_gorm_demo

## 开发及发布
### 工具
* golang(1.9+)

#### 如何开始
1. 下载并安装[Golang](https://golang.org/dl/)

#### 项目说明

    1）项目路径应该在 本机go的src目录下 {gopath}/src/go_gorm_demo, 
    如果项目路径正确的路径,会导致不让访问第三方包资源
    2）sql目录为示例脚本

#### 启动模板前的准备

1.项目配置文件在config目录下, 本地启动项目使用的是local.json这个配置文件

#### 启动命令:
```
go run main.go -conf config/local
```

#### 启动遇到缺包问题

解决方式
1.通过go get下载缺少的包(部分包需要翻墙) 

```
eg: go get http://github.com/xxxxxxx
```

2.有自己维护过的包放在vendor目录，运行时会先从vendor目录中查找，没找到再从goath中查找

#### proto生成go文件
```
protoc --go_out=. *.proto 
```
    
