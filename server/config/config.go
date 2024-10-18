package config

// Config 系统全局配置
type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	Jwt    Jwt    `mapstructure:"jwt"`
	File   File   `mapstructure:"file"`
	Mail   Mail   `mapstructure:"mail"`
	Alipay Alipay `mapstructure:"alipay"`
}

// Server 服务器配置
type Server struct {
	Port   int    `mapstructure:"port"`
	RunEnv string `mapstructure:"run_env"`
}

// Mysql 数据库配置
type Mysql struct {
	DNS             string `mapstructure:"dns"`
	MinIdleConn     int    `mapstructre:"MinIdleConn"`
	MinOpenConn     int    `mapstructre:"MinOpenConn"`
	ConnMaxLifetime int    `mapstructure:"ConnMaxLifetime"`
	DBFile          string `mapstructure:"dbFile"`
}

// Redis redis配置
type Redis struct {
	Host        string `mapstructure:"host"`
	Port        int    `mapstructure:"port"`
	Password    string `mapstructure:"password"`
	DB          int    `mapstructure:"db"`
	PoolSize    int    `mapstructure:"poolSize"`
	MinIdleConn int    `mapstructure:"minIdleConn"`
}

// Jwt 用户认证配置
type Jwt struct {
	SigningKey  string `mapstructure:"signingKey"`
	ExpiredTime int    `mapstucture:"expiredTime"`
}

// File 文件存储配置
type File struct {
	Path string `mapstructure:"path"`
}

// Mail 邮件配置
type Mail struct {
	Smtp   string `mapstructure:"smtp"`
	Secret string `mapstructure:"secret"`
	Sender string `mapstructure:"sender"`
}

// 支付宝支付服务配置
type Alipay struct {
	AppId            string `mapstructure:"appId"`
	PrivateKey       string `mapstructure:"privateKey"`
	AppPublicCert    string `mapstructure:"appPublicCert"`
	AlipayRootCert   string `mapstructure:"alipayRootCert"`
	AlipayPublicCert string `mapstructure:"alipayPublicCert"`
	ReturnURL        string `mapstructure:"returnURL"`
	NotifyURL        string `mapstructure:"notifyURL"`
}
