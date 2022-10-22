func combine(n int, k int) [][]int {
    res := [][]int{}    //двумерный массив для ответа
    
    if k == 0 || n == 0{    //1 шаг обрабатываем случи когда все данные 0
        return res
    }
    
    cur := []int{}
    backtrack(n, k, 1, cur, &res) //изменим res
    return res
    
    
}
/*
Встроенная функция copy копирует элементы в целевой срез dst из исходного среза src.


func copy(dst, src []Type) int

var b = make([]byte, 5)
copy(b, "Hello, world!") 
// b == []byte("Hello")
*/

func backtrack(n int, k int, start int, cur []int, res *[][] int){
    if len(cur) == k { //если длина равна комбинации 
        temp := make([]int, k)//создадим временный массив с размером комбинации 
        copy(temp, cur) // копируем в темп модержимое cur
        *res = append(*res, temp)//изменим массив по ссылке добавив в него темп
        return 
    }
    for i := start; i <= n; i++{
        //add option
        cur = append(cur, i) //добавим вариянты
    //recursive call backtrack 
         backtrack(n, k, i+1, cur, res) //вызываем уже для след числа
        //remove option
        cur = cur[:len(cur)-1] //типо последни элементқа келгенде все уже повторено 
    }
}

/*
func combine(n int, k int) [][]int {
    var mas []int
    var masMini []int
    
    
    for i := 0; i < n; i++{
        for j := 0; j < n; j++{
            
        }
    }
}

*/