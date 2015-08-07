package tpl

import (
	"errors"
	"github.com/spf13/hugo/tpl"
	"github.com/spf13/viper"
	"io/ioutil"
)

func init() {
	viper.Reset()
	viper.Set("BaseURL", "http://fuzz/")
	viper.Set("WorkDir", "/tmp/")
}

func Fuzz(data []byte) int {

	t := tpl.New()
	err := t.AddTemplate("fuzz", string(data))

	if err != nil {
		return 0
	}

	d := &Data{
		A: 42,
		B: "foo",
		C: []int{1, 2, 3},
		D: map[int]string{1: "foo", 2: "bar"},
		E: Data1{42, "foo"},
		F: []string{"a", "b", "c"},
		G: []string{"a", "b", "c", "d", "e"},
		H: "a,b,c,d,e,f",
	}

	err = t.ExecuteTemplate(ioutil.Discard, "fuzz", d)

	if err != nil {
		return 0
	}

	return 1
}

type Data struct {
	A int
	B string
	C []int
	D map[int]string
	E Data1
	F []string
	G []string
	H string
}

type Data1 struct {
	A int
	B string
}

func (Data1) Q() string {
	return "foo"
}

func (Data1) W() (string, error) {
	return "foo", nil
}

func (Data1) E() (string, error) {
	return "foo", errors.New("Data.E error")
}

func (Data1) R(v int) (string, error) {
	return "foo", nil
}

func (Data1) T(s string) (string, error) {
	return s, nil
}
