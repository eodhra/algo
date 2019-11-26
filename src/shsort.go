package main
import "fmt"

var arrl int

func bachinsrt(arr []int, gap int) {    
    for h := 0; h < gap; h += 1 {
        for l := 1; l * gap + h < arrl; l += 1 {
            for j := l; j > 0; j -= 1 {
                var rj int = j * gap + h
                if arr[rj - 1] > arr[rj] {
                    arr[rj - 1], arr[rj] = arr[rj], arr[rj - 1]
                }
            }
        }
    }
}

func shsort(arr []int) {
    var gap int = arrl / 2
    for gap >= 1 {
        bachinsrt(arr, gap)
        gap /= 2
    }
}

func main() {
    var arr []int = []int{1, 9, 1, 9, 8, 1, 0}
    arrl = len(arr)
    shsort(arr)
    
    fmt.Printf("%v\n", arr)
}
