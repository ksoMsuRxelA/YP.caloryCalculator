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

	duration, err := time.ParseDuration(parts[2])
	if err != nil || duration <= 0 {
		err = fmt.Errorf("there was a problem when calling \"parseTraining()\" with this argument's part \"%s\": %w\n", parts[1], ErrInvalidInput)
		return 0, "", 0, err
	}

	return steps, parts[1], duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := stepLengthCoefficient * height
	distInM := stepLength * float64(steps)
	return distInM / mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	dist := distance(steps, height)
	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, actionType, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	switch actionType {
	case "Бег":
		dist, speed := distance(steps, height), meanSpeed(steps, height, duration)
		burnedCals, err := RunningSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", actionType, duration.Hours(), dist, speed, burnedCals)
		return result, nil
	case "Ходьба":
		dist, speed := distance(steps, height), meanSpeed(steps, height, duration)
		burnedCals, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", actionType, duration.Hours(), dist, speed, burnedCals)
		return result, nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func isValidInputs(title string, steps int, weight, height float64, duration time.Duration) (bool, error) {
	var err error
	switch {
	case steps <= 0:
		err = fmt.Errorf("there was a problem when calling \"%s\" with this argument \"%d\": %w\n", title, steps, ErrSubZeroValue)
		return false, err
	case weight <= 0:
		err = fmt.Errorf("there was a problem when calling \"%s\" with this argument \"%f\": %w\n", title, weight, ErrSubZeroValue)
		return false, err
	case height <= 0:
		err = fmt.Errorf("there was a problem when calling \"%s\" with this argument \"%f\": %w\n", title, height, ErrSubZeroValue)
		return false, err
	case duration <= 0:
		err = fmt.Errorf("there was a problem when calling \"%s\" with this argument \"%d\": %w\n", title, duration, ErrSubZeroValue)
		return false, err
	}
	return true, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	_, err := isValidInputs("RunningSpenCalories()", steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	speed := meanSpeed(steps, height, duration)
	mins := duration.Minutes()
	burnedCals := (speed * mins * weight) / minInH
	return burnedCals, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	_, err := isValidInputs("WalkingSpentCalories()", steps, weight, height, duration)
	if err != nil {
		return 0, err
	}

	speed := meanSpeed(steps, height, duration)
	mins := duration.Minutes()
	burnedCals := (speed * mins * weight) / minInH

	return burnedCals * walkingCaloriesCoefficient, nil
}
