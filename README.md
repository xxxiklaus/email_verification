### 利用Gin框架实现的新用户进行邮箱注册登录验证与激活等流程

## 主页面
![home](https://user-images.githubusercontent.com/124338898/233028794-f7a76fcd-ef29-450f-8c9b-e211bb28f9f2.png)

## 获取信息
![Apikey](https://user-images.githubusercontent.com/124338898/233029473-4887c065-6658-45c3-93a1-d8b011a14b07.png)

## 发送失败（注册邮箱不存在的情况）
<br>![failed](https://user-images.githubusercontent.com/124338898/233029774-70e92c0e-3398-4a10-a727-72845d8cc4c5.png)

## 发送成功 PC端&手机端
![email pc](https://user-images.githubusercontent.com/124338898/233030075-466a1693-c014-44c8-8d67-5321930876b9.png)

![phone](https://user-images.githubusercontent.com/124338898/233030107-63447e96-b864-4f27-8c1e-44b86519142f.jpg)

## 验证激活成功

![vsucceed](https://user-images.githubusercontent.com/124338898/233030228-86b425c5-1fe6-434b-aaae-b67f7912ba13.png)


## Features

-  数据持久化 MongoDB
-  验证码缓存 Redis
-  邮件发送 SMTP
-  邮件模板 HTML
-  接口文档 Swagger

### Swagger

有些前端不喜欢在电脑上装客户端，swagger会是一个很好的选择

``` go
type PingResponse struct {
	CommonResponse
	Msg string `json:"Message" example:"boom"`
}

// Ping godoc
// @Summary Test if server is alive
// @Schemes
// @Tags Common
// @Produce json
// @Success 200 {object}  models.PingResponse
// @Router /ping [get]
func Ping(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "boom",
    })
}
```

### JWT

其实jwt服务端做很麻烦，但是客户端调用简单。

使用时需要注意的点：

- 除签发时间到期外，没有其他办法让已经生成的JWT失效，除非服务器端换算法。
- JWT不应该存储敏感的信息
- 如果一意孤行的存放敏感信息，请再次加密。
- 最好设置较短的过期时间，防止被盗用后一直有效，降低损失。
- Payload也可以存储一些业务信息，以便减少数据库的压力。

``` go
func InitAuthMiddlewares(controllers controllers.UserControllers) (*jwt.GinJWTMiddleware, error) {
    return jwt.New(&jwt.GinJWTMiddleware{
        IdentityKey:      "id",
        Realm:            "email-verification",
        SigningAlgorithm: "HS256",
        Key:              []byte(config.GetConfig().JwtKey),
        Timeout:          time.Hour * time.Duration(config.GetConfig().JwtAccessAge),
        MaxRefresh:       time.Hour * time.Duration(config.GetConfig().JwtRefreshAge),
        TokenLookup:      "header: Authorization, query: token, cookie: jwt",
        TokenHeadName:    "Bearer",
        TimeFunc:         time.Now,
        Authenticator:    controllers.Login,
        Authorizator:     authorizedFunc,
        PayloadFunc:      payloadHandle,
        LoginResponse:    loginResponse,
        Unauthorized:     unauthorizedFunc,
        IdentityHandler:  identityHandler,
    })
}

...

authMiddleware, err := middles.InitAuthMiddlewares(controllersImpl)
if err != nil { panic(err) }
router.POST("/login", authMiddleware.LoginHandler)
router.GET("/refresh_token", authMiddleware.RefreshHandler)
router.GET("/logout", authMiddleware.LogoutHandler)
router.GET("/user", authMiddleware.MiddlewareFunc(), controllersImpl.Info)
```

### requestID

- 如何将客户端请求与服务端日志关联
- 微服务架构下，访问日志如何查询
- 不同项目交互出现异常，如何做日志关联

``` go
r := gin.New()
r.Use(requestid.New())

r.GET("/ping", func(c *gin.Context) {
   c.String(http.StatusOK, "pong "+fmt.Sprint(time.Now().Unix()))
})

r.Run(":8080")
```



### 不足之处
- 很多邮件客户端对HMTL的全局style支持一般，所以邮件的展示实现不够好
- 暂时只考虑到手机端邮件的情况，记住邮件的layout一定要可响应式的，平板端的UI部分不逞现..

## 参考的相关Guide

### Swagger

[Golang and MongoDB using the official mongo-driver](https://wb.id.au/computer/golang-and-mongodb-using-the-mongo-go-driver/)

[Gin middleware with Swagger 2.0](https://github.com/swaggo/gin-swagger)

[使用swag自动生成Restful API文档](https://razeen.me/posts/go-swagger)

### JWT

[Issue:jwt in swagger not include `Bearer`](https://github.com/swaggo/gin-swagger/issues/90)

[如何在Gin框架中使用JWT实现认证机制](https://juejin.cn/post/7042520107976753165)

[JWT Middleware for Gin Framework](https://github.com/appleboy/gin-jwt)

[Request ID middleware for Gin Framework](https://github.com/gin-contrib/requestid)

### Email

[在线预览和内联email-html-style](https://htmlemail.io/inline/)

[响应式邮件模板](https://github.com/leemunroe/responsive-html-email-template)
