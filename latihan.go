package main

import (
	"fmt"
)

func superPrima(n int) bool{
    var b1, b2 bool
    var i, hasil, p int
    b2 = true
    b1 = true
    hasil = 0
    i = 2
    for i < n {
        if n%i == 0{
            b1 = false
        }
        i++
    }
    i = 2
    for n > 0{
        p = n%10
        hasil += p
        n = n/10
    }
    for i < hasil{
        if hasil%i == 0{
            b2 = false
        }
        i++
    }
    if b1 == true && b2 == true{
        return true
    }else{
        return false
    }

}
func main(){
    var i int
    fmt.Scan(&i)
    fmt.Print(superPrima(i))
}