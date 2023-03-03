package config

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Dbhost         string
	Dbport         string
	Dbuser         string
	Dbpass         string
	Dbname         string
	JwtSecret      string
	TemplatePath   string
	MailHost       string
	MailPort       int64
	MailSSL        bool
	MailUser       string
	MailPass       string
	MailAlias      string
	MsClientId     string
	MsClientSecret string
	MsRedirectUri  string
	QqRpcSecret    string
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("读取配置文件失败：%v", err)
	}

	Dbhost = viper.Get("mysql.hostname").(string)
	Dbport = viper.Get("mysql.port").(string)
	Dbuser = viper.Get("mysql.username").(string)
	Dbpass = viper.Get("mysql.password").(string)
	Dbname = viper.Get("mysql.database").(string)
	JwtSecret = viper.Get("jwt.secret").(string)
	TemplatePath = viper.Get("template.path").(string)
	MailHost = viper.Get("mail.host").(string)
	MailPort = viper.Get("mail.port").(int64)
	MailSSL = viper.Get("mail.ssl").(bool)
	MailUser = viper.Get("mail.username").(string)
	MailPass = viper.Get("mail.password").(string)
	MailAlias = viper.Get("mail.alias").(string)
	MsClientId = viper.Get("ms.client_id").(string)
	MsClientSecret = viper.Get("ms.client_secret").(string)
	MsRedirectUri = viper.Get("ms.redirect_uri").(string)
	QqRpcSecret = viper.Get("qq.rpc_secret").(string)
}
