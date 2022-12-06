package main

import "github.com/clementejuliana/grupoDprojetoGO/server"

func main() {
	server := server.NewServer()

	server.Run()

}