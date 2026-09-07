package minitask

import "fmt"

func PrintJendela(p, l int8) {
	for i := range p {
		for j := range l {
			if i == 0 || i == p-1 {
				fmt.Print("*" + " ")
				if j == l-1 {
					fmt.Print("\n")
				}
			} else {
				if j == 0 || j == l-1 {
					fmt.Print("*" + " ")
				} else {
					fmt.Print("  ")
				}
				if j == l-1 {
					fmt.Print("\n")
				}
			}
		}
	}
}
