//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func initMovieRepository(fileName string) *MovieRepository {
	wire.Build(
		NewMovieRepository,
		NewColonDelimitedMovieFinder,
	)

	return nil
}
