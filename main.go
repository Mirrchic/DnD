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
	bot, err := tgbotapi.NewBotAPI("1047071352:AAG_C5IxQip5ZZmBrG3LjTKDQAsZTLyLXAk")
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
	//creating  buttons  body
	var Menu1 = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Humanid"),
			tgbotapi.NewKeyboardButton("Goblinid"),
			tgbotapi.NewKeyboardButton("Beasts"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Dragons"),
			tgbotapi.NewKeyboardButton("Plant"),
			tgbotapi.NewKeyboardButton("Construct"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("About bot"),
			tgbotapi.NewKeyboardButton("Next page (2)"),
		),
	)

	var Menu2 = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Dragons"),
			tgbotapi.NewKeyboardButton("Jin's"),
			tgbotapi.NewKeyboardButton("somethink"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Previos page (1)"),
			tgbotapi.NewKeyboardButton("Next page(1)"),
		),
	)

	var MenuHuman = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Elphs"),
			tgbotapi.NewKeyboardButton("Dwarfs"),
			tgbotapi.NewKeyboardButton("Gnom's"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Humans"),
			tgbotapi.NewKeyboardButton("Half-Orks"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Main menu"),
			tgbotapi.NewKeyboardButton("Move to Goblinid"),
		),
	)

	var MenuGoblinid = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Ogrs"),
			tgbotapi.NewKeyboardButton("Orks"),
			tgbotapi.NewKeyboardButton("Goblin's"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Main menu"),
			tgbotapi.NewKeyboardButton("Beasts"),
		),
	)

	var MenuBeast = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Fire elemental"),
			tgbotapi.NewKeyboardButton("Ice elemental"),
			tgbotapi.NewKeyboardButton("Eath elemental"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Main menu"),
			tgbotapi.NewKeyboardButton("Dragons"),
		),
	)

	var MenuPlants = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Dragons"),
			tgbotapi.NewKeyboardButton("Jin's"),
			tgbotapi.NewKeyboardButton("somethink"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Main menu"),
			tgbotapi.NewKeyboardButton("Next page(1)"),
		),
	)

	var MenuDragons = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Metall"),
			tgbotapi.NewKeyboardButton("Collor"),
			tgbotapi.NewKeyboardButton("Legend"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Main menu"),
		),
	)
	//readiing json file
	Lists := List{}
	readJson, _ := ioutil.ReadFile("monsters.json")
	//returning jsons data to Jists struct
	_ = json.Unmarshal(readJson, &Lists)
	//loop thats handle all updates thats bot taking
	for update := range updates {

		cmd := update.Message.Command()
		//Creats logic for Main menu buttons
		if cmd == "start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to main menu")
			msg.ReplyMarkup = Menu1
			bot.Send(msg)
		}
		//humanid button
		if update.Message.Text == Menu1.Keyboard[0][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Humanoids are the main peoples of a fantasy gaming world, both civilized and savage, including humans and a tremendous variety of other species. They have language and culture, few if any innate magical abilities (though most humanoids can learn spellcasting), and a bipedal form. The most common humanoid races are the ones most suitable as player characters: humans, dwarves, elves, and halflings.")
			msg.ReplyMarkup = MenuHuman
			bot.Send(msg)
		}
		//goblinid button
		if update.Message.Text == Menu1.Keyboard[0][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "They are considered humanoid creatures, however, they are distinguished by increased aggressiveness, savagery and insidiousness. The goblinnidrv group includes orcs, goblins, ogres. They have a skin color between yellow, green and gray. Distinguished by an evil worldview, extremely insidious.")
			msg.ReplyMarkup = MenuGoblinid
			bot.Send(msg)
		}
		//Beests button
		if update.Message.Text == Menu1.Keyboard[0][2].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Beasts are nonhumanoid creatures that are a natural part of the fantasy ecology. Some of them have magical powers, but most are unintelligent and lack any society or language. Beasts include all varieties of ordinary animals, dinosaurs, and giant versions of animals.")
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
			bot.Send(msg)
		}

		//About botton
		if update.Message.Text == Menu1.Keyboard[2][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "This is very first version of pocket DnD bestiary")
			bot.Send(msg)
		}
		//Next Page burron
		if update.Message.Text == Menu1.Keyboard[2][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Page 2:")
			msg.ReplyMarkup = Menu2
			bot.Send(msg)
		}
		//for second menu
		if update.Message.Text == Menu2.Keyboard[0][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			bot.Send(msg)
		}
		//
		if update.Message.Text == Menu2.Keyboard[0][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			bot.Send(msg)
		}
		//
		if update.Message.Text == Menu2.Keyboard[0][2].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			bot.Send(msg)
		}
		//
		if update.Message.Text == Menu2.Keyboard[1][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			bot.Send(msg)
		}

		if update.Message.Text == Menu2.Keyboard[1][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			bot.Send(msg)
		}

		//varibals for answer on buttons that moves us to another monsters types or main menu
		var mainAnswer = "we hope it helped you"
		var nextType = "Previous creatior type was:"
		//Login for Humanid Menu
		if update.Message.Text == MenuHuman.Keyboard[0][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			msg.Text = "Name: " + Lists.ListOfMon[0].MonstersInfo[0].Name + "\n\n Character:  " + Lists.ListOfMon[0].MonstersInfo[0].Char + "\n\n Size:  " + Lists.ListOfMon[0].MonstersInfo[0].Skills + "\n\n Speed: " + Lists.ListOfMon[0].MonstersInfo[0].Speed + "\n\n\n" + Lists.ListOfMon[0].MonstersInfo[0].Desc
			bot.Send(msg)
		}
		if update.Message.Text == MenuHuman.Keyboard[0][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			msg.Text = "Name: " + Lists.ListOfMon[0].MonstersInfo[1].Name + "\n\n Character:  " + Lists.ListOfMon[0].MonstersInfo[0].Char + "\n\n Size:  " + Lists.ListOfMon[0].MonstersInfo[0].Skills + "\n\n Speed: " + Lists.ListOfMon[0].MonstersInfo[1].Speed + "\n\n\n" + Lists.ListOfMon[0].MonstersInfo[1].Desc
			bot.Send(msg)
		}
		if update.Message.Text == MenuHuman.Keyboard[0][2].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			msg.Text = "Name: " + Lists.ListOfMon[0].MonstersInfo[2].Name + "\n\n Character:  " + Lists.ListOfMon[0].MonstersInfo[2].Char + "\n\n Size:  " + Lists.ListOfMon[0].MonstersInfo[2].Skills + "\n\n Speed: " + Lists.ListOfMon[0].MonstersInfo[2].Speed + "\n\n\n" + Lists.ListOfMon[0].MonstersInfo[2].Desc
			bot.Send(msg)
		}
		if update.Message.Text == MenuHuman.Keyboard[1][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			msg.Text = "Name: " + Lists.ListOfMon[0].MonstersInfo[3].Name + "\n\n Character:  " + Lists.ListOfMon[0].MonstersInfo[3].Char + "\n\n Size:  " + Lists.ListOfMon[0].MonstersInfo[3].Skills + "\n\n Speed: " + Lists.ListOfMon[0].MonstersInfo[3].Speed + "\n\n\n" + Lists.ListOfMon[0].MonstersInfo[3].Desc
			bot.Send(msg)
		}
		if update.Message.Text == MenuHuman.Keyboard[1][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			msg.Text = "Name: " + Lists.ListOfMon[0].MonstersInfo[4].Name + "\n\n Character:  " + Lists.ListOfMon[0].MonstersInfo[4].Char + "\n\n Size:  " + Lists.ListOfMon[0].MonstersInfo[4].Skills + "\n\n Speed: " + Lists.ListOfMon[0].MonstersInfo[4].Speed + "\n\n\n" + Lists.ListOfMon[0].MonstersInfo[4].Desc
			bot.Send(msg)
		}

		if update.Message.Text == MenuHuman.Keyboard[2][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, mainAnswer)
			msg.ReplyMarkup = Menu1
			bot.Send(msg)
		}
		if update.Message.Text == MenuHuman.Keyboard[2][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, nextType+"Humans")
			msg.ReplyMarkup = MenuGoblinid
			bot.Send(msg)
		}

		//Logic for Goblinid menu

		if update.Message.Text == MenuGoblinid.Keyboard[0][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			msg.Text = "Name: " + Lists.ListOfMon[1].MonstersInfo[0].Name + "\n\n Character:  " + Lists.ListOfMon[1].MonstersInfo[0].Char + "\n\n Size:  " + Lists.ListOfMon[1].MonstersInfo[0].Skills + "\n\n Speed: " + Lists.ListOfMon[1].MonstersInfo[0].Speed + "\n\n\n" + Lists.ListOfMon[1].MonstersInfo[0].Desc
			bot.Send(msg)
		}
		if update.Message.Text == MenuGoblinid.Keyboard[0][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			msg.Text = "Name: " + Lists.ListOfMon[1].MonstersInfo[1].Name + "\n\n Character:  " + Lists.ListOfMon[1].MonstersInfo[1].Char + "\n\n Size:  " + Lists.ListOfMon[1].MonstersInfo[1].Skills + "\n\n Speed: " + Lists.ListOfMon[1].MonstersInfo[1].Speed + "\n\n Hit Points: " + Lists.ListOfMon[1].MonstersInfo[1].HitPoints + "\n\n Armor class: " + Lists.ListOfMon[1].MonstersInfo[1].ArmorClass + "\n\n\n" + Lists.ListOfMon[1].MonstersInfo[1].Desc
			bot.Send(msg)
		}
		if update.Message.Text == MenuGoblinid.Keyboard[0][2].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			msg.Text = "Name: " + Lists.ListOfMon[1].MonstersInfo[2].Name + "\n\n Character:  " + Lists.ListOfMon[1].MonstersInfo[2].Char + "\n\n Size:  " + Lists.ListOfMon[1].MonstersInfo[2].Skills + "\n\n Speed: " + Lists.ListOfMon[1].MonstersInfo[2].Speed + "\n\n Hit Points: " + Lists.ListOfMon[1].MonstersInfo[2].HitPoints + "\n\n Armor class: " + Lists.ListOfMon[1].MonstersInfo[2].ArmorClass + "\n\n\n" + Lists.ListOfMon[1].MonstersInfo[2].Desc
			bot.Send(msg)
		}
		if update.Message.Text == MenuGoblinid.Keyboard[1][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, mainAnswer)
			msg.ReplyMarkup = Menu1
			bot.Send(msg)
		}
		if update.Message.Text == MenuGoblinid.Keyboard[1][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, nextType+"Govlinid")
			msg.ReplyMarkup = MenuBeast
			bot.Send(msg)
		}
		/* Todo: Have to finish menus that's left
		 */
		if update.Message.Text == MenuBeast.Keyboard[1][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "welcome back")
			msg.ReplyMarkup = Menu1
			bot.Send(msg)
		}
		if update.Message.Text == MenuBeast.Keyboard[1][1].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, nextType+"Beasts")
			msg.ReplyMarkup = MenuDragons
			bot.Send(msg)
		}

		//Dragons
		if update.Message.Text == MenuDragons.Keyboard[1][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, mainAnswer)
			msg.ReplyMarkup = Menu1
			bot.Send(msg)
		}
		//plant
		//construct
		if update.Message.Text == MenuPlants.Keyboard[0][0].Text {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Yes it is just buttons")
			bot.Send(msg)
		}

	}

}
