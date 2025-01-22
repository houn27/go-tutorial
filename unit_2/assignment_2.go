package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func FloatAss() {
	coins := []float64{0.05, 0.10, 0.25}
	piggy := 0.0
	fmt.Printf("Initial piggy bank: $%4.2f\n", piggy)
	for piggy < 20 {
		tmp_coin := coins[rand.Intn(len(coins))]
		piggy += tmp_coin
		fmt.Printf("Current piggy bank: $%4.2f \n", piggy)
	}
}

func BigNumAss() {
	const light_year_unit = 9.4607e12
	const distance float64 = 2360000000000000000

	fmt.Printf("type of distance: %T \n", distance)
	fmt.Printf("distance from earth is %v light year", distance/light_year_unit)
}

func StringAss() {
	old_str := "L fdph,L vdz,L frqtxhuhg"
	new_str := ""
	for _, tmp_char := range old_str {
		if tmp_char >= 'a' && tmp_char <= 'z' {
			tmp_char += 3
			if tmp_char > 'z' {
				tmp_char -= 26
			}
		} else if tmp_char >= 'A' && tmp_char <= 'Z' {
			tmp_char += 3
			if tmp_char > 'Z' {
				tmp_char -= 26
			}
		}

		new_str += string(tmp_char)
	}
	fmt.Println(new_str)
}

func TransferAss(str string) (bool, error) {
	if str == "true" || str == "yes" || str == "1" {
		return true, nil
	} else if str == "false" || str == "no" || str == "0" {
		return false, nil
	} else {
		return false, errors.New("invalid input")
	}
}

// float
func FloatTest() {
	num := 12.222
	pi_val := math.Pi
	var empty_val float64
	fmt.Printf("Type %T of num: %[1]v \n", num)
	fmt.Printf("Type %T of pi: %[1]v \n", pi_val)
	fmt.Printf("Typen %T of empty float variable: %[1]v \n", empty_val)
}

// float compare
func FloatEqualTest() {
	num_1 := 0.1
	num_1 += 0.2
	fmt.Printf("Value of num_1: %v \n", num_1)
	fmt.Printf("== compare num_1 with o.3: %v \n", num_1 == 0.3)
	fmt.Printf("Minus compare num_1 with o.3: %v \n", math.Abs(num_1-0.3) < 0.00001)
}

// integer
func IntegerTest() {
	now_time := time.Now()
	now_timestamp := time.Now().Unix()
	fmt.Println("Now time: ", now_time)
	fmt.Printf("Type %T of now timestamp: %[1]v \n", now_timestamp)
}

// big number: should declare as float
func BigTest() {
	// const num = 24000000000000000000
	const num float64 = 24000000000000000000
	fmt.Printf("Type %T of num: %[1]v \n", num)

	const num_1 = 240000
	fmt.Printf("Type %T of num_1: %[1]v \n", num_1)
}

// string
func StringTest() {
	c := 'a'
	c = c + 3
	fmt.Printf("Character '%c' refer to value %[1]v \n", c)

	// len give number of byte
	massage := "shalom"
	fmt.Printf("Length of \"%v\": %v \n", massage, len(massage))

	//string transfer in chracter array
	for idx, c := range massage {
		fmt.Printf("%v-th character is:%[2]c, with number: %[2]v in utf-8\n", idx, c)
	}
}

func main() {

}
