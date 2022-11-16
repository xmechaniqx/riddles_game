package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {

	var firstOrCont string
	fmt.Println("*ВЕСЕЛЫЕ ЗАГАДКИ*(нет)")
	fmt.Println("Ну привет, искатель")
	fmt.Println("В первый раз играем?(введи да/нет)")
	fmt.Scanln(&firstOrCont)
	if firstOrCont == "да" || firstOrCont == "Да" || firstOrCont == "ДА" {
		begin()
	}
	if firstOrCont == "нет" || firstOrCont == "Нет" {
		fmt.Println("Напиши цифрой на какой загадке ты остановилась")
		contiNue()
	}
}
func contiNue() {
	var level int
again:
	fmt.Scanln(&level)
	switch {
	case level == 1:
		firstQ()
	case level == 2:
		secondQ()
	case level == 3:
		thirdQ()
	case level == 4:
		fourthQ()
	case level == 5:
		fivethQ()
	default:
		fmt.Println("Что-то тут не так, попробуй снова")
		goto again
	}
}
func begin() {
	var stroka string
repeat:
	fmt.Println("Сыграем в игру? (введи да/нет)")
	fmt.Scanln(&stroka)

	if stroka == "нет" {
		for stroka == "нет" || stroka == "Нет" {
			fmt.Println("Ты не понял, у тебя нет выбора")
			fmt.Println("Сыграем в игру? (введи да/нет)")
			fmt.Scanln(&stroka)
		}
	}
	if stroka == "да" || stroka == "Да" || stroka == "ДА" {
		gameStart()
	} else {
		fmt.Println("Ну внимательнее, попробуй еще раз")
		goto repeat
	}
}
func gameStart() {
	var stroka1 string
	fmt.Println("Отлично")
	time.Sleep(2 * time.Second)
	fmt.Println("Задача очень простая, я загадываю тебе загадки")
	time.Sleep(4 * time.Second)
	fmt.Println("Если ты угадываешь то получаешь ключ подсказку")
	time.Sleep(4 * time.Second)
	fmt.Println("Собрав все подсказки и разобравшись что с ними делать")
	time.Sleep(4 * time.Second)
	fmt.Println("Ты увидишь забавную картинку")
	time.Sleep(4 * time.Second)
	fmt.Println("И да, я очень тебя прошу не искать ответы в интернете")
	time.Sleep(4 * time.Second)
	fmt.Println("Начнем?)     (введи да/нет)")
again:
	fmt.Scanln(&stroka1)
	if stroka1 == "нет" {
		for stroka1 == "нет" || stroka1 == "Нет" {
			fmt.Println("Не принемается")
			time.Sleep(1 * time.Second)
			fmt.Println("Пиши 'да' и погнали)")
			fmt.Scanln(&stroka1)
		}
	}
	if stroka1 == "да" || stroka1 == "Да" || stroka1 == "ДА" {
		firstQ()
	} else {
		fmt.Println("Звязывай, пиши 'да'")
		goto again
	}
}
func firstQ() {
	var answer1 string
	fmt.Println(`"Первая загадка:
	Стоит дуб,
	В нем двенадцать гнезд,
	В каждом гнезде
	По четыре яйца,
	В каждом яйце
	По семи цыпленков."
	Что это такое?`)
wrong1:
	fmt.Scanln(&answer1)
	if answer1 == "Год" || answer1 == "год" {
		fmt.Println("Ну это было не трудно")
		time.Sleep(2 * time.Second)
		fmt.Println(`Держи ключ - "A"`)
		time.Sleep(2 * time.Second)
		fmt.Println("Так а теперь посложнее")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println(no((rand.Intn(7))))
		goto wrong1
	}
	secondQ()
}
func secondQ() {
	var answer2 string
	fmt.Println(`Вторая загадка сейчас актуальная и звучит так:
	Если день нахмурится,
	Если дождь пойдет —
	Выйдет он на улицу,
	Надо мной вспорхнет.`)
wrong2:
	fmt.Scanln(&answer2)
	if answer2 == "Зонт" || answer2 == "зонт" || answer2 == "Зонтик" || answer2 == "зонтик" {
		fmt.Println("А ты не промах")
		time.Sleep(2 * time.Second)
		fmt.Println(`Держи ключ - "L"`)
		time.Sleep(2 * time.Second)
		fmt.Println("Так а теперь")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println(no((rand.Intn(7))))
		goto wrong2
	}
	thirdQ()
}
func thirdQ() {
	var answer3 string
	fmt.Println(`Третья:
	На листочке,
	На страничке —
	То ли точки,
	То ли птички.
	Все сидят на лесенке,
	Все щебечут песенки.`)
wrong3:
	fmt.Scanln(&answer3)
	if answer3 == "Ноты" || answer3 == "ноты" {
		fmt.Println("Хорошо, молодец")
		time.Sleep(2 * time.Second)
		fmt.Println(`Держи ключ - "T"`)
		time.Sleep(2 * time.Second)
		fmt.Println("Тогда следующая")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println(no((rand.Intn(7))))
		goto wrong3
	}
	fourthQ()
}
func fourthQ() {
	var answer4 string
	fmt.Println(`Четвертая загадка:
	Мудрец в нем видел мудреца,
	Глупец — глупца,
	Баран — барана,
	Овцу в нем видела овца,
	И обезьяну — обезьяна,
	Но вот подвели к нему Федю Баратова,
	И Федя неряху увидел лохматого.`)
wrong4:
	fmt.Scanln(&answer4)
	if answer4 == "Зеркало" || answer4 == "зеркало" {
		fmt.Println("Да уж, эрудит...")
		time.Sleep(2 * time.Second)
		fmt.Println(`Держи ключ - "+"`)
		time.Sleep(2 * time.Second)
		fmt.Println("Ты почти на финише")
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println(no((rand.Intn(7))))
		goto wrong4
	}
	fivethQ()
}
func fivethQ() {
	var answer5 string
	fmt.Println(`И наконец последняя загадка, 
думаю она дастся тебе с трудом, звучит она так:
	Орехов не ест,
	Сахара не просит,
	А щипцы с собой носит.`)
wrong5:
	fmt.Scanln(&answer5)
	if answer5 == "Рак" || answer5 == "рак" {
		fmt.Println("Что же, твоя взяла")
		time.Sleep(2 * time.Second)
		fmt.Println(`Держи ключ - "1"`)
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println(no((rand.Intn(7))))
		goto wrong5
	}
	end()
}
func end() {
	fmt.Println(`Надеюсь ты записал все ключи, 
	осталось понять как их связать...и причем тут блокнот?...`)
	runNotepad()
}
func no(val int) string {
	var ret string
	switch {
	case val == 1:
		ret = "Ну неееее"
	case val == 2:
		ret = "Чушь какая-то,нет"
	case val == 3:
		ret = "Нет, не правильно"
	case val == 4:
		ret = "Совершенно........не верно"
	case val == 5:
		ret = "Подумай хорошенько"
	case val == 6:
		ret = "Такое себе, нет, ерунда"
	case val == 7:
		ret = "Неа, внимательнее!"
	}
	return ret
}
func runNotepad() {
	time.Sleep(5 * time.Second)

	c := exec.Command("notepad")
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}
