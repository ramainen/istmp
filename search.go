package intersector

type TermQuery struct {
	Key    string
	Values []string
}

type NumericQuery struct {
	Key     string
	Minimum int
	Maximum int
}

type Search struct {
	Table          string                    /* имя таблицы (бакета)*/
	TermFilters    []TermQuery               /* массив фильтров перед фасетами*/
	TermFacets     []TermQuery               /* массив фасетов, по которым надо искать*/
	NumericFilters []NumericQuery            /* массив чисовых запросов */
	FilterResult   []int                     /* промежуточное значение списка id, до применения фасетов */
	Result         []int                     /* итоговый результат */
	FacetsResults  map[string]map[string]int /* результаты фасетов (фасет:значение:количество), например [цвет][красный]=50  */

	SortField                string /* не используется (поле для сортировки) */
	InvertedSort             bool   /* не используется (инвертируемая сортировка) */
	FacetResultCache         []int  /* кеш результатов фасетов*/
	FacetResultCachePrepared bool   /* кеш результатов фасетов*/
}

func NewSearch(table string) *Search {
	new_search := new(Search)
	new_search.Table = table
	new_search.TermFilters = make([]TermQuery, 0, 50)
	new_search.TermFacets = make([]TermQuery, 0, 50)
	new_search.NumericFilters = make([]NumericQuery, 0, 50)

	new_search.FacetsResults = make(map[string]map[string]int)

	return new_search
}

func (s *Search) AddTermFilter(property string, values []string) {
	s.TermFilters = append(s.TermFilters, TermQuery{Key: property, Values: values})
}

func (s *Search) AddNumericFilter(property string, minimum int, maximum int) {
	s.NumericFilters = append(s.NumericFilters, NumericQuery{Key: property, Minimum: minimum, Maximum: maximum})
}

func (s *Search) AddTermFacet(property string, values []string) {
	s.TermFacets = append(s.TermFacets, TermQuery{Key: property, Values: values})
}

/*
	AggregateFacet Обработка одного фасета, при условии, что данные уже получены
*/
func (s *Search) AggregateFacet(facet *TermQuery) {
	s.FacetsResults[facet.Key] = make(map[string]int)
	//Делаем выборку по id, исключая текущий фасет
	//is_empty := false
	prepared_ids := []int{}
	if len(facet.Values) == 0 && s.FacetResultCachePrepared {
		prepared_ids = s.FacetResultCache
	} else {
		tmp_intersection := make([][]int, 0, len(s.TermFacets)+1)
		tmp_intersection = append(tmp_intersection, s.FilterResult)
		for _, i := range s.TermFacets {
			if i.Key != facet.Key {
				switch len(i.Values) {
				case 0:
					//Ничего не выбрано. Результат будет пустой.
					//Примечаение:так выбираются "все варианты", бдет выбрано все
				case 1:
					//Один вариант, добавляем к пересечению

					tmp_intersection = append(tmp_intersection, global_index.TermIndex[s.Table][i.Key][i.Values[0]])

				default:
					// два или более массива, предварительно объединяем через ИЛИ

					tmpSlices := make([][]int, 0, len(i.Values))
					for _, j := range i.Values {
						tmpSlices = append(tmpSlices, global_index.TermIndex[s.Table][i.Key][j])

					}

					tmp_intersection = append(tmp_intersection, Merge(tmpSlices...))
				}
			}
		}

		prepared_ids = Intersect(tmp_intersection...)
		if len(facet.Values) == 0 {
			s.FacetResultCachePrepared = true
			s.FacetResultCache = prepared_ids
		}
	}
	//fmt.Println(prepared_ids)
	/* в tmpFacetsResults подсчитывается прямым поиском и пересчетом все значения из свойств, сохраненных в global_index.TermValues */
	tmpFacetsResults := make(map[string]int)
	ln := len(prepared_ids)
	for j := 0; j <= ln-1; j++ {
		i := prepared_ids[j]
		//if _, is_exists := global_index.TermValues[s.Table][facet.Key][i]; is_exists {
		for _, val := range global_index.TermValues[s.Table][facet.Key][i] {
			if _, is_existstmpFacetsResults := tmpFacetsResults[val]; is_existstmpFacetsResults {
				tmpFacetsResults[val]++
			} else {
				tmpFacetsResults[val] = 1
			}
		}
		//}
	}
	s.FacetsResults[facet.Key] = tmpFacetsResults
}

