package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	Location    string    `"binding:"required" json:"location"`
	DateTime    time.Time `binding:"required" json:"date_time"`
	UserID      int       `json:"user_id"`

}


var events = []Event{

}

func (e Event) SaveData() {
	events = append(events, e)
}


func GetAllEvents() []Event {
	return events
}