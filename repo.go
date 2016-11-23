package main

import (
	"fmt"
	"strconv"
)

var id int

var profiles Profiles

func init() {
	SaveProfile(Profile{
		Name:      "Ivan",
		Surname:   "BoobsLover",
		AvatarUri: "https://cdn0.iconfinder.com/data/icons/iconshock_guys/512/andrew.png",
		Age:       28})
	SaveProfile(Profile{
		Name:      "Anna",
		Surname:   "ZeroBoobs",
		AvatarUri: "https://cdn.meme.am/cache/instances/folder422/500x/57658422.jpg",
		Age:       17})
	SaveProfile(Profile{
		Name:      "TJ",
		Surname:   "SorryGuy",
		AvatarUri: "http://www.sklep.selected.pl/images/obrazki_przedmiotow/MEMY/MEM%2015.png",
		Age:       7})

}

func SaveProfile(p Profile) Profile {
	id += 1
	p.Id = strconv.Itoa(id)
	profiles = append(profiles, p)
	return p
}

func UpdateRepoProfile(p Profile) Profile {
	cur := GetProfile(p.Id)
	cur.Name = p.Name
	cur.Surname = p.Surname
	cur.AvatarUri = p.AvatarUri
	cur.Age = p.Age

	return cur
}

func RemoveProfile(id string) error {
	for i, p := range profiles {
		if p.Id == id {
			profiles = append(profiles[:i], profiles[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("No such profile")
}

func GetProfile(id string) Profile {

	for _, p := range profiles {
		if p.Id == id {
			return p
		}
	}
	return Profile{}
}