/*
Обработка одного фасета, при условии, что данные уже получены.
Получение методом пересечения массивов
*/
func (s *Search) AggregateFacetOverIntersect(facet *TermQuery) {
	s.FacetsResults[facet.Key] = make(map[string]int)
	//Делаем выборку по id, исключая текущий фасет
	tmp_intersection := make([][]int, 0, len(s.TermFacets)+1)
	tmp_intersection = append(tmp_intersection, s.FilterResult)
	for _, i := range s.TermFacets {
		if i.Key != facet.Key {
			switch len(i.Values) {
			case 0:
				//Ничего не выбрано. Результат будет пустой.
				//Примечаение:так выбираются "все варианты", бдет выбрано все
			case 1:
				//Один вариант, добавляем к пересечению

				tmp_intersection = append(tmp_intersection, global_index.TermIndex[s.Table][i.Key][i.Values[0]])

			default:
				// два или более массива, предварительно объединяем через ИЛИ

				tmpSlices := make([][]int, 0, len(i.Values))
				for _, j := range i.Values {
					tmpSlices = append(tmpSlices, global_index.TermIndex[s.Table][i.Key][j])

				}

				tmp_intersection = append(tmp_intersection, Merge(tmpSlices...))
			}
		}
	}
	prepared_ids := Intersect(tmp_intersection...)
	/* в tmpFacetsResults подсчитывается прямым поиском и пересчетом все значения из свойств, сохраненных в global_index.TermValues */
	tmpFacetsResults := make(map[string]int)
	for k, _ := range global_index.TermIndex[s.Table][facet.Key] {
		faceted_values := global_index.TermIndex[s.Table][facet.Key][k]
		cnt := Intersect(prepared_ids, faceted_values)
		if len(cnt) > 0 {
			tmpFacetsResults[k] = len(cnt)
		}
	}

	s.FacetsResults[facet.Key] = tmpFacetsResults
}

func (s *Search) Run() {
	/*
		TODO:
		✓ фильтр по числам
		✓ аггрегация по фасетам (числа в скобках)
		фасеты числовые
		загрзука из БОЛТ
		аггрегации по двум значениям
	*/
	if len(s.TermFilters) > 0 {
		first_intersect := make([][]int, 0, 50)
		//first_intersect = append(first_intersect, global_index.IdsList[s.Table])
		//проходим по фильтрам, получаем массив пересечений которые надо сделать в обязательном порядке
		for _, i := range s.TermFilters {
			//fmt.Println("intersect_with", i, i.Values)
			switch len(i.Values) {
			case 0:
				//Ничего не выбрано. Результат будет пустой. Нечего было в фильтре пустой массив передавать.
				first_intersect = append(first_intersect, []int{})
			case 1:
				//Один вариант, добавляем к пересечению

				first_intersect = append(first_intersect, global_index.TermIndex[s.Table][i.Key][i.Values[0]])

			default:
				// два или более массива, предварительно объединяем через ИЛИ

				tmpSlices := make([][]int, 0, len(i.Values))
				for _, j := range i.Values {
					tmpSlices = append(tmpSlices, global_index.TermIndex[s.Table][i.Key][j])

				}

				first_intersect = append(first_intersect, Merge(tmpSlices...))
			}

		}
		s.FilterResult = Intersect(first_intersect...)
	} else {
		s.FilterResult = global_index.IdsList[s.Table]

	}
	//fmt.Println("FilterResult:", s.FilterResult)

	//Этап второй: фильтруем по NumericFilters
	if len(s.NumericFilters) > 0 {
		numeric_intersect := make([][]int, 0, 50)
		numeric_intersect = append(numeric_intersect, s.FilterResult)
		for _, i := range s.NumericFilters {
			numeric_intersect = append(numeric_intersect, SearchBetween(global_index.NumericIndex[s.Table][i.Key], global_index.NumericSort[s.Table][i.Key], i.Minimum, i.Maximum))
		}

		s.FilterResult = Intersect(numeric_intersect...)

	}
	// Превариетльный фильтры готовы
	//fmt.Println("FilterResult with numeric:", s.FilterResult)

	//ВТорой этап фильтры по фасетам.

	if len(s.TermFacets) > 0 {
		second_intersect := make([][]int, 0, 50)
		second_intersect = append(second_intersect, s.FilterResult)
		//проходим по фильтрам, получаем массив пересечений которые надо сделать в обязательном порядке
		for _, i := range s.TermFacets {
			//fmt.Println("Facet intersect_with", i, i.Values)
			switch len(i.Values) {
			case 0:
				//Ничего не выбрано. Результат будет пустой.
				//second_intersect = append(second_intersect, []int{})
				//Здесь нет ничего!
			case 1:
				//Один вариант, добавляем к пересечению

				second_intersect = append(second_intersect, global_index.TermIndex[s.Table][i.Key][i.Values[0]])

			default:
				// два или более массива, предварительно объединяем через ИЛИ

				tmpFSlices := make([][]int, 0, len(i.Values))
				for _, j := range i.Values {
					tmpFSlices = append(tmpFSlices, global_index.TermIndex[s.Table][i.Key][j])

				}

				second_intersect = append(second_intersect, Merge(tmpFSlices...))
			}

		}
		s.Result = Intersect(second_intersect...)
	} else {
		s.Result = s.FilterResult

	}
	//fmt.Println("Result:", s.Result)

	//Аггрегации по фасетам
	if len(s.TermFacets) > 0 {
		for _, i := range s.TermFacets {
			s.AggregateFacet(&i)
		}
	}

}
