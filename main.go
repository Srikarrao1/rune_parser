package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

func ParseOP_RETURN(opReturnData string) (string, error) {
	if len(opReturnData) < 2 || opReturnData[:2] != "6a" {
		return "", errors.New("not an OP_RETURN data")
	}

	//Extract the data part
	dataHex := opReturnData[2:]
	data, err := hex.DecodeString(dataHex)
	if err != nil {
		return "", err
	}

	return string(data), nil
} 

func main() {
	opReturnData := "6a28f7377a2065897e429287682bd2ed67ed7d0f5ebfc35cf4b1575756c3bebf0000b84002004b4d4400" //example OP_RETURN data
	decodedData, err := ParseOP_RETURN(opReturnData)
	if err != nil {
		log.Fatalf("Failed to parse OP_RETURN data: %v", err)
	}
	fmt.Println("Decoded runes data:", decodedData)
}