// Package util contains utility functions
package util

import "errors"

func StoreImage(image []byte) (sha string, overridden bool, err error) {
	return "", false, errors.New("Not implemented")
}

func RetrieveImage(sha string) (image []byte, err error) {
    return []byte{0}, errors.New("Not implemented")
}
