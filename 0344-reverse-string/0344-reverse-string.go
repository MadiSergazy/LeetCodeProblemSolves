

func reverseString(s []byte)  {
    leftIdx, rightIdx := 0, len(s)-1
    for leftIdx <= rightIdx {
        s[leftIdx], s[rightIdx] = s[rightIdx], s[leftIdx]
        leftIdx += 1
        rightIdx -= 1
    }
}

/*
func reverseString(s []byte)  {
   var answer []byte
   
    for i := len(s) -1; i >= 0; i --{
       
        answer = append(answer, s[i])
    }
    for i := 0; i < len(s); i++{
         s[i] = answer[i] 
       
    }
}*/