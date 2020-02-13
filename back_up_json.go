package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//structur getting numbers,  bases, algebratic char,
//and base that result should be converted to

type Monsters struct {
	Name, Char, Size, Speed, Skills, HitPoints, Rolls, ArmorClass, Desc, Language, Sense string
}

type Monster struct {
	ArType       string
	MonstersInfo []Monsters
}

type List struct {
	ListOfMon []Monster
}

//simple Json write method
func main() {
	data := List{
		ListOfMon: []Monster{
			Monster{
				ArType: "Humanoid",
				MonstersInfo: []Monsters{
					Monsters{
						Name:       "Elves",
						Char:       "Any",
						Size:       "Medium",
						Speed:      "30 ft.",
						Skills:     "Darkvision 60 ft., Keen Senses, Trance, Fey Ancestry",
						HitPoints:  "Individual",
						ArmorClass: "Individual",
						Desc:       "Elves, much like their living ancestors, the eladrin, who were also commonly called elves, were a long-lived race of the Tel-quessir found most commonly in forests, shrublands, and other wildernesses. From time to time the elves organized strong nations, though with far less frequency than eladrin, in some cases adopting even a nomadic lifestyle. Almost all elves worshiped the gods of the Seldarine, and the elves were generally, though not always, good in nature",
					},
					Monsters{
						Name:       "Dwarfs",
						Char:       "Good",
						Size:       "4-5 ft.",
						Speed:      "25 ft.",
						Skills:     "Darkvision 60 ft., Dwarven Resilience, Dwarven Combat Training, Stonecunning",
						HitPoints:  "Individual",
						ArmorClass: "Individual",
						Desc:       "Dwarves are a short and stocky race, and stand about a foot shorter than most humans, with wide, compact bodies that account for their burly appearance. Male and female dwarves pride themselves on the long length of their hair, and men often decorate their beards with a variety of clasps and intricate braids. Clean-shavenness on a male dwarf is a sure sign of madness, or worse—no one familiar with their race trusts a beardless dwarven man.",
					},
					Monsters{
						Name:       "Gnomes",
						Char:       "Neutral Good",
						Size:       "3-4 ft.",
						Speed:      "20 ft..",
						Skills:     "Darkvision, Gnome Cunning",
						HitPoints:  "Individuall",
						ArmorClass: "Individuall",
						Desc:       "Skinny and flaxen-haired, his skin walnut brown and his eyes a startling turquoise, Burgell stood half as tall as Aeron and had to climb up on a stool to look out the peephole. Like most habitations in Oeble, that particular tenement had been built for humans, and smaller residents coped with the resulting awkwardness as best they could.",
					},
					Monsters{
						Name:       "Humans",
						Char:       "Any",
						Size:       "5-6 ft.",
						Speed:      "30 ft.",
						Skills:     "Individual",
						HitPoints:  "Individual",
						ArmorClass: "Individual",
						Desc:       "The physical characteristics of humans are as varied as the world’s climes. From the dark-skinned tribesmen of the southern continents to the pale and barbaric raiders of the northern lands, humans possess a wide variety of skin colors, body types, and facial features. Generally speaking, humans’ skin color assumes a darker hue the closer to the equator they live. At the same time, bone structure, hair color and texture, eye color, and a host of facial and bodily phenotypic characteristics vary immensely from one locale to another. Cheekbones may be high or broad, noses aquiline or flat, and lips full or thin; eyes range wildly in hue, some deep set in their sockets, and others with full epicanthic folds. Appearance is hardly random, of course, and familial, tribal, or national commonalities often allow the knowledgeable to identify a human’s place of origin on sight, or at least to hazard a good guess. Humans’ origins are also indicated through their traditional styles of bodily decoration, not only in the clothing or jewelry worn, but also in elaborate hairstyles, piercing, tattooing, and even scarification.",
					},
					Monsters{
						Name:       "Half-Orks",
						Char:       "Any",
						Size:       "5-6 ft.",
						Speed:      "30 ft.",
						Skills:     "Darkvision, Menacing, Relentless Endurance, Savage Attacks",
						HitPoints:  "Individual",
						ArmorClass: "Individual",
						Desc:       "Whether united under the leadership of a mighty warlock or having fought to a standstill after years of conflict, orc and human tribes sometimes form alliances, joining forces into a larger horde to the terror of civilized lands nearby. When these alliances are sealed by marriages, half-orcs are born. Some half-orcs rise to become proud chiefs of orc tribes, their human blood giving them an edge over their full-blooded orc rivals. Some venture into the world to prove their worth among humans and other more civilized races. Many of these become adventurers, achieving greatness for their mighty deeds and notoriety for their barbaric customs and savage fury.",
					},
				},
			},
			Monster{
				ArType: "Goblinids",
				MonstersInfo: []Monsters{
					Monsters{
						Name:       "Orcs",
						Char:       "Chaotic evil",
						Size:       "5-6 ft.",
						Speed:      "30 ft.",
						Skills:     "Darkvision 60 ft., Passive Perception 10 ",
						HitPoints:  "Individual",
						ArmorClass: "Individual",
						Desc:       "Orcs were a race of humanoids that had been a threat to the civilized cultures of Toril, particularly Faerûn for as long as any could remember. This changed somewhat in the years preceding and immediately after the Spellplague, when a horde of mountain orcs under the command of King Obould Many-Arrows unified into a single kingdom, one that was remarkably civilized.",
					},
					Monsters{
						Name:       "Goblins",
						Char:       "Chaotic evil",
						Size:       "Medium",
						Speed:      "30 ft.",
						Skills:     "Darkvision 60 ft., Passive Perception 8 ",
						HitPoints:  "7 (2d6)",
						ArmorClass: "15 leather armor",
						Desc:       "Goblins were a race of small and numerous goblinoids common throughout Toril, often living in underground caverns near the surface known as lairs. The race was often, though not always, dominated by other goblinoids, most commonly hobgoblins. Goblins may have, in fact, been initially created by this related race to serve as scouts and infiltrators",
					},
					Monsters{
						Name:       "Ogrs",
						Char:       "Any",
						Size:       "Giant",
						Speed:      "40 ft.",
						Skills:     "Darkvision 60 ft.,",
						HitPoints:  "59 (7d10 + 21)",
						ArmorClass: "11 Hide armor",
						Desc:       "Ogres appeared as giant humanoids with very muscular bodies and large heads. They stood between 9 and 10 feet tall and could weigh 600 to 650 lbs",
					},
				},
			},
			Monster{
				ArType:       "Beasts",
				MonstersInfo: []Monsters{},
			},
			Monster{
				ArType:       "Dragons",
				MonstersInfo: []Monsters{},
			},
			Monster{
				ArType:       "Plant",
				MonstersInfo: []Monsters{},
			},
			Monster{
				ArType:       "Construct",
				MonstersInfo: []Monsters{},
			},
		},
	}
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("monsters.json", file, 0644)
	fmt.Println(string(file))
}
