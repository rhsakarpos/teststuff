package main

import (
	"fmt"
	"os/exec"
	"regexp"
)

func main(){

	resp, err := exec.Command("iostat").Output()
	strResp := string(resp)
	if err!=nil {
		fmt.Print(err)
		return
	}




	a := regexp.MustCompile(" *")

	strArr := a.Split(strResp,-1)

	fmt.Println(strArr)
}
