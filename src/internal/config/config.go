package config

import "github.com/jinzhu/configor"

var Config struct {
	DB struct {
		SQLDriver     string `default:"mysql"`
		Database      string `default:"are-you-free"`
		User          string `default:"test_user"`
		Password      string `default:""`
		Root_password string `default:""`
		Path          string `default:""`
	}
	AUTH struct {
		AccessTokenCookieName  string `default:"atcn"`
		RefreshTokenCookieName string `default:"rtcn"`
		JWTSecretKey           string `default:""`
		JWTRefreshSecretKey    string `default:""`
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
