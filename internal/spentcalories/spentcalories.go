package spentcalories

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	ErrInvalidInput = errors.New("invalid format of input was provided")
	ErrConvToInt    = errors.New("couldn't convert to an integer")
	ErrSubZeroValue = errors.New("can't be less or equal to zero")
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		err := fmt.Errorf("there was a problem when calling \"parseTraining()\" with argument \"%s\": %w\n", data, ErrInvalidInput)
		return 0, "", 0, err
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		err = fmt.Errorf("there was a problem when calling \"parseTraining()\" with this argument's part \"%s\": %w\n", parts[0], ErrConvToInt)
		return 0, "", 0, err
	}
	if steps <= 0 {
		err = fmt.Errorf("there was a problem when calling \"parseTraining()\" with this argument's part \"%s\": %w\n", parts[0], ErrSubZeroValue)
		return 0, "", 0, err
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		err = fmt.Errorf("there was a problem when calling \"parseTraining()\" with this argument's part \"%s\": %w\n", parts[1], ErrInvalidInput)
		return 0, "", 0, err
	}

	return steps, parts[1], duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}
