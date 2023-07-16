package maps

import "errors"

type Dictionary map[string] string

var ErrorNotFound = errors.New("not found") 

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}