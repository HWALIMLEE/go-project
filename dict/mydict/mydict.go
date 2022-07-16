package mydict

import "errors"

//Dictionary type
//type 은 method를 가질 수 있음
type Dictionary map[string]string //Dictionary 는 여기서 alias(별명) 같은 것

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Can't update non-existing word")
	errCantDelete = errors.New("Can't delete non-existing word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	// if err == errNotFound {
	// 	d[word] = def
	// } else if err == nil {
	// 	return errWordExists
	// }
	// return nil
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, new_def string) error {
	_, err := d.Search(word)
	// if err == nil { // 단어 존재
	// 	d[word] = new_def
	// } else {
	// 	return errNotFound
	// }
	// return nil
	switch err {
	case nil:
		d[word] = new_def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}

/*
type Money int

...

Money(1) -> 1
*/
