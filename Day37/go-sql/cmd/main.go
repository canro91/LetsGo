package main

import (
	"log"
	"strings"
	"os"
	"fmt"
	"bufio"
	"github.com/canro91/30DaysOfGo/Day37/go-sql"
)

func main() {
	backend := gosql.NewMemoryBackend()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to gosql.")
	for {
		fmt.Print("# ")
		text, err := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		ast, err := gosql.Parse(text)
		if err != nil {
			log.Fatal(err)
		}

		for _, stmt := range ast.Statements {
			switch stmt.Kind {
			case gosql.CreateTableKind:
				err = backend.CreateTable(ast.Statements[0].CreateTableStatement)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("ok")

			case gosql.InsertKind:
				err = backend.Insert(stmt.InsertStatement)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("ok")

			case gosql.SelectKind:
				results, err := backend.Select(stmt.SelectStatement)
				if err != nil {
					log.Fatal(err)
				}

				for _, col := range results.Columns {
					fmt.Printf("| %s", col.Name)
				}
				fmt.Println("|")

				for i := 0; i < 20; i++ {
					fmt.Printf("=")
				}
				fmt.Println()
				
				for _, result := range results.Rows {
					fmt.Printf("|")

					for i, cell := range result {
						typ := results.Columns[i].Type
						s := ""
						switch typ {
						case gosql.IntType:
							s = fmt.Sprintf("%d", cell.AsInt())
						case gosql.TextType:
							s = cell.AsText()
						}

						fmt.Printf(" %s | ", s)
					}

					fmt.Println()
				}

				fmt.Println()
				fmt.Printf("%d results\n", len(results.Rows))

				fmt.Println("ok")
			}
		}
	}
}