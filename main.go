package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	humanTime := formatHumanTime(now)
	fmt.Println(humanTime)
}

// formatHumanTime converts a time.Time to human-readable format
func formatHumanTime(t time.Time) string {
	hour := t.Hour()
	minute := t.Minute()

	// Convert 24-hour to 12-hour format
	displayHour := hour
	if displayHour == 0 {
		displayHour = 12
	} else if displayHour > 12 {
		displayHour = displayHour - 12
	}

	// Determine time of day suffix
	timeOfDay := getTimeOfDay(hour)

	// Round minutes to nearest 5 for more natural speech
	roundedMinute := ((minute + 2) / 5) * 5
	if roundedMinute >= 60 {
		roundedMinute = 0
		displayHour++
		if displayHour > 12 {
			displayHour = 1
		}
	}

	// Generate human-readable time string
	return generateTimeString(displayHour, roundedMinute, timeOfDay)
}

// getTimeOfDay returns appropriate time of day description
func getTimeOfDay(hour int) string {
	switch {
	case hour >= 5 && hour < 12:
		return "morning"
	case hour >= 12 && hour < 17:
		return "afternoon"
	case hour >= 17 && hour < 21:
		return "evening"
	default:
		return "night"
	}
}

// generateTimeString creates the natural language time string
func generateTimeString(hour int, minute int, timeOfDay string) string {
	hourWord := numberToWord(hour)

	// Handle special cases
	switch minute {
	case 0:
		// On the hour
		if hour == 12 {
			return getVariation([]string{
				"It's noon.",
				"It's twelve o'clock.",
				"It's midday.",
			})
		}
		return getVariation([]string{
			fmt.Sprintf("It's %s o'clock.", hourWord),
			fmt.Sprintf("It's %s.", hourWord),
			fmt.Sprintf("About %s o'clock.", hourWord),
		})

	case 15:
		// Quarter past
		if hour == 12 && timeOfDay == "afternoon" {
			return getVariation([]string{
				"A quarter past noon.",
				"About a quarter past twelve.",
				"Fifteen minutes past noon.",
			})
		}
		return getVariation([]string{
			fmt.Sprintf("A quarter past %s.", hourWord),
			fmt.Sprintf("About a quarter past %s.", hourWord),
			fmt.Sprintf("Fifteen minutes past %s.", hourWord),
		})

	case 30:
		// Half past
		if hour == 12 && timeOfDay == "afternoon" {
			return getVariation([]string{
				"Half past noon.",
				"About half past twelve.",
				"Thirty minutes past noon.",
			})
		}
		return getVariation([]string{
			fmt.Sprintf("Half past %s.", hourWord),
			fmt.Sprintf("About half past %s.", hourWord),
			fmt.Sprintf("Almost half past %s.", hourWord),
		})

	case 45:
		// Quarter to next hour
		nextHour := hour + 1
		if nextHour > 12 {
			nextHour = 1
		}
		nextHourWord := numberToWord(nextHour)

		if nextHour == 12 {
			return getVariation([]string{
				"A quarter to noon.",
				"About a quarter to twelve.",
				"Fifteen minutes to noon.",
			})
		}
		return getVariation([]string{
			fmt.Sprintf("A quarter to %s.", nextHourWord),
			fmt.Sprintf("About a quarter to %s.", nextHourWord),
			fmt.Sprintf("Almost a quarter to %s.", nextHourWord),
		})

	default:
		// Other minutes
		if minute < 30 {
			// Minutes past the hour
			minuteWord := minutesToWords(minute)
			if hour == 12 && timeOfDay == "afternoon" {
				return getVariation([]string{
					fmt.Sprintf("%s past noon.", minuteWord),
					fmt.Sprintf("About %s past twelve.", minuteWord),
				})
			}
			return getVariation([]string{
				fmt.Sprintf("%s past %s.", minuteWord, hourWord),
				fmt.Sprintf("About %s past %s.", minuteWord, hourWord),
				fmt.Sprintf("Just %s past %s.", minuteWord, hourWord),
			})
		} else {
			// Minutes to the next hour
			minutesToNext := 60 - minute
			minuteWord := minutesToWords(minutesToNext)
			nextHour := hour + 1
			if nextHour > 12 {
				nextHour = 1
			}
			nextHourWord := numberToWord(nextHour)

			if nextHour == 12 {
				return getVariation([]string{
					fmt.Sprintf("%s to noon.", minuteWord),
					fmt.Sprintf("About %s to twelve.", minuteWord),
				})
			}
			return getVariation([]string{
				fmt.Sprintf("%s to %s.", minuteWord, nextHourWord),
				fmt.Sprintf("About %s to %s.", minuteWord, nextHourWord),
				fmt.Sprintf("Almost %s to %s.", minuteWord, nextHourWord),
			})
		}
	}
}

// numberToWord converts hour numbers to words
func numberToWord(num int) string {
	words := map[int]string{
		1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six",
		7: "seven", 8: "eight", 9: "nine", 10: "ten", 11: "eleven", 12: "twelve",
	}
	return words[num]
}

// minutesToWords converts minute values to descriptive words
func minutesToWords(minutes int) string {
	switch minutes {
	case 5:
		return "Five minutes"
	case 10:
		return "Ten minutes"
	case 20:
		return "Twenty minutes"
	case 25:
		return "Twenty-five minutes"
	case 35:
		return "Twenty-five minutes"
	case 40:
		return "Twenty minutes"
	case 50:
		return "Ten minutes"
	case 55:
		return "Five minutes"
	default:
		return fmt.Sprintf("%d minutes", minutes)
	}
}

// getVariation returns a random variation from the provided options
func getVariation(options []string) string {
	if len(options) == 0 {
		return ""
	}
	return options[rand.Intn(len(options))]
}

func init() {
	// Seed the random number generator for variation in output
	rand.Seed(time.Now().UnixNano())
}
