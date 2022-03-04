package ascii

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type ASCII struct {
	Args         string
	SplittedArgs []string
	Fileinfo     []string
	Reverse      bool
}

func Ascii(art ASCII) error {
	if !isValidSym(art.Args) {
		return errors.New("Unregistered symbols")
	}
	argStr := strings.Split(strings.ReplaceAll(art.Args, "\n", "\\n"), "\\n") // разделил аргумент на строки
	file, err := os.Open("fonts/standard.txt")                                //открываем файл
	if err != nil {
		return err
	}
	defer file.Close()                // закрываем
	scanner := bufio.NewScanner(file) //создаем переменную сканнер, сканируем файл
	for scanner.Scan() {
		art.Fileinfo = append(art.Fileinfo, scanner.Text()) // построчно записываем в массив строк . Каждая строка, как отдельный элемент массива
	}
	ConsoleOutput(argStr, art.Fileinfo)
	return nil
}

func Reverse(art ASCII) error {

	return nil
}

//ConsoleOutput - Вывод символов на консоль
func ConsoleOutput(argStr, fileinfo []string) {
	for _, s := range argStr { // цикл в аргументе, s - это строка в массиве
		if s != "" { // если строка не пустая, как раз на случай если передавать \n\n
			for i := 0; i < 8; i++ { // цикл на 8 строк
				for j := 0; j < len(s); j++ {
					fmt.Print(fileinfo[int(s[j]-32)*9+1+i])
				}
				fmt.Println()
			}
		} else { //если передали \n\n
			fmt.Println()
		}
	}
}

//isValidSym - проверяем, аргументы на невалидные символы
func isValidSym(s string) bool {
	for _, s := range s { // Цикл на случай, если в агрументах будут символы, которых нет в нашем стандарте
		if s < 32 || s > 126 {
			return false
		}
	}
	return true
}
