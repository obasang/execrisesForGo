/*
练习 12.5：hello_who.go 写一个"Hello World"的变种程序：把人的名字作为程序命令行执行的一个参数，
比如：hello_who Evan Michael Laura 那么会输出Hello Evan Michael Laura!
*/

package main
 
import (
	"fmt"
	"os"      
	"strings"
)

func main(){    
	who := ""
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("Hello %s!\n",who)
}
