package daysteps

import (
	"time"
	"strings"
	"strconv"
	"errors"
)

var (
	ErrInvalidInput = errors.New("invalid format of input was provided")
	ErrConvToInt = errors.New("couldn't convert to an integer")
	ErrZeroValue = errors.New("can't be zero")
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("there was a problem when calling \"parsePackage()\" with argument \"%s\": %w\n", data, ErrInvalidInput)
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("there was a problem when calling \"parsePackage()\" with this argument's part \"%s\": %w\n", parts[0], ErrConvToInt)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("there was a problem when calling \"parsePackage()\" with this argument's part \"%s\": %w\n", parts[0], ErrZeroValue)
	}
	
	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("there was a problem when calling \"parsePackage()\" with this argument's part \"%s\": %w\n", parts[1], ErrInvalidInput)
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
}
