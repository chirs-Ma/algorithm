//方法一：使用自带函数
func reverseWords(s string) string {
	list :=  strings.Split(s ," ")
	var res []string
	for i:=len(list)-1;i>=0;i--{
		if len(list[i])>0{
			res = append(res,list[i])
	   }
	}
	s =strings.Join(res," ")
	return s
}
//方法二：自己编写函数
func reverseWords(s string) string {
    if s == "" {
        return ""
    }
    res := []byte{}
    queue := []string{}

    word := []byte{}
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            if len(word) > 0 {
                queue = append(queue, string(word))
                word = []byte{}
            }
        } else {
            word = append(word, s[i])
        }
    }
    if len(word) > 0 {
        queue = append(queue, string(word))
    }
    
    if len(queue) <= 0 {
        return ""
    }

    for i := len(queue) - 1; i >= 0; i-- {
        res = append(res, []byte(queue[i])...)
        res = append(res, ' ')
    }

    return string(res[:len(res)-1])
}
