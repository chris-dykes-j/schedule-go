package main

import "fmt"

func main() {
    const TimeDivision = 24 * 2 // Half hour blocks.
    sched := Schedule{}
    sched.TimeAvailable = TimeDivision
    fmt.Println(sched.TimeAvailable)
}

type Day int

const (
    Monday Day = iota
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
)

func (d Day) String() string {
    switch d {
    case Monday:
        return "Monday"
    case Tuesday:
        return "Tuesday"
    case Wednesday:
        return "Wednesday"
    case Thursday:
        return "Thursday"
    case Friday:
        return "Friday"
    case Saturday:
        return "Saturday"
    case Sunday:
        return "Sunday"
    }
    return "Unknown"
}

type Activity struct {
    Name string
    Description string
}

type TimeBlock struct {
    Label string
    Value int
}

type ActivityBlock struct {
    TimeBlock TimeBlock
    Activity Activity
}

type Schedule struct {
    Type string // Week schedule, or Saturday, Sunday?
    Days []Day // What days the schedule occurs.
    TimeAvailable int
    Activities []ActivityBlock // Activities for the day.
}

