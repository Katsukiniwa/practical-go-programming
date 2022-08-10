package main

import (
	"reflect"
	"testing"
)

func TestFindAll(t *testing.T) {
	n := NewColonDelimitedMovieFinder("foo")
	r := NewMovieRepository(n)
	movies := r.finder.FindAll()
	want := []Movie{{Director: "foo"}, {Director: "bar"}}
	if !reflect.DeepEqual(movies, want) {
		t.Errorf("got %v want %v", movies, want)
	}
}
