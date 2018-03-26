package main

import (
	"encoding/json"
	"net/http"
)

func listSpendings(w http.ResponseWriter, r *http.Request) {
	var days = parse("spendings.xlsx")
	var daysB, _ = json.Marshal(days)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(daysB)
}

func total(w http.ResponseWriter, r *http.Request) {
	var days = parse("spendings.xlsx")
	var total = 0
	for _, day := range days {
		for _, spending := range day.Spendings {
			total += spending.Value
		}
	}
	var totalB, _ = json.Marshal(total)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(totalB)
}

func average(w http.ResponseWriter, r *http.Request) {
	var days = parse("spendings.xlsx")
	var total = 0
	for _, day := range days {
		for _, spending := range day.Spendings {
			if day.Weekday != "Fix" {
				total += spending.Value
			}
		}
	}
	var avgB, _ = json.Marshal(total / len(days))
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(avgB)
}

func favouritePlace(w http.ResponseWriter, r *http.Request) {
	var days = parse("spendings.xlsx")
	var spendingsByPlace = make(map[string]int)
	var favouritePlace Spending
	for _, day := range days {
		for _, spending := range day.Spendings {
			if day.Weekday != "Fix" {
				spendingsByPlace[spending.Location] += spending.Value
			}
		}
	}

	for place, value := range spendingsByPlace {
		if favouritePlace.Value == 0 || value > favouritePlace.Value {
			favouritePlace.Location = place
			favouritePlace.Value = value
		}
	}
	var favouriteB, _ = json.Marshal(favouritePlace)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write(favouriteB)
}

func main() {
	http.HandleFunc("/", listSpendings)
	http.HandleFunc("/favouritePlace", favouritePlace)
	http.HandleFunc("/averagePerDay", average)
	http.HandleFunc("/total", total)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
