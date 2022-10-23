
      
func isPalindrome(x int) bool {
    strX := strconv.Itoa(x)
    masR := []rune(strX)
    
    for i, j := 0, len(masR)-1; i < j; i, j = i+1, j-1{
        if (masR[i] != masR[j]){
            return false
        }
    }
    return true
    
    
    
}