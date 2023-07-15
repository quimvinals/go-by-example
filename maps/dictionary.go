package maps

type Dictionary map[string] string

func Search(d Dictionary, key string) string {
	return d[key]
}