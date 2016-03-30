package actiongenerator

import (
	"fmt"
	"math/rand"
	"net/http"
)

type Skill struct {
	Name    string
	Actions []string
}

type Race struct {
	Name        string
	Descriptors []string
	Article     string
}

var acrobatics = Skill{Name: "Acrobatics", Actions: []string{"leap over a pit", "walk across a rope", "scale a cliff"}}
var animalHandling = Skill{Name: "Animal Handling", Actions: []string{"calm a horse", "train a dog", "coax a cat", "intimidate a wildcat"}}
var athletics = Skill{Name: "Athletics", Actions: []string{"leap over a pit", "climb a rope", "scale a cliff", "open a door", "cling to a troll's back"}}
var arcana = Skill{Name: "Arcana", Actions: []string{"read a scroll", "examine a spellbook", "recall a spell's details"}}
var deception = Skill{Name: "Deception", Actions: []string{"fool a shopkeeper", "befuddle a guard", "hoodwink a noble", "pretend not to notice a tail"}}
var history = Skill{Name: "History", Actions: []string{"recall a battle", "recognize an ancient crest", "date a ruin"}}
var investigation = Skill{Name: "Investigation", Actions: []string{"search a closet", "deduce a trap", "recreate a battle"}}
var nature = Skill{Name: "Nature", Actions: []string{"spot a sinkhole", "identify a creature's weakness", "predict a storm"}}
var medicine = Skill{Name: "Medicine", Actions: []string{"diagnose a poison", "treat a shallow wound", "advise a villager"}}
var perception = Skill{Name: "Perception", Actions: []string{"spot an ambush", "notice a pursuer", "hear someone in hiding"}}
var performance = Skill{Name: "Performance", Actions: []string{"play the mandolin", "tell jokes", "impress a child with magic", "recite an epic poem"}}
var persuasion = Skill{Name: "Persuasion", Actions: []string{"wheedle a barkeeper", "talk down a scared soldier", "reconcile a feud"}}

//TODO: "Survival", "Religion", "Sleight of hand", "Stealth"}
var races = []string{"Dragonborn", "Dwarf", "Elf", "Gnome", "Halfling", "Human", "Tiefling"}
var classes = []string{"Bard", "Cleric", "Druid", "Fighter", "Paladin", "Rogue", "Sorceror", "Warlock", "Wizard"}
var skills = []Skill{animalHandling, athletics, arcana, deception, history, investigation, nature, medicine}

func Randomize(w http.ResponseWriter, r *http.Request) {
	race := races[rand.Intn(len(races))]
	class := classes[rand.Intn(len(classes))]
	skill := skills[rand.Intn(len(skills))]
	action := skill.Actions[rand.Intn(len(skill.Actions))]
	fmt.Fprintf(w, "A %s %s uses %s to %s", race, class, skill.Name, action)
}

func RandomizeSimple(w http.ResponseWriter, r *http.Request) {
	race := races[rand.Intn(len(races))]
	class := classes[rand.Intn(len(classes))]
	skill := skills[rand.Intn(len(skills))]
	fmt.Fprintf(w, "A %s %s uses %s", race, class, skill.Name)
}
