package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/snaproute-mino/snapcucumber/parser"
	"io/ioutil"
)

type GListener struct {
	*parser.BaseGherkinListener
}

func ReadFile() string {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(b)
}

func main() {

	str := ReadFile()

	// Setup the input
	is := antlr.NewInputStream(str)

	// Create the Lexer
	lexer := parser.NewGherkinLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewGherkinParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&GListener{}, p.Feature())
}
