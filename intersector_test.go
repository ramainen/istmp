package intersector

/*
	16.02.2017
	License: MIT

*/
import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func BenchmarkRegexp(b *testing.B) {
	r, _ := regexp.Compile("p([a-z]+)ch")
	for n := 0; n < b.N; n++ {
		r.MatchString("peach")
	}

}
func BenchmarkMergeHeating(b *testing.B) {
	/* этот бенчмарк идет первым. его задача буквально нагреть прооцессор на ноутбуках, чтобы выключить режим экономии энергии */
	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{5, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17},
		[]int{99, 98, 88, 88, 67, 45, 2},
		[]int{1},
	}
	for n := 0; n < b.N; n++ {
		Merge(myarr...)
		Merge(myarr...)
		Merge(myarr...)
		Merge(myarr...)
	}

	for n := 0; n < b.N; n++ {
		Merge(myarr...)
		Merge(myarr...)
		Merge(myarr...)
		Merge(myarr...)
	}

	for n := 0; n < b.N; n++ {
		Merge(myarr...)
		Merge(myarr...)
		Merge(myarr...)
		Merge(myarr...)
	}
}

func BenchmarkIntersect(b *testing.B) {

	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},

		//[]int{2, 3, 4, 5, 6},
		//	[]int{1},
	}
	for n := 0; n < b.N; n++ {
		IntersectOverMap(myarr...)
	}
}

func BenchmarkFor(b *testing.B) {

	myarr := make([]int, 5000, 5000)
	for i := 0; i <= 4500; i++ {
		myarr[i] = i
	}

	for n := 0; n < b.N; n++ {
		s := 0
		for i := 0; i <= 4500; i++ {
			s += myarr[i]
		}
	}
}

func BenchmarkRange(b *testing.B) {

	myarr := make([]int, 5000, 5000)
	for i := 0; i <= 4500; i++ {
		myarr[i] = i
	}

	for n := 0; n < b.N; n++ {
		s := 0
		for _, i := range myarr {
			s += i
		}
	}
}

func BenchmarkSortedIntersect(b *testing.B) {

	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},

		//[]int{2, 3, 4, 5, 6},
		//	[]int{1},
	}

	for n := 0; n < b.N; n++ {
		Intersect(myarr...)
	}
}

func BenchmarkMerge(b *testing.B) {

	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{5, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17},
		[]int{99, 98, 88, 88, 67, 45, 2},
		[]int{1},
	}
	for n := 0; n < b.N; n++ {
		Merge(myarr...)
	}
}

func BenchmarkMergeFast(b *testing.B) {

	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{5, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17},
		[]int{2, 3, 98, 99, 100},
		[]int{1, 2, 3, 4, 5},
	}
	for n := 0; n < b.N; n++ {
		MergeFast(myarr...)
	}
}
func BenchmarkMergeFast32(b *testing.B) {

	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{5, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17},
		[]int{2, 3, 98, 99, 100},
		[]int{1, 2, 3, 4, 5},
	}
	for n := 0; n < b.N; n++ {
		MergeFast32(myarr...)
	}
}

func BenchmarkMergeFast32bool(b *testing.B) {

	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{5, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17},
		[]int{2, 3, 98, 99, 100},
		[]int{1, 2, 3, 4, 5},
	}
	for n := 0; n < b.N; n++ {
		MergeFast32bool(myarr...)
	}
}
func BenchmarkMergeFast32boolArray(b *testing.B) {

	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{5, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17},
		[]int{2, 3, 98, 99, 100},
		[]int{1, 2, 3, 4, 5},
	}
	for n := 0; n < b.N; n++ {
		MergeFast32boolArray(myarr...)
	}
}
func BenchmarkMergeFast32boolWoID(b *testing.B) {

	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{5, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17},
		[]int{2, 3, 98, 99, 100},
		[]int{1, 2, 3, 4, 5},
	}
	for n := 0; n < b.N; n++ {
		MergeFast32boolWoID(myarr...)
	}
}
func BenchmarkMerge32overDistinct(b *testing.B) {

	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{5, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17},
		[]int{2, 3, 98, 99, 100},
		[]int{1, 2, 3, 4, 5},
	}
	for n := 0; n < b.N; n++ {
		MergeFast32boolWoID(myarr...)
	}

}

func TestMerge(t *testing.T) {

	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{5, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17},
		[]int{2, 3, 98, 99, 100},
		[]int{1, 2, 3, 4, 5},
	}

	res := Merge(myarr...)

	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 98, 99, 100}) {
		t.Error("Eroror testions ", res)
	}

	myarr = [][]int{
		[]int{1},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5},
	}

	res = Merge(myarr...)

	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}) {
		t.Error("Eroror testions ", res)
	}

	myarr = [][]int{
		[]int{},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5},
	}

	res = Merge(myarr...)

	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}) {
		t.Error("Eroror testions ", res)
	}
	myarr = [][]int{

		[]int{99},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5},
	}

	res = Merge(myarr...)

	if !reflect.DeepEqual(res, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 99}) {
		t.Error("Eroror testions ", res)
	}

	myarr = [][]int{}

	res = Merge(myarr...)

	if !reflect.DeepEqual(res, []int{}) {
		t.Error("Eroror testions ", res)
	}

	myarr = [][]int{

		[]int{1},
		[]int{99},
		[]int{2},
	}

	res = Merge(myarr...)

	if !reflect.DeepEqual(res, []int{1, 2, 99}) {
		t.Error("Eroror testions ", res)
	}

	myarr = [][]int{

		[]int{},
		[]int{},
		[]int{},
	}

	res = Merge(myarr...)

	if !reflect.DeepEqual(res, []int{}) {
		t.Error("Eroror testions ", res)
	}

}
func BenchmarkMergeSorted(b *testing.B) {

	myarr := [][]int{
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		[]int{5, 6, 7, 8, 9, 11, 13, 14, 15, 16, 17},
		[]int{2, 3, 98, 99, 100},
		[]int{1, 2, 3, 4, 5},
	}
	for n := 0; n < b.N; n++ {
		Merge(myarr...)
	}

}

