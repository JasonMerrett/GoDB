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
	PrepareSyntaxError
)

type StatementType int

const (
	StatementInsert StatementType = iota
	StatementSelect
)

type Row struct {
	id       uint32
	username string
	email    string
}

type Statement struct {
	statementType StatementType
	rowToInsert   Row
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

		statement := &Statement{
			rowToInsert: Row{},
		}
		switch prepareStatement(text, statement) {
		case PrepareSuccess:
			fmt.Println("statement prepared")
		case PrepareUnrecognizedStatement:
			fmt.Printf("Unrecognized keyword %s\n", text)
			continue
		case PrepareSyntaxError:
			fmt.Printf("Syntax error %s\n", text)
			continue
		}

		fmt.Printf("%+v\n", statement)

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
		arg_count, err := fmt.Sscanf(command, "insert %d %s %s", &statement.rowToInsert.id, &statement.rowToInsert.username, &statement.rowToInsert.email)
		if err != nil {
			panic(err)
		}
		if arg_count < 3 {
			return PrepareSyntaxError
		}
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
