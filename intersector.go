package intersector

/*
	16.02.2017
	License: MIT

*/

func Distinct(arg []int) []int {

	temp := make(map[int]int8)
	result := make([]int, 0, len(arg))
	//for id := range arg {
	for id := 0; id <= len(arg)-1; id++ {
		temp[arg[id]] = 0
	}
	for key := range temp {
		result = append(result, key)
	}
	return result

}

func IntersectOverMap(args ...[]int) []int {
	arrLength := len(args)
	tempMap := make(map[int]int)
	maxlen := 0
	for _, arg := range args {
		tempArr := arg // Distinct(arg)
		if maxlen < len(tempArr) {
			maxlen = len(tempArr)
		}
		for idx := range tempArr {
			if _, ok := tempMap[tempArr[idx]]; ok {
				tempMap[tempArr[idx]]++
			} else {
				tempMap[tempArr[idx]] = 1
			}
		}
	}
	tempArray := make([]int, 0, maxlen+1)
	for key, val := range tempMap {
		if val == arrLength {
			tempArray = append(tempArray, key)
		}
	}

	return tempArray
}
func Intersect(args ...[]int) []int {

	ns := make([]int, len(args))
	is := make([]int, len(args))
	n := len(args)
	f := true
	max := 0
	//fmt.Println("sda")
	v := 0

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
		if ns[i] == 0 {
			f = false
		}
	}
	result_ids := make([]int, 0, maxlen)
	for f {

		max = args[0][is[0]]
		for i := 1; i < n; i++ {
			v = args[i][is[i]]
			if max < v {
				max = v
			}
		}
		for i := 0; i < n; i++ {
			for is[i] < ns[i] && args[i][is[i]] < max {
				is[i]++
			}
			if is[i] == ns[i] {
				f = false
			}
		}
		if f {
			fl := true
			for i := 0; i < n; i++ {
				if args[i][is[i]] != max {
					fl = false
				}
			}

			if fl {
				result_ids = append(result_ids, max)
				for i := 0; i < n && f; i++ {
					is[i]++
					if is[i] == ns[i] {
						f = false
					}
				}
			}
		}
	}
	return result_ids
}
func BinarySearch(target []int, value_min int, value_max int) (int, int) {

	start_index := int(0)
	end_index := int(len(target)) - 1
	start_index_max := int(0)
	end_index_max := int(len(target)) - int(1)

	if value_min > value_max {
		return -1, -1
	}
	for start_index <= end_index {
		median := (start_index + end_index) / 2

		if target[median] < value_min {
			start_index = median + 1
		} else {
			end_index = median - 1
		}

	}

	for start_index_max <= end_index_max {
		median := (start_index_max + end_index_max) / 2

		if target[median] <= value_max {
			start_index_max = median + 1
		} else {
			end_index_max = median - 1
		}

	}

	if start_index == int(len(target)) {
		start_index = -1
	}
	if start_index_max == int(len(target))+1 {
		start_index_max = 0
	}

	return start_index, start_index_max - 1

}

func SearchBetween(values []int, results []int, min_value int, max_value int) []int {
	start_index, end_index := BinarySearch(values, min_value, max_value)
	//fmt.Println(start_index, end_index)
	var result []int
	if start_index != -1 {
		if end_index != -1 {
			result = results[start_index : end_index+1]
		} else {
			result = results[start_index:]
		}
	} else {
		if end_index != -1 {
			result = results[:end_index+1]
		}
		//else - пустой массив
	}
	return result
}

func GetData(table string, ids []int, fields []string) map[int]map[string][]string {
	response := make(map[int]map[string][]string)
	for _, i := range ids {
		response[i] = make(map[string][]string)

		for _, v := range fields {

			response[i][v] = global_index.TermValues[table][v][i]
		}
	}
	return response
	//TermValues map[string]map[string][][]string       /* Прямоей кеш значений по первичному ключу table.field.id.values. */
	//global_index.TermIndex[table]

	//[facet.Key][k]
}
