package main

import "github.com/tidwall/fastlane"

func main() {
	//var wg
	var ch fastlane.Chan

	go func() {
		for {
			ch.Send("ping")
		}
	}()

	for {
		v := ch.Recv()
		println(v.(string))
	}
}
