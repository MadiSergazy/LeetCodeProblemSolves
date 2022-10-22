/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*func middleNode(head *ListNode)   *ListNode {
    counter := 0
    link := []*ListNode{}
    
    for i := head; i != nil; i = i.Next{
        counter++
        link = append(link, i)
    }
    fmt.Println(counter)
    counter = counter/2
     fmt.Println(counter)
    return link[len(link)/2]
}*/
/*
func middleNode(head *ListNode)   *ListNode {
    counter := 0
    link := [100]*ListNode{}
    
    for i := head; i != nil; i = i.Next{
        link[counter] = i
        counter++
        
    }
    fmt.Println(counter)
    //counter = counter/2
   // fmt.Println(counter)
 //   link2 := link[:counter]
   // return link2[counter/2]
      return link[counter/2]
}*/

/*
func middleNode(head *ListNode)   *ListNode {
    out := []*ListNode{}
    for head != nil{
        out = append(out, head)
        head = head.Next
        
    }
    return out[len(out)/2]
    
    
}*/
func middleNode(head *ListNode)   *ListNode {
    slow, fast := head, head   //представим что два чеоека двигаются противоположно друг другу в напровлении 
    for fast != nil && fast.Next != nil && slow != nil{ // fast движется в два раза быстрее 
        slow = slow.Next
        fast = fast.Next.Next  //в любом случае когда фаст придет к концу слоу будет посередине
        
    }
    return slow
    
    
}
