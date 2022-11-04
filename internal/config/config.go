package config

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	BaseURL string `env:"base_url,default=http://localhost:8080"`
	Host    string `env:"host,default=0.0.0.0"`
	Port    int    `env:"port,default=8080"`
	DB      DBConfig
}

func (c Config) ListenAddr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type DBConfig struct {
	DSN string `env:"edgedb_dsn"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		lookuper := UpcaseLookuper(envconfig.OsLookuper())

		if err := envconfig.ProcessWith(context.Background(), &cfg, lookuper); err != nil {
			log.Fatal(err)
		}
	})

	return cfg
}

type upcaseLookuper struct {
	Next envconfig.Lookuper
}

func (l *upcaseLookuper) Lookup(key string) (string, bool) {
	return l.Next.Lookup(strings.ToUpper(key))
}

func UpcaseLookuper(next envconfig.Lookuper) *upcaseLookuper {
	return &upcaseLookuper{
		Next: next,
	}
}
