package steamcmd

import "encoding/json"

type apiResponse struct {
	Data   map[string]json.RawMessage `json:"data"`
	Status string                     `json:"status"`
}

type gameResponse struct {
	ChangeNumber int    `json:"_change_number"`
	MissingToken bool   `json:"_missing_token"`
	Sha          string `json:"_sha"`
	Size         int    `json:"_size"`
	Appid        string `json:"appid"`
	Common       struct {
		Gameid string `json:"gameid"`
		Name   string `json:"name"`
		Oslist string `json:"oslist"`
		Type   string `json:"type"`
	} `json:"common"`
	Extended struct {
		Developer                 string `json:"developer"`
		Gamedir                   string `json:"gamedir"`
		Homepage                  string `json:"homepage"`
		Icon                      string `json:"icon"`
		Noservers                 string `json:"noservers"`
		Primarycache              string `json:"primarycache"`
		Sourcegame                string `json:"sourcegame"`
		State                     string `json:"state"`
		Visibleonlywheninstalled  string `json:"visibleonlywheninstalled"`
		Visibleonlywhensubscribed string `json:"visibleonlywhensubscribed"`
	} `json:"extended"`
}
