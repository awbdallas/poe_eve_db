/*
 * GO POE DB
 * Created in order to make requests.
 *
 *
 */

package main

import (
	"net/http"
	"time"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

const POE_API_ENDPOINT = `http://api.pathofexile.com/public-stash-tabs`
const POE_API_PUBLIC_STASH_TABS = `public-stash-tabs`
const TIME_BETWEEN_REQUESTS = 5;
const THREADS = 4;

type ItemCategories struct {
	Gems[] string				`json:"gems"`
	Weapons[] string			`json:"weapons"`
	Jewels[] string				`json:"jewels"`
	Accessories[] string		`json:"accessories"`
	Flasks[] string				`json:"flasks"`
	Armour[] string				`json:"armour"`
}

type ItemProperty struct {
	Name string			`json:"name"`
	Values[]interface{}  	`json:"values"`
	DisplayMode int 	`json:"displayMode"`
	Type int			`json:"type"`
}

type ItemSocket struct {
	Group int 				`json:"group"`
	Attr  string 			`json:"attr"`
	SocketColour string 		`json:"sColour"`
}

type StashItems struct {
	Verified bool				`json:"verified,omitempty"`
	Width int					`json:"w,omitempty"`
	Height int					`json:"h,omitempty"`
	Ilvl int					`json:"ilvl,omitempty"`
	Icon string					`json:"icon,omitempty"`
	League string				`json:"league,omitempty"`
	Support bool				`json:"support,omitempty"`
	Id string					`json:"id,omitempty"`
	Sockets[] ItemSocket 		`json:"sockets,omitempty"`
	Name string					`json:"name,omitempty"`
	TypeLine string				`json:"typeLine,omitempty"`
	Identified bool 			`json:"typeLine,omitempty"`
	Note string					`json:"note,omitempty"`
	Properties[] ItemProperty	`json:"properties,omitempty"`
	AdditionalProperties[] ItemProperty	`json:"additionalProperties,omitempty"`
	Requirements[] ItemProperty	`json:"requirements,omitempty"`
	SecDescriptionText string					`json:"secDescrText,omitempty"`
	ExplicitMods[] string		`json:"explicitMods,omitempty"`
	ItemCategory ItemCategories `json:"category,omitempty"`
	FlavourText[] string		`json:"flavourText,omitempty"`
	FrameType int				`json:"frameType,omitempty"`
	Xcord int					`json:"x,omitempty"`
	YCord int					`json:"y,omitempty"`
	InventoryId string			`json:"inventoryId,omitempty"`
}

type Stash struct {
	Id string				`json:"id,omitempty"`
	Public bool				`json:"public,omitempty"`
	AccountName string		`json:"accountName,omitempty"`
	LastCharacterName  string	`json:"lastCharacterName,omitempty"`
	Stash string			`json:"stash,omitempty"`
	StashType string		`json:"stashType,omitempty"`
	Items[] StashItems		`json:"items,omitempty"`
}

type PublicStashesRequest struct {
	NextChangeId string     `json:"next_change_id"`
	Stashes[] Stash 		`json:"stashes"`
}


func main() {
	var publicStashRequest PublicStashesRequest

	resp := reliableGet(POE_API_ENDPOINT, 5)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		err = json.Unmarshal(body, &publicStashRequest)
		fmt.Println(publicStashRequest.Stashes[20])
	}

}

func reliableGet(url string, tries int) *http.Response {
	timeout := time.Duration(30 * time.Second)

	client := http.Client{
		Timeout: timeout,
	}

	for i := 0; i < tries; i++ {
		resp, err := client.Get(url)

		if err != nil || resp.StatusCode != 200 {
			continue
		} else {
			return resp
		}
	}

	return nil
}
