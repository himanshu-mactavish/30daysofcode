package main

import "fmt"


func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
 var testNum int
 fmt.Scanf("%d",&testNum)
 for i:=0;i<int(testNum);i++{
     var mystring string
     fmt.Scanf("%s",&mystring)
     start:=""
     end:=""
     for j:=0;j<len(mystring);j++{
         if j%2==0{
            start+=string(mystring[j]);
         } else{
            end+=string(mystring[j])    
         }
         
     }
     fmt.Print(start+" ")
     fmt.Println(end)
 }

}