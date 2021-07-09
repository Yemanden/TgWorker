package main

import (
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
)

func main() {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("1862189298:AAE30MUSMtsPWFqMuRb6NQU1_rjXB-iRJ_g")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// инициализируем канал, куда будут прилетать обновления от API
	var ucfg tgbotapi.UpdateConfig = tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60

	//err = bot.UpdatesChan(ucfg)
	updChan, err := bot.GetUpdatesChan(ucfg)
	// читаем обновления из канала
	for {
		select {
		case update := <-updChan:
			// Пользователь, который написал боту
			UserName := update.Message.From.UserName

			// ID чата/диалога.
			// Может быть идентификатором как чата с пользователем
			// (тогда он равен UserID) так и публичного чата/канала
			ChatID := update.Message.Chat.ID

			// Текст сообщения
			Text := update.Message.Text

			log.Printf("[%s] %d %s", UserName, ChatID, Text)

			// Ответим пользователю его же сообщением
			reply := Text
			// Созадаем сообщение
			msg := tgbotapi.NewMessage(ChatID, reply)
			// и отправляем его
			bot.Send(msg)
		}

	}
	//	fmt.Println(utf8.RuneCountInString(`Анатомическая подушка для кормления Nutribaby разработана для кормящих женщин и грудных детей. Подушка в меру упругая, точно повторяет изгибы тела и создает персональный комфорт. Незаменима для кормления и развивающих игр с малышами. Создает безопасную среду для ребенка.
	//Nutribaby отличается хорошей воздухопроницаемостью, что исключает появление и развитие вредных микроорганизмов.
	//Серые звезды на белом фоне добавляют подушке выразительности. Ручка делает Nutribaby удобной в использовании.
	//Важно знать: чехол съемный, на молнии.
	//Достоинства Nutribaby:
	//подходит людям с чувствительной кожей, склонным к аллергии, детям.
	//Важно знать: чехол съемный, на молнии.
	//Достоинства Nutribaby:
	//подходит людям с чувствительной кожей, склонным к аллергии, детям.Важно знать: чехол съемный, на молнии.
	//Достоинства Nutribaby:
	//подходит людям с чувствительной кожей, склонным к аллергии, детямс.ммвываывафывцвцвцвцвцвцвцвцвцвцвцвццвцвцвцвцвцвцвццццццццццццццццццццццццццццццццццццццццццццццццццццццццццццццу4545454`))
}
