# duty-grpc-test

# 这是duty在学习grpc后的小尝试，虽然在教学视频的带领下进行的代码工作，但一字一句也是本人自己敲得，每行代码也是看懂了一二的。

# 这个test旨在运用grpc服务以及gin框架来搭建一个用户登录注册以及查询用户的后端服务，其中的通过MySQL进行存储用户的信息，并且对用户的登录密码进行加密存储，保证了用户信息的安全，

# web部分采用gin框架进行编写，加入了登录时的验证码验证，利用JWT中间件进行用户登录状态的检验。还设置了一个解决前后端跨域问题的中间件，虽然还没有看懂。
