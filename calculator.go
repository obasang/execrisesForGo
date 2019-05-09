/*
练习 12.2: calculator.go

编写一个简单的逆波兰式计算器，它接受用户输入的整型数（最大值 999999）和运算符 +、-、*、/。
输入的格式为：number1 ENTER number2 ENTER operator ENTER --> 显示结果
当用户输入字符 'q' 时，程序结束。请使用您在练习11.3中开发的 stack 包
*/

package main
 
import (
	"fmt"      
	"os"
	"bufio"
	"strings"
	"errors"
	"strconv"
)

func parseArgs(c []string)(float64, float64, error) {
	num1, err := strconv.ParseFloat(c[0], 64)
	if err != nil {
		fmt.Println(err)
		return 0.0, 0.0, err
	}
	num2, err := strconv.ParseFloat(c[1], 64)
	if err != nil {
		fmt.Println(err)
		return 0.0, 0.0, err
	}
	return num1, num2, nil
}

func MyStack(e string) (float64, error) {
	result := 0.0
	c := strings.Split(e, " ")
	if len(c) - 1 < 2 {
		if c[0] == "q" {
			return 0.0, errors.New("quit")
		}
		return 0.0, errors.New("error: some arguments are not supplied")
	}
	num1, num2, err := parseArgs(c)
	if err != nil {
		fmt.Println(err)
		return 0.0, err
	}
	if num1 > 999999 || num2 > 999999{
		return 0.0, errors.New("error: arguments can not large than 999999")
	}
	switch c[2] {
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0.0 {
			return 0.0, errors.New("error: you tried to divide by zero")
		}
		result = num1 / num2
	case "+":
		fmt.Println(num1, "+", num2," =", num1+num2)
		result = num1 + num2
	case "-":
		result = num1 - num2
	}
	return result, nil
}
 
func main(){    
	for {
		inputReader := bufio.NewReader(os.Stdin)
		fmt.Print("calc>")
		input, err := inputReader.ReadString('\n')

		fmt.Println(input)
		res, err := MyStack(strings.Replace(string(input), "\n", "", -1))
		if err != nil {
			if err.Error() == "quit"{
				fmt.Println("quit")
				return
			}
			fmt.Println(err)
		} else {
			fmt.Println(res)
		}
	}
}
