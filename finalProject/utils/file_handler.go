package service

import "os"

func ReadFromFile() ([]byte, error) {
	file, err := os.Open("tickets.csv")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	file.Read(buffer)
	return buffer, nil
}

func WriteToFile(data []byte) error {
	file, err := os.Create("tickets.json")
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(data)
	return nil
}
