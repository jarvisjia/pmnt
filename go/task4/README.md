使用方法：

1.注册 请求方式：post URL：/user/register 参数类型：form表单 参数示例： { "username":"", "password":"" }

2.登录 请求方式：post URL：/user/login 参数类型：form表单 请求示例： { "username":"", "password":"" }

3.文章列表 请求方式：get token类型：Bearer URL：/post/list

4.查询文章详情 请求方式：get token类型：Bearer URL：/post/detail/:id  参数类型：path

5.发表文章 请求方式：post token类型：Bearer URL：/post/create  参数类型：form表单 请求示例： { "title":"", "content":"" }

6.修改文章 请求方式：post token类型：Bearer URL：/post/edit 参数类型：form表单 请求示例： { "id":"", "title":"", "content":"" }

7.删除文章 请求方式：post token类型：Bearer URL：/post/delete/:id 参数类型：path

8.发表评论 请求方式：post token类型：Bearer URL：/comment/create 参数类型：form表单 请求示例： { "content":"", "postid":"" }

9.查询文章所有评论 请求方式：get token类型：Bearer URL：/comment/list/:postid 参数类型：path

postman文件版本v2.1 postman版本v11.60.4