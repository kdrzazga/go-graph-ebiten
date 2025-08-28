package main

/*
go get github.com/faiface/beep/vorbis
go get github.com/faiface/beep/speaker@v1.1.0
*/
import (
    "github.com/faiface/beep"
    "github.com/faiface/beep/vorbis"
    "github.com/faiface/beep/speaker"
    "os"
    "log"
    "time"
)

func main() {
    f, err := os.Open("sound.ogg")
    if err != nil {
        log.Fatal(err)
    }
    stream, format, err := vorbis.Decode(f)
    if err != nil {
        log.Fatal(err)
    }
    defer stream.Close()

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
    speaker.Play(beep.Seq(stream, beep.Callback(func() {
        // done
    })))
    select{}
}
