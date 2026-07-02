package main

import "go.mod/gadget"

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	mixtape := []string{"Jessie's Girls", "Whip It", "9 to 5"}
	// player := gadget.TapePlayer{}
	var player Player = gadget.TapePlayer{}
	playList(player, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
}
