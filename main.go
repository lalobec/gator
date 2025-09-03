package main

import (
	"github.com/lalobec/gator/internal/config"
	"log"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	programState := &state{
		cfg: &cfg,
	}

	cmds := commands {
		registeredCommands : make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	
	if len(os.Args) < 2 {
		log.Fatal("usage: cli <command> [args...]")
	}

	typed_cmd := command {
		name: os.Args[1],
		arguments: os.Args[2:],
	}

	err = cmds.run(programState, typed_cmd)
	if err != nil {
		log.Fatalf("Something bad happened running the command, %v\n", err)
	}
}


	/*
	if err := cfg.SetUser("lalobg"); err != nil {
		log.Fatalf("Could not set user: %v", err)
	}
	*/
