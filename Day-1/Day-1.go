package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	
	var i uint64=4;
	var d float64=4.0;
	var s string="HackerRank "
	
	scanner:=bufio.NewScanner(os.Stdin)
	//My Code
	
	var i2 uint64;
	var d2 float64;
	var s2 string;

	fmt.Scanf("%d",&i2)
	fmt.Scanf("%f",&d2)
	scanner.Scan()
	s2=scanner.Text()
	fmt.Printf("%d\n%0.1f\n%s",i+i2,d+d2,s+s2)
}