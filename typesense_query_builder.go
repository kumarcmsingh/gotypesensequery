package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Constants for field types and operators
const (
	TextType     = "text"
	NumberType   = "number"
	DateType     = "date"
	DateTimeType = "datetime"
	// Add other field types as needed
)

const (
	// Add other operators as needed
	OperatorIs       = "is"
	OperatorIsNot    = "isnot"
	OperatorIn       = "in"
	OperatorContains = "contains"
	OperatorLike     = "like"
	OperatorNoData   = "nodata"
	OperatorHasData  = "hasdata"
	OperatorBetween  = "between"
)

// FilterRequest represents a filter request for Typesense query
type FilterRequest struct {
	Field           string    `json:"field"`
	Operator        string    `json:"operator"`
	TextValue       string    `json:"text_value"`
	TextArray       []string  `json:"text_array"`
	NumberValue     float64   `json:"number_value"`
	NumberRange     []float64 `json:"number_range"`
	CustomDateRange []float64 `json:"custom_date_range"`
	HumanDate       string    `json:"human_date"`
}

// GenerateTypesenseQuery generates a Typesense-supported query string based on filter requests
func GenerateTypesenseQuery(filterRequests []FilterRequest) (string, error) {
	var conditions []string

	// Iterate through filter requests and build conditions
	for _, filter := range filterRequests {
		condition := generateCondition(filter)
		if condition != "" {
			conditions = append(conditions, condition)
		}
	}

	// Combine conditions with AND
	queryString := strings.Join(conditions, " AND ")
	log.Printf("Generated Typesense Query: %s", queryString)

	return queryString, nil
}

// generateCondition generates a condition for a single filter
func generateCondition(filter FilterRequest) string {
	switch filter.Field {
	case TextType:
		return handleTextConditions(filter)
	case NumberType:
		return handleNumberConditions(filter)
	case DateType, DateTimeType:
		return handleDateConditions(filter)
	// Add other field types as needed
	default:
		log.Printf("Unsupported field type: %s", filter.Field)
		return ""
	}
}

// handleTextConditions handles text field conditions
func handleTextConditions(filter FilterRequest) string {
	switch filter.Operator {
	case OperatorIs, OperatorIsNot, OperatorIn, OperatorContains, OperatorLike, OperatorNoData, OperatorHasData:
		return fmt.Sprintf("%s:%s:%s", filter.Field, filter.Operator, filter.TextValue)
	default:
		log.Printf("Unsupported operator for text field: %s", filter.Operator)
		return ""
	}
}

// handleNumberConditions handles number field conditions
func handleNumberConditions(filter FilterRequest) string {
	switch filter.Operator {
	case OperatorBetween, OperatorHasData, OperatorNoData:
		if filter.Operator == OperatorBetween {
			return fmt.Sprintf("%s:%s:[%.2f,%.2f]", filter.Field, filter.Operator, filter.NumberRange[0], filter.NumberRange[1])
		}
		return fmt.Sprintf("%s%s%.2f", filter.Field, filter.Operator, filter.NumberValue)
	default:
		log.Printf("Unsupported operator for number field: %s", filter.Operator)
		return ""
	}
}

// handleDateConditions handles date field conditions
func handleDateConditions(filter FilterRequest) string {
	switch filter.Operator {
	case OperatorNoData, OperatorHasData:
		return fmt.Sprintf("%s:%s", filter.Field, filter.Operator)
	default:
		// Convert human date to Typesense datetime
		datetimeCondition := convertHumanDateToTypesenseDatetime(filter.Field, filter.HumanDate)

		// Check if conversion was successful
		if datetimeCondition != "" {
			return datetimeCondition
		}

		// Log an error for unsupported operator
		log.Printf("Unsupported operator for date field: %s", filter.Operator)
		return ""
	}
}

