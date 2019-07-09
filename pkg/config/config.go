package config

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Auth struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type NumPolicy struct {
	Num int `yaml:"number"`
}

type Policy struct {
	// IncludePublic indicates whether to clean public projects
	IncludePublic bool `yaml:"includePublic"`
	// RetainNum configure policy to retain given number tags in repo
	NumPolicy *NumPolicy `yaml:"numberPolicy,omitempty"`
	// RetainTags is tag patterns to be retained
	RetainTags []string `yaml:"retainTags"`
}

type C struct {
	Host     string   `yaml:"host"`
	Auth     Auth     `yaml:"auth"`
	Projects []string `yaml:"projects"`
	Policy   Policy   `yaml:"policy"`
}

var Config = C{}

func Load(configFile string) error {
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		logrus.WithField("f", configFile).Error("Read config file error: ", err)
		return err
	}

	err = yaml.Unmarshal(b, &Config)
	if err != nil {
		logrus.WithField("f", configFile).Error("Unmarshal config file error: ", err)
		return err
	}

	return nil
}