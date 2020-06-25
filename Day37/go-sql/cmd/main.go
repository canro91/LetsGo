package main

import (
	"fmt"
	"github.com/canro91/30DaysOfGo/Day37/go-sql"
	"github.com/chzyer/readline"
	"github.com/olekukonko/tablewriter"
	"io"
	"log"
	"os"
)

func doSelect(mb gosql.Backend, slct *gosql.SelectStatement) error {
	results, err := mb.Select(slct)
	if err != nil {
		return err
	}

	if len(results.Rows) == 0 {
		fmt.Println("(no results)")
		return nil
	}

	table := tablewriter.NewWriter(os.Stdout)
	headers := []string{}
	for _, col := range results.Columns {
		headers = append(headers, col.Name)
	}
	table.SetHeader(headers)
	table.SetAutoFormatHeaders(false)

	rows := [][]string{}
	for _, result := range results.Rows {
		row := []string{}
		for i, cell := range result {
			typ := results.Columns[i].Type
			r := ""
			switch typ {
			case gosql.IntType:
				i := cell.AsInt()
				r = fmt.Sprintf("%d", i)
			case gosql.TextType:
				s := cell.AsText()
				r = s
			case gosql.BoolType:
				b := cell.AsBool()
				r = "t"
				if !b {
					r = "f"
				}
			}

			row = append(row, r)
		}

		rows = append(rows, row)
	}

	table.SetBorder(false)
	table.AppendBulk(rows)
	table.Render()

	if len(rows) == 1 {
		fmt.Println("(1 result)")
	} else {
		fmt.Printf("(%d results)\n", len(rows))
	}

	return nil
}

func main() {
	backend := gosql.NewMemoryBackend()

	l, err := readline.NewEx(&readline.Config{
		Prompt:          "# ",
		HistoryFile:     "/tmp/gosql.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})
	if err != nil {
		panic(err)
	}
	defer l.Close()

	fmt.Println("Welcome to gosql.")

repl:
	for {
		fmt.Print("# ")
		line, err := l.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue repl
			}
		} else if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading line:", err)
			continue repl
		}

		ast, err := gosql.Parse(line)
		if err != nil {
			log.Fatal(err)
		}

		for _, stmt := range ast.Statements {
			switch stmt.Kind {
			case gosql.CreateTableKind:
				err = backend.CreateTable(ast.Statements[0].CreateTableStatement)
				if err != nil {
					fmt.Println("Error creating table:", err)
					continue repl
				}
				fmt.Println("ok")

			case gosql.InsertKind:
				err = backend.Insert(stmt.InsertStatement)
				if err != nil {
					fmt.Println("Error inserting values:", err)
					continue repl
				}
				fmt.Println("ok")

			case gosql.SelectKind:
				err := doSelect(backend, stmt.SelectStatement)
				if err != nil {
					fmt.Println("Error selecting table:", err)
					continue repl
				}

				fmt.Println("ok")
			}
		}
	}
}
