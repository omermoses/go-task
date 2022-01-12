package structures

type City struct {
	name            string
	numberOfOffices int
	PostOffices     map[string]*PostOffice
}

func CreateNewCity(cityName string) *City {
	newCity := &City{
		name:            cityName,
		numberOfOffices: 0,
		PostOffices:     make(map[string]*PostOffice),
	}
	return newCity
}
