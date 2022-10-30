func isValid(s string) bool {
   /* mapa := make(map[string]int)
    for _, v := range s{
        mapa[string(v)]++
    }*/
    m := map[rune]rune {
    '(' : ')',
    '[': ']',
    '{': '}',
    }
   
   var q []rune
    for _, c := range s {
        switch c {
            
    case '[', '{', '(':
             q = append(q, m[c])
             default:
            if len(q) == 0 || c != q[len(q)-1]{
             return false
            } else {
                q = q[:len(q)-1]
            }
        }
        
    }
     return len(q)==0
    /*
	for _, c := range s {
        switch c {
            
        case '[':
            q = append(q, ']')
            
        case '(':
            q = append(q, ')')
            
        case '{':
            q = append(q, '}')
        default:
            if len(q) == 0 || c != q[len(q)-1]{
             return false
            } else {
                q = q[:len(q)-1]
            }
        }   
    }
    return len(q)==0
*/
}