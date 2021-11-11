package app

type Config struct {
	HTTPPort    string `env:"PORT" envDefault:"8080"`
	Environment string `env:"ENV"`
	MemoliPath  string `env:"MEMOLI_PATH" envDefault:"/tmp"`
}
