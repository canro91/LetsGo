package maps

type DictionaryErr string
type Dictionary map[string]string

const (
	ErrKeyNotFound = DictionaryErr("Key not found")
	ErrKeyAlreadyAddedd = DictionaryErr("Key already added")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", 	ErrKeyNotFound
	}
	
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotFound:
		d[key] = value
		return nil

	case nil:
		return ErrKeyAlreadyAddedd

	default:
		return err
	}
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	if err == nil {
		d[key] = value
		return nil
	}

	return err
}
func (d Dictionary) Delete(key string) {
	delete(d, key)
}