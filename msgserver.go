package msgserver

//Hub of hubs lives here
var hubs *hubMap = &hubMap{m: map[string]*hub{}}

//Package-global config lives here:
var cfg *config
