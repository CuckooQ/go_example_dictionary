package dict

import "errors"

// Dictionary
type Dictionary map[string]string

// error
var (
	errNotFound   error = errors.New("Not Found")
	errWordExists error = errors.New("That word already exists")
	errCantUpdate error = errors.New("Can't update non-existing word")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, isExists := d[word]

	if isExists {
		return value, nil
	}

	return "", errNotFound
}

// Add a word to the Dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	if err == errNotFound {
		d[word] = def
		return nil
	}

	return errWordExists
}

// Update a word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	if err == nil {
		d[word] = def
		return nil
	}

	return errCantUpdate
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
