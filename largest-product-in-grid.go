package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type gridElement struct {
	value            int
	northProduct     int
	southProduct     int
	eastProduct      int
	westProduct      int
	southEastProduct int
	northEastProduct int
	southWestProduct int
	northWestProduct int
}

var grid = make(map[int]gridElement)

func main() {

	if len(os.Args) == 2 {
		itemsInProduct := os.Args[1]

		if _, err := strconv.Atoi(itemsInProduct); err == nil {

			itemsInProductNumber, _ := strconv.Atoi(itemsInProduct)
			columns, rows, err := loadGrid("grid.txt", itemsInProductNumber)

			if err == nil {
				workOutLargestProduct(itemsInProductNumber, columns, rows)
			}

		} else {
			fmt.Println("Required argument is not numeric.")
		}
	} else {
		fmt.Println("Please enter the number of product numbers.")
	}
}

//Read the grid in the grid.txt file.
//Check that the grid is well formed.
//For a grid to be well formed, it must:
//		1. Have the same number of columns in all rows.
//		2. Be entirely composed of numbers.
//		3. The grid rows and columns must be at least as big
//		   as the number of elements making up the product.
func loadGrid(gridFile string, countInProduct int) (int, int, error) {

	var spacePosition int
	var firstRowElements int
	var otherRowElements int
	var firstRowCounted bool
	var elementString string
	var rowCount int

	file, err := os.Open(gridFile)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		otherRowElements = 0
		rowCount++

		lineInFile := scanner.Text()

		spacePosition = strings.Index(lineInFile, " ")

		for spacePosition >= 0 {

			elementString = lineInFile[0:spacePosition]

			if _, err := strconv.Atoi(elementString); err == nil {

				elementNumber, _ := strconv.Atoi(elementString)

				appendGridElementToGrid(elementNumber)

				lineInFile = lineInFile[spacePosition+1:]
				spacePosition = strings.Index(lineInFile, " ")

				if !firstRowCounted {
					firstRowElements++
				} else {
					otherRowElements++
				}
			} else {
				fmt.Printf("Non numeric element (%s) found.\n", elementString)
			}
		}

		if _, err := strconv.Atoi(lineInFile); err == nil {
			elementNumber, _ := strconv.Atoi(lineInFile)
			appendGridElementToGrid(elementNumber)

			//Count the number of columns.
			if !firstRowCounted {
				firstRowElements++
			} else {
				otherRowElements++
			}
		} else {
			fmt.Printf("Non numeric element (%s) found.\n", lineInFile)
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

func appendGridElementToGrid(elementValue int) {

	var newGridElement gridElement

	newGridElement.value = elementValue

	grid[len(grid)] = newGridElement
}

func workOutLargestProduct(elementsInProduct int, columns int, rows int) {

	/*
		Look:

		North, if row >= elements in product.
		South, if row <= rows - elements in product.
		East, if column <= columns - elements in product.
		West, if column >= elements in product.
		South East, if row <= rows - elements in product and column <= columns - elements in product.
		North East, if row >= elements in product and column <= columns - elements in product.
		South West, if row <= rows - elements in product and column >= elements in product.
		North West, if row >= elements in product and column >= elements in product.
	*/

	var row int
	var column int
	var largestProduct int
	var largestProductRow int
	var largestProductColumn int
	var largestProductDirection string

	//Calculate the product for each number in the grid.
	//Save the number into the appropriate place in the map.
	//Store the largest product details.
	for i := 0; i < len(grid); i++ {

		row, column = convertToRowColumns(i, columns)

		if row >= elementsInProduct {
			//Look North.
			newgrid := grid[i]
			newgrid.northProduct = calculateProduct(elementsInProduct, row, column, columns, "n")
			grid[i] = newgrid

			if newgrid.northProduct > largestProduct {
				largestProduct = newgrid.northProduct
				largestProductRow = row
				largestProductColumn = column
				largestProductDirection = "North"
			}
		}

		if row <= rows-elementsInProduct {
			//Look South.
			newgrid := grid[i]
			newgrid.southProduct = calculateProduct(elementsInProduct, row, column, columns, "s")
			grid[i] = newgrid

			if newgrid.southProduct > largestProduct {
				largestProduct = newgrid.southProduct
				largestProductRow = row
				largestProductColumn = column
				largestProductDirection = "South"
			}
		}

		if column <= columns-elementsInProduct {
			//Look East.
			newgrid := grid[i]
			newgrid.eastProduct = calculateProduct(elementsInProduct, row, column, columns, "e")
			grid[i] = newgrid

			if newgrid.eastProduct > largestProduct {
				largestProduct = newgrid.eastProduct
				largestProductRow = row
				largestProductColumn = column
				largestProductDirection = "East"
			}
		}

		if column >= elementsInProduct {
			//Look West.
			newgrid := grid[i]
			newgrid.westProduct = calculateProduct(elementsInProduct, row, column, columns, "w")
			grid[i] = newgrid

			if newgrid.westProduct > largestProduct {
				largestProduct = newgrid.westProduct
				largestProductRow = row
				largestProductColumn = column
				largestProductDirection = "West"
			}
		}

		if row <= rows-elementsInProduct && column <= columns-elementsInProduct {
			//Look South East.
			newgrid := grid[i]
			newgrid.southEastProduct = calculateProduct(elementsInProduct, row, column, columns, "se")
			grid[i] = newgrid

			if newgrid.southEastProduct > largestProduct {
				largestProduct = newgrid.southEastProduct
				largestProductRow = row
				largestProductColumn = column
				largestProductDirection = "South East"
			}
		}

		if row >= elementsInProduct && column <= columns-elementsInProduct {
			//Look North East.
			newgrid := grid[i]
			newgrid.northEastProduct = calculateProduct(elementsInProduct, row, column, columns, "ne")
			grid[i] = newgrid

			if newgrid.northEastProduct > largestProduct {
				largestProduct = newgrid.northEastProduct
				largestProductRow = row
				largestProductColumn = column
				largestProductDirection = "North East"
			}
		}

		if row <= rows-elementsInProduct && column >= elementsInProduct {
			//Look South West.
			newgrid := grid[i]
			newgrid.southWestProduct = calculateProduct(elementsInProduct, row, column, columns, "sw")
			grid[i] = newgrid

			if newgrid.southWestProduct > largestProduct {
				largestProduct = newgrid.southWestProduct
				largestProductRow = row
				largestProductColumn = column
				largestProductDirection = "South West"
			}
		}

		if row >= elementsInProduct && column >= elementsInProduct {
			//Look North West.
			newgrid := grid[i]
			newgrid.northWestProduct = calculateProduct(elementsInProduct, row, column, columns, "nw")
			grid[i] = newgrid

			if newgrid.northWestProduct > largestProduct {
				largestProduct = newgrid.northWestProduct
				largestProductRow = row
				largestProductColumn = column
				largestProductDirection = "North West"
			}
		}
	}

	//Output result.
	fmt.Printf("Largest product = %d, row = %d, column = %d, direction = %s\n", largestProduct, largestProductRow, largestProductColumn, largestProductDirection)
}

func convertToRowColumns(index int, columns int) (int, int) {
	return (index / columns) + 1, (index % columns) + 1
}

func convertToIndex(row int, column int, columns int) int {
	return ((row - 1) * columns) + (column - 1)
}

//Scan the adjacent elements in the grid required to make up the product.
func calculateProduct(elementsInProduct int, row int, column int, columns int, direction string) int {

	var product int

	product = 1

	for i := 0; i < elementsInProduct; i++ {
		switch direction {
		case "n":
			//Going North.
			product *= grid[convertToIndex(row-i, column, columns)].value
		case "s":
			//Going South.
			product *= grid[convertToIndex(row+i, column, columns)].value
		case "e":
			//Going East.
			product *= grid[convertToIndex(row, column+i, columns)].value
		case "w":
			//Going West.
			product *= grid[convertToIndex(row, column-i, columns)].value
		case "ne":
			//Going North East.
			product *= grid[convertToIndex(row-i, column+i, columns)].value
		case "se":
			//Going South East.
			product *= grid[convertToIndex(row+i, column+i, columns)].value
		case "nw":
			//Going North West.
			product *= grid[convertToIndex(row-i, column-i, columns)].value
		case "sw":
			//Going South West.
			product *= grid[convertToIndex(row+i, column-i, columns)].value
		}
	}

	return product
}
