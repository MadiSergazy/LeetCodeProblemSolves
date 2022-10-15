func moveZeroes(nums []int)  {   //sucses
    coun := 0
    
    for i := 0; i < len(nums); i++{
        if nums[i] != 0{
            nums[coun] = nums[i]
            coun++
        }
    }
    
    for i := coun; i < len(nums); i++{
        nums[i] = 0
    }
    
}


/*func moveZeroes(nums []int)  {   //sucses
    for j := 0; j < len(nums); j++ {
          for i := 0; i < len(nums)-1; i++{
        if nums[i] == 0 && nums[i+1] != 0{
            temp := nums[i]
            nums[i] = nums[i+1]
            nums[i+1] = temp
        }
    } 
    }
 
}*/


/*   nerabotaet 
func moveZeroes(nums []int)  {
    if len(nums) == 1 {
        return
    }
    if  len(nums) == 2 {
        if nums[0] == 0 && nums [1] != 0{
             temp := nums[0]
            nums[0] = nums[0+1]
            nums[0+1] = temp  
        }
        if nums[0] == 0 && nums[1] == 0{
            return
        }
    }
    str, stp := 0, len(nums) -1 
        for i := 0; i < len(nums); i++{
         for j := 0; j < len(nums)-1; j++{
             if nums[j] > nums[j+1] {
                  temp := nums[j]
            nums[j] = nums[j+1]
            nums[j+1] = temp  
             }
         }
    }
    for true {
        if str == stp {
            break
        }
        if nums[str] == 0 || nums[stp] == 0{
            temp := nums[str]
            nums[str] = nums[stp]
            nums[stp] = temp  
            str++
            stp--
        }else {
            str++
            stp--
        }
    }
    
         for i := 0; i < len(nums)-2; i++{
         for j := 0; j < len(nums)-3; j++{
             if nums[j] > nums[j+1] {
                  temp := nums[j]
            nums[j] = nums[j+1]
            nums[j+1] = temp  
             }
         }
    }
    

}*/