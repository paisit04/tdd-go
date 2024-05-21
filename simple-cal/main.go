package main

import (
	"flag"
	"log"
	"simple-cal/calculator"
	"simple-cal/input"
)

func main() {
	expr := flag.String("expression", "", "mathematical expression to parse")
	flag.Parse()

	engine := calculator.NewEngine()
	validator := input.NewValidator(engine.GetNumOperands(), engine.GetValidOperators())
	parser := input.NewParser(engine, validator)
	result, err := parser.ProcessExpression(*expr)
	log.Println(*expr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(*result)
}
