package main

import (
	"log"
	"mutants/app/presenters"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	presenters.RunRestServer()
}
