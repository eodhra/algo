package main
import "fmt"

func qksort(arr []int, l int, r int) {
    if l >= r {
        return
    }
    
    fmt.Printf("%v\n", arr)
    
    var b, i, j int = arr[l], l, r
    for i < j {
        for i < j && arr[j] >= b {
            j -= 1
        }
        arr[i] = arr[j]
        
        for i < j && arr[i] < b {
            i += 1
        }
        arr[j] = arr[i]
    }
    
    arr[i] = b
    qksort(arr, 0, i - 1)
    qksort(arr, i + 1, r)
}

func main() {
    var arr []int = []int{1, 9, 1, 9, 8, 1, 0}
    qksort(arr, 0, len(arr) - 1)
    
    fmt.Printf("%v\n", arr)
}
