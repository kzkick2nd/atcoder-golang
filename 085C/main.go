package main

import (
	"fmt"
	"os"
)

func main() {
	var N, Y int
	fmt.Scan(&N, &Y)
	for ix := 0; ix <= N; ix++ {
		for iy := 0; iy <= N; iy++ {
			iz := N - ix - iy
			if iz < 0 {
				continue
			}
			if Y == 10000*ix+5000*iy+1000*iz {
				fmt.Println(ix, iy, iz)
				os.Exit(0)
			}
		}
	}
	fmt.Println(-1, -1, -1)
}
