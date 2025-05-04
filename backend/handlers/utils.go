package handlers

import (
	"strconv"
	"time"
)

func parseDate(value string) *time.Time {
	if value == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", value)
	if err != nil {
		return nil
	}
	return &t
}

func parseInt(value string) *int {
	if value == "" {
		return nil
	}
	i, err := strconv.Atoi(value)
	if err != nil {
		return nil
	}
	return &i
}

func parseFloat(value string) *float64 {
	if value == "" {
		return nil
	}
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return nil
	}
	return &f
}

func parseString(value string) *string {
	if value == "" {
		return nil
	}
	return &value
}
