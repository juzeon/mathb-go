package conf

import (
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
	"os"
)

type conf struct {
	RecaptchaSitekey string `yaml:"recaptcha_sitekey"`
	RecaptchaSecret  string `yaml:"recaptcha_secret"`
}

var Conf conf

func init() {
	v := lo.Must(os.ReadFile("config.yml"))
	lo.Must0(yaml.Unmarshal(v, &Conf))
}
