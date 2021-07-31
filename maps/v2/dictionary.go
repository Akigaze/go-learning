package maps

import "errors"

// Dictionary 声明一个map使用 map[keyType]valueType, 并将其定义为一种类型
type Dictionary map[string]string

var (
	ErrNotFound = errors.New("could not find the word you were looking for")
)

func (d Dictionary) Search(word string) (string, error) {
	// map[key] 会返回两个参数，第一个是key对于的值，第一个是key在map中是否存在定义bool
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
