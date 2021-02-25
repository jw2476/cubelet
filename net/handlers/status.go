package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/jw2476/cubelet/client"
	"github.com/jw2476/cubelet/net/prot"
)

type Response struct {
	Version struct{
		Name string `json:"name"`
		Protocol int `json:"protocol"`
	} `json:"version"`
	Players struct{
		Max int `json:"max"`
		Online int `json:"online"`
	} `json:"players"`
	Description struct{
		Text string `json:"text"`
	} `json:"description"`
}

func handleRequest(request prot.Request) {
	response := Response{
		Version: struct{
			Name string `json:"name"`
			Protocol int `json:"protocol"`
		}{
			Name: "1.16.5",
			Protocol: 754,
		},
		Players: struct {
			Max int `json:"max"`
			Online int `json:"online"`
		}{
			Max: 100,
			Online: 5,
		},
		Description: struct {
			Text string `json:"text"`
		}{
			Text: "Hello Cubelet",
		},
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}

	packet := client.NewPacket(0)
	fmt.Println(response)
	packet.WriteString(string(responseJSON))
	err = request.Client.Send(packet)
	if err != nil {
		fmt.Println(err)
		return
	}
}

