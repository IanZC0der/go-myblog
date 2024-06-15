package conf_test

import (
	"testing"

	"github.com/IanZC0der/go-myblog/conf"
)

func TestLoadConfigFromToml(t *testing.T) {

	err := conf.LoadConfigFromToml("test/config_test.toml")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(conf.C())

}
