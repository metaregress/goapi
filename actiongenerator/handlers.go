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
var athletics = Skill{Name: "Atheltics", Actions: []string{"leap over a pit", "climb a rope", "scale a cliff", "open a door", "cling to a troll's back"}}
var arcana = Skill{Name: "Arcana", Actions: []string{"read a scroll", "examine a spellbook", "recall a spell's details"}}
var deception = Skill{Name: "Deception", Actions: []string{"fool a shopkeeper", "wheedle a barkeeper", "befuddle a guard", "hoodwink a noble"}}
var history = Skill{Name: "History", Actions: []string{"recall a battle", "recognize an ancient crest", "date a ruin"}}
var investigation = Skill{Name: "Investigation", Actions: []string{"search a closet", "deduce a trap", "recreate a battle"}}
var nature = Skill{Name: "Nature", Actions: []string{"spot a sinkhole", "identify a creature's weakness", "predict a storm"}}
var medicine = Skill{Name: "Medicine", Actions: []string{"diagnose a poison", "treat a shallow wound", "advise a villager"}}

var races = []string{"Dragonborn", "Dwarf", "Elf", "Gnome", "Halfling", "Human", "Tiefling"}
var classes = []string{"Bard", "Cleric", "Druid", "Fighter", "Paladin", "Rogue", "Sorceror", "Warlock", "Wizard"}
var skills = []Skill{animalHandling, athletics, arcana, deception, history, investigation, nature, medicine} // "Perception", "Performance", "Persuasion", "Survival", "Religion", "Sleight of hand", "Stealth"}

func Randomize(w http.ResponseWriter, r *http.Request) {
	race := races[rand.Intn(len(races))]
	class := classes[rand.Intn(len(classes))]
	skill := skills[rand.Intn(len(skills))]
	action := skill.Actions[rand.Intn(len(skill.Actions))]
	fmt.Fprintf(w, "A %s %s uses %s to %s", race, class, skill.Name, action)
}
