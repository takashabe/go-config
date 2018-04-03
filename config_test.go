package config

import (
	"testing"

	"github.com/k0kubun/pp"
)

func TestLoad(t *testing.T) {
	load()
	pp.Println(Config)
}
