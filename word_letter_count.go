/*
练习 12.1: word_letter_count.go

编写一个程序，从键盘读取输入。当用户输入 'S' 的时候表示输入结束，这时程序输出 3 个数字：
i) 输入的字符的个数，包括空格，但不包括 '\r' 和 '\n'
ii) 输入的单词的个数
iii) 输入的行数
*/

package main
 
import (
	"fmt"      
	"os"
	"bufio"
)
 
func main(){    
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter what you want to input, type S to stop:")
	input, err := inputReader.ReadString('S')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Println("Your input message is %s", input)
	var cbyte, cword, cline int = 0, 0, 1
	for _, c := range input {
		switch c {
		case '\r': fallthrough
		case '\n':
			cword++
			cline++
		case ' ':
			cbyte++
			cword++
		case 'S':
			cbyte++
			cword++ 
		default:
			cbyte++
		}
	}
	fmt.Println("Here are the counts:")
	fmt.Printf("Number of characters: %d\n", cbyte)
	fmt.Printf("Number of words: %d\n", cword)
	fmt.Printf("Number of lines: %d\n", cline)
}
