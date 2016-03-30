package actiongenerator

import (
	"fmt"
	"math/rand"
	"net/http"
)

var races = []string{"Dragonborn", "Dwarf", "Elf", "Gnome", "Halfling", "Human", "Tiefling"}
var classes = []string{"Bard", "Cleric", "Druid", "Fighter", "Paladin", "Rogue", "Sorceror", "Warlock", "Wizard"}
var skills = []string{"Acrobatics", "Athletics", "Arcana", "Heal", "Nature", "Religion", "Stealth"}

func Randomize(w http.ResponseWriter, r *http.Request) {
	race := races[rand.Intn(len(races))]
	class := classes[rand.Intn(len(classes))]
	skill := skills[rand.Intn(len(skills))]
	fmt.Fprintf(w, "A %s %s uses %s", race, class, skill)
}
