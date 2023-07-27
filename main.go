package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type WorkoutRecord struct {
	Workout_Name       string    `csv:"Workout Name,omitempty"`
	Workout_Start_Time time.Time `csv:"Workout Start Time,omitempty"`
	Workout_End_Time   time.Time `csv:"Workout End Time,omitempty"`
	Exercise_Name      string    `csv:"Exercise Name,omitempty"`
	Reps               int       `csv:"Reps,omitempty"`
	Date               string    `csv:"Date,omitempty"`
}

func parseDate(dateStr string) (time.Time, error) {
	parts := strings.Split(dateStr, " ")
	if len(parts) < 6 {
		return time.Time{}, fmt.Errorf("invalid date format: %s", dateStr)
	}

	dateTimeStr := parts[0] + " " + parts[1] + " " + parts[2] + " " + parts[3] + " " + parts[4] + " "
	fmt.Println(dateTimeStr)
	return time.Parse("Mon Jan 2 2006 15:04:05 -0700", dateTimeStr)
}

func createWorkoutList(data [][]string) []WorkoutRecord {
	var workout []WorkoutRecord
	var invalidRecords []int
	for i, line := range data {
		if i > 0 {
			reps, err := strconv.Atoi(line[4])
			if err != nil {
				invalidRecords = append(invalidRecords, i)
			}
			starttime, err := parseDate(line[1])
			if err != nil {
				invalidRecords = append(invalidRecords, i)
				continue
			}
			endtime, err := parseDate(line[2])
			if err != nil {
				invalidRecords = append(invalidRecords, i)
				continue
			}
			rec := WorkoutRecord{
				Workout_Name:       line[0],
				Workout_Start_Time: starttime,
				Workout_End_Time:   endtime,
				Exercise_Name:      line[3],
				Reps:               reps,
				Date:               line[5],
			}
			log.Printf("Constructed record: %+v", rec)

			workout = append(workout, rec)
		}
	}
	if len(invalidRecords) > 0 {
		log.Printf("Invalid records at line numbers: %v", invalidRecords)
	}
	return workout
}

func main() {
	f, err := os.Open("/home/wintermute/Downloads/workouts.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	workouts := createWorkoutList(data)

	fmt.Printf("%+v\n", workouts)
}
