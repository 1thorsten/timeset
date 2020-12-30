package handleTime

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

// SetSystemDate sets the system date (windows)
func SetSystemDate(newTime time.Time) error {
	year2digits := strconv.Itoa(newTime.Year())[2:4]
	dateString := fmt.Sprintf("%02d-%02d-%s", newTime.Day(), newTime.Month(), year2digits)
	fmt.Printf("Setting system date to: %s\n", dateString)

	out, err := exec.Command("cmd.exe", "/c", "date", dateString).Output()
	if err != nil {
		return fmt.Errorf("could not set date. %s\n", err.Error())
	}

	out, _ = exec.Command("cmd.exe", "/c", "date", "/T", dateString).Output()
	fmt.Printf("Current date is: %s\n", out)

	return nil
}

// SetSystemTime sets the system time (windows)
func SetSystemTime(newTime time.Time) error {
	timeString := fmt.Sprintf("%02d:%02d:%02d", newTime.Hour(), newTime.Minute(), newTime.Second())
	fmt.Printf("Setting system time to: %s\n", timeString)

	out, err := exec.Command("cmd.exe", "/c", "time", timeString).Output()
	if err != nil {
		return fmt.Errorf("could not set time. %s\n", err.Error())
	}

	out, _ = exec.Command("cmd.exe", "/c", "time", "/T", timeString).Output()
	fmt.Printf("Current time is: %s", out)

	return nil
}
