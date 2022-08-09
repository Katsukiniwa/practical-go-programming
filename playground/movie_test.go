package main

import (
	"testing"
	"reflect"
)

func TestFindAll(t *testing.T) {
	r := NewMovieRepository()
	movies := r.finder.FindAll()
    want := []Movie{{Director: "foo"}, {Director: "bar"}}
	if !reflect.DeepEqual(movies, want) {
		t.Errorf("got %v want %v", movies, want)
	}
}
