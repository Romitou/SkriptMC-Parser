package config

import (
	"github.com/romitou/skriptmc-parser/structures"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

func LoadConfig(ctx *structures.ParserContext) {
	ctx.Config = &structures.Config{}

	var configFileName string
	if _, err := os.Stat("./config.local.yaml"); err == nil {
		configFileName = "config.local.yaml"
	} else if os.IsNotExist(err) {
		configFileName = "config.yaml"
	} else {
		return
	}

	viper.SetConfigFile(configFileName)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("[Skript-MC] An error occurred while reading the configuration: ", err)
	}

	viper.SetEnvPrefix("SKMC")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	toBind := []string{"redis.address", "redis.username", "redis.password"}
	for i := range toBind {
		err = viper.BindEnv(toBind[i])
		if err != nil {
			log.Fatal("[Skript-MC] An error occurred while reading the environment variable: ", err)
			return
		}
	}

	err = viper.Unmarshal(ctx.Config)
	if err != nil {
		log.Fatal("[Skript-MC] An error occurred during the unmarshal of the config: ", err)
	}

	log.Println("[Skript-MC] Configuration successfully loaded.")
}
