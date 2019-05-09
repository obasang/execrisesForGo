/*
练习 12.3：read_csv.go

文件 products.txt 的内容如下：

"The ABC of Go";25.5;1500
"Functional Programming with Go";56;280
"Go for It";45.9;356
"The Go Way";55;500
每行的第一个字段为 title，第二个字段为 price，第三个字段为 quantity。内容的格式基本与 示例 12.3c 的相同，除了分隔符改成了分号。请读取出文件的内容，创建一个结构用于存取一行的数据，然后使用结构的切片，并把数据打印出来。

关于解析 CSV 文件，encoding/csv 包提供了相应的功能。具体请参考 http://golang.org/pkg/encoding/csv/
*/

package main
 
import (
	"fmt"      
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"
	"io"
)

type Book struct {
	title string
	price float64
	quantity int
}
 
func main(){    
	bks := make([]Book, 1)
	file, err := os.Open("products.txt")
	if err != nil {
		log.Fatalf("Error %s opening file products.txt: ", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err:= reader.ReadString('\n')
		if err != nil{
			if err == io.EOF {
				break
			}
			log.Fatalf("Error %s reading file: ", err)
			return
		}

		line = strings.Replace(line, "\n", "", -1)
		strs := strings.Split(line, ";")
		book := new(Book)
		book.title = strs[0]
		book.price, err = strconv.ParseFloat(strs[1], 64)
		if err != nil {
			log.Fatalf("Error in file: %v", err)
		}
		book.quantity, err = strconv.Atoi(strs[2])
		if err != nil {
			log.Fatalf("Error in file: %v", err)
		}
		if bks[0].title == ""{
			bks[0] = *book
		} else {
			bks = append(bks, *book)
		}
	}
	fmt.Println("We have read the following books from the file: ")
	for _, bk := range bks {
		fmt.Println(bk)
	}
}
