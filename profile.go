package main

type Profile struct {
	Id        string `json:"id,omitempty"`
	Name      string `json: "name,omitempty"`
	Surname   string `json: "surname,omitempty"`
	AvatarUri string `json: "avatarUri,omitempty"`
	Age       int    `json: "age,omitempty"`
}

type Boobs struct {
	HasBoobs *bool `json:"hasboobs,omitempty"`
}

type Profiles []Profile
