

func rob(nums []int) int {
    leng := len(nums)
    
    if nums == nil || leng == 0{
return 0
    }
    if leng == 1{
return nums[0] 
    } else if leng == 2 {
        return max(nums[0], nums[1])
    } else {
        dp := make([]int, leng)
        
        dp[0] = nums[0]
        dp[1] = max(nums[0], nums[1])
        
        for i := 2; i < leng; i++{
            dp[i] = max(nums[i] + dp[i-2], dp[i-1])
        }
        fmt.Println(dp)
         return  dp[leng-1]
    }
    
 
}

func max(a, b int) int{
    if a > b {
        return a
    } else {
        return b
    }
   
}



/* ///эквивалент
func rob(nums []int) int {
	n := len(nums)
	if n<=1 {
		return nums[0]
	}
	dp := make([]int, n)

	dp[0] = nums[0]
	if nums[1]>nums[0]{
		dp[1]=nums[1]
	} else {
		dp[1] = nums[0]
	}

	for i:= 2;i<n;i++{
		if (nums[i]+dp[i-2])>dp[i-1]{
			dp[i]= nums[i]+dp[i-2]
		} else {
			dp[i]= dp[i-1]
		}
	}

	return dp[n-1]
}*/

/*
func rob(nums []int) int {
    rob1, rob2 := 0, 0
    fmt.Println("ROB1")
    fmt.Println(rob1)
    fmt.Println("ROB2")
    fmt.Println(rob2)
    for _, n := range nums{
       // temp := max(n+rob1, rob2)
       // rob1 = rob2
       // rob2 = temp
         rob1 = rob2
        rob2 =  max(n+rob1, rob2)
    fmt.Println("ROB1")
    fmt.Println(rob1)
    fmt.Println("ROB2")
    fmt.Println(rob2)
    }
    
    
    
    return rob2
}

func max(a, b int) int{
    if a > b {
        return a
    } else {
        return b
    }
   
}*/
/*
func rob(nums []int) int {
    
    ans := 0
    if len(nums) == 1{
        return nums[0]
    }
    if len(nums) <= 3 {
       
            for i := 0; i < len(nums)-1; i++{
             if nums[i] >= nums[i+1]{
           ans += nums[i]
           
        } else {
            ans += nums[i+1]
           
        }
        } 
       
       
        
         return ans
    }
    
    for i, v := range nums{
        fmt.Println((i+1)%2)
        if (i+1) % 2 == 1 {
            ans += v
        }
    }
    return ans
}*/