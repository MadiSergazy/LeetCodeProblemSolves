type Nums struct {
    Ans [][]int
    Ns []int
    Used []bool
    Tmp []int
    Lenght int
}


func (n *Nums) backTrack(){ //(n *Nums) показывает что это функция структуры пля которых мы можем менять
    if len(n.Tmp) == n.Lenght{
      
        n.Ans = append(n.Ans, append([]int{}, n.Tmp...))//n.Tmp... это приводит его к двумерному массиву
        mat := append([]int{}, n.Tmp...)  //делает одномерный массив котрый можно добавит к двумерному 
        fmt.Println(mat)
        return //делаем остановку потомушто чель достигнута 
         fmt.Println(n.Ans)
    }
    
    for i := 0; i < n.Lenght; i++{
        if !n.Used[i] {
            n.Used[i] = true
            
            n.Tmp = append(n.Tmp, n.Ns[i])
            
            n.backTrack()//повторно вызываем чтоб остальные вариянты проверить
            
            n.Used[i] = false
            n.Tmp = n.Tmp[:len(n.Tmp)-1]
        }
    }
}


func permute(nums []int) [][]int {
    n := new(Nums)    //с помощю new мы создали поля которые просто инициализированные 
    //var n New //new эквивалентен 
	//p := &n
    
    n.Ans = [][]int{}
    n.Ns = nums
    n.Used = make([]bool, len(nums))
    n.Tmp = make([]int, 0, len(nums))
    n.Lenght = len(nums)
    
    n.backTrack()
    
    return n.Ans
}






/*
Runtime: 6 ms, faster than 17.93% of Go online submissions for Permutations.
Memory Usage: 2.6 MB, less than 69.20% of Go online submissions for Permutations.
*/
/*
type Nums struct {
	Ans  [][]int      //для ответа 
	Ns   []int
	Used []bool       //для проверки использовали ли м уже рание 
	Tmp  []int
	Len  int
}

// nums = [1,2,3]
func permute(nums []int) [][]int {

	//init инициализация
	n := new(Nums)
	n.Ns = nums
	n.Len = len(nums)
	n.Tmp = make([]int, 0, n.Len)//временный массив с длиной 0 и вместимостю 3
    fmt.Println("TMP")
    fmt.Println(n.Tmp)
	n.Used = make([]bool, n.Len)   //булевой массви для проверки то были ли у нас эти числа
    fmt.Println("USED")
    fmt.Println(n.Used)

	n.backtrack()

	return n.Ans
}

func (n *Nums) backtrack() {

	//Goal  ( Constraints2 )
	if len(n.Tmp) == n.Len {
		n.Ans = append(n.Ans, append([]int{}, n.Tmp...))
		return
	}

	for i := 0; i < n.Len; i++ {

		//Constraints1
		if !n.Used[i] { //если мы еще неиспользовали это число 

			//Choice
			n.Used[i] = true
			n.Tmp = append(n.Tmp, n.Ns[i]) //Go добавим в временный массив нейспользованное число 

			n.backtrack()

			n.Used[i] = false   //укажем что в этом индексе элемент испоользован 
			n.Tmp = n.Tmp[:len(n.Tmp)-1] //Back и удаляем последни элемент массива 
		}
	}

}
*/

/*
Runtime: 865 ms, faster than 9.54% of Go online submissions for Permutations.
Memory Usage: 7.7 MB, less than 10.34% of Go online submissions for Permutations.
*/
/*
func permute(nums []int) [][]int {
    if len(nums) == 1 {
        fmt.Println("IF NUMS == 1")
        fmt.Println(nums)
        return [][]int{nums}
    } else {
        var result [][]int
        for i,n := range nums {
            nums2 := make([]int,len(nums))  // создаем временный массив 
            copy(nums2,nums)                //кпируем во временный массив значения нашего переданного массива
            temp := permute(append(nums2[:i],nums2[i+1:]...))  //
            fmt.Println("TEMP:")
            fmt.Println(temp)
            for _,m := range temp {
                m = append(m,n)
                result = append(result,m)
                fmt.Println("result:")
                fmt.Println(result)
            }
        }
        return result
    }
}
*/



/*
Runtime: 5 ms, faster than 29.88% of Go online submissions for Permutations.
Memory Usage: 2.6 MB, less than 94.25% of Go online submissions for Permutations.
*/
/*
func permute(nums []int) [][]int {
    var res [][]int
    backtrack(&res, len(nums), nums, 0)
    return res
}

func backtrack(res *[][]int, n int, nums []int, first int) {
    if first == n {
        tmp := make([]int, len(nums))
        copy(tmp, nums)
        *res = append(*res, tmp[:])
        return
    }

    for i:=first; i<n ;i++ {
        swap(nums, first, i)
        backtrack(res, n, nums, first+1)
        swap(nums, first, i)
    }
}

func swap(nums []int, a, b int) {
    nums[a], nums[b] = nums[b], nums[a]
}
*/


