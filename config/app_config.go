package config

import (
	"github.com/suulaav/altbackup/constants"
	"github.com/suulaav/altbackup/pkg/altutils"
	"gopkg.in/yaml.v3"
)

var AppConfig AppConfigs

type AppConfigs struct {
	Db    Db    `yaml:"db" json:"db"`
	OAuth OAuth `yaml:"oAuth" json:"oAuth"`
}
type Db struct {
	DbType string `yaml:"dbType" json:"dbType"`
	DbUrl  string `yaml:"dbUrl" json:"dbUrl"`
}

type OAuth struct {
	ClientID                string   `yaml:"clientId" json:"client_id"`
	ClientSecret            string   `yaml:"clientSecret" json:"client_secret"`
	RedirectURIs            []string `yaml:"redirectUris" json:"redirect_uris"`
	AuthURI                 string   `yaml:"authUri" json:"auth_uri"`
	TokenURI                string   `yaml:"tokenUri" json:"token_uri"`
	ProjectId               string   `yaml:"projectId" json:"project_id"`
	AuthProviderX509CertUrl string   `yaml:"authProviderX509CertUrl" json:"auth_provider_x509_cert_url"`
}

func GetConfig() AppConfigs {
	appConfigMap := *altutils.ReadYaml(constants.ConfigFilePath)
	yl, _ := yaml.Marshal(appConfigMap)
	err := yaml.Unmarshal(yl, &AppConfig)
	altutils.CheckError(err)
	return AppConfig
}
