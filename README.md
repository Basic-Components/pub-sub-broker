# pub-sub-broker

发布订阅的中间人,固定两端端点以供发布端和订阅端连接.通常这种结构的发布端是一个`n->1`结构,而订阅端是一个`1->n`结构.因此前端使用`PUSH-PULL`对而后端使用`PUB-SUB`.

## 使用方法

使用命令行`./pub-sub-broker`启动组件,下面是可选的参数:

| 标志            | 类型   | 默认值         | 说明                                         |
| --------------- | ------ | -------------- | -------------------------------------------- |
| `-help`         | bool   | `false`        | 帮助命令                                     |
| `-debug`        | bool   | `false`        | 是否使用debug模式启动                        |
| `-stack_name`   | string | `unknown`      | 服务堆名                                   |
| `-frontend_url` | string | `tcp://*:5569` | 前端连接的地址                               |
| `-backend_url`  | string | `tcp://*:5570` | 后端绑定的地址                               |
| `-log_format`   | string | `json`         | 设定log的形式                                |
| `-log_output`   | string | `空字符串`     | 设定log输出的流位置                          |
| `-config_path`  | string | `空字符串`     | 设定读取配置文件地址                         |
| `-conflate`     | bool   | `false`        | 描述是否缓存只保留最近的消息                 |
| `-receive_hwm`  | int    | `1000`         | 描述接收端(前端)缓存的条数限制,`0`代表不限制 |
| `-send_hwm`     | int    | `1000`         | 描述发送端(后端)缓存的条数限制,`0`代表不限制 |

启动的时候按需求填入参数.

配置文件为json格式,以下为默认配置的配置文件形式:

```json
{
	"stack_name":"unknown",
	"frontend_url":"tcp://*:5569",
	"backend_url":"tcp://*:5570",
	"debug":false,
	"log_format":"json",
	"log_output":"",
	"conflate":false,
	"receive_hwm":1000,
	"send_hwm":1000
}
```

**注意,`-1`将跳过设置,因此会使用zmq的默认值**

配置的优先级为: `命令行参数>配置文件>默认`

例子可以看`example`文件夹下的例子

## 通过docker使用

镜像为:`hsz1273327/pub-sub-broker`,一个可以参考的使用方式是执行:`docker run -p 5559:5559 -p 5560:5560  hsz1273327/pub-sub-broker ./pub-sub-broker -debug`

但通常这个组件是一个服务群的对外端点,使用`docer-compose.yml`进行编排,细节不表,这边给出一个参考配置文件:

```yml
version: '3'
services:

  # ############################################代理
  broker:
    image: hsz1273327/pub-sub-broker:latest
    networks:
      - out 
      - server-group
    command: ./pub-sub-broker

 # ############################################实际的服务
  pub1:
    image: xxx:latest
    networks:
      - server-group
    command: python pub.py

  pub2:
    image: xxx:latest
    networks:
      - server-group
    command: ./pub

  sub1:
    image: xxx:latest
    networks:
      - server-group
    command: python sub.py

  sub2:
    image: xxx:latest
    networks:
      - server-group
    command: ./sub

# ############################################配置网络
networks:
  out:
    external: true
  server-group:
    external: true
```
