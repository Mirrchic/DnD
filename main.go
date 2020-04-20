package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// We going to return jsons data in this structs
type Monsters struct {
	Name       string `json:"Name"`
	Char       string `json:"Char`
	Size       string `json:"Size"`
	Speed      string `json:"Speed"`
	Skills     string `json:"Skills"`
	HitPoints  string `json:"HitPoints"`
	Rolls      string `json:"Rolls"`
	ArmorClass string `json:"ArmorClass"`
	Desc       string `json:"Desc"`
	Language   string `json:"Language"`
	Sense      string `json:"Sense"`
}

//htis struct taking monsters typr and slices of Monster struct
type Monster struct {
	ArType       string     `json:"ArType"`
	MonstersInfo []Monsters `json:"MonstersInfo"`
}

//this strutct is taking slice of monster struct
type List struct {
	ListOfMon []Monster `json:"ListOfMon"`
}

func main() {
	//Init of bot
	var err error
	bot, err := tgbotapi.NewBotAPI("1070281604:AAGBJn6VtB8dSqo_--9HKZHWBon2ra0kzFc")
	if err != nil {
		log.Panic("bot init error:", err)
		return
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	u.Limit = 10
	u.Offset = 0
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic("update error:", err)
	}

	MainResult := MainScan()
	HumanButtons := MonsterScan(0)
	GoblinidButtons := MonsterScan(1)
	BeastButtons := MonsterScan(2)
	DragonButtons := MonsterScan(3)
	UndeadButtons := MonsterScan(4)
	ConstructButtons := MonsterScan(5)

	var Menu1 = tgbotapi.NewReplyKeyboard(
		MainResult[:3],
		MainResult[3:6],
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Help"),
			tgbotapi.NewKeyboardButton("About"),
		),
	)
	//Init keuboards for different types of monsters
	var MenuHumanoid = tgbotapi.NewReplyKeyboard(
		HumanButtons...,
	)

	var MenuGoblinid = tgbotapi.NewReplyKeyboard(
		GoblinidButtons...,
	)

	var MenuBeasts = tgbotapi.NewReplyKeyboard(
		BeastButtons...,
	)
	var MenuDragons = tgbotapi.NewReplyKeyboard(
		DragonButtons...,
	)
	var MenuUndead = tgbotapi.NewReplyKeyboard(
		UndeadButtons...,
	)
	var MenuConstruct = tgbotapi.NewReplyKeyboard(
		ConstructButtons...,
	)

	//loop thats handle all updates thats bot taking
	for update := range updates {

		cmd := update.Message.Command()
		//Creats logic for Main menu buttons
		if cmd == "start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to main menu")
			msg.ReplyMarkup = Menu1
			bot.Send(msg)
		}

		if update.Message.Text == Menu1.Keyboard[0][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Humanoids are the main peoples of a fantasy gaming world, both civilized and savage, including humans and a tremendous variety of other species. They have language and culture, few if any innate magical abilities (though most humanoids can learn spellcasting), and a bipedal form. The most common humanoid races are the ones most suitable as player characters: humans, dwarves, elves, and halflings.")
			msg.ReplyMarkup = MenuHumanoid
			bot.Send(msg)
		}
		//goblinid button
		if update.Message.Text == Menu1.Keyboard[0][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "They are considered humanoid creatures, however, they are distinguished by increased aggressiveness, savagery and insidiousness. The goblinnidrv group includes orcs, goblins, ogres. They have a skin color between yellow, green and gray. Distinguished by an evil worldview, extremely insidious.")
			msg.ReplyMarkup = MenuGoblinid
			bot.Send(msg)
		}

		if update.Message.Text == "Main Menu" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to main menu")
			msg.ReplyMarkup = Menu1
			bot.Send(msg)
			log.Print("works")
		}
		//Beests button
		if update.Message.Text == Menu1.Keyboard[0][2].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Beasts are nonhumanoid creatures that are a natural part of the fantasy ecology. Some of them have magical powers, but most are unintelligent and lack any society or language. Beasts include all varieties of ordinary animals, dinosaurs, and giant versions of animals.")
			msg.ReplyMarkup = MenuBeasts
			bot.Send(msg)
		}
		//Dragon button
		if update.Message.Text == Menu1.Keyboard[1][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Dragons are large reptilian creatures of ancient origin and tremendous power. True dragons, including the good metallic dragons and the evil chromatic dragons, are highly intelligent and have innate magic. Also in this category are creatures distantly related to true dragons, but less powerful, less intelligent, and less magical, such as wyverns and pseudodragons.")
			msg.ReplyMarkup = MenuDragons
			bot.Send(msg)
		}
		//Plant button
		if update.Message.Text == Menu1.Keyboard[1][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Not ready yet")
			msg.ReplyMarkup = MenuUndead
			bot.Send(msg)
		}
		//Construct button
		if update.Message.Text == Menu1.Keyboard[1][2].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Not ready yet")
			msg.ReplyMarkup = MenuConstruct
			bot.Send(msg)
		}

		//About button
		if update.Message.Text == Menu1.Keyboard[2][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You can use buttons to navigate, or just write the monsters name")
			bot.Send(msg)
		}
		//Help button
		if update.Message.Text == Menu1.Keyboard[2][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "The bot is written to facilitate the search for monsters in the bestiary during the game.")
			bot.Send(msg)
		}

		if update.Message.Text != "" && update.Message.Text != "Main Menu" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, AnswerMenu(update.Message.Text))
			//chech this place
			log.Print("Message: " + update.Message.Text + "\n  Username: " + update.Message.Chat.UserName)
			bot.Send(msg)
		}

	}

}

