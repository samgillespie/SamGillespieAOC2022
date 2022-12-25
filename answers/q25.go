package answers

import "fmt"

func Day25() []interface{} {
	data := ReadInputAsStr(25)
	return []interface{}{q25part1(data), q25part2(data)}

}

func q25part1(data []string) string {
	sum := 0
	for _, row := range data {
		value := ConvertSnafuToInteger(row)

		// fmt.Println(row, "    ", value, "    ", snafu)
		sum += value
	}
	fmt.Println(sum)
	snafu := ConvertIntegerToSnafu(sum)
	return snafu
}

func q25part2(data []string) int {
	// This is a freebee,  Have a good number:
	return 42069
}

func Power(value int, power int) int {
	result := 1
	for i := 0; i < power; i++ {
		result = result * value
	}
	if value < 0 {
		return -result
	} else {
		return result
	}
}

func ConvertSnafuToInteger(snafu string) int {
	solution := 0
	power := 0
	for i := len(snafu) - 1; i >= 0; i-- {

		digit := snafu[i]
		var value int
		if digit == '1' {
			value = 1
		} else if digit == '2' {
			value = 2
		} else if digit == '0' {
			value = 0
		} else if digit == '-' {
			value = -1
		} else if digit == '=' {
			value = -2
		}
		solution += value * Power(5, power)
		power++
	}
	return solution
}

func DivideAndGiveQuotientAndRemainder(value int, divisor int) (int, int) {
	quotient := value / divisor
	remainder := value % divisor
	return quotient, remainder
}

func ConvertIntegerToSnafu(value int) string {
	toProcess := value
	output := ""
	carry := 0
	for toProcess > 0 || carry > 0 {
		remainder := 0
		if carry > 0 {
			toProcess = toProcess + carry
			carry = 0
		}
		toProcess, remainder = DivideAndGiveQuotientAndRemainder(toProcess, 5)
		if remainder == 1 {
			output = "1" + output
		} else if remainder == 2 {
			output = "2" + output
		} else if remainder == 0 {
			output = "0" + output
		} else if remainder == 3 {
			output = "=" + output
			carry = 1
		} else if remainder == 4 {
			output = "-" + output
			carry = 1
		}
	}
	return output
}

// Incorrect 34561628468940
