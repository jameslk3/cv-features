package data

import (
	"encoding/json"
	"strconv"
	"fmt"
	"io"
	"os"
)

// Struct for JSON schedule file that is used to get days a player is playing
type WeekSchedule struct {
	StartDate     string           	   	  	 `json:"startDate"`
	EndDate       string           	      	 `json:"endDate"`
	GameSpan  	  int                     	 `json:"gameSpan"`
	TeamSchedules map[string]map[string]bool `json:"games"`
}

// Struct to organize the season schedule
type SeasonSchedule struct {
	Schedule map[string]WeekSchedule `json:"schedule"`
}

var ScheduleMap SeasonSchedule

func InitSchedule(path string) {
	LoadSchedule(path)
}

// Function to load schedule from JSON file into memory
func LoadSchedule(path string) {
	
	// Load JSON schedule file
	json_schedule, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening json schedule:", err)
	}
	defer json_schedule.Close()

	// Read the contents of the json_schedule file
	jsonBytes, err := io.ReadAll(json_schedule)
	if err != nil {
		fmt.Println("Error reading json_schedule:", err)
	}

	// Unmarshal the JSON data into ScheduleMap
	err = json.Unmarshal(jsonBytes, &ScheduleMap)
	if err != nil {
		fmt.Println("Error turning jsonBytes into map:", err)
	}
}

// Function to get the schedule for a specific week
func (s *SeasonSchedule) GetWeekSchedule(week string) WeekSchedule {
	return s.Schedule[week]
}

// Function to check if a team is playing on a specific day
func (s *SeasonSchedule) IsPlaying(week string, day int, team string) bool {
	if _, ok := s.Schedule[week].TeamSchedules[team][strconv.Itoa(day)]; ok {
		return true
	} else {
		return false
	}
}

func (w *WeekSchedule) GetStartDate() string {
	return w.StartDate
}

func (w *WeekSchedule) GetEndDate() string {
	return w.EndDate
}

func (w *WeekSchedule) GetGameSpan() int {
	return w.GameSpan
}