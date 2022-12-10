package windirection

var cardinalEs = [17]Cardinal{
	{"N", "Norte"},
	{"NNO", "Nor Noreste"},
	{"NE", "Noreste"},
	{"ENE", "Este Noreste"},
	{"E", "Este"},
	{"ESE", "Este Sureste"},
	{"SE", "Sureste"},
	{"SSE", "Sur Sureste"},
	{"S", "Sur"},
	{"SSO", "Sur Suroeste"},
	{"SO", "Suroeste"},
	{"OSO", "Oeste Suroeste"},
	{"O", "Oeste"},
	{"ONO", "Oeste Noroeste"},
	{"NO", "Noroeste"},
	{"NNO", "Nor Noroeste"},
}

var cardinalEn = [17]Cardinal{
	{"N", "North"},
	{"NNE", "North-northeast"},
	{"NE", "Northeast"},
	{"ENE", "East-northeast"},
	{"E", "East"},
	{"ESE", "East-southeast"},
	{"SE", "Southeast"},
	{"SSE", "South-southeast"},
	{"S", "South"},
	{"SSW", "South-southwest"},
	{"SW", "Southwest"},
	{"WSW", "West-southwest"},
	{"W", "West"},
	{"NWN", "West-northwest"},
	{"NW", "Northwest"},
	{"NNW", "North-northwest"},
}

func getLangCardinals(lang string) *[17]Cardinal {
	if lang == "es" {
		return &cardinalEs
	}
	return &cardinalEn
}
