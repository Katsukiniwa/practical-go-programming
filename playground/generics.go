package playground

import (
	"fmt"
	"strconv"
)

func TryGenerics[T comparable](x T) T {
	return x
}

func SampleFuncGenerics4[T int | string](x T) T {
	return x
}

func F1[T any](v T) T {
	return v
}

func F2[T interface{}](v T) T {
	return v
}

type OldInterface interface {
	SomeMethod(v interface{})
}

type SomeStruct struct{}

func (s SomeStruct) SomeMethod(v any) {}

func CallSomeStructMethod() {
	var v OldInterface = SomeStruct{}
	fmt.Println(v)

	a := []int{1, 2, 3}
	first := First(a)
	fmt.Println(first)

	Store("a", 1)
	NewStore("a", 1)
	NewStore[int64]("a", 1)
	// NewStore[string]("a", 1) // compile error
	b := NewLoad[string]("a")
	fmt.Println(b)
}

// ⭕: 任意の型のスライスを受け取って、
//
//	その要素型で値を返すことが出来る
func First[S ~[]Elem, Elem any](s S) Elem {
	return s[0]
}

var storage = map[string]interface{}{}

func Store(key string, value any) {
	storage[key] = value
}

func NewStore[T any](key string, value T) {
	storage[key] = value
}

func Load(key string) any {
	return storage[key]
}

func NewLoad[T any](key string) T {
	v := storage[key]
	return v.(T) // `T` 型で型アサーションする
}

type Stringer interface {
	String() string
}

func F(xs []Stringer) []string {
	var result []string
	for _, x := range xs {
		result = append(result, x.String())
	}
	return result
}

// fは型パラメータを持つ関数
// Tは型パラメータ
// インタフェースStringerは、Tに対する型制約として使われている
func FGood[T Stringer](xs []T) []string {
	var result []string
	for _, x := range xs {
		// xは型制約StringerによりString()メソッドが使える
		result = append(result, x.String())
	}
	return result
}

type MyInt int

func (i MyInt) String() string {
	return strconv.Itoa(int(i))
}

type bar struct {
	name string
}

func (bar bar) String() string {
	return bar.name
}

func (bar bar) Call() string {
	return bar.name
}

func Foo() {
	xs := []MyInt{0, 1, 2}
	// F(xs) //  Fは[]Stringerを受け付ける compile error
	FGood(xs)

	barList := []bar{{name: "katsuki"}, {name: "niwa"}}
	FGood(barList)
}