//Function is returning buttons for main menu
func MainScan() (result []tgbotapi.KeyboardButton) {
	v := ListsWriter()
	for i := range v.ListOfMon {
		result = append(result, tgbotapi.NewKeyboardButton(v.ListOfMon[i].ArType))
	}
	return
}

//Function returning buttons for the monster menu
func MonsterScan(n int) (result [][]tgbotapi.KeyboardButton) {
	v := ListsWriter()
	var mainMenu []tgbotapi.KeyboardButton
	mainMenu = append(mainMenu, tgbotapi.NewKeyboardButton("Main Menu"))
	result = append(result, mainMenu)
	for i := range v.ListOfMon[n].MonstersInfo {
		var grownSlice []tgbotapi.KeyboardButton
		grownSlice = append(grownSlice, tgbotapi.NewKeyboardButton(v.ListOfMon[n].MonstersInfo[i].Name))
		result = append(result, grownSlice)
	}
	return
}

//function that returning data from JSON
func ListsWriter() (Lists List) {
	Lists = List{}
	readJson, err := ioutil.ReadFile("monsters.json")
	if err != nil {
		log.Print(err, "Problem with reading file")
	}
	_ = json.Unmarshal(readJson, &Lists)
	return Lists
}

//The function accepts a user request and looks for field in JSON to answer
func AnswerMenu(m string) (result string) {
	v := ListsWriter()
	for i := range v.ListOfMon {
		for r := range v.ListOfMon[i].MonstersInfo {
			if m == v.ListOfMon[i].MonstersInfo[r].Name {
				result = v.ListOfMon[i].MonstersInfo[r].Name + "\n" + v.ListOfMon[i].MonstersInfo[r].Desc + "\n\n Hitpoints: " + v.ListOfMon[i].MonstersInfo[r].HitPoints + "\n\n ArmorClass " + v.ListOfMon[i].MonstersInfo[r].ArmorClass + "\n\n Speed: " + v.ListOfMon[i].MonstersInfo[r].Speed + "\n\n Size: " + v.ListOfMon[i].MonstersInfo[r].Size + "\n\n Skills: " + v.ListOfMon[i].MonstersInfo[r].Skills + "\n\n Character: " + v.ListOfMon[i].MonstersInfo[r].Char

			}
		}
	}
	return
}
