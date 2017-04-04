package intersector

// "fmt"
import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/boltdb/bolt"
)

type Index struct {
	TermIndex  map[string]map[string]map[string][]int /*Обратный индекс строковых значений table.field.value.ids*/
	TermValues map[string]map[string][][]string       /* Прямоей кеш значений по первичному ключу table.field.id.values. */
	TermSort   map[string]map[string][]int            /* список полей отсортиорованных от меньшего к большему table.field.ids. */

	NumericIndex  map[string]map[string][]int       /* список значений отсортированных по возрастанию table.field.values */
	NumericSort   map[string]map[string][]int       /* список ids отсортированных по возрастанию table.field.ids. он же sort */
	NumericValues map[string]map[string]map[int]int /* значение поля по id table.field.id.value*/
	IdsList       map[string][]int                  /*список айдишников в products*/
}

var global_index *Index

func NewEmptyIndex() *Index {
	fmt.Print("new")
	new_index := new(Index)
	global_index = new_index
	return new_index
}

func NewIndex() *Index {
	new_index := new(Index)

	//СОздаем color и  size
	new_index.TermIndex = make(map[string]map[string]map[string][]int)

	new_index.TermIndex["products"] = make(map[string]map[string][]int)

	new_index.TermIndex["products"]["color"] = make(map[string][]int)
	new_index.TermIndex["products"]["color"]["red"] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	new_index.TermIndex["products"]["color"]["green"] = []int{3, 4, 5, 6, 7, 2, 4}
	new_index.TermIndex["products"]["color"]["blue"] = []int{3, 4, 5, 6, 7, 2, 4}
	new_index.TermIndex["products"]["color"]["black"] = []int{10}

	new_index.TermIndex["products"]["width"] = make(map[string][]int)
	new_index.TermIndex["products"]["width"]["100"] = []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	new_index.TermIndex["products"]["width"]["200"] = []int{2, 4, 6, 8, 10, 12, 14, 16}
	new_index.TermIndex["products"]["width"]["300"] = []int{3, 6, 9, 12, 15}
	new_index.TermIndex["products"]["width"]["400"] = []int{10}
	/*
		new_index.TermValues = make(map[string]map[string]map[int][]string)
		new_index.TermValues["products"] = make(map[string]map[int][]string)
	*/
	//global_index.TermValues[s.Table][facet.Key][i]
	/*
		new_index.TermValues["products"]["color"] = map[int][]string{
			1:  []string{"red"},
			2:  []string{"red", "green", "blue"},
			3:  []string{"red", "green", "blue"},
			4:  []string{"red", "green", "blue"},
			5:  []string{"red", "green", "blue"},
			6:  []string{"red", "green", "blue"},
			7:  []string{"red", "green", "blue"},
			8:  []string{"red"},
			9:  []string{"red"},
			10: []string{"black"},
		}
		new_index.TermValues["products"]["width"] = map[int][]string{
			1:  []string{"100"},
			2:  []string{"200"},
			3:  []string{"100", "300"},
			4:  []string{"200"},
			5:  []string{"100"},
			6:  []string{"200", "300"},
			7:  []string{"100"},
			8:  []string{"200"},
			9:  []string{"100", "300"},
			10: []string{"200", "400"},
			11: []string{"100"},
			12: []string{"200", "300"},
			13: []string{"100"},
			14: []string{"200"},
			15: []string{"100", "300"},
			16: []string{"200"},
			17: []string{"100"},
		}
	*/
	new_index.NumericIndex = make(map[string]map[string][]int)
	new_index.NumericIndex["products"] = make(map[string][]int)
	new_index.NumericIndex["products"]["price"] = []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1000, 2000}

	new_index.NumericSort = make(map[string]map[string][]int)
	new_index.NumericSort["products"] = make(map[string][]int)
	new_index.NumericSort["products"]["price"] = []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	new_index.NumericValues = make(map[string]map[string]map[int]int)
	new_index.NumericValues["products"] = make(map[string]map[int]int)
	new_index.NumericValues["products"]["price"] = map[int]int{12: 100, 11: 200, 10: 300, 9: 400, 8: 500, 7: 600, 6: 700, 5: 800, 4: 900, 3: 1000, 2: 1000, 1: 2000}

	new_index.IdsList = make(map[string][]int)
	new_index.IdsList["products"] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	global_index = new_index
	return new_index

}
func InitFromBolt() {
	db, err := bolt.Open(`C:\Gosrc\src\minipure\bolt.db`, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	bucket := "products"
	db.View(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(bucket))

		global_index.TermIndex = make(map[string]map[string]map[string][]int)
		global_index.TermIndex[bucket] = make(map[string]map[string][]int)

		global_index.TermValues = make(map[string]map[string][][]string)
		global_index.TermValues[bucket] = make(map[string][][]string)
		global_index.IdsList = make(map[string][]int)
		global_index.IdsList[bucket] = []int{}
		c := b.Cursor()
		i := 0
		maxcnt := 0
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
			/*buff := make(map[string][]string)
			json.Unmarshal([]byte(v), &buff)*/
			key, _ := strconv.Atoi(string(k))
			if maxcnt < key {
				maxcnt = key
			}

		}
		for k, v := c.First(); k != nil; k, v = c.Next() {
			key, _ := strconv.Atoi(string(k))
			global_index.IdsList[bucket] = append(global_index.IdsList[bucket], key)
			i++

			//fmt.Printf("key=%s, value=%s\n", k, v)
			buff := make(map[string][]string)
			json.Unmarshal([]byte(v), &buff)

			for prop, value := range buff {
				//Заполняем массив TermIndex (инвертированный индекс)
				if _, exists := global_index.TermIndex[bucket][prop]; !exists {
					global_index.TermIndex[bucket][prop] = make(map[string][]int)

				}
				//Заполняем массив TermValues (быстрый кеш возможных значений)
				if _, exists := global_index.TermValues[bucket][prop]; !exists {
					global_index.TermValues[bucket][prop] = make([][]string, maxcnt+1)
				}
				for j := 0; j <= len(value)-1; j++ {
					if _, exists := global_index.TermIndex[bucket][prop][value[j]]; !exists {
						global_index.TermIndex[bucket][prop][value[j]] = []int{}
					}
					global_index.TermIndex[bucket][prop][value[j]] = append(global_index.TermIndex[bucket][prop][value[j]], key)

				}
				global_index.TermValues[bucket][prop][key] = value

			}

		}

		for sk, _ := range global_index.TermIndex[bucket] {
			for sk2, _ := range global_index.TermIndex[bucket][sk] {
				sort.Ints(global_index.TermIndex[bucket][sk][sk2])
			}
		}
		return nil
	})
	//	fmt.Println("%#v", global_index)
	fmt.Println("OK")
	fmt.Println(len(global_index.TermIndex["products"]))
	//global_index = new (In)
}
