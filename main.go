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

var mc = flag.Int("mc", 1, "几块内存, -mc=1")
var m = flag.Int("m", 1, "每块内存大小（默认是1Gb）: -m=1")
var t = flag.Int("t", 1, "占用几个Cpu, -t=1")
var mid = flag.Int("mid", 1, "废弃的")
var sleep = flag.Int("sleep", 1, "sleep Microsecond, -s=1")
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
