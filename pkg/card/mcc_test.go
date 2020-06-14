package card_test

import (
	"fmt"
	"github.com/artrey/bgo-std-card/pkg/card"
)

func ExampleTranslateMCC() {
	fmt.Println(card.TranslateMCC("5411"))
	fmt.Println(card.TranslateMCC("5533"))
	fmt.Println(card.TranslateMCC("5812"))
	fmt.Println(card.TranslateMCC("5912"))
	fmt.Println(card.TranslateMCC("7788"))
	// Output:
	// Супермаркеты
	// Автоуслуги
	// Рестораны
	// Аптеки
	// Категория не указана
}
