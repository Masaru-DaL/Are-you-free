package config

import "github.com/jinzhu/configor"

var Config struct {
	AUTH struct {
		JWTSecretKey string `default:""`
	}
}

/*
	 config.yamlを読み込む関数,
		読み込んだ後は構造体にアクセスして環境変数を読み込む
*/
func LoadConfigForYaml() {
	err := configor.Load(&Config, "./internal/config/config.yaml")
	if err != nil {
		panic("ERROR: config cannot be loaded")
	}
}
