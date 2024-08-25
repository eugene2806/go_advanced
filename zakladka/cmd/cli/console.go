package cli

import (
	"fmt"
	"zakladka/model"
	"zakladka/service"
)

func ConsoleInput() {
	fmt.Println("Приложение для закладок")

	variant := getMenu()
	bookmarkService := service.BookmarkService{}
	bookMark := model.Bookmark{
		Name: "",
		Text: "",
	}
menu:
	for {
		switch variant {
		case 1:
			fmt.Println("Просмотр закладок")
			if len(model.DB) <= 0 {
				fmt.Println("Закладок нет")
				variant = getMenu()
				break
			}
			bookmarkService.ViewBookMarks(model.DB)
			variant = getMenu()
		case 2:
			fmt.Println("Введите название закладки")
			fmt.Scan(&bookMark.Name)
			fmt.Println("Введите тест закладки")
			fmt.Scan(&bookMark.Text)
			msg, _ := bookmarkService.AddBookMarks(model.DB, bookMark.Name, bookMark.Text)
			fmt.Println(msg)
			variant = getMenu()
		case 3:
			fmt.Println("Введите название закладки")
			fmt.Scan(&bookMark.Name)
			msg, _ := bookmarkService.DeleteBookMarks(model.DB, bookMark.Name)
			fmt.Println(msg)
			variant = getMenu()
		case 4:
			break menu

		}
	}

}

func getMenu() int {
	fmt.Println()
	fmt.Println("Введите значение")
	fmt.Println("- 1. Посмотреть закладки")
	fmt.Println("- 2. Добавить закладку")
	fmt.Println("- 3. Удалить закладку")
	fmt.Println("- 4. Выход")
	var number int
	fmt.Scan(&number)
	return number
}
