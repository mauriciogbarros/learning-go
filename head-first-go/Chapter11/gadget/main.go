package main

import "_/Users/mauricio/Desktop/learning-go/head-first-go/Chapter11/gadget"

func playList(device gadget.TapePlayer, song []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girls", "Whip It", "9 to 5"}
	playList(player, mixtape)
}
