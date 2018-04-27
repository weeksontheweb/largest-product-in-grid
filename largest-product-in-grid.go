package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var grid []int

func main() {

	if len(os.Args) == 2 {
		itemsInProduct := os.Args[1]

		if _, err := strconv.Atoi(itemsInProduct); err == nil {
			fmt.Println("Argument = ", itemsInProduct)

			itemsInProductNumber, _ := strconv.Atoi(itemsInProduct)
			x, y, err := loadGrid("grid.txt", itemsInProductNumber)

			if err == nil {
				fmt.Printf("%d %d\n", x, y)

				fmt.Println("Point element = ", grid[2], grid[6])
			}

		} else {
			fmt.Println("Required argument is not numeric.")
		}
	} else {
		//Check to see if the number is numeric.

		fmt.Println("Please enter the number of product numbers.")
	}
}

func loadGrid(gridFile string, countInProduct int) (int, int, error) {

	//Return an error if the grid is not uniform.
	//Could be 20 x 10, 15 x 4 etc.
	//Error if
	var spacePosition int
	var firstRowElements int
	var otherRowElements int
	var firstRowCounted bool
	var elementString string
	//var elementNumber int
	var rowCount int

	file, err := os.Open(gridFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())

		otherRowElements = 0
		rowCount++

		lineInFile := scanner.Text()

		spacePosition = strings.Index(lineInFile, " ")

		for spacePosition >= 0 {
			//fmt.Println("First position = ", spacePosition)

			elementString = lineInFile[0:spacePosition]
			fmt.Println("element = ", elementString)

			if _, err := strconv.Atoi(elementString); err == nil {

				elementNumber, _ := strconv.Atoi(elementString)

				fmt.Println("elementNumber = ", elementNumber)
				grid = append(grid, elementNumber)

				fmt.Println("Length = ", len(grid))

				lineInFile = lineInFile[spacePosition+1:]
				fmt.Println("New lineInFile =", lineInFile)
				spacePosition = strings.Index(lineInFile, " ")
				//fmt.Println("Second position = ", spacePosition)

				if !firstRowCounted {
					firstRowElements++
				} else {
					otherRowElements++
				}
			} else {
				fmt.Printf("Non numeric element (%s) found.\n", elementString)
			}
		}

		if !firstRowCounted {
			firstRowCounted = true
		} else {
			if firstRowElements != otherRowElements {
				fmt.Println("Row element mismatch.")
			}
		}
	}

	return firstRowElements, rowCount, nil

}
