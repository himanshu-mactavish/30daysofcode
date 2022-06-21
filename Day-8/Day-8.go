package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var entries int;
	fmt.Scanf("%d",&entries)
	myphoneBook:=make(map[string]string,entries)
	scanner:=bufio.NewScanner(os.Stdin)

	for i:=0;i<entries;i++{
		var key,value string
		fmt.Scanf("%s %s",&key,&value)
		myphoneBook[key]=value
	}

	for scanner.Scan(){
		query:=scanner.Text()
		if val,ok:=myphoneBook[query];ok{
			fmt.Printf("%s=%s\n",query,val)
		} else{
			fmt.Println("Not found");
		}
	}
}