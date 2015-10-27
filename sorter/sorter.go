package main

import (
	"bufio" //bufio 包实现了带缓存的 I/O 操作
	"flag"  //解析命令行参数
	"fmt"
	"io" //Reader和Writer接口
	"os"
	"strconv" //与字符串相关的类型转换都是通过 strconv 包实现的
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("failed to open the input file ", infile)
		return
	}

	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0) //make用于slice，map，和channel的初始化。

	for {

		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line)
		value, err1 := strconv.Atoi(str)

		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return
}

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithm = ", *algorithm)
	}
	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read values:", values)
	} else {
		fmt.Println(err)
	}

}
