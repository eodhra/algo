package main
import "fmt"

func max(arr []int) int {
    var m int = 0
    
    for _, e := range arr {
        if e > m {
            m = e
        }
    }
    
    return m
}

func cntsort(arr []int) {
    var m int = max(arr)
    var b []int = make([]int, m + 1)
    
    for _, e := range arr {
        b[e] += 1
    }
    
    fmt.Printf("%v\n", b)
    
    var i int = 0
    for x, e := range b {
        for e > 0 {
            arr[i] = x
            i += 1
            e -= 1
        }
    }
}

func main() {
    var arr []int = []int{1, 9, 1, 9, 8, 1, 0}
    cntsort(arr)
    fmt.Printf("%v\n", arr)
}
