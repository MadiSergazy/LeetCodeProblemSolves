


func twoSum(numbers []int, target int) []int {
    start , stop := 0, len(numbers)-1
   // var answer []int
    for true {
        if numbers[start] + numbers[stop] > target{
            stop--
        }else if numbers[start] + numbers[stop] < target {
            start++
        }else {
            //answer = append(answer, start+1)
            // answer = append(answer, stop+1)
             return [] int{start+1, stop+1}
        }
    }
    return [] int{start+1, stop+1}
    
}

/*
func twoSum(numbers []int, target int) []int {
    var answer []int
    
    for i := 0; i < len(numbers); i++{
        for j := 0; j < len(numbers); j++{
            if i == j {
                continue 
            }
            if numbers[i] + numbers[j] == target {
                answer = append(answer, i+1)
                answer = append(answer,  j+1)
                return answer
            }
        }
    }
    return answer
}
*/
