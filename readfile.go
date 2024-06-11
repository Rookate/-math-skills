package skills

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Read_File(number string) []int {
	file, err := os.Open(number)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file_scanner := bufio.NewScanner(file)

	var data []int

	for file_scanner.Scan() {
		New_data, err := strconv.Atoi(file_scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, New_data)
	}
	return data
}
