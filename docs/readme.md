# 目录结构

```
➜  power tree -L 2
.
├── app
│     ├── admin           # 管理服务
│     ├── base            # 基础服务(gid,code,token, ...)
│     ├── game            # 游戏服务
│     ├── gateway         # 网关服务 
│     ├── user            # 会员服务
│     ├── merchant        # 商户服务
│     ├── trade           # 交易服务
│     ├── pay             # 支付服务
|     |-- bill            # 账单服务
│     └── wallet          # 钱包服务
|     |-- promotion       #促销活动服务
├── deploy # 部署脚本
├── pb # grpc生成文件
├── go.mod
├── go.sum
├── pkg  # 公共包
│     ├── gorm.go
│     ├── kafka.go
│     └── redis.go
├── readme.md
└── scripts  # sql脚本
    ├── member
    ├── merchant
    ├── order
    ├── pay
    └── account
```

# 生成rpc应用模版(不生成grpc client)

goctl rpc new user -c=false

# 生成api应用模版

goctl api new apigateway

goctl api go -api apigateway.api -dir . --style goZero

# 通过 -m 指定 goctl 生成分组的代码

goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style goZero -m

# 生成服务代码

cd user
goctl rpc protoc user.proto -v -c=false --go_out=./../../pb --go-grpc_out=./../../pb --zrpc_out=. --style goZero

cd merchant

goctl rpc protoc merchant.proto -c=false --go_out=./../../pb --go-grpc_out=./../../pb --zrpc_out=. --style goZero

goctl rpc protoc pay.proto -c=false --go_out=./../../pb --go-grpc_out=./../../pb --zrpc_out=. --style goZero

goctl rpc protoc wallet.proto -m -c=false --go_out=./../../pb --go-grpc_out=./../../pb --zrpc_out=. --style goZero

# 生成Dockerfile

```
    goctl docker --go reports.go --exe reports
```

# 设置go module的私有库地址（在go env中添加，不配置此env，本地编译会报错）

```
    GOPRIVATE='git.168tzi.com'
```

# model 生成

```
gentool -dsn "root:tRyf6U58C1EI@tcp(192.168.86.179:3306)/cash_sys?charset=utf8mb4&parseTime=true" -tables "reports"  -onlyModel
```

# docker build(切记：在.gitignore中将vendor/排除)

```
cd reports/
go mod vendor # 将第三方依赖内化，依赖本地代码库，无法直接通过go mod download 下载依赖库 
docker build -t reports:v1 .
```

go mod replace powergame.com/powergame/pb => ./pb

商户

    支付渠道配置（支付渠道，提款渠道）

    汇率配置（充值汇率，提款汇率）

    手续费配置（充值手续费，提款手续费）

业务下单余额支付 order->pay->wallet

充值 pay->bill-> paygateway->wallet

提款 pay->bill->wallet->paygateway

用户注册

选择注册方式(email,phone)-> 发送pincode->验证pincode -> 生成用户id->获取RSA公钥->设置登录密码->设置支付密码->完善用户基本信息

用户注册成功后，会异步创建钱包

支付订单

如何保证幂等？

一个业务交易单只能有一个支付订单？

todo

引入公共proto？

goctl rpc protoc paygateway.proto --proto_path=./ --proto_path=./../../pkg/third_party_types -m -c=false
--go_out=./../../pb --go-grpc_out=./../../pb --zrpc_out=. --style goZero
