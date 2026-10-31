package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse записывает данные прогулки из строки формата "678,0h50m".
// При ошибке поля структуры остаются без изменений.
func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return fmt.Errorf("daysteps: Parse: ожидалось 2 поля, получено %d", len(data))
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("daysteps: Parse: не удалось разобрать количество шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("daysteps: Parse: количество шагов должно быть больше нуля")
	}

	duration, err := time.ParseDuration(data[1])
	if err != nil {
		return fmt.Errorf("daysteps: Parse: не удалось разобрать длительность: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("daysteps: Parse: длительность должна быть больше нуля")
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

// ActionInfo возвращает отчет о прогулке или пустую строку и ошибку расчета калорий.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("daysteps: ActionInfo: не удалось рассчитать калории: %w", err)
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
