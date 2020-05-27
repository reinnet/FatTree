package common

import (
	"os"

	"github.com/reinnet/topology/model"
	"gopkg.in/yaml.v3"
)

// Write writes a given configuration into topology.yaml.
func Write(cfg model.Config) error {
	b, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	f, err := os.Create("topology.yaml")
	if err != nil {
		return err
	}

	if _, err := f.Write(b); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
