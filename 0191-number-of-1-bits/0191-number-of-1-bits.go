func hammingWeight(num uint32) int {
    counter := uint32(0)
    
    //  4 = 100
    
    for it := num ; it > uint32(0); it = it >> 1 { // смещение на бит вправо тоесть / 2 при каждой итерации 
        counter += it&1 // 0000001 самый правый битті қосамыз по битово
        
    }
    return int(counter)
    
}
/*
func hammingWeight(num uint32) int {
    
    
    
    fmt.Println(num)
    return bits.OnesCount32(num)
}*/