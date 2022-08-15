package playground

type Dictionary map[string]string

func Search(dictionary Dictionary, word string) string {
	return dictionary[word]
}
