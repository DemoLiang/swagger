# go swagger 自动生成文档demo

## 安装依赖支持
```$xslt
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag
```

## 示例代码注释
```golang
// @Summary 脚本修改接口
// @Tags script
// @Title script gets
// @Version 1.0.0
// @Description 脚本修改接口
// @BasePath /api/v1
// @Host 127.0.0.1:8004
// @Accept json
// @Produce json
// @Param uuid path string true "uuid:要修改的脚本UUID"
// @Param namesapce query string true "namesapce:namesapce"
// @Param script body taskScriptPutForm true "script info"
// @Success 200 {object} models.TaskScript "script info"
// @Failure 400 object object "{"data": null, "message": "post script fail","code":1000111}"
// @Router /api/job-ce/script/{uuid} [put]
func scriptPut(ctx *gin.Context) {
    //TODO 
}

//注意以上uuid 是path的，注意router中的uuid需要用{}括起来
```

## 自动生成命令
将$GOPATH/bin 添加到环境变了$PATH
在项目目录下执行: ```swag init -o src/modules/job/swagger/docs -g job.go  -d src/modules/job/ --parseDependency src/models/ ```
- -o 定义自动生成的swagger文档输出目录
- -g 指定main函数的文件
- -d 指定扫描的目录
- --parseDependency 指定查找依赖的路径

## 文档语法说明

- Summary: 简单描述API的功能
- Description： API详细描述
- Tags: API 所属分类
- Accept: API 接收参数的格式
- Produce: 输出的数据格式，这里是json格式
- Param: 参数，分为6 个字段，其中第6个字段是可选的，各个字段含义位：
    1) 参数名称
    2) 参数在HTTP请求中的位置,body,path,query,header 
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


## 代码引入
- 引入文档包： "XXXX/swagger/docs"
- 加入控制开关和路由配置：
```golang 
	if config.Config.Logger.Level == "DEBUG" {
		docs.SwaggerInfo.Title = "Job Service Auto Generate Api"
		docs.SwaggerInfo.Description = "This is a job server."
		docs.SwaggerInfo.Version = "1.0"
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
```


## API 接口访问

```
http://127.0.0.1:12002/swagger/index.html
```
   
## refer

[https://github.com/swaggo](https://github.com/swaggo)
[https://github.com/swaggo/swag](https://github.com/swaggo/swag)
