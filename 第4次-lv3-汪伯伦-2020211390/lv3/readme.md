## 实现功能

- 注册
- 登录
- 上传博客
- 点赞
- 查看博客

## 接口文档

### 注册

请求方式：POST

请求地址：/register

请求参数

|  参数名   | 类型  |  说明   | 
|  ----  | ----  |  ----  | 
| username  | string | 用户名  | 
| password | string | 密码 | 

### 登录

请求方式：GET

请求地址：/login

请求参数

|  参数名   | 类型  |  说明   | 
|  ----  | ----  |  ----  | 
| username  | string | 用户名  | 
| password | string | 密码 | 

### 上传博客

请求方式：POST

请求地址：/blog/upload

请求参数(json)

如:
```json
{
    "title": "test",
    "content": "hello world"
}
```


|  参数名   | 类型  |  说明   | 
|  ----  | ----  |  ----  | 
| title  | string | 标题  | 
| content | string | 内容|

### 点赞

请求方式：PUT

请求地址：/blog/like/id:

请求参数(query)

|  参数名   | 类型  |  说明   | 
|  ----  | ----  |  ----  | 
| id  | int | blog唯一表示符  | 


### 查看博客

请求方式：GET

请求地址：/blog/show/:id

请求参数(query)

|  参数名   | 类型  |  说明   | 
|  ----  | ----  |  ----  | 
| id  | int | blog唯一表示符  | 


返回博客接口

```json
{
    "code": 1005,
    "data": {
        "title": "test",
        "content": "hello world",
        "author": "mx",
        "likes": "2"
    },
    "message": "成功"
}
```