/********************/

func BenchmarkPrepare(b *testing.B) {
	NewEmptyIndex()
	InitFromBolt()
}

func BenchmarkSpeed(b *testing.B) {

	/*NewIndex()
	for n := 0; n < b.N; n++ {
		search := NewSearch("products")

		search.AddTermFacet("width", []string{})

		search.AddTermFacet("color", []string{"black", "blue"})

		search.AddNumericFilter("price", 400, 1000)

		search.Run()
	}*/

	for n := 0; n < b.N; n++ {

		search := NewSearch("products")

		search.AddTermFilter("category_id", []string{"259"}) //смесители
		search.AddTermFilter("field83", []string{"230", "275"})
		search.AddTermFacet("field2", []string{})
		search.AddTermFacet("field20", []string{})
		search.AddTermFacet("field151", []string{})
		search.AddTermFacet("field17", []string{})
		search.AddTermFacet("field141", []string{})
		search.AddTermFacet("field29", []string{})
		search.AddTermFacet("field88", []string{})
		search.AddTermFacet("field7", []string{})
		search.AddTermFacet("field155", []string{})
		search.AddTermFacet("field156", []string{})
		search.AddTermFacet("field86", []string{})
		search.AddTermFacet("field187", []string{})
		search.AddTermFacet("field226", []string{})
		search.AddTermFacet("field80", []string{})
		search.AddTermFacet("field227", []string{})
		search.AddTermFacet("field228", []string{})
		search.AddTermFacet("field229", []string{})
		search.AddTermFacet("field231", []string{})
		search.AddTermFacet("field232", []string{})
		search.AddTermFacet("field233", []string{})
		search.AddTermFacet("field239", []string{})
		search.AddTermFacet("field238", []string{})
		search.AddTermFacet("field379", []string{})
		search.AddTermFacet("field79", []string{})
		search.AddTermFacet("field154", []string{})
		search.Run()

		/*
			search := NewSearch("products")
			search.AddTermFilter("category_id", []string{"42"}) //смесители
			search.AddTermFacet("field372", []string{})
				search.Run()
		*/

	}

}

func BenchmarkSpeedWOFacet(b *testing.B) {

	for n := 0; n < b.N; n++ {

		search := NewSearch("products")

		search.AddTermFilter("category_id", []string{"259"}) //смесители
		search.AddTermFilter("field83", []string{"230", "275"})

		/*
			search := NewSearch("products")
			search.AddTermFilter("category_id", []string{"42"}) //смесители
			search.AddTermFacet("field372", []string{})
		*/
		search.Run()

	}

}

/********************/
func BenchmarkSpeed2(b *testing.B) {

	/*NewIndex()
	for n := 0; n < b.N; n++ {
		search := NewSearch("products")

		search.AddTermFacet("width", []string{})

		search.AddTermFacet("color", []string{"black", "blue"})

		search.AddNumericFilter("price", 400, 1000)

		search.Run()
	}*/

	search := NewSearch("products")

	search.AddTermFilter("category_id", []string{"259"}) //смесители
	search.AddTermFilter("field83", []string{"230", "275"})
	search.AddTermFacet("field2", []string{})
	search.AddTermFacet("field20", []string{})
	search.AddTermFacet("field151", []string{})
	search.AddTermFacet("field17", []string{})
	search.AddTermFacet("field141", []string{})
	search.AddTermFacet("field29", []string{})
	search.AddTermFacet("field88", []string{})
	search.AddTermFacet("field7", []string{})
	search.AddTermFacet("field155", []string{})
	search.AddTermFacet("field156", []string{})
	search.AddTermFacet("field86", []string{})
	search.AddTermFacet("field187", []string{})
	search.AddTermFacet("field226", []string{})
	search.AddTermFacet("field80", []string{})
	search.AddTermFacet("field227", []string{})
	search.AddTermFacet("field228", []string{})
	search.AddTermFacet("field229", []string{})
	search.AddTermFacet("field231", []string{})
	search.AddTermFacet("field232", []string{})
	search.AddTermFacet("field233", []string{})
	search.AddTermFacet("field239", []string{})
	search.AddTermFacet("field238", []string{})
	search.AddTermFacet("field379", []string{})
	search.AddTermFacet("field79", []string{})
	search.AddTermFacet("field154", []string{})

	/*search.AddTermFilter("category_id", []string{"42"}) //смесители
	search.AddTermFacet("field372", []string{})
	*/
	search.Run()

	fmt.Println("%#v", search)
}
