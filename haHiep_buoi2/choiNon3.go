package main

import "fmt"

func main() {
	var a,b int
	fmt.Scanf("%d %d",&a,&b)

	if a==b {
		fmt.Printf("%d bang %d\n",a,b)
	} else if a>b {
		fmt.Printf("%d lon hon %d\n",a,b)
	} else {
		fmt.Printf("%d nho hon %d\n",a,b)
	}
}
