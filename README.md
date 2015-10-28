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
├── sorter
│   ├── algorithms
│   │   ├── bubblesort.go
│   │   ├── bubblesort_test.go
│   │   ├── qsort.go
│   │   └── qsort_test.go
│   ├── sorted.dat
│   ├── sorter
│   ├── sorter.go
│   └── unsorted.dat
└── src
```

