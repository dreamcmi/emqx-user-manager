# EMQX-User-Manager

一个从现行的管理系统分离重构出的用户管理中间件系统。

本系统用于直接管理EMQX的鉴权数据库，当然说的是外置数据库。

本系统使用以下工具

1、`GO v1.18`

2、`GoFrame v1.16.7`

**运行流程如下**

```shell
---------------获取token-----------------      
POST   /api/cert -----> emq auth --Y--> 返回token
(http basic auth)                --N--> 返回错误原因
---------------添加用户-------------------
POST   /api/user --Y--> 返回OK
                 --N--> 返回失败原因
---------------删除用户-------------------       
DELETE /api/user --Y--> 返回OK
                 --N--> 返回失败原因
```

## 详细说明

1. 访问`api/cert`时需要使用`HTTP BASIC AUTH`，所使用的`username`和`password`来自于`EMQX`(即API秘钥)
2. 访问`api/user`时需要携带header `auth:token`,`token`来自于第一步
3. `token`有效时间为1h
4. 当前版本使用`EMQX:5.0.9` `PGSQL:15.0` `REDIS:6.2.7`
5. 关于数据库建表请参考EMQX官方文档
6. **当前版本删除了DAO(后面可能会更新回来)**
7. ~~**当前版本只做了SHA256的加密,后续版本会支持自定义加密协议**~~
8. **此系统仅适合作为中间件使用，而不是直接开放给终端用户**

## 访问示例

### 1、申请token

```shell
curl --request POST \
  --basic -u usrename:password \
  --url http://127.0.0.1:5555/api/cert
```

成功返回

```json
{
  "code": 0,
  "date": "2022-11-09T00:36:11.8254273+08:00",
  "msg": "OK",
  "token": "2iw7g90co7275tfm1ow100pnlzzqc26a"
}
```

失败返回

http表示EMQX的httpcode返回值

```json
{
  "code": 4001,
  "http": 401,
  "msg": "EMQ Auth Failed"
}
```

### 2、添加用户

```shell
curl --request POST \
  --url http://127.0.0.1:5555/api/user \
  --header '2iw7g90co7275tfm1ow100pnlzzqc26a' \
  --data '{
    "username":"test111111",
    "password":"12346",
    "salt":"kkk"
}'
```

成功返回

```json
{
  "code": 0,
  "msg": "OK"
}
```

失败返回

1. 鉴权失败

```json
{
  "code": 3001,
  "msg": "Auth Failed"
}
```

2. 数据库获取数据错误 请检查配置

```json
{
  "code": 2001,
  "msg": "SQL GET ERROR"
}
```

3. 数据库设置数据错误 请检查配置

```json
{
  "code": 2002,
  "msg": "SQL SET ERROR"
}
```

4. 用户已存在

```json
{
  "code": 2003,
  "msg": "The User is already in the database"
}
```

### 3、删除用户

```shell
curl --request DELETE \
  --url http://127.0.0.1:5555/api/user \
  --header 'auth: 2iw7g90co7275tfm1ow100pnlzzqc26a' \
  --data '{
    "username":"test111111"
}'
```

成功返回

```json
{
  "code": 0,
  "msg": "OK"
}
```

失败返回

1. 鉴权失败

```json
{
  "code": 3001,
  "msg": "Auth Failed"
}
```

2. 数据库获取数据错误 请检查配置

```json
{
  "code": 2001,
  "msg": "SQL GET ERROR"
}
```

3. 数据库设置数据错误 请检查配置

```json
{
  "code": 2002,
  "msg": "SQL SET ERROR"
}
```

## LICENSE

Apache 2.0
