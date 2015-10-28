package main

import (
	"algorithms/bubblesort"
	"algorithms/qsort"
	"bufio" //bufio 包实现了带缓存的 I/O 操作
	"flag"  //解析命令行参数
	"fmt"
	"io" //Reader和Writer接口
	"os"
	"strconv" //与字符串相关的类型转换都是通过 strconv 包实现的
	"time"
)

var infile *string = flag.String("i", "infile", "File contains values for sorting")
var outfile *string = flag.String("o", "outfile", "File to receive sorted values")
var algorithms *string = flag.String("a", "qsort", "Sort algorithm")

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

func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file ", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}
	return nil

}

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile = ", *infile, "outfile = ", *outfile, "algorithms = ", *algorithms)
	}

	values, err := readValues(*infile)
	if err == nil {
		fmt.Println("Read values:", values)
		t1 := time.Now()
		switch *algorithms {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithms", *algorithms, "is  either unknown or unsupported")
		}
		t2 := time.Now()
		fmt.Println("the sorting procsss costs", t2.Sub(t1), "to complete!")
		writeValues(values, *outfile)

	} else {
		fmt.Println(err)
	}

}
