package subrectangleQueries

//https://leetcode-cn.com/problems/subrectangle-queries
type SubrectangleQueries struct {
	rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	result := SubrectangleQueries{
		rectangle,
	}
	return result
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.rectangle[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	if len(this.rectangle) == 0 {
		return 0
	}

	return this.rectangle[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
