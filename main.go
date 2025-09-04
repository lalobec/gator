package main

import (
	"database/sql"
	"github.com/lalobec/gator/internal/config"
	"github.com/lalobec/gator/internal/database"
	"log"
	"os"
	_ "github.com/lib/pq"
)


type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db: dbQueries,
		cfg: &cfg,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)

	if len(os.Args) < 2 {
		log.Fatal("usage: cli <command> [args...]")
	}

	typed_cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	err = cmds.run(programState, typed_cmd)
	if err != nil {
		log.Fatal(err)
	}
}

/*
	if err := cfg.SetUser("lalobg"); err != nil {
		log.Fatalf("Could not set user: %v", err)
	}
*/
