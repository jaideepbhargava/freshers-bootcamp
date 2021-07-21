package main

import "fmt"

type Matrix struct{
	numRows int
	numColumn int
	mat [][] int
}

func (mat *Matrix) getRows(){
	fmt.Scan(&mat.numRows)

}

func (mat *Matrix ) getColumn() {
	fmt.Scan(&mat.numColumn)
}

func (mat *Matrix) setValues(){
	var posRow, posColumn, input int
	fmt.Scan(&posRow)
	fmt.Scan(&posColumn)
	fmt.Scan(&input)
	(*mat).mat[posRow][posColumn] = input
}

func (mat *Matrix) intitializeMatrix(){
	(*mat).mat = make([][]int,(*mat).numRows )
	for i:=range (*mat).mat{
		(*mat).mat[i] = make([]int, (*mat).numColumn)
	}
}
func (mat *Matrix) addTwoMatrix(different [][]  int){

	for index, row := range (*mat).mat{
		for counter, value:= range row{
			(*mat).mat[index][counter] = value+different[index][counter]
		}
	}

}

func (mat *Matrix) printJSON(){
	for _, row := range (*mat).mat{
		for index, value:= range row{
			fmt.Println("{", index, " :",  value, "}")
		}
	}
}

func main(){
	var inputMat Matrix

	inputMat.getRows()
	inputMat.getColumn()
	inputMat.intitializeMatrix()
	inputMat.setValues()
	inputMat.printJSON()

}