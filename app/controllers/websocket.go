package controllers

import (
	"code.google.com/p/go.net/websocket"
	"github.com/robfig/revel"
	"github.com/mati1979/go-revel-mobile-cars-adstream/app/adstream"
	//"fmt"
)

type WebSocket struct {
	*revel.Controller
}

func (c WebSocket) AdStreamSocket(ws *websocket.Conn) revel.Result {
	sub := adstream.Subscribe()

	defer sub.Cancel()

	// Send down the archive.
	//fmt.Println("Archive...");
	for _, event := range sub.Archive {
		//fmt.Println("event:" + event.AdId);
		if websocket.JSON.Send(ws, &event) != nil {
			// They disconnected
			return nil
		}
	}

	for {
		select {
		case event := <-sub.New:
			if websocket.JSON.Send(ws, &event) != nil {
				// They disconnected.
				return nil
			}
		}
	}

	return nil
}
