package mysql

import "github.com/go-ini/ini"

type DbConfig struct {
	User     string
	Password string
	HostName string
	Port     string
	DbName   string
}

func  NewDbConfig () DbConfig{
	config ,err := ini.Load("./pkg/external/config/config.ini")
	passwordConf ,err := ini.Load("./pkg/external/config/password.ini")
	if err != nil {
		panic(err)
	}
	return DbConfig{
		User : config.Section("db").Key("MYSQL_USER").MustString("root"),
		Password : passwordConf.Section("db").Key("MYSQL_PASSWORD").MustString("hogehoge"),
		HostName : config.Section("db").Key("MYSQL_HOSTNAME").MustString("localhost"),
		Port : config.Section("db").Key("MYSQL_PORT").String(),
		DbName : config.Section("db").Key("MYSQL_DBNAME").String(),
	}
}
