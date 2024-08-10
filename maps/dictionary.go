package maps

type Dictionary map[string]string

type DictionaryErr string

const (
	ErrNotFound     = DictionaryErr("word not found")
	ErrExists       = DictionaryErr("word already exists")
	ErrDoesNotExist = DictionaryErr("word does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	def, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return def, nil
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = def
	case nil:
		return ErrExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrDoesNotExist
	case nil:
		d[word] = def
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
