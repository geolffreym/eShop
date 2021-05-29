package time

import (
	"fmt"
	"time"
)

func Wait(timer time.Duration) {
	//Before wait
	//fmt.Println(int32(time.Now().Unix()))
	//Waiting
	fmt.Printf("Waiting %s \n", timer)
	time.Sleep(timer)
	//After wait
	//fmt.Println(int32(time.Now().Unix()))
}
