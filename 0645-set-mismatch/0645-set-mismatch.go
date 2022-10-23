func findErrorNums(nums []int) []int {
    
	result := make([]int, 2)
	size := len(nums)
	sum := 0
	existNums := make([]bool, size+1)
	for _, num := range nums {
		sum += num
		if !existNums[num] {
			existNums[num] = true
			continue
		}
		result[0] = num
	}

	result[1] = size*(size+1)/2 - sum + result[0]

	return result
}
    
   /* var ans []int
    sort.Ints(nums)
    dup, missing := -1, -1
    for i := 1; i < len(nums); i++{
        if nums[i] == nums[i-1]{
            dup = nums[i]
        } else if nums[i] > nums[i-1]+1{
            missing = nums[i-1]+1
        }
    }
    
    if nums[len(nums)-1] != len(nums) {
    ans = append(ans, dup)
             ans = append(ans, len(nums))
            return ans
    } else {
       
         ans = append(ans, dup)
             ans = append(ans, missing)
            return ans
            
    }

    */





/*func findErrorNums(nums []int) []int {
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
    
  
}*/