package tool

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	AppName     string         `json:"app_name"`
	AppMode     string         `json:"app_mode"`
	AppHost     string         `json:"app_host"`
	AppPort     string         `json:"app_port"`
	Sms         SmsConfig      `json:"sms"`
	Database    DatabaseConfig `json:"database"`
	RedisConfig RedisConfig    `json:"redis_config"`
}

type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	RegionId     string `json:"region_id"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DbName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSql  bool   `json:"show_sql"`
}

type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     string `json:"port"`
	Password string `json:"password"`
	Db       int    `json:"db"`
}

//var _cfg *Config = nil
//var _cfg Config // 零值为nil
var _cfg *Config // 定义结构体指针类型对象；零值为nil

func GetConfig() *Config {
	return _cfg
}

/*
	这里展示了一个为什么在返回结构体类型的时候比较喜欢用指针类型对象
	个人理解：
		返回一个结构体的指针对象，有助于该结构体的对象的属性内容在其它源文件被更改，
		然后在结构体原来的源文件内的其它地方调用该结构体对象时
		表现的是实时改变得值
*/
//func ParseConfig(path string) (Config, error) {
//	file, err := os.Open(path)
//	if err != nil {
//		panic(err)
//	}
//	defer file.Close()
//	reader := bufio.NewReader(file) // Reader指针对象
//	decoder := json.NewDecoder(reader) // 生成一个Decoder解析器
//	// 注意: 如果遇到函数参数是interface类型，需要传入一个地址值
//	if err := decoder.Decode(&_cfg); err != nil {
//		return Config{}, err // 看似主要这样不同！
//	}
//	return _cfg, nil
//}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)    // Reader指针对象
	decoder := json.NewDecoder(reader) // 生成一个Decoder解析器
	// 注意: 如果遇到函数参数是interface类型，需要传入一个地址值
	if err := decoder.Decode(&_cfg); err != nil {
		return nil, err
	}
	return _cfg, nil
}

func Test() {
	fmt.Println(_cfg)
}
