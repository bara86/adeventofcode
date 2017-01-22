package common

import "os"

func OpenFile(fileName string) *os.File {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	return file
}
