package slovo

import "fmt"

const VERSION = "2024.01.05"
const CODENAME = "U+2C16 GLAGOLITIC CAPITAL LETTER UKU (Ⱆ)"

type Config struct {
	Debug      bool
	ConfigFile string
	Serve      ServeConfig
}

type ServeConfig struct {
	Port int
}

var DefaultConfig Config

func init() {
	DefaultConfig = Config{
		Debug:      false,
		ConfigFile: "etc/config.yaml",
		Serve:      ServeConfig{Port: 3000},
	}
}
func ServeCGI() {

	fmt.Println("in slovo.ServeCGI()")
}

func Serve() {

	fmt.Println("in slovo.Serve()")
}
