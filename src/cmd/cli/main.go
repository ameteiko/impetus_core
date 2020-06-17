package main

import (
	"fmt"
	"log"

	"github.com/ameteiko/mindnet/src/internal/config"
	"github.com/ameteiko/mindnet/src/internal/di"
	"github.com/ameteiko/mindnet/src/platform/neo4j"
)

func main() {
	cfg := config.NewConfig()

	dbSession, err := di.Neo4JSession(cfg.DB.URI, cfg.DB.User, cfg.DB.Pwd)
	if err != nil {
		log.Fatal("Cannot start application due to database connection error", err)
	}
	defer dbSession.Close()

	db := neo4j.NewDB(dbSession)
	println(db)

	if err := db.CreateNode("Artifact", "ShellTest", "shtest.md"); err != nil {
		fmt.Println("Error: ", err)
	}
	//
	// // Put the terminal into an uncooked mode.
	// if err := (&readline.RawMode{}).Enter(); err != nil {
	// 	log.Fatal("Cannot put terminal into a uncooked mode")
	// }
	//
	// // Display an initial terminal caret.
	// fmt.Print("\033c")
	// input := cmd.NewInput()
	// fmt.Print(input.ShellCaret())
	//
	// // Spin-up a goroutine to read user's input.
	// c := make(chan byte)
	// b := make([]byte, 1)
	// go func() {
	// 	for {
	// 		_, _ = os.Stdin.Read(b)
	// 		c <- b[0]
	// 	}
	// }()
	//
	// for {
	// 	input.Process(<-c)
	// 	fmt.Print("\033c")
	// 	fmt.Printf("%s", input.ShellCaret())
	// }
}
