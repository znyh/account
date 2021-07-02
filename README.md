# 账号数据管理服务

## 项目简介
#### account服务用于管理用户账号相关内容。

1.创建普通账号和创建机器人账号。

2.facebook绑定，line绑定，手机绑定，OPPO绑定管理。

3.冻结，封号和禁言管理。

本服务只用到redis存储数据

------

```
|-- Dockerfile
|-- account-sa.yaml
|-- account.yaml
|-- api
|-- cmd
|-- configs
|   |-- bindreward.json         # 绑定赠送配置文件
|   `-- littleacclimit.txt      # 小号限制配置文件
|-- docker-compose.yml
|-- internal
|   |-- code                    # 服务自定义错误码
|   |   |-- account.ecode.go    # 自定义错误码绑定说明
|   |   |-- account.proto       # 自定义错误码
|   |   `-- errmsg.go           # 错误码的proto文件
|   |-- config                  # 加载自定义配置文件逻辑
|   |   |-- bindreward.go
|   |   |-- config.go
|   |   `-- littleacclimit.go
|   |-- dao
|   |   |-- bind.go             # 绑定的redis处理逻辑
|   |   |-- blank.go            # 封号禁言冻结的redis处理逻辑
|   |   |-- limit.go            # 小号限制逻辑
|   |   |-- mongo.go            # mongodb初始化逻辑
|   |   `-- account.go          # 账号创建获取redis处理逻辑
|   |-- di
|   |-- model                   # 定义的相关redis数据处理结构
|   |   |-- limit.go            # 小号限制rediskey定义
|   |   |-- http.go             # facebook，line的数据结构
|   |   |-- account.go
|   |   |-- bind.go
|   |   |-- blank.go
|   |   |-- model.go
|   |   `-- register.go
|   |-- server
|   `-- service
|       |-- http.go             # facebook和line获取数据逻辑
|       |-- bind.go             # 绑定相关逻辑处理
|       |-- blank.go            # 封号禁言冻结相关逻辑处理
|       |-- query.go            # 获取信息相关逻辑处理
|       |-- register.go         # 注册相关逻辑处理
|       |-- service.go          # 接口所在地方
|       `-- service_test.go     # 接口单元测试
`-- pkg                         # 服务启动逻辑包
    `-- pkg.go
```

