# Instructions and Notes

## The Solution

I decided that I would expand on the original basic request, as denoted in Project Euler. 

The grid itself is contained in the grid.txt file, and the length of elements that are required to make up the product, are passed in via the command line. Any size grid can be used, with the code, working out the rows and columns.  The code checks that the grid is well-formed by comparing the following:

* The number of rows and columns must be equal to, or greater, than the number of elements used to make up the product.

* All grid elements must be numeric.

* All rows in the grid must contain the same number of columns.

## Design Considerations

The reason for using a Map rather than an array, was so that I could index the individual grid cells directly. This makes the processing of the diagonal (NE,NW,SE,SW) a lot easier.  

The reason for using a struct is so I could save the 'product' calculations for each 'direction', against the grid cell itself. Although this is not required in this iteration of the solution, it means that the solution can be expanded to display highest, lowest, average etc products.

## Future Plans (if I ever get round to it)

* Add more command line switches to allow multiple cases to be returned (largest, smallest, average product).

* Make the code more efficient. A 'South East' search from row 2, column 2, would produce the same result as a 'North West' search from row 6, column 6, with 4 elements. There is no point doing this twice, so maybe searched 'South', 'East', 'South East' and 'South West' searches, are all that is required.

* Investigate the prospect of using 'goroutines' and 'channels' to search the grid more quickly, by delagating each row to a seperate 'goroutine'.
