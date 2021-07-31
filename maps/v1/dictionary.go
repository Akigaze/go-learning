package maps

// Dictionary 声明一个map使用 map[keyType]valueType, 并将其定义为一种类型
type Dictionary map[string]string

func (d Dictionary) Search(word string) string {
	return d[word]
}
