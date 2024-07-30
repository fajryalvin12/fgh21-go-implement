package mainprocess

import "fmt"

type Bio struct {
	name     string
	skillSet []string
}
type SkillSet struct {
	name string
}

func Struct() {
	var bio []Bio = []Bio{
		{
			name: "Anies",
			skillSet: []string{
				"Javascript",
				"PHP",
			},
		},
		{
			name: "Prabowo",
			skillSet: []string{
				"JavaScript",
				"PHP",
				"Golang",
				"React",
			},
		},
		{
			name: "Ganjar",
			skillSet: []string{
				"JavaScript",
				"PHP",
				"Golang",
				"React",
			},
		},
		{
			name: "Jokowi",
			skillSet: []string{
				"JavaScript",
				"PHP",
				"Golang",
				"React",
			},
		},
		{
			name: "Luhut",
			skillSet: []string{
				"JavaScript",
				"Golang",
				"PHP",
				"React",
			},
		},
	}
	fmt.Println(bio[4].skillSet[2])
	fmt.Println(bio[0].skillSet[0])
}