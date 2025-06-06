// Пакет daysteps отвечает за учёт активности в течение дня.
// Он собирает переданную информацию в виде строк, парсит их и выводит информацию о количестве шагов, пройденной дистанции и потраченных калориях.
package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

const (
	StepLength = 0.65 // длина шага в метрах
	mInKm      = 1000 // метров в километре
)

func parsePackage(data string) (int, time.Duration, error) {
	parts := strings.Split(data, ",")

	if len(parts) != 2 {
		return 0, 0, errors.New("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("step count convertation failed: %w", err)
	}

	if steps <= 0 {
		return 0, 0, errors.New("negative number of steps")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("time convertation failed: %w", err)
	}

	return steps, duration, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	distanceMeters := float64(steps) * StepLength

	distanceKilometers := distanceMeters / mInKm

	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps, distanceKilometers, calories)

	return result
}
