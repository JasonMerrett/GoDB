package main

import (
	"bufio"
	"fmt"
	"os"
)

type MetaCommandResult int

const (
	MetaCommandSuccess MetaCommandResult = iota
	MetaCommandUnrecognizedCommand
)

type PrepareResult int

const (
	PrepareSuccess PrepareResult = iota
	PrepareUnrecognizedStatement
)

type StatementType int

const (
	StatementInsert StatementType = iota
	StatementSelect
)

type Statement struct {
	statementType StatementType
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("GoDB> ")

		scanner.Scan()
		text := scanner.Text()
		if text[0] == '.' {
			switch doMetaCommand(text) {
			case MetaCommandSuccess:
				os.Exit(0)
			case MetaCommandUnrecognizedCommand:
				fmt.Printf("Unrecognized command %s\n", text)
				continue
			}
		}

		statement := &Statement{}
		switch prepareStatement(text, statement) {
		case PrepareSuccess:
			fmt.Println("statement prepared")
		case PrepareUnrecognizedStatement:
			fmt.Printf("Unrecognized keyword %s\n", text)
			continue
		}

		executeStatement(statement)
		fmt.Println("Executed")
	}
}

func doMetaCommand(command string) MetaCommandResult {
	if command == ".exit" {
		return MetaCommandSuccess
	} else {
		return MetaCommandUnrecognizedCommand
	}
}

func prepareStatement(command string, statement *Statement) PrepareResult {
	if command[:6] == "insert" {
		statement.statementType = StatementInsert
		return PrepareSuccess
	}

	if command[:6] == "select" {
		statement.statementType = StatementSelect
		return PrepareSuccess
	}

	return PrepareUnrecognizedStatement
}

func executeStatement(statement *Statement) {
	switch statement.statementType {
	case StatementInsert:
		fmt.Println("Do insert")
	case StatementSelect:
		fmt.Println("Do select")
	}
}
