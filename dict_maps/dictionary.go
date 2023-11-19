package dictmaps

import (
	"errors"
)

var ErrNotFound = errors.New("could not find the word you were looking for")
type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
  definition, ok := d[key]

  if !ok {
    return "", ErrNotFound
  }
  return definition, nil
}

func (d Dictionary) Add(key, value string) (string, error) {
  d[key] = value
  return value, nil
}