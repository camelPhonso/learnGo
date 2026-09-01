package db

import (
	"os"
	"bufio"
	"fmt"
	"encoding/json"
)

func Read(fileName string) ([]Movie, error) {
	var result []Movie
	
	content, err := os.Open(fileName)
	if err != nil {
		return nil, Error{ErrBadFile}
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		var newMovie Movie

		line := []byte(fmt.Sprintf("{%s}", scanner.Text()))
		err := json.Unmarshal(line, &newMovie)
		if err != nil {
			return nil, Error{ErrBadFile}
		}
		
		result = append(result, newMovie)
	}
	
	return result, nil
}