package main

import (
	"fmt"
	wapi "github.com/iamacarpet/go-win64api"
	"time"
)

func main() {
	for true {
		list, _ := wapi.ProcessList()
		for _, i := range list {
			fmt.Println(i.Executable)
			if i.Executable == "LeagueClient.exe" {
				wapi.ProcessKill(uint32(i.Pid))
				fomt.Printf("SAVED YOU! GOOD LUCK MY FRIEND!\n")
			}
		}
		time.Sleep(2500)
	}
}
