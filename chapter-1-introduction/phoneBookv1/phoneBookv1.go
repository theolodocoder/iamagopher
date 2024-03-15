package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

type Entry struct {
	Name    string
	Age     int
	Surname string
	Tel     string
}

var data = []Entry{}
var max int = 20
var min int = 5

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func genString(len int64) string {
	tmp := ""
	startChar := "A"
	var i int64 = 1
	for {
		myRand := random(min, max)
		newChar := string(startChar[0] + byte(myRand))
		tmp = tmp + newChar
		if i == len {
			break
		}
		i++
	}
	return tmp
}

func populate(n int) {
	for i := 0; i < n; i++ {
		name := genString(4)
		surname := genString(4)
		tel := strconv.Itoa(random(100, 189))
		age := random(18, 30)
		data = append(data, Entry{name, age, surname, tel})
	}

}

func search(key string) *Entry {
	for i, v := range data {
		if v.Tel == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	arg := os.Args
	if len(arg) == 1 {
		// exe holds the value of the file to be executed
		exe := path.Base(arg[0])
		fmt.Printf("Usage: %s search | list <arguments>\n", exe)
		return
	}
	data = append(data, Entry{"Mihalis", 20, "Tsoukalos", "2109416471"})
	data = append(data, Entry{"Mary", 33, "Doe", "2109416871"})
	data = append(data, Entry{"John", 34, "Black", "2109416123"})

	populate(80)
	fmt.Printf("There are %d entries\n", len(data))

	SEED := time.Now().Unix()
	rand.New(rand.NewSource(SEED))

	// switch to differenciate based on the arguments
	switch arg[1] {
	case "search":
		if len(arg) != 3 {
			fmt.Println("Usage: search Surname")
			return
		}
		result := search(arg[2])

		if result == nil {
			fmt.Println("Entry not found: ", arg[2])
		}
		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("Arguments are not valid")
	}

}
