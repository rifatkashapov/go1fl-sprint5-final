package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse парсит строку с данными формата "3456,Ходьба,3h00m" и записывает данные в
// соответствующие поля структуры Training.
func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("trainings: Parse: ожидалось 3 поля, получено %d", len(data))
	}
	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("trainings: Parse: не удалось разобрать количество шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("trainings: Parse: количество шагов должно быть больше нуля")
	}
	duration, err := time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("trainings: Parse: не удалось разобрать длительность: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("trainings: Parse: длительность должна быть больше нуля")
	}
	t.Steps = steps
	t.TrainingType = data[1]
	t.Duration = duration
	return nil
}

// ActionInfo возвращает сведения о тренировке. Для неизвестного типа возвращается
// ошибка «неизвестный тип тренировки», а ошибки расчета калорий содержат причину
// из spentenergy: неположительное число шагов, вес, рост или длительность.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", fmt.Errorf("trainings: ActionInfo: неизвестный тип тренировки")
	}
	if err != nil {
		return "", fmt.Errorf("trainings: ActionInfo: не удалось рассчитать калории: %w", err)
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
