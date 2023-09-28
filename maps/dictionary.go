package maps

type Dictionary map[string] string

var (
	ErrorNotFound = DictionaryErr("not found")
	ErrWordExists = DictionaryErr("word already exists")
)

type DictionaryErr string

// https://dave.cheney.net/2016/04/07/constant-errors
func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]

	if !ok {
		return "", ErrorNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)
	
	switch err {
		case ErrorNotFound:
			d[key] = value
		case nil:
			return ErrWordExists
		default:
			return err
	}

	return nil
}