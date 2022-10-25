func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    s1, s2 := "", ""
    for _, v := range word1{
        s1 += string(v)
    }
    
      for _, v := range word2{
        s2 += string(v)
    }
    if s1 == s2 {
        return true 
    } else {
        return false
    }
}