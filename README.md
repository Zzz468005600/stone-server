## [stone-android的服务端](https://github.com/Zzz468005600/stone-android)
### 使用到的一些库：
1. [pgx，数据库PostgreSQL的驱动和工具框架](https://github.com/jackc/pgx)
2. [echo，web框架](https://github.com/labstack/echo)
3. [grace，热重启golang服务器](https://github.com/facebookgo/grace)
### 接口
1. 注册接口：   

    POST请求   

    参数：   

    ​     name          用户名     类型：string   

    ​     mobile        手机号     类型：string       

    ​     password   密码         类型：string

    返回结果：

    ​	成功：{"code":0,  "result": {}}

    ​	失败：{"code":999,  "msg":失败原因}



**注：服务器ip和数据库相关信息在配置文件中，没有上传，配置信息相关字段解析在`global`这个package下**

**持续更新**