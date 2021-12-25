package main

import (
	interpreter "chefgi/interpreter"
	utils "chefgi/utils"
	"fmt"
	"os"
)

func main() {
	filePath, err := utils.GetFilePath()

	if err != nil {
		panic("Could not get source file! :(")
	}

	rawSource, err := os.ReadFile(filePath)

	if err != nil {
		panic("Could not read the file passed as an argument, does it exist?")
	}

	source := string(rawSource)
	tokens := interpreter.Tokenize(source)
	message := interpreter.ParseTokensToString(tokens)

	fmt.Println(message)
}

