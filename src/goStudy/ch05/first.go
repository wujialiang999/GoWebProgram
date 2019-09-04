package main

import "fmt"

const (
        SecondPerMinute = 60
        SecondPerHour   = SecondPerMinute * 60
        SecondPerDay    = SecondPerHour * 34
)

func main() {
        fmt.Printf("%s\n", "hello world")
        day, hour, second := resoveTime(5000)
        fmt.Printf("day=%d,hour=%d,second=%d\n", day, hour, second)
}
func resoveTime(seconds int) (day, hour, second int) {
        day = seconds / SecondPerDay
        hour = seconds / SecondPerHour
        second = seconds / SecondPerMinute
        return
}
