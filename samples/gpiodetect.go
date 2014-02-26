package main

import (
	"time"

	"github.com/kidoman/embd"
	"github.com/kidoman/embd/host"
)

func main() {
	h, _, err := host.Detect()
	if err != nil {
		return
	}

	var pinNo interface{}

	switch h {
	case host.BBB:
		pinNo = "P9_31"
	case host.RPi:
		pinNo = 10
	default:
		panic("host not supported (yet :P)")
	}

	gpio, err := embd.NewGPIO()
	if err != nil {
		panic(err)
	}
	defer gpio.Close()

	led, err := gpio.DigitalPin(pinNo)
	if err != nil {
		panic(err)
	}
	defer led.Close()

	if err := led.SetDir(embd.Out); err != nil {
		panic(err)
	}
	if err := led.Write(embd.High); err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	if err := led.SetDir(embd.In); err != nil {
		panic(err)
	}
}
