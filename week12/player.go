package main

import "gadget"

// func playList(device gadget.TapePlayer, songs []string) {
// 	for _, song := range songs {
// 		device.Play(song)
// 	}
// 	device.Stop()
// }

type Player interface {
	Play(string)
	// Record()
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	// device.Record()
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	recordPlayer := gadget.TapeRecorder{}
	mixtape := []string{"거짓말", "너랑 나", "농담곰송"}
	mcrtap := []string{"Welcome to my PlayGround", "담곰송", "BetterBetter"}

	playList(player, mixtape)
	playList(recordPlayer, mcrtap)
}
