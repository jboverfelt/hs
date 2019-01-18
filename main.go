package main

import (
	"net/http"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	resp, err := http.Get(os.Args[1])
	checkErr(err)
	defer resp.Body.Close()

	s, format, err := mp3.Decode(resp.Body)
	checkErr(err)

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	checkErr(err)

	done := make(chan struct{})
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))

	<-done
}
