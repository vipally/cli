package cli

import (
	"flag"
	"fmt"
)

type Flag struct {
	LogicName   string
	Name        string
	Aliases     []string
	Usage       string
	DefaultText string
	Required    bool
	Hidden      bool
	EnvVars     []string
	FilePath    string
	Value       Value
	names       []string
	logicName   string
}
