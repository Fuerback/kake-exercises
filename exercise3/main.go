package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/*
This exercise involves analyzing server request latency logs. Each log entry is a string in the format:
"YYYY-MM-DD HH:MM:SS,Latency: XXXms"

You need to:
1. Fix the bug in the parseEntry function that's causing test failures
2. Implement the medianLatency function to calculate median latency values

The medianLatency function should:
- Take a slice of log entries and a timestamp
- If the timestamp is zero, calculate median latency for all logs
- If timestamp is provided, calculate median latency only for logs after that timestamp
- Return the median latency value in milliseconds

Note: The input logs are guaranteed to be sorted by timestamp.

## Example

```go
logs := []string{
    "2023-10-27 10:15:30,Latency: 120ms",
    "2023-10-27 10:16:15,Latency: 85ms",
    "2023-10-27 10:45:12,Latency: 92ms",
}

// Calculate median for all logs
median, _ := medianLatency(logs, time.Time{})  // returns 92

// Calculate median for logs after 10:16:00
since := time.Date(2023, 10, 27, 10, 16, 0, 0, time.UTC)
median, _ := medianLatency(logs, since)  // returns 92
```
*/

var latencyLogs = []string{
	"2023-10-27 10:15:30,Latency: 120ms",
	"2023-10-27 10:16:15,Latency: 85ms",
	"2023-10-27 10:45:12,Latency: 92ms",
	"2023-10-27 11:02:40,Latency: 250ms",
	"2023-10-27 11:35:20,Latency: 95ms",
	"2023-10-27 12:01:10,Latency: 310ms",
}

// parseEntry parses a latency log entry and returns the timestamp and latency.
func parseEntry(log string) (time.Time, int, error) {
	parts := strings.Split(log, ",")
	if len(parts) != 2 {
		return time.Time{}, 0, fmt.Errorf("invalid log format: %s", log)
	}

	timestampStr := strings.TrimSpace(parts[0])
	latencyStr := strings.TrimSpace(strings.TrimPrefix(parts[1], "Latency: "))

	timestamp, err := time.Parse("2006-01-02 15:04:05", timestampStr)
	if err != nil {
		return time.Time{}, 0, fmt.Errorf("invalid timestamp: %s", timestampStr)
	}

	latency, err := strconv.Atoi(latencyStr)
	if err != nil {
		return time.Time{}, 0, fmt.Errorf("invalid latency: %s", latencyStr)
	}

	return timestamp, latency, nil
}

// averageLatency calculates the average latency of all logs.
func averageLatency(logs []string) (int, error) {
	totalLatency := 0
	for _, log := range logs {
		_, latency, err := parseEntry(log)
		if err != nil {
			return 0, err
		}
		totalLatency += latency
	}
	return totalLatency / len(logs), nil
}

// medianLatency calculates the median latency for requests.
// If since is zero, calculate median for all logs.
// If since is provided, calculate median only for logs after that timestamp.
func medianLatency(logs []string, since time.Time) (int, error) {
	// TODO: Implement this function
	return 0, nil
}

func runTests() {
	fmt.Println("Running parseEntry and averageLatency tests...")
	testAverageLatency()
	fmt.Println("\nRunning medianLatency tests...")
	testMedianLatency()
}

func testAverageLatency() {
	expected := 158
	actual, err := averageLatency(latencyLogs)
	if err != nil {
		fmt.Printf("❌ Test Failed: %v\n", err)
		return
	}
	if actual != expected {
		fmt.Printf("❌ Test Failed: expected %d, got %d\n", expected, actual)
		return
	}
	fmt.Println("✅ Test Passed")
}

func testMedianLatency() {
	testCases := []struct {
		name     string
		since    time.Time
		expected int
	}{
		{
			name:     "all logs",
			since:    time.Time{},
			expected: 107,
		},
		{
			name:     "after 10:15:30",
			since:    time.Date(2023, 10, 27, 10, 15, 30, 0, time.UTC),
			expected: 95,
		},
	}

	for _, tc := range testCases {
		actual, err := medianLatency(latencyLogs, tc.since)
		if err != nil {
			fmt.Printf("❌ Test '%s' Failed: %v\n", tc.name, err)
			continue
		}
		if actual != tc.expected {
			fmt.Printf("❌ Test '%s' Failed: expected %d, got %d\n", tc.name, tc.expected, actual)
			continue
		}
		fmt.Printf("✅ Test '%s' Passed\n", tc.name)
	}
}

func main() {
	runTests()
}
