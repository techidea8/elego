package config

type System struct {
	Env         string `mapstructure:"env" json:"env" yaml:"env"`
	Port        int    `mapstructure:"port" json:"port" yaml:"port"`
	OssType     string `mapstructure:"oss-type" json:"ossType" yaml:"oss-type"`
	Name        string `mapstructure:"name" json:"name" yaml:"name"`
	Logo        string `mapstructure:"logo" json:"logo" yaml:"logo"`
	Dashboard   string `mapstructure:"dashboard" json:"dashboard" yaml:"dashboard"`
	Vodhost     string `mapstructure:"vodhost" json:"vodhost" yaml:"vodhost"`
	Description string `mapstructure:"description" json:"description" yaml:"description"`
	Apikey      string `mapstructure:"apikey" json:"apikey" yaml:"apikey"`
	QQmapKey    string   `mapstructure:"qqmapkey" json:"qqmapkey" yaml:"qqmapkey"`
}
