//go:build wireinject
// +build wireinject

package playground

import "github.com/google/wire"

func initMovieRepository(fileName string) *MovieRepository {
	wire.Build(
		NewMovieRepository,
		NewColonDelimitedMovieFinder,
	)

	return nil
}
