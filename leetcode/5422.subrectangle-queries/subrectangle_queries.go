package leetcode

// SubrectangleQueries ...
type SubrectangleQueries struct {
	data [][]int
}

// Constructor ...
func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{data: rectangle}
}

// UpdateSubrectangle ...
func (q *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			q.data[i][j] = newValue
		}
	}
}

// GetValue ...
func (q *SubrectangleQueries) GetValue(row int, col int) int {
	return q.data[row][col]
}
