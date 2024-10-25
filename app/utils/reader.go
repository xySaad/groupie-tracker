package utils

import (
	"errors"
	"io"
)

// Read reads from the provided io.ReadCloser and returns all the data as a byte slice.
// It takes a buffer size to control the amount of data read in each iteration.
func Read(readCloser io.ReadCloser, bufferSize int) ([]byte, error) {
	if bufferSize <= 0 {
		return nil, errors.New("bufferSize must be greater than zero")
	}

	buffer := make([]byte, bufferSize)
	var data []byte

	for {
		n, err := readCloser.Read(buffer)

		if err == io.EOF {
			data = append(data, buffer[:n]...)
			break
		} else if err != nil {
			return nil, err
		}

		data = append(data, buffer[:n]...)
	}

	return data, nil
}
