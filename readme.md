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
```
3. ```go run main.go```小跑一下
4. ```go build sso_gin```编译出执行文件
5. 你可以双击打开，但是```log```颜色无法被```cmd```正常解析，所以请使用```Windows Terminal```

## 2.怎么玩啊
> 另请参阅ApiFox

1. 一个请求打到```localhost:3000/register```

2. ```body```为```json```格式，要包含以下字段
```
{
"username": "你几把谁啊",
"password": "所有可见字符都能做密码",
"nickname": "女特工艾米莉",
"email": "请尝试绕过邮箱正则",
"code": "以后的验证码，现在只要填了就给过"
}
```
3. 然后你会看到响应
```
{
    "code": 200,
    "message": "注册成功"
}
```
4. 你可以试试删字段或者弄点非法输入，正在测试可靠性。