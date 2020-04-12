package main

import (
	"sort"
	"bufio"
	"flag"
	"fmt"
	"github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing/object"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

var ignore = [...]string{".jekyll-cache", "_site"}

func shouldIgnore(path string) bool {
	for _, pattern := range ignore {
		if pattern == path {
			return true
		}
	}
	return false
}

func scan(path string) {
	folders := recursivelyScanFolder(path)
	dotfile := findDotFile()
	addFoldersToDotfile(folders, dotfile)
	fmt.Println(folders)
	fmt.Println(dotfile)
}

func recursivelyScanFolder(path string) []string {
	var folders = make([]string, 0)
	return scanFolders(folders, path)
}

func scanFolders(folders []string, folder string) []string {
	var path string

	folder = strings.TrimSuffix(folder, "/")
	files, _ := ioutil.ReadDir(folder)
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		if shouldIgnore(file.Name()) {
			continue
		}

		path = folder + "/" + file.Name()
		if file.Name() == ".git" {
			path = strings.TrimSuffix(path, ".git")
			folders = append(folders, path)
			continue
		}
		folders = scanFolders(folders, path)
	}
	return folders
}

func findDotFile() string {
	homeDir, _ := os.UserHomeDir()
	return homeDir + "/" + ".gogitlocalstats.txt"
}

func addFoldersToDotfile(folders []string, dotfile string) {
	f, _ := os.OpenFile(dotfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	for _, folder := range folders {
		f.WriteString(folder)
		f.WriteString("\n")
	}
	f.Close()
}

func statistics(email string) {
	commits := processRepositories(email)
	printStats(commits)
}

const daysInLastMonth = 30
const weeksInLastMonth = 4

func processRepositories(email string) map[int]int {
	dotfile := findDotFile()
	repos := parseDotFile(dotfile)
	daysInMap := daysInLastMonth

	commits := make(map[int]int, daysInMap)
	for i := 0; i < daysInMap; i++ {
		commits[i] = 0
	}

	for _, repo := range repos {
		commits = fillCommits(email, repo, commits)
	}

	return commits
}

func parseDotFile(path string) []string {
	var repos []string

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		repos = append(repos, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return repos
}

func today() time.Time {
	now := time.Now()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return today
}

func daysAgo(days int) time.Time {
	since := today().Add(-time.Hour * 24 * daysInLastMonth)
	return since
}

func fillCommits(email, path string, commits map[int]int) map[int]int {
	repo, _ := git.PlainOpen(path)
	ref, _ := repo.Head()

	since := daysAgo(daysInLastMonth)
	cIter, _ := repo.Log(&git.LogOptions{From: ref.Hash(), Since: &since})

	offset := calcOffset()
	_ = cIter.ForEach(func(c *object.Commit) error {
		if c.Author.Email != email {
			return nil
		}

		daysAgo := countDaysSince(c.Author.When) + offset
		if daysAgo != -1 {
			commits[daysAgo]++
		}

		return nil
	})

	return commits
}

func calcOffset() int {
	var offset int
	weekday := time.Now().Weekday()

	switch weekday {
	case time.Sunday:
		offset = 7
	case time.Monday:
		offset = 6
	case time.Tuesday:
		offset = 5
	case time.Wednesday:
		offset = 4
	case time.Thursday:
		offset = 3
	case time.Friday:
		offset = 2
	case time.Saturday:
		offset = 1
	}

	return offset
}

func countDaysSince(date time.Time) int {
	days := 0

	today := today()
	for date.Before(today) {
		date = date.Add(time.Hour * 24)
		days++
		if days > daysInLastMonth {
			return -1
		}
	}

	return days
}

func printStats(commits map[int]int) {
	keys := sortMapIntoSlice(commits)
    cols := buildCols(keys, commits)
    printCells(cols)
}

func sortMapIntoSlice(m map[int]int) []int {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	return keys
}

type column []int

func buildCols(keys []int, commits map[int]int) map[int]column {
	cols := make(map[int]column) 
	col := column{}

	for _, k := range keys {
		week := int(k / 7)
		dayInWeek := k % 7

		if dayInWeek == 0 {
			col = column{}
		}

		col = append(col, commits[k])

		if dayInWeek == 6 {
			cols[week] = col
		}
	}
	return cols
}

func printCells(cols map[int]column) {
	printMonths()
	for j := 6; j >= 0; j-- {
		for i := weeksInLastMonth + 1; i >= 0; i-- {
			if i == weeksInLastMonth+1 {
				printDayCol(j)
			}
			if col, ok := cols[i]; ok {
				//special case today
				if i == 0 && j == calcOffset()-1 {
					printCell(col[j], true)
					continue
				} else {
					if len(col) > j {
						printCell(col[j], false)
						continue
					}
				}
			}
			printCell(0, false)
		}
		fmt.Printf("\n")
	}
}

func printMonths() {
	week := daysAgo(daysInLastMonth)
	month := week.Month()
	fmt.Printf("         ")
	for {
		if week.Month() != month {
			fmt.Printf("%s ", week.Month().String()[:3])
			month = week.Month()
		} else {
			fmt.Printf("    ")
		}

		week = week.Add(7 * time.Hour * 24)
		if week.After(time.Now()) {
			break
		}
	}
	fmt.Printf("\n")
}

func printDayCol(day int) {
	out := "     "
	switch day {
	case 1:
		out = " Mon "
	case 3:
		out = " Wed "
	case 5:
		out = " Fri "
	}

	fmt.Printf(out)
}

func printCell(val int, today bool) {
    escape := "\033[0;37;30m"
    switch {
    case val > 0 && val < 5:
        escape = "\033[1;30;47m"
    case val >= 5 && val < 10:
        escape = "\033[1;30;43m"
    case val >= 10:
        escape = "\033[1;30;42m"
    }

    if today {
        escape = "\033[1;37;45m"
    }

    if val == 0 {
        fmt.Printf(escape + "  - " + "\033[0m")
        return
    }

    str := "  %d "
    switch {
    case val >= 10:
        str = " %d "
    case val >= 100:
        str = "%d "
    }

    fmt.Printf(escape+str+"\033[0m", val)
}

func main() {
	var folder, email string
	flag.StringVar(&folder, "add", "", "Add a folder to scan")
	flag.StringVar(&email, "email", "johndoe@hotmail.com", "User email to scan")
	flag.Parse()

	if len(folder) != 0 {
		scan(folder)
		return
	}

	statistics(email)
}
