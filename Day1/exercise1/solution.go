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

func (mat *Matrix) intitializeMat(){
	(*mat).mat = make([][]int,(*mat).numRows )
	for i:=range (*mat).mat{
		(*mat).mat[i] = make([]int, (*mat).numColumn)
	}
}
func (mat *Matrix) addTwo(different [][]  int){

	for index, row := range (*mat).mat{
		for j, col:= range row{
			(*mat).mat[index][j] = col+different[index][j]
		}
	}

}

func (mat *Matrix) printJSON(){
	for index, row := range (*mat).mat{
		for j, col:= range row{
			fmt.Println("{", j, " :",  col, "}")
		}
	}
}

func main(){
	var inputMat Matrix

	inputMat.getRows()
	inputMat.getColumn()
	inputMat.intitializeMat()
	inputMat.setValues()

}