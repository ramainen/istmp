package intersector

// Различные реализации метода Merge, замеры по скорости
/*
BenchmarkMerge-4                  	  100000	     13322 ns/op
BenchmarkMergeFast-4              	  100000	     12683 ns/op
BenchmarkMergeFast32-4            	  200000	     12056 ns/op
BenchmarkMergeFast32bool-4        	  100000	     12062 ns/op
BenchmarkMergeFast32boolArray-4   	  100000	     12062 ns/op
BenchmarkMergeFast32boolWoID-4    	  100000	     12153 ns/op
BenchmarkMerge32overDistinct-4    	  100000	     12149 ns/op

*/

//MergeSorted
func Merge(args ...[]int) []int {
	ns := make([]int, len(args))
	is := make([]int, len(args))
	n := len(args)
	f := true

	protect := 0 //Защита от бесконечного зацикливания
	cprotect := 0

	stopped := make([]bool, len(args))
	stopped_cnt := 0
	min := 0

	if n < 2 {
		if n == 1 {
			return args[0]
		}
		f = false
	}
	maxlen := 0
	for i := 0; f && i < n; i++ {

		if maxlen < len(args[i]) {
			maxlen = len(args[i])
		}
		ns[i] = len(args[i])
		is[i] = 0
		protect += ns[i]
		if ns[i] == 0 {
			stopped[i] = true
			stopped_cnt++
		} else {
			stopped[i] = false
		}

	}
	protect += len(args)
	result_ids := make([]int, 0, maxlen*2)
	//min = args[0][0]
	f = (stopped_cnt < n)
	for f {
		cprotect++
		if cprotect > protect {
			/**/
			panic("inifinity cycled")
			break
		}
		for i := 0; i < n; i++ {
			if stopped[i] == false {
				min = args[i][is[i]]
				break
			}
		}

		for i := 0; i < n; i++ {
			for stopped[i] == false && args[i][is[i]] < min {
				min = args[i][is[i]]
			}
		}

		result_ids = append(result_ids, min)

		for i := 0; i < n; i++ {
			if stopped[i] == false && args[i][is[i]] == min && is[i] <= ns[i]-2 {
				is[i]++
			}

			if stopped[i] == false && is[i] >= ns[i]-1 && args[i][is[i]] == min {
				stopped[i] = true
				stopped_cnt++
			}
		}

		f = (stopped_cnt < n)
	}
	return result_ids
}

func MergeFast32bool(args ...[]int) []int {
	tempMap := make(map[int]bool)
	for _, arg := range args {
		for id := range arg {
			tempMap[arg[id]] = true
		}
	}
	tempArray := make([]int, 0, len(tempMap))
	for key := range tempMap {
		tempArray = append(tempArray, key) //TODO: протестировать умный append
	}
	return tempArray
}

func Merge64(args ...[]int) []int {
	tempMap := make(map[int]uint8)
	for _, arg := range args {
		for id := range arg {
			tempMap[arg[id]] = 0
		}
	}
	tempArray := make([]int, 0)
	for key := range tempMap {
		tempArray = append(tempArray, key) //TODO: протестировать умный append
	}
	return tempArray
}
func MergeFast(args ...[]int) []int {
	tempMap := make(map[int]uint8)
	for _, arg := range args {
		for id := range arg {
			tempMap[arg[id]] = 0
		}
	}
	tempArray := make([]int, 0, len(tempMap))
	for key := range tempMap {
		tempArray = append(tempArray, key) //TODO: протестировать умный append
	}
	return tempArray
}

func MergeFast32(args ...[]int) []int {
	tempMap := make(map[int]uint8)
	for _, arg := range args {
		for id := range arg {
			tempMap[arg[id]] = 0
		}
	}
	tempArray := make([]int, 0, len(tempMap))
	for key := range tempMap {
		tempArray = append(tempArray, key) //TODO: протестировать умный append
	}
	return tempArray
}

func MergeFast32boolArray(args ...[]int) []int {
	tempMap := make(map[int]bool)
	for _, arg := range args {
		for id := range arg {
			tempMap[arg[id]] = true
		}
	}
	tempArray := make([]int, len(tempMap))

	i := 0
	for key := range tempMap {
		tempArray[i] = key
		i++
	}
	return tempArray
}

func MergeFast32boolWoID(args ...[]int) []int {
	tempMap := make(map[int]bool)
	for _, arg := range args {
		for _, v := range arg {
			tempMap[v] = true
		}
	}
	tempArray := make([]int, 0, len(tempMap))
	for key := range tempMap {
		tempArray = append(tempArray, key) //TODO: протестировать умный append
	}
	return tempArray
}
