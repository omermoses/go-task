package structures

type PostOffice struct {
	name        string
	max_weight  int
	min_weight  int
	mailPackage map[int]*MailPackage
}

func CreateNewPostOffice(cityName string) *PostOffice {
	newPostOffice := &PostOffice{
		name:        cityName,
		max_weight:  0,
		min_weight:  0,
		mailPackage: make(map[int]*MailPackage),
	}
	return newPostOffice
}
