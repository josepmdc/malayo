package main

import (
	"log"
	"malayo/cmd"
)

func main()  {
	if err := cmd.RootCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
