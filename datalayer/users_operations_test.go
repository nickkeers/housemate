package datalayer

import (
    "fmt"
    "testing"
    "time"
)

func TestGetUserById(t *testing.T) {
    db, err := NewPostgresDataAdapter("postgres://localhost:5432/housemate?sslmode=disable")
    if err != nil {
        t.Error(err)
    }

    member, err := db.GetHouseholdMemberById(1)
    if err != nil {
        t.Error(err)
    }

    expected := HouseholdMember{id: 1, Email: "nick@housemate.com", Birthday: time.Date(1994, 1, 13, 0, 0, 0, 0, time.UTC),
        Name: "Nick"}

    if *member != expected {
        fmt.Printf("%+v vs %+v", *member, expected)
        t.Error("expected values not returned")
    }
}
