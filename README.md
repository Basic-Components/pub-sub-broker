# pub-sub-broker

发布订阅的中间人,固定两端端点以供发布端和订阅端连接.通常这种结构的发布端是一个`n->1`结构,而订阅端是一个`1->n`结构.因此前端使用`PUSH-PULL`对而后端使用`PUB-SUB`.

可以配置的参数有:

参数|类型|说明
---|---|---
`-stack_name`|string|默认为`unknown`,描述服务群的名称
`-frontend_url`|string|默认为`tcp://*:5569`,推送端接收的url
`backend_url`|string|默认为`tcp://*:5570`,广播端发送的url
`debug`|bool|默认`false`,log等级是否为debug级
`log_format`|string|默认为`json`,描述log的形式
`log_output`|string|默认为`空字符串`,描述是否将log输出为文件,以及输出文件的位置
`conflate`|bool|默认为`false`,描述是否缓存只保留最近的消息
`receive_hwm`|int|默认为`1000`,描述接收端(前端)缓存的条数限制,`0`代表不限制
`send_hwm`|int|默认为`1000`,描述发送端(后端)缓存的条数限制,`0`代表不限制
