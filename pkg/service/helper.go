package service

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func find(data int, s *[]int) int {
	for _, v := range *s {
		if v == data {
			return data
		}
	}
	return -1
}

func remove(data int, s *[]int) int {
	for iterator, v := range *s {
		if v == data {
			t := (*s)[iterator+1 : len((*s))]
			(*s) = (*s)[:iterator]
			(*s) = append(*s, t...)
			return data
		}
	}
	return -1
}
func sendRange(index int, size int, total_length int, loss *[]int) ([]int, bool) {
	result := true
	lossPackageIndex := make([]int, 0)
	if !(size+index > total_length) {
		for i := index; i <= index+size; i++ {
			result := send(i, loss)
			if result != -1 {
				lossPackageIndex = append(lossPackageIndex, i)
				remove(i, loss)
			}
		}

	} else {
		result = false
	}

	if len(lossPackageIndex) > 0 {
		result = false
	}
	return lossPackageIndex, result
}

func send(index int, loss *[]int) int {
	result := find(index, loss)
	if result == -1 {
		/*
			//fmt.Print("	Sent ")
			//fmt.Print(index)
			//fmt.Println(" nd package.")*/
	} else {
		//fmt.Print("	Unsent ")
		//fmt.Print(index)
		//fmt.Println(" nd package.")
	}
	return result
}

func getLossPackages() (int, *[]int) {
	f, err := os.Open("../../lost-packages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	size, _ := strconv.ParseInt(scanner.Text(), 10, 0)
	var packages []int
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		packages = append(packages, number)
	}

	return int(size), &packages
}
