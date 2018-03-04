package main

import (
	"log"
	"net/http"
	"io"
)

func Hello(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Oh, hai there!")
}

func Start(res http.ResponseWriter, req *http.Request) {
	log.Print("START REQUEST")

	//data, err := NewStartRequest(req)
	//if err != nil {
	//	log.Printf("Bad start request: %v", err)
	//}
	//dump(data)

	respond(res, StartResponse{
		Taunt:          "Tssssssssss!",
		Color:          "#75CEDD",
		HeadURL:        "https://forums.androidcentral.com/images/forum_icons/200.png",
		HeadType:       HEAD_TONGUE,
		TailType:       TAIL_FRECKLED,
		SecondaryColor: "#F7D3A2",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	log.Printf("MOVE REQUEST")

	data, err := NewMoveRequest(req)

	if err != nil {
		log.Printf("Bad move request: %v", err)
	}
	//dump(data)

	respond(res, MoveResponse{
		Move: getMove(data),
	})
}

func End(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("200 - Game Over :("))
}