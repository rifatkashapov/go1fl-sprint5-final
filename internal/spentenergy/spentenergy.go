package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
	errTemplate                = "spentenergy: %s: значение должно быть больше нуля"
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	runningCalories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return runningCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf(errTemplate, "количество шагов")
	}
	if weight <= 0 {
		return 0, fmt.Errorf(errTemplate, "вес")
	}
	if height <= 0 {
		return 0, fmt.Errorf(errTemplate, "рост")
	}
	if duration <= 0 {
		return 0, fmt.Errorf(errTemplate, "длительность")
	}
	distance := MeanSpeed(steps, height, duration) * duration.Minutes()
	return (weight * distance) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}

	dst := Distance(steps, height) / duration.Hours()

	if dst == 0 {
		return 0
	}

	return dst
}

func Distance(steps int, height float64) float64 {
	if steps <= 0 {
		return 0
	}
	if height <= 0 {
		return 0
	}
	return (height * stepLengthCoefficient * float64(steps)) / mInKm
}
