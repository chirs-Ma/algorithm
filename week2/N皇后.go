func solveNQueens(n int) [][]string {
    //初始化棋盘
    board := make([][]string,n)
    for i := range board{
        board[i] = make([]string,n)
        for j := range board[i] {
            board[i][j] = "."
        }
    }

    result := [][]string{}
    cols := map[int]bool{}
    dig1 := map[int]bool{}
    dig2 := map[int]bool{}
    helper(n,0,board,&result,cols,dig1,dig2)
    return result
}
func helper(n int, r int, board [][]string,res *[][]string,cols,dig1,dig2 map[int]bool) {
    if r == n {
        temp := make([]string,len(board))
        for i:= 0; i < n; i++ {
            temp[i] = strings.Join(board[i],"")
        }
        *res = append(*res, temp)
        return
    }
    for c := 0; c < n; c++ {
        if !cols[c] && !dig1[r+c] && !dig2[r-c] {
            board[r][c] = "Q"
            cols[c] = true
            dig1[r+c] = true
            dig2[r-c] = true
            helper(n,r+1,board,res,cols,dig1,dig2)
            board[r][c] = "."
            cols[c] = false
            dig1[r+c] = false
            dig2[r-c] = false

        }
    }
}