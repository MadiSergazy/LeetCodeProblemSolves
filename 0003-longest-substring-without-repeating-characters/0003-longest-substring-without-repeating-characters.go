
//метод скользящего окна
func lengthOfLongestSubstring(s string) int {
    
    if s == "" || len(s) == 0{
        return 0
    }
    
    i, j, max := 0,0,0
    hm := map[byte]bool{}  
    
    for j < len(s){
        if _, ok := hm[s[j]]; !ok { //если фолс 
          
                 hm[s[j]] = true
        j++ //движем j на позицию 
            if max < j-i{
                max = j-i
            }
        } else { //иначе удалим 
            delete (hm, s[i])// удаляем элемент с ключем i
             i++ //движем i на позицию 
        }
   
    }
    
    return max
}


/*
//брутфорс 
func lengthOfLongestSubstring(s string) int {
    max := 0
    
    hm := map[byte]bool{}  
    for i := 0; i < len(s); i++{
         hm = map[byte]bool{}     //обноаляем мапу 
        for j := i; j < len(s); j++{
            
            if _, ok := hm[s[j]]; !ok {  //ok - это булевая переменная которая либо тру либо фолс 
                hm[s[j]] = true
            if max < j-i + 1 {  //если он даст фолс то мы изменим его на тру и изменим макс 
                max = j-i + 1
            }
            } else{
                break
            }
            
        }
    }
    
    return max
    
}*/




/*  //метод индуса
func lengthOfLongestSubstring(s string) int {

    start, end, max := 0,0,0
    if len(s) != 0 {
        charMap := make(map[rune]int)
        for i, v := range s{
            if k, ok := charMap[v]; ok && start <= k {
                start = k + 1
            }
            charMap[v] = i
            
            if max < i-start + 1{
                max = i-start + 1
            }
        }
    }
    
    return max
}*/

/* метод Мадо несработал
func lengthOfLongestSubstring(s string) int {

    dic := make(map[rune]int)
var lastword string
    
    ifOneLet :=  make(map[rune]int)
     for _, v := range s{
     ifOneLet[v]++  
     }
      for _, v := range s{
     if  ifOneLet[v] == len(s){
        return 1
    }  
        
     }
   
    
    
    for _, v := range s{
        
        fmt.Println(string(v))
         fmt.Println(string(lastword))
        dic[v]++
        if dic[v] > 1 && lastword != string(v){ 
        return len(dic)
        } else if  dic[v] > 1 {//&& lastword == string(v) {
            fmt.Println(string(v))
            return len(dic)+1
        }
            
             
        lastword = string(v)
         
       
    }
   
         return len(dic)
  
    
}*/

/*
func lengthOfLongestSubstring(s string) int {
	
	sLength := len(s)   //берем длину 
	maxLength := 0
	nowLength := 0
	leftIndex := 0
	rightWord := ""
	dic := make(map[string]int)  //словар для неповторяющихся слов 

	for rightIndex := 0; rightIndex <= sLength; rightIndex++ {

		if rightIndex != sLength {
			rightWord = string(s[rightIndex])  //пока мы не в конце строки добавим слово 
		} else {
			rightWord = " " //когда мы в коце обновим 
		}

		if d, ok := dic[rightWord]; ok {
			if leftIndex == 0 {
				nowLength = rightIndex
			} else {
				nowLength = rightIndex - leftIndex
			}
			if d >= leftIndex {
				leftIndex = d + 1
			}

		} else {
			nowLength = rightIndex - leftIndex
		}

		dic[rightWord] = rightIndex
		if nowLength > maxLength {
			maxLength = nowLength
		}

	}

	return maxLength
}*/

