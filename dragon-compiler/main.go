package main

import (
	"pda"
)

func main() {
	pdaParser := pda.NewBracketPDA()
	pdaParser.Parse("(()")
}
