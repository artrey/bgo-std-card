package card

func TranslateMCC(code MCC) string {
	mcc := map[MCC]string{
		"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"5812": "Рестораны",
		"5912": "Аптеки",
	}

	if category, ok := mcc[code]; ok {
		return category
	}
	return "Категория не указана"
}
