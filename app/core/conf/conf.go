package conf

import (
	"github.com/spf13/viper"
)

type Model struct {
	App      App
	PanelApp App
	PGSql    map[string]PGSql
	Redis    map[string]Redis
	SMS      SMS
	Token    Token
}

func (m *Model) SetUp() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return viper.Unmarshal(m)
}
