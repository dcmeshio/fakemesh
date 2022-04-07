□ 待测试
  □ Get 请求的 Param 测试；
  □ 对待使用请求进行服务器应答测试；
    □ SpringBoot 服务器
    □ Nginx 服务器

■ 请求应答模式 [A终端、B浏览器]

■ A RequestFirstLine
  ■ 四个路径随时切换、GET、HTTP1.1
  ■ Get 请求参数切换，加入随机数参数，不通的路径使用不通的参数名；
  ■ 样例：GET /outer/checkQuote?quotecode=1111&namespacecode=2222 HTTP/1.1
■ A RequestHeader
  ■ User-Agent: 移动端：使用 dort 默认 [获取方式：发送请求]，电脑端：使用 默认 [获取方式：发送请求]；
  ■ X-Token: 参考老方式；
  ■ Accept-Encoding: gzip, deflate [固定内容]
  ■ Host: 实际机器的 域名:端口号、如：idimesh:helmsnets.com:5855
  ■ Connection: keep-alive [固定内容]
  ■ Pragma: no-cache
■ A ResponseHeader
  ■ Content-Type: application/octet-stream | video/mpeg | video/mpeg4 | audio/wav [获取方式：查找视频、音频常用格式，并添加到列表中来]
  ■ Transfer-Encoding: chunked
  ■ Date: 特定格式的时间
  ■ Connection: keep-alive
  ■ Server: nginx/1.20.0 [需要跟配套版本一致]
  
■ B RequestFirstLine
  ■ 六个路径随时切换、GET、HTTP1.1
  ■ Get 请求参数切换，加入随机数参数，不通的路径使用不通的参数名；
■ B RequestHeader
  ■ User-Agent: 收集N多的浏览器的 Agent 20个以上、动态切换、并隔一段时间更新一次内容，尽量有差异化[Chrome/Firefox/Safri/IE]；
  ■ X-Token: 改变名称，并将内容使用 AES 对称加密后传输；
  ■ Accept-Encoding: gzip, deflate, br [固定内容]
  ■ Host: 实际机器的 域名:端口号、如：idimesh:helmsnets.com:5855
  ■ Connection: keep-alive [固定内容]
  ■ Accept: 各个浏览器自带的
■ B ResponseHeader
  ■ Conetent-Type: application/octet-stream | video/mpeg | video/mpeg4 | audio/wav [获取方式：查找视频、音频常用格式，并添加到列表中来]
  ■ Transfer-Encoding: chunked
  ■ Date: 特定格式的时间
  ■ Connection: keep-alive
  ■ Server: nginx/1.20.0 [需要跟配套版本一致]


■ 应答的方式 -> 查看 Nginx 的返回内容
  ■ Conetent-Type: octet-stream
  ■ Transfer-Encoding: chunked
  ■ Date: 特定格式的时间
  ■ Connection: keep-alive
  ■ Server: nginx/1.20.0 [需要跟配套版本一致]

■ 设计需求，能动态通过配置调整：(key、value)；
■ 设计需求，呈现组织出现的问题；如：chrome 的 agent 和 Accept、firefox 的 agent 和 Accept；
■ 设计需求，有的内容需要每次都改变：如 Host 为：idimesh.helmsnets.com；
■ 防止重放，最好用时间戳，从中央服务器获取时间戳，保证服务器和代理服务器的时间一致，将时间戳放入 token 中，时间超过 60 秒，放弃；
■ 作为一个独立项目，使用 Git module 模式引入到项目中；

■ 路径
  ■ 统一构建方式；
  ■ 完成 sample 工程；
    ■ go -> go；
    ■ rust -> go；
  ■ 项目改造 go -> go；
  ■ 项目改造 rust -> go；



□ 待完成周边
  □ 服务器获取时间戳
  □ 所有服务器进行统一的始终校准
  
  