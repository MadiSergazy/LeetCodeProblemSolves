func findErrorNums(nums []int) []int {
    var ans []int
   
    dup, missing := -1, -1
    
    for i := 1; i <= len(nums); i++{ //начинаем с того числа который будет в начале массива тоесть 1 и закончиваем длиной потомушто пропущенное число моежт быть в самом конце 
        count := 0
        for j := 0; j < len(nums); j++{
            if nums[j] == i{
                count++
            }
        }
        if count == 2{
            dup = i
        } else if count == 0{
            missing = i
        }
        if dup > 0 && missing > 0{
            ans = append(ans, dup)
             ans = append(ans, missing)
            return ans
        }
        
        
    }
     ans = append(ans, dup)
             ans = append(ans, missing)
            return ans
    
  
}