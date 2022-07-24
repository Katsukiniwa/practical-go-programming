package main

type Movie struct {
	Director string
}

type MovieRepository struct {
	finder MovieFinder
}

type MovieFinder interface {
	FindAll() []Movie
}

type ColonDelimitedMovieFinder struct {
	file string
}

func (ml *MovieRepository) MoviesDirectedBy(director string) []Movie {
	movieList := ml.finder.FindAll()
	result := make([]Movie, 0, len(movieList))

	for _, m := range movieList {
		if director == m.Director {
			result = append(result, m)
		}
	}

	return result
}

func NewMovieRepository(finder MovieFinder) *MovieRepository {
	return &MovieRepository{
		finder: finder,
	}
}

func NewColonDelimitedMovieFinder(arg string) MovieFinder {
	return ColonDelimitedMovieFinder{
		file: arg,
	}
}

func (c ColonDelimitedMovieFinder) FindAll() []Movie {
	movies := []Movie{{Director: "foo"}, {Director: "bar"}}

	return movies
}
