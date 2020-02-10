/**
 * Auth :   liubo
 * Date :   2020/2/10 20:45
 * Comment:
 */

package main

import (
	"flag"
	"fmt"
)

var s = flag.Int("s", 1, "memory size")
var t = flag.Int("t", 1, "thread count")
func main() {
	flag.Parse()
	fmt.Println(*s, *t)

	var size = 1024 * 1024 * 1024 * *s
	var mem = make([]byte, size, size)
	var counter = 0
	if *t > 1 {
		for i:=0; i<*t-1; i++ {
			go func() {
				for {
					counter++
					mem[0] = 1
				}
			}()
		}
	}

	for {
		counter++
		mem[0] = 1
	}

}
