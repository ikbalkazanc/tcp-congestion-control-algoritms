package service

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type PackageService struct {
}

func NewPackageService() PackageService {
	return PackageService{}
}

func (p *PackageService) GetLostPackages() (int, *[]int) {
	f, err := os.Open("../../lost-packages.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	size, _ := strconv.ParseInt(scanner.Text(), 10, 0)
	var packages []int
	scanner.Scan()
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		packages = append(packages, number)
	}

	return int(size), &packages
}

func (p *PackageService) Generate(_length int, _rate float64, _id_rate float64) (int, *[]int) {
	length := _length

	rate := _rate

	//fmt.Println("Gönderilecek dosya paket sayısı : " + strconv.Itoa(length))
	//fmt.Sprintf("Paket kayıp başlangıç oranı : %f", rate)

	id_rate := _id_rate

	printIncreaseDecreaseRate(id_rate)

	f, err := os.Create("../../lost-packages.txt")
	if err != nil {
		log.Fatal(err)
	}

	writeFileIndexLength(f, strconv.Itoa(length))
	defer f.Close()
	var packages []int
	n := 1
	lossNumber := 0
	for n < int(length) {
		rate = rate + id_rate*rate
		if rate != rate+id_rate*rate {
			fmt.Print("Oran : ")
			fmt.Println(rate)
		}
		if getRateResult(rate) {
			_, err3 := f.WriteString(strconv.Itoa(n) + "\n")
			packages = append(packages, n)
			if err3 != nil {
				log.Fatal(err3)
			}
			lossNumber++
		}
		n++
	}

	fmt.Println("Kayıp paket sayısı : " + strconv.Itoa(lossNumber))

	return lossNumber, &packages
}

//get boolean rate according to arg[2]
func getRateResult(x float64) bool {
	rand.Seed(time.Now().UnixNano() * time.Now().UnixNano())
	random := rand.Float64()
	if x > random {
		return true
	}
	return false
}

func printIncreaseDecreaseRate(x float64) {
	if x > 0 {
		fmt.Println("Kayıp Paket Değişim Oranı (+): " + strconv.FormatFloat(x, 'f', 6, 64))
	} else if x < 0 {
		fmt.Println("Kayıp Paket Değişim Oranı (-): " + strconv.FormatFloat(x, 'f', 6, 64))
	} else {
		fmt.Println("Kayıp Paket Değişim Oranı (): 0")
	}
}

func writeFileIndexLength(file *os.File, length string) {
	//first index package length
	_, err2 := file.WriteString(length + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}

}
