package maps

import "errors"

var errNotFound = errors.New("could not find word you were looking for")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errNotFound
	}

	return definition, nil
}
