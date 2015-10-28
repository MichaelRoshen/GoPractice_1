```ruby
sorter.go是主程序,
qsort.go用于实现快速排序,
bubblesort.go用于实现冒泡排序。 
下面我们先定义一下排序算法函数的函数原型:

func QuickSort(in []int)[]int 
func BubbleSort(in []int)[]int
 
.
├── README.md
├── bin
├── pkg
│   └── darwin_amd64
│       └── algorithms
│           ├── bubblesort.a
│           └── qsort.a
├── sorted.dat
├── sorter
├── src
│   ├── algorithms
│   │   ├── bubblesort
│   │   │   ├── bubblesort.go
│   │   │   └── bubblesort_test.go
│   │   └── qsort
│   │       ├── qsort.go
│   │       └── qsort_test.go
│   └── sorter
│       └── sorter.go
└── unsorted.dat

构建与执行

$ echo $GOPATH
~/goproject/sorter
$ go build algorithm/qsort
$ go build algorithm/bubblesort
$ go test algorithm/qsort
ok algorithm/qsort0.007s
$ go test algorithm/bubblesort 9
ok  algorithm/bubblesort0.013s
$ go install algorithm/qsort
$ go install algorithm/bubblesort
$ go build sorter
$ go install sorter

测试：
./sorter -i unsorted.dat -o sorted.dat -a bubblesort
    infile = unsorted.dat outfile = sorted.dat algorithm = bubblesort

```

