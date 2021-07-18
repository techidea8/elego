package config

type Local struct {
	Path     string `mapstructure:"path" json:"path" yaml:"path" `
	Location string `mapstructure:"location" json:"location" yaml:"location"`
}

type Qiniu struct {
	Zone          string `mapstructure:"zone" json:"zone" yaml:"zone"`
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Location      string `mapstructure:"location" json:"location" yaml:"location"`
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"`
}

type Aliyun struct {
	Endpoint      string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	Bucket        string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
	Location      string `mapstructure:"location" json:"location" yaml:"location"`
	UseHTTPS      bool   `mapstructure:"use-https" json:"useHttps" yaml:"use-https"`
	AccessKey     string `mapstructure:"access-key" json:"accessKey" yaml:"access-key"`
	SecretKey     string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	UseCdnDomains bool   `mapstructure:"use-cdn-domains" json:"useCdnDomains" yaml:"use-cdn-domains"`
}
