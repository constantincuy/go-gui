package theme

type Sheet struct {
	elements map[string]map[string]Property
}

func (sheet *Sheet) AddToElement(element string, prop Property) {
	if sheet.elements[element] == nil {
		sheet.elements[element] = make(map[string]Property)
	}

	sheet.elements[element][prop.Name] = prop
}

func (sheet *Sheet) Select(element string) map[string]Property {
	return sheet.elements[element]
}

func NewSheet() Sheet {
	return Sheet{elements: make(map[string]map[string]Property)}
}
