package cli

import (
	"bookmarks/model"
	"bookmarks/service"
	"fmt"
)

type MyCliConsole struct {
	bookMarkService *service.BookmarkService
}

func (mc *MyCliConsole) GetAllBookMarks() (map[string]string, error) {
	return mc.bookMarkService.ViewBookMarks()
}

func (mc *MyCliConsole) AddBookMark(name, text string) error {
	return mc.bookMarkService.AddBookMarks(name, text)
}

func (mc *MyCliConsole) RemoveBookMark(name string) error {
	return mc.bookMarkService.DeleteBookMarks(name)
}

func ConsoleInput() {
	fmt.Println("Приложение для закладок")

	myCli := MyCliConsole{bookMarkService: &service.BookmarkService{}}

	bookMark := model.Bookmark{}

menu:
	for {
		variant := myCli.getMenu()
		switch variant {

		case 1:

			fmt.Println("Просмотр закладок")
			m, err := myCli.GetAllBookMarks()
			if err != nil {
				fmt.Println(err)

				break
			}

			if len(m) <= 0 {
				fmt.Println("Закладок нет")

				break
			}

			fmt.Println("-----------")
			for key, value := range m {
				fmt.Printf("Название закладки: %s | Содержимое закладки: %s\n", key, value)
			}

			fmt.Println("-----------")

		case 2:
			fmt.Println("Введите название закладки")
			fmt.Scan(&bookMark.Name)
			fmt.Println("Введите тест закладки")
			fmt.Scan(&bookMark.Text)
			err := myCli.AddBookMark(bookMark.Name, bookMark.Text)
			if err != nil {
				fmt.Println("Закладка с таким названием уже есть")

				break
			}

			fmt.Println("Закладка добавлена")

		case 3:
			fmt.Println("Введите название закладки")
			fmt.Scan(&bookMark.Name)
			err := myCli.RemoveBookMark(bookMark.Name)
			if err != nil {
				fmt.Println("Закладки с таким названием нет")

				break
			}

			fmt.Println("Закладка удалена")

		case 4:
			break menu

		}
	}

}

func (mc *MyCliConsole) getMenu() int {
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
