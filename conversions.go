package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func getCurrentTime() string {
    current_time := time.Now()
    fmt.Println("current time is: ", current_time.Format("2006-01-02 15:04:05 Monday"))
    return current_time.Format("2006-01-02 15:04:05")
}

func getTimeWithZone() string {
    location, _ := time.LoadLocation("America/New_York")
    current_time := time.Now().In(location)
    fmt.Println("current time is: ", time.Date(2020, 10, 10, 23, 0, 0, 0, time.UTC));
    return current_time.Format("2006-01-02 15:04:05")
}

func main() {
    getTimeWithZone();
    getCurrentTime();
   println("welcome to my shopp");
   print("please enter rating 1 to 5: ");
    rating  :=bufio.NewReader(os.Stdin);
    result,_ := rating.ReadString('\n');

    numRatin,err  := strconv.ParseFloat(strings.TrimSpace(result), 64);
  
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(numRatin+1)
    }
   
}