// convertHumanDateToTypesenseDatetime converts a human-readable date to Typesense datetime
// Returns the generated datetime condition string
func convertHumanDateToTypesenseDatetime(field string, humanDate string) string {
	currentTime := time.Now()

	switch humanDate {
	case "yesterday":
		yesterday := currentTime.AddDate(0, 0, -1)
		return generateDatetimeCondition(field, yesterday, yesterday.AddDate(0, 0, 1))
	case "tomorrow":
		tomorrow := currentTime.AddDate(0, 0, 1)
		return generateDatetimeCondition(field, tomorrow.AddDate(0, 0, -1), tomorrow)
	case "last_week":
		startOfWeek := getStartOfWeek(currentTime)
		endOfWeek := startOfWeek.AddDate(0, 0, 7).Add(-time.Second)
		return generateDatetimeCondition(field, startOfWeek, endOfWeek)
	case "next_week":
		startOfWeek := getStartOfWeek(currentTime.AddDate(0, 0, 7))
		endOfWeek := startOfWeek.AddDate(0, 0, 7).Add(-time.Second)
		return generateDatetimeCondition(field, startOfWeek, endOfWeek)
	case "last_month":
		startOfMonth := time.Date(currentTime.Year(), currentTime.Month()-1, 1, 0, 0, 0, 0, currentTime.Location())
		endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)
		return generateDatetimeCondition(field, startOfMonth, endOfMonth)
	case "next_month":
		startOfMonth := time.Date(currentTime.Year(), currentTime.Month()+1, 1, 0, 0, 0, 0, currentTime.Location())
		endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)
		return generateDatetimeCondition(field, startOfMonth, endOfMonth)
	case "last_year":
		startOfYear := time.Date(currentTime.Year()-1, 1, 1, 0, 0, 0, 0, currentTime.Location())
		endOfYear := startOfYear.AddDate(1, 0, 0).Add(-time.Second)
		return generateDatetimeCondition(field, startOfYear, endOfYear)
	case "next_year":
		startOfYear := time.Date(currentTime.Year()+1, 1, 1, 0, 0, 0, 0, currentTime.Location())
		endOfYear := startOfYear.AddDate(1, 0, 0).Add(-time.Second)
		return generateDatetimeCondition(field, startOfYear, endOfYear)
	case "today":
		startOfDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
		endOfDay := startOfDay.AddDate(0, 0, 1).Add(-time.Second)
		return generateDatetimeCondition(field, startOfDay, endOfDay)
	case "this_week":
		startOfWeek := getStartOfWeek(currentTime)
		endOfWeek := startOfWeek.AddDate(0, 0, 7).Add(-time.Second)
		return generateDatetimeCondition(field, startOfWeek, endOfWeek)
	case "this_month":
		startOfMonth := time.Date(currentTime.Year(), currentTime.Month(), 1, 0, 0, 0, 0, currentTime.Location())
		endOfMonth := startOfMonth.AddDate(0, 1, 0).Add(-time.Second)
		return generateDatetimeCondition(field, startOfMonth, endOfMonth)
	case "this_year":
		startOfYear := time.Date(currentTime.Year(), 1, 1, 0, 0, 0, 0, currentTime.Location())
		endOfYear := startOfYear.AddDate(1, 0, 0).Add(-time.Second)
		return generateDatetimeCondition(field, startOfYear, endOfYear)
	default:
		log.Printf("Unsupported human_date value: %s", humanDate)
		return ""
	}
}

// generateDatetimeCondition generates a Typesense datetime condition string
func generateDatetimeCondition(field string, start time.Time, end time.Time) string {
	return fmt.Sprintf("%s:>=:%s AND %s:<=:%s", field, start.Format(time.RFC3339), field, end.Format(time.RFC3339))
}

// getStartOfWeek returns the start of the week for a given time
func getStartOfWeek(t time.Time) time.Time {
	weekday := t.Weekday()
	startOfWeek := t.AddDate(0, 0, -int(weekday))
	return time.Date(startOfWeek.Year(), startOfWeek.Month(), startOfWeek.Day(), 0, 0, 0, 0, startOfWeek.Location())
}
