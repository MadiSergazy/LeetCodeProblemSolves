func isPowerOfTwo(n int) bool {

	if n <= 0 {
		return false
	} else if n == 1 || n == 2 {
		return true
	}
    for n != 2 {
        if n % 2 != 0{
            return false
        }
        n = n >> 1 // n/2^1
    }
    return true
    
}

/*
func isPowerOfTwo(n int) bool {

	if n <= 0 {
		return false
	} else if n == 1 || n == 2 {
		return true
	}

	for ; n > 2; n = n>>1 {
		if n%2 != 0 {
			return false
		}
	}
	return true
}
*/
/*
func isPowerOfTwo(n int) bool {
    if n == 1 || n == 2{
        return true 
    }
    it := 2
    for it < n{
       
        it  *= 2 
        fmt.Println(it)
        if it == n {
            return true
        }
}
    fmt.Println(it)
    return false 
    
    
}*/