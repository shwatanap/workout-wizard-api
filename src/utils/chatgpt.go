package utils

import (
	"regexp"
	"strings"

	"github.com/shwatanap/workout-wizard-api/src/domain/entity"
)

func ParseMenu(text string) *entity.Menu {
	me := &entity.Menu{Target: "arm", Comment: "HelloWorld"}
	re := regexp.MustCompile(`\d+\.\s`)
	workout_strs := re.Split(text, -1)[1:]
	for _, workout_str := range workout_strs {
		we := parseWorkout(workout_str)
		me.Workouts = append(me.Workouts, we)
	}
	return me
}

func parseWorkout(workout_str string) *entity.Workout {
	we := &entity.Workout{}
	tokens := strings.Split(workout_str, "\n- ")
	we.Name = tokens[0]
	we.Detail = strings.Join(tokens[1:], "\n")
	return we
}
