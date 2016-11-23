package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"ProfileInfo",
		"GET",
		"/profile/{profileId}",
		ProfileInfo,
	},
	Route{
		"ProfileAvatar",
		"GET",
		"/profile/{profileId}/avatar",
		ProfileAvatar,
	},
	Route{
		"ProfileBoobs",
		"GET",
		"/profile/{profileId}/hasboobs",
		ProfileBoobs,
	},
	Route{
		"UpdateProfile",
		"POST",
		"/profile/{profileId}",
		UpdateProfile,
	},
	Route{
		"CreateProfile",
		"POST",
		"/profile",
		CreateProfile,
	},
}
