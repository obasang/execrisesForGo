/*
练习 12.4：wiki_part1.go

（这是一个独立的练习，但是同时也是为章节15.4做准备）

程序中的数据结构如下，是一个包含以下字段的结构:

type Page struct {
    Title string
    Body  []byte
}
请给这个结构编写一个 save 方法，将 Title 作为文件名、Body作为文件内容，写入到文本文件中。

再编写一个 load 函数，接收的参数是字符串 title，该函数读取出与 title 对应的文本文件。请使用 *Page 做为参数，因为这个结构可能相当巨大，我们不想在内存中拷贝它。请使用 ioutil 包里的函数（参考章节12.2.1）
*/

package main
 
import (
	"fmt"      
	"io/ioutil"
)

type Page struct {
    Title string
    Body  []byte
}

func (this *Page)save() (err error) {
	return ioutil.WriteFile(this.Title, this.Body, 0666)
}

func (this *Page)load(title string) (err error) {
	this.Title = title
	this.Body, err = ioutil.ReadFile(this.Title)
	return err
}
 
func main(){    
	page := Page{
		"pageTest.txt",
		[]byte("# Test\n## Section1\nThis is section1."),
	}
	page.save()

	// load from Page.md
	var new_page Page
	new_page.load("pageTest.txt")
	fmt.Println(string(new_page.Body))
}
