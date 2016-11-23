package main

type Profile struct {
	Id        string `json:"id"`
	Name      string `json: "name"`
	Surname   string `json: "surname"`
	AvatarUri string `json: "avatarUri"`
	Age       int    `json: "age"`
}

type Profiles []Profile
