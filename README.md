# fast-gin是一个基于gin框架和gin-example改造的的可以快速上手的后端开发框架，框架采用MYSQL+GORM

整个框架分为以下几个部分：
- config:
    用于配置一些mysql，server等其他配置的参数和初始化
  
- dao:
    dao层，用于写一些和数据交互相关的方法，主要为mysql的curd
  
- model:
    model层，用于定义mysql的表结构体和方法请求及返回的结构体
  
- router:
    server路由层，用于定义接口的的路由和入口方法
  
- service:
    service层，主要用来写一些业务逻辑
  
- library:
    公共库，包含一些通用的工具方法和中间件，比如这里的md5和jwt

通过这个框架可以几分钟内快速的开发自己的业务。

## 后续TODO：
- 接入redis 
- 接入消息队列
- 支持rpc

"# fast-gin" 
