package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	AluaUrl  string `split_words:"true" required:"true"`
	BbarUrl  string `split_words:"true" required:"true"`
	BmaUrl   string `split_words:"true" required:"true"`
	BymaUrl  string `split_words:"true" required:"true"`
	CepuUrl  string `split_words:"true" required:"true"`
	ComeUrl  string `split_words:"true" required:"true"`
	CresUrl  string `split_words:"true" required:"true"`
	EdnUrl   string `split_words:"true" required:"true"`
	GgalUrl  string `split_words:"true" required:"true"`
	IrsaUrl  string `split_words:"true" required:"true"`
	LomaUrl  string `split_words:"true" required:"true"`
	MirgUrl  string `split_words:"true" required:"true"`
	PampUrl  string `split_words:"true" required:"true"`
	SupvUrl  string `split_words:"true" required:"true"`
	Teco2Url string `split_words:"true" required:"true"`
	Tgno4Url string `split_words:"true" required:"true"`
	Tgsu2Url string `split_words:"true" required:"true"`
	TranUrl  string `split_words:"true" required:"true"`
	TxarUrl  string `split_words:"true" required:"true"`
	ValoUrl  string `split_words:"true" required:"true"`
	YpfdUrl  string `split_words:"true" required:"true"`
	HttpPort string `split_words:"true" default:"8080"`
}

func NewConfig() *Config {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		panic(err)
	}

	return &c
}
