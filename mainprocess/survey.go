package mainprocess

import "fmt"

type Respondent struct {
	fullName     string
	age          int
	gender       string
	isSmoker     bool
	cigarVariant []string
}

func Survey() {
	var bio []Respondent = []Respondent{
		{
			fullName: "Admin",
			age:      21,
			gender:   "Perempuan",
			isSmoker: true,
			cigarVariant: []string{
				"Marlboro",
				"Gudang Garam",
				"Djarum",
			},
		},
		{
			fullName: "Admin juga",
			age:      21,
			gender:   "Perempuan",
			isSmoker: true,
			cigarVariant: []string{
				"Wismilak",
				"Sampoerna",
			},
		},
		{
			fullName: "Aku juga admin",
			age:      21,
			gender:   "Perempuan",
			isSmoker: true,
			cigarVariant: []string{
				"Gudang Garam",
				"Sampoerna",
			},
		},
	}
	for i := 0; i < len(bio); i++ {
		fmt.Println("Responden:", i+1)
		fmt.Println("Nama:", bio[i].fullName)
		fmt.Println("Umur:", bio[i].age)
		fmt.Println("Jenis Kelamin:", bio[i].gender)
		fmt.Println("Merokok tidak:", bio[i].isSmoker)
		fmt.Println("Rokok apa hayo:", bio[i].cigarVariant)
	}
}