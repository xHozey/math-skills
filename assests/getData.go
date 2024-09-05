package math_skills

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func GetData() []float64 {
	arg := os.Args[1:]
	if len(arg) != 1 {
		log.Fatal("wrong arguments")
	}
	filePath := arg[0]
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	var intSlice []float64
	number := strings.Split(string(data), "\n")
	for i := 0; i < len(number); i++ {
		if number[i] != "" {
			nb, err := strconv.Atoi(number[i])
			if err != nil {
				fmt.Println(number[i])
				log.Fatal(err)
			}
			intSlice = append(intSlice, float64(nb))
		}
	}

	return intSlice
}
