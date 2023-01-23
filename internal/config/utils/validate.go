package config

import (
	"io/ioutil"
	"log"
	"strings"

	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/encoding/yaml"
)

func ValidateRedfile(config string) {
	redFile, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(err)
	}

	r := strings.NewReader(string(redFile))
	b, _ := ioutil.ReadAll(r)
	cue, _ := ioutil.ReadFile("internal/config/scheme/redfile.cue")

	// Cue API for Go
	c := cuecontext.New()
	v := c.CompileBytes(cue)
	err = yaml.Validate(b, v)

	if err != nil {
		log.Fatal(err)
	}
}
