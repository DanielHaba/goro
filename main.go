package main

import (
	// "bufio"
	"fmt"
	"math/rand"
	"os"

	"github.com/DanielHaba/goro/component"

	// "os/exec"
	"time"

	"github.com/stianeikeland/go-rpio"
)

func random(from, to int64) int64 {
	return int64(float32(to-from)*rand.Float32()) + from
}

func randomDuration(from, to time.Duration) time.Duration {
	return time.Duration(random(int64(from), int64(to)))
}

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()
	motor := component.Motor{rpio.Pin(13), rpio.Pin(12), rpio.Pin(6)}
	motor.Setup()
	for i := 0; i < 1000000; i++ {
		motor.Forward()
		motor.On()
		time.Sleep(500 * time.Millisecond)
		motor.Off()
		motor.Backward()
		motor.On()
		time.Sleep(300 * time.Millisecond)
		motor.Off()
		time.Sleep(200 * time.Millisecond)
	}
}

/*
func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer rpio.Close()

	leftMotor := component.Motor{rpio.Pin(20), rpio.Pin(21), rpio.Pin(26)}
	rightMotor := component.Motor{rpio.Pin(13), rpio.Pin(12), rpio.Pin(6)}

	diff := component.Differential{leftMotor, rightMotor}

	diff.Setup()

	r := bufio.NewReader(os.Stdin)
	// disable input buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	for {
		ch, _, err := r.ReadRune()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		switch ch {
		case 'w':
			diff.Forward()
		case 'a':
			diff.Left()
		case 's':
			diff.Backward()
		case 'd':
			diff.Right()
		case 'e':
			return
		default:
			continue
		}
		diff.On()
		time.Sleep(5 * time.Millisecond)
		diff.Off()
	}
}
*/
