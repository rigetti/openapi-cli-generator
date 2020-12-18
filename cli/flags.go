package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

type flagDef struct {
	name         string
	short        string
	description  string
	defaultValue interface{}
}

var flagRegistry = make(map[string][]*flagDef)

// AddFlag registers a new custom flag for the command path. Use the
// `RegisterBefore` and `RegisterAfter` functions to register a handler that
// can check the value of this flag.
func AddFlag(path, name, short, description string, defaultValue interface{}) {
	if _, ok := flagRegistry[path]; !ok {
		flagRegistry[path] = make([]*flagDef, 0, 1)
	}

	flagRegistry[path] = append(flagRegistry[path], &flagDef{
		name:         name,
		short:        short,
		description:  description,
		defaultValue: defaultValue,
	})
}

// SetCustomFlags sets up the command with additional registered flags.
func SetCustomFlags(cmd *cobra.Command) {
	path := commandPath(cmd)

	if flags, ok := flagRegistry[path]; ok {
		for _, f := range flags {
			switch v := f.defaultValue.(type) {
			case bool:
				cmd.Flags().BoolP(f.name, f.short, v, f.description)
			case int:
				cmd.Flags().Int64P(f.name, f.short, int64(v), f.description)
			case int32:
				cmd.Flags().Int64P(f.name, f.short, int64(v), f.description)
			case int64:
				cmd.Flags().Int64P(f.name, f.short, v, f.description)
			case float32:
				cmd.Flags().Float64P(f.name, f.short, float64(v), f.description)
			case float64:
				cmd.Flags().Float64P(f.name, f.short, v, f.description)
			default:
				cmd.Flags().StringP(f.name, f.short, fmt.Sprintf("%v", v), f.description)
			}
		}
	}
}
