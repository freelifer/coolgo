# 接口说明

### 登录

##### 请求路径

GET /wx/login

##### 请求参数

|参数名称|必须|类型|描述|默认|
|:-:|:-:|:-:|:-:|:-:|:-:|
|code|是|string|微信登录code||

##### 响应参数
|参数名称|类型|描述|默认|
|:-:|:-:|:-:|:-:|:-:|
|sessionid|string|登录成功的sessionKey||

|——main.go         入口文件
|——conf            配置文件和处理模块
|——controllers     控制器入口
|——models          数据库处理模块
|——utils           辅助函数库
|——static          静态文件目录
|——views           视图库