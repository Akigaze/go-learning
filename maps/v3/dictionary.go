package maps

// Dictionary 声明一个map使用 map[keyType]valueType, 并将其定义为一种类型
type Dictionary map[string]string

type DictionaryError string

// 通过实现 Error 方法来集成error接口
func (d DictionaryError) Error() string {
	return string(d)
}

var (
	ErrNotFound   = DictionaryError("could not find the word you were looking for")
	ErrWordExists = DictionaryError("cannot add word because it already exists")
)

// Search map 是引用类型的数据，所以作为参数是传递的就是引用本身
func (d Dictionary) Search(word string) (string, error) {
	// map[key] 会返回两个参数，第一个是key对于的值，第一个是key在map中是否存在定义bool
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	// GO 的switch case 似乎是不用break的
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
