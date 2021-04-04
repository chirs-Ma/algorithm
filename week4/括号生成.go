func generateParenthesis(n int) []string {
    result := []string{}
    generate(0,0,n,"",&result)
    return result

}
func generate(left int,right int,n int, t string, s *[]string){
    if left == n && right == n {
        *s = append(*s, t)
        return
    }
    if left < n {
        generate(left+1,right,n,t+"(",s)
    }
    if right < left {
        generate(left,right+1,n,t+")",s)
    }
}

//