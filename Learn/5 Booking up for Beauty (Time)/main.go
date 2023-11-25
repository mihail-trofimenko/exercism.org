package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(Schedule("11/28/1984 2:02:02"))
}

// Schedule returns a time.Time from a string containing a date.
// Schedule("7/25/2019 13:45:00")
// => 2019-07-25 13:45:00 +0000 UTC
func Schedule(date string) time.Time {
	time, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		fmt.Println("Incoming time: ", date)
		fmt.Println("Error detected: ", err)
	}
	return time

}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	dt, err := time.Parse("January 2, 2006 3:04:05", date)
	if err != nil {
		fmt.Println("Incoming time: ", date)
		fmt.Println("Error detected: ", err)
	}
	return !dt.After(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// Парсим строку времени
	fmt.Println("Incoming string:", date)
	appointmentTime, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	// Определяем временной диапазон для послеполуденного времени (12:00 PM - 5:59:59 PM)

	afternoonStartTime := time.Date(appointmentTime.Year(), appointmentTime.Month(), appointmentTime.Day(), 12, 0, 0, 0, time.UTC)
	afternoonEndTime := afternoonStartTime.Add(6 * time.Hour) // 12 + 6 = 18

	// Проверяем, находится ли время записи в послеполуденной временном диапазоне
	isAfternoon := appointmentTime.After(afternoonStartTime) && appointmentTime.Before(afternoonEndTime)

	return isAfternoon
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// Парсим строку времени
	fmt.Println("Incoming string:", date)
	appointmentTime, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}

	// Возвращаем описание даты и времени
	return fmt.Sprintf("You have an appointment on %s", appointmentTime.Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	// Задаем дату открытия салона
	openingDate := time.Date(2020, time.September, 15, 0, 0, 0, 0, time.UTC)

	// Получаем текущий год
	currentYear := time.Now().Year()

	// Возвращаем дату открытия салона в текущем году
	return time.Date(currentYear, openingDate.Month(), openingDate.Day(), 0, 0, 0, 0, time.UTC)
}
