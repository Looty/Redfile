package config

import (
	"fmt"
	"log"
	"os/exec"
)

func RunCmd(config string, command string) {
	f = RedfileToStruct(config)

	// TODO: validate if cmd exists before executing

	fmt.Printf("Executing %s:\n", command)

	if f.Env != nil {
		SetGlobalEnvVariable(config)
	}

	if len(f.Tasks[command].Env) > 0 {
		SetCmdEnv(f.Tasks[command].Env)
	}

	for c := range f.Tasks[command].Cmds {

		out, err := exec.Command("bash", "-c", f.Tasks[command].Cmds[c]).Output()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", out)
	}
}
