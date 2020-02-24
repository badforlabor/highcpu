/**
 * Auth :   liubo
 * Date :   2020/2/10 20:45
 * Comment:
 */

package main

import (
	"flag"
	"fmt"
	"time"
)

var m = flag.Int("m", 1, "memory size")
var mc = flag.Int("mc", 1, "memory count")
var t = flag.Int("t", 1, "thread count")
var mid = flag.Int("mid", 1, "mid cpu")
var sleep = flag.Int("sleep", 1, "mid cpu")
func main() {
	flag.Parse()
	fmt.Println(*m, *t)

	var size = 1024 * 1024 * 1024 * *m
	var memList = make([][]byte, *mc)
	for i,_ := range memList {
		memList[i] = make([]byte, size, size)
	}
	var counter = 0
	if *t > 1 {
		for i:=0; i<*t-1; i++ {
			go func() {
				for {
					counter++
					for im, _ := range memList {
						memList[im][0] = 1
					}
				}
			}()
		}
	}

	for {
		var cnt = 0
		for cnt < 1024 * 1024 * 1 * *mid {
			counter++
			for im, _ := range memList {
				memList[im][0] = 1
			}
			cnt++
		}
		time.Sleep(time.Microsecond * time.Duration(*sleep))
	}

}
