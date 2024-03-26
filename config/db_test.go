package config_test

import (
	"miniprojek3/config"
	"testing"
)

func TestConnection(t *testing.T) {
	testing.Init()
	config.OpenDB()
}
