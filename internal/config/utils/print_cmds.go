package config

import (
	"fmt"
)

func PrintCmds(config string) {
	f = RedfileToStruct(config)

	fmt.Printf("Tasks:\n")

	for t := range f.Tasks {
		if f.Tasks[t].Description != "" {
			fmt.Printf("%s: %s\n", t, f.Tasks[t].Description)
		} else {
			fmt.Printf("%s: <nil>\n", t)
		}
	}
}
