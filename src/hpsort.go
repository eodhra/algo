package main
import "fmt"

var arrl int

func swap(arr []int, i1 int, i2 int) {
    arr[i1], arr[i2] = arr[i2], arr[i1]
}

func maxheap(arr []int, r int) {
    var f int = r * 2 + 1
    for f < arrl {
        fmt.Printf("r = %d ", r)
        if f + 1 < arrl && arr[f] < arr[f + 1] {
            f += 1
        }
        
        if arr[f] > arr[r] {
            swap(arr, f, r)
            r = f
            f = f * 2 + 1
        } else {
            fmt.Printf("break\n")
            break
        }
        fmt.Printf("%v\n", arr)
    }
}

func main() {
    var arr []int = []int{1, 9, 1, 9, 8, 1, 0}
    arrl = len(arr)
    
    for r := arrl / 2 - 1; r >= 0; r -= 1 {
        maxheap(arr, r)
        
    }
    
    fmt.Printf("----------\n")
    
    for x := arrl - 1; x > 0; x -= 1 {
        swap(arr, x, 0)
        arrl -= 1
        maxheap(arr, 0)
    }
    
    fmt.Printf("%v\n", arr)
}
