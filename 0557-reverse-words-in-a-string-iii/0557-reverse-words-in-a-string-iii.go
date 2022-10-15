func reverseWords(s string) string {
 var answer bytes.Buffer 
    it := 0
    
    for it < len(s){
        if s[it] == ' '{
            answer.WriteByte(s[it])
            it++
            continue 
        }
        j := it 
        for j < len(s) && s[j] != ' '{ //сөзді алып аламыз керекті 
            j++
        }
        for k := j-1; k >= it; k--{
            answer.WriteByte(s[k])
        }
        it = j
    }
    return answer.String()
    
}

//Буфер — это буфер байтов переменного размера с методами чтения и записи. Нулевое значение Buffer означает 
//пустой буфер, готовый к использованию.

//WriteByte добавляет байт c к буферу, увеличивая буфер по мере необходимости. Возвращаемая ошибка всегда равна нулю, но включается, чтобы соответствовать WriteByte bufio.Writer. Если буфер становится слишком большим, WriteByte будет паниковать с ошибкой ErrTooLarge
/*
func reverseWords(s string) string {
    var res bytes.Buffer  
    i := 0
    for i < len(s) {   
        if s[i] == ' ' {    //если найдем пробел 
            res.WriteByte(s[i])  //записываем его в переменную
            i++//переходим на след 
            continue  //пропустим 
        }
        j := i   //начнем с позиции i
        for j < len(s) && s[j] != ' ' {
            j++   //пока мы неуперлис в конетц либо ненашли пробел 
        }
        for k := j - 1; k >= i; k-- {   //кода мы нашли пробел или конетц то будем записывать до позиции i
            res.WriteByte(s[k])
        } 
        i = j  //дадим позицию в которой остановилис 
    }
    return res.String()  //вернем резулютатт в виде строки 
}*/


/*
func reverseWords(s string) string {
    var splitString []string
    word := ""
    
    for i := 0; i < len(s); i++{
        if s[i] == ' ' {
            splitString = append(splitString, word) 
            word = ""
            continue
        }
        word += string(s[i])
    }
    if word != ""{
      splitString = append(splitString, word)   
    }
    answer := ""
    for index, v := range splitString{
        for i := len(v)-1; i >= 0; i--{
            answer += string(v[i])
            
        }
        if index != len(splitString)-1{
             answer += " "
        }
       
    }
    

    return answer 
}*/