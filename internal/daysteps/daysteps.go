package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

var (
	ErrInvalidInput = errors.New("invalid format of input was provided")
	ErrConvToInt    = errors.New("couldn't convert to an integer")
	ErrSubZeroValue = errors.New("can't be less or equal to zero")
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
		err := fmt.Errorf("there was a problem when calling \"parsePackage()\" with argument \"%s\": %w\n", data, ErrInvalidInput)
		return 0, 0, err
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		err = fmt.Errorf("there was a problem when calling \"parsePackage()\" with this argument's part \"%s\": %w\n", parts[0], ErrConvToInt)
		return 0, 0, err
	}
	if steps <= 0 {
		err = fmt.Errorf("there was a problem when calling \"parsePackage()\" with this argument's part \"%s\": %w\n", parts[0], ErrSubZeroValue)
		return 0, 0, err
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		err = fmt.Errorf("there was a problem when calling \"parsePackage()\" with this argument's part \"%s\": %w\n", parts[1], ErrInvalidInput)
		return 0, 0, err
	}
	if duration <= 0 {
		err = fmt.Errorf("there was a problem when calling \"parsePackage()\" with this argument's part \"%s\": %w\n", parts[1], ErrSubZeroValue)
		return 0, 0, err
	}

	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distInM := stepLength * float64(steps)
	distInKm := distInM / mInKm

	burnedCals, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distInKm, burnedCals)
}
