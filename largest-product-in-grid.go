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
			fmt.Println("Argument = ", itemsInProduct)

			itemsInProductNumber, _ := strconv.Atoi(itemsInProduct)
			columns, rows, err := loadGrid("grid.txt", itemsInProductNumber)

			if err == nil {
				fmt.Printf("%d %d\n", columns, rows)

				fmt.Println("Point element = ", grid[2].value, grid[6].value, grid[8].value)

				workOutLargestProduct(itemsInProductNumber, columns, rows)
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
			//fmt.Println("element = ", elementString)

			if _, err := strconv.Atoi(elementString); err == nil {

				elementNumber, _ := strconv.Atoi(elementString)

				//fmt.Println("elementNumber = ", elementNumber)

				appendGridElementToGrid(elementNumber)

				lineInFile = lineInFile[spacePosition+1:]
				//fmt.Println("New lineInFile =", lineInFile)
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

		//fmt.Println("Last element = ", lineInFile)
		if _, err := strconv.Atoi(lineInFile); err == nil {
			elementNumber, _ := strconv.Atoi(lineInFile)
			appendGridElementToGrid(elementNumber)

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

	fmt.Println("firstRowElements = ", firstRowElements)
	return firstRowElements, rowCount, nil
}

func appendGridElementToGrid(elementValue int) {

	var newGridElement gridElement

	//fmt.Println("1111")
	newGridElement.value = elementValue
	//fmt.Println("2222")

	grid[len(grid)] = newGridElement
}

func workOutLargestProduct(elementsInProduct int, columns int, rows int) {

	//Look north, if row >= elements in product.
	//Look south, if row <= rows - elements in product.
	//Look east, if column <= columns - elements in product.
	//Look west, if column >= elements in product.
	//Look south east,
	//Look north east,
	//Look south west,
	//Look north west,

	var row int
	var column int
	var largestProduct int
	var largestRow int
	var largestColumn int
	var largestDirection string

	for i := 0; i < len(grid); i++ {

		row, column = convertToRowColumns(i, columns)

		if row >= elementsInProduct {
			//Look north.
			newgrid := grid[i]
			newgrid.northProduct = calculateProduct(elementsInProduct, row, column, columns, "n")
			grid[i] = newgrid
		}

		if row <= rows-elementsInProduct {
			//Look south.
			newgrid := grid[i]
			newgrid.southProduct = calculateProduct(elementsInProduct, row, column, columns, "s")
			grid[i] = newgrid
		}

		if column <= columns-elementsInProduct {
			//Look east.
			newgrid := grid[i]
			newgrid.eastProduct = calculateProduct(elementsInProduct, row, column, columns, "e")
			grid[i] = newgrid
		}

		if column >= elementsInProduct {
			//Look west.
			newgrid := grid[i]
			newgrid.westProduct = calculateProduct(elementsInProduct, row, column, columns, "w")
			grid[i] = newgrid
		}

		if row <= rows-elementsInProduct && column <= columns-elementsInProduct {
			//Look south east.
			newgrid := grid[i]
			newgrid.southEastProduct = calculateProduct(elementsInProduct, row, column, columns, "se")
			grid[i] = newgrid
		}

		if row >= elementsInProduct && column <= columns-elementsInProduct {
			//Look north east.
			newgrid := grid[i]
			newgrid.northEastProduct = calculateProduct(elementsInProduct, row, column, columns, "ne")
			grid[i] = newgrid
		}

		if row <= rows-elementsInProduct && column >= elementsInProduct {
			//Look south west.
			newgrid := grid[i]
			newgrid.southWestProduct = calculateProduct(elementsInProduct, row, column, columns, "sw")
			grid[i] = newgrid
		}

		if row >= elementsInProduct && column >= elementsInProduct {
			//Look north west.
			newgrid := grid[i]
			newgrid.northWestProduct = calculateProduct(elementsInProduct, row, column, columns, "nw")
			grid[i] = newgrid
		}
	}

	for i := 0; i < len(grid); i++ {
	}

	fmt.Printf("Largest product = %d, row = %d, column = %d, direction = %s\n", largestProduct, largestRow, largestColumn, largestDirection)
}

func convertToRowColumns(index int, columns int) (int, int) {
	return (index / columns) + 1, (index % columns) + 1
}

func convertToIndex(row int, column int, columns int) int {

	if row == 7 && column == 9 {
		//fmt.Printf("Converting row %d, column %d, to index %d\n", row, column, ((row-1)*columns)+(column-1))
		//fmt.Println("Value = ", grid[((row-1)*columns)+(column-1)].value)
	}
	return ((row - 1) * columns) + (column - 1)
}

func calculateProduct(elementsInProduct int, row int, column int, columns int, direction string) int {

	var product int

	product = 1

	//fmt.Println("Going in - product = ", product)
	for i := 0; i < elementsInProduct; i++ {
		switch direction {
		case "n":
			//fmt.Println("Going north.")
			product *= grid[convertToIndex(row-i, column, columns)].value
		case "s":
			//fmt.Println("Going south.")
			product *= grid[convertToIndex(row+i, column, columns)].value
		case "e":
			//fmt.Println("Going east.")
			product *= grid[convertToIndex(row, column+i, columns)].value
		case "w":
			//fmt.Println("Going west.")
			product *= grid[convertToIndex(row, column-i, columns)].value
		case "ne":
			//fmt.Println("Going north east.")
			product *= grid[convertToIndex(row-i, column+i, columns)].value
		case "se":
			//fmt.Println("Going south east.")
			product *= grid[convertToIndex(row+i, column+i, columns)].value
		case "nw":
			//fmt.Println("Going north west.")
			product *= grid[convertToIndex(row-i, column-i, columns)].value
		case "sw":
			//fmt.Println("Going south west.")
			product *= grid[convertToIndex(row+i, column-i, columns)].value
		}
	}

	if row == 7 && column == 9 {
		//fmt.Println("Actual product = ", product)
	}

	return product
}
