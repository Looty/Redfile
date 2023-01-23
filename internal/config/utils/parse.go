package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func RedfileToStruct(config string) Redfile {
	f, err := os.ReadFile(config)
	if err != nil {
		log.Fatal(err)
	}

	// Create an empty Redfile to be are target of unmarshalling
	var r Redfile

	// Unmarshal our input YAML file into empty Redfile (var r)
	if err := yaml.Unmarshal(f, &r); err != nil {
		log.Fatal(err)
	}

	// Print out the new struct
	// fmt.Printf("%+v\n", r)
	return r
}
