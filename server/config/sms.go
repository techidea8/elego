package config

type Sms struct {
	Tplcode  string `mapstructure:"tplcode" json:"tplcode" yaml:"tplcode" `
	SignName string `mapstructure:"signname" json:"signname" yaml:"signname"`
}
