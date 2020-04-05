package datalayer

import "time"

type WorkBreak struct {
    Description string
    Duration time.Duration
}

// WorkDayTimePeriod Represents a period of time at work to support normal 9-5
// and split shifts
type WorkDayTimePeriod struct {
    // StartAt is the time of starting work
    StartAt time.Time
    // FinishAt is the time of finishing work
    FinishAt time.Time
    Breaks []*WorkBreak
}

type WorkDay struct {
    Hours []WorkDayTimePeriod
    Date time.Time
}

// WorkSchedule represents a household members work schedule
type WorkSchedule struct {
    Schedule [7]*WorkDay
}

type HouseholdMember struct {
    id int
    Email string
    Name string
    Birthday time.Time
    WorkSchedule *WorkSchedule
}

