package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/FrancescoIlario/grower/cmd/shutter/conf"
	"github.com/sirupsen/logrus"
	"github.com/stianeikeland/go-rpio/v4"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableLevelTruncation: true,
		DisableTimestamp:       false,
		ForceColors:            true,
		FullTimestamp:          true,
		DisableColors:          false,
	})

	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	c, err := conf.GetConfigurationFromEnv()
	if err != nil {
		return err
	}

	if err := rpio.Open(); err != nil {
		return err
	}
	defer rpio.Close()

	prepare(c)

	pin := rpio.Pin(c.InputPin)
	pin.Detect(rpio.FallEdge)

	fmt.Println("press the button!")

	for {
		if pin.EdgeDetected() { // check if event occured
			fmt.Println("button pressed")
			shutdown(rpio.Pin(c.OutputPin))
		}
		time.Sleep(200 * time.Millisecond)
	}
}

func shutdown(opin rpio.Pin) {
	fmt.Println("shutting down the system")

	opin.High()
}

func prepare(c conf.Configuration) {
	ipin := rpio.Pin(c.InputPin)
	ipin.Input()
	ipin.PullDown()

	opin := rpio.Pin(c.OutputPin)
	opin.Output()
	opin.Low()
}
