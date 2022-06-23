package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)



func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)
	binValue:=fmt.Sprintf("%b",n)
	max:=-1;
	count:=0

	for i:=0;i<len(binValue);i++{

		if binValue[i]=='1'{
			count++
			if count>max{
				max=count
			}
		}
		if binValue[i]=='0'{
			count=0
		}
	}
	fmt.Println(max)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
