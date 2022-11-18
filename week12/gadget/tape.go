package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

type TapeRecorder struct {
	Microphones int
}

func (t TapePlayer) Play(song string) {
	fmt.Println("재생 중", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("정지")
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("재생 중", song)
}
func (t TapeRecorder) Record() {
	fmt.Println("녹음 중")
}

func (t TapeRecorder) Stop() {
	fmt.Println("정지")
}
