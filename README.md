# go swagger 自动生成文档demo

## 安装依赖支持
```$xslt
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag
```

将$GOPATH/bin 添加到环境变了$PATH
在项目目录下执行：swag init

## 文档语法说明

- Summary: 简单描述API的功能
- Description： API详细描述
- Tags: API 所属分类
- Accept: API 接收参数的格式
- Produce: 输出的数据格式，这里是json格式
- Param: 参数，分为6 个字段，其中第6个字段是可选的，各个字段含义位：
    1) 参数名称
    2) 参数在HTTP请求中的位置,body,path,query
    3) 参数类型 string,int,bool,struct 等
    4）是否必须 true,false
    5) 参数描述
    6) 选项，这里用的是default() 用来指定默认值
- Success: 成功返回数据格式，分为四个字段
    1) HTTP 返回 Code
    2) 返回数据类型
    3) 返回数据模型
    
## 说明

- 路由格式，分为两个字段：
   1) API 路径
   2) HTTP 方法
   
API文档有更新时，需要重新执行swag init 并重新编译代码运行 API Server   

## API 接口访问

```
http://127.0.0.1:12002/swagger/index.html
```
   
