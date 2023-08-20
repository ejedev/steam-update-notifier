package steamcmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rs/zerolog/log"
)

func GetChangeNumber(id int) int {
	resp, err := http.Get(fmt.Sprintf("%s%d", url, id))
	if err != nil {
		log.Error().
			Msg(fmt.Sprintf("There was an error fetching data for App %d. Returning a default value of 0.", id))
		return 0
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result apiResponse
	var gameResult gameResponse
	if err := json.Unmarshal(body, &result); err != nil {
		log.Error().
			Msg(fmt.Sprintf("Error unmarshalling data for app %d. Returning a default value of 0.", id))
		return 0
	}
	if err := json.Unmarshal(result.Data[fmt.Sprintf("%d", id)], &gameResult); err != nil {
		log.Error().
			Msg(fmt.Sprintf("Error unmarshalling data for app %d. Returning a default value of 0.", id))
		return 0
	}
	return gameResult.ChangeNumber
}
