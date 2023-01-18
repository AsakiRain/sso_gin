# SSO_gin

## 0.什么东西啊

```gin``` + ```gorm``` + ```mysql```

```
crypto/bcrypt 
# go get -u golang.org/x/crypto/bcrypt
```
用于密码加盐

```
viper
# go get -u github.com/spf13/viper
```
用于读取配置

## 1.怎么起啊
1. ```go mod tidy```一下
2. 在项目目录创建```config.toml```并填写以下配置
```
[mysql]
hostname = "主机"
port = "端口"
username = "用户名"
password = "密码"
database = "数据库"

[jwt]
secret = "草神小小的，香香的🤤"

[template]
path = "template"

[mail]
host = "smtp服务器地址"
port = 465
ssl = true
username = "邮箱"
password = "密码"
alias = "发件人别名"

[ms]
client_id = "创建应用的id"
client_secret = "密钥不知道为什么暂时用不上"
redirect_uri = "前端回调链接"
```
3. ```go run main.go```小跑一下
4. ```go build sso_gin```编译出执行文件
5. 你可以双击打开，但是```log```颜色无法被```cmd```正常解析，所以请使用```Windows Terminal```

## 2.怎么玩啊
> 去看Apifox，里面有很多成功用例

## 3.好多错误码啊
```
  code: {
    20000: "成功",
    40000: "请求有误",
    40001: "参数错误",
    40002: "参数检查不通过",
    40100: "未授权",
    40101: "token不存在",
    40102: "token校验出错",
    40103: "token已过期",
    42200: "无法处理",
    42201: "用户不存在",
    42202: "用户名已存在",
    42203: "密码错误",
    42204: "邮箱不存在",
    42205: "邮箱已经注册",
    42206: "邮箱验证码错误",
    42207: "操作频繁",
    42208: "参数不存在",
    42209: "参数已存在",
    42210: "参数不匹配",
    50000: "服务器错误",
  },
```