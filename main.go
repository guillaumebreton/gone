package main

import (
	"flag"
	"fmt"
	"github.com/gosuri/uilive"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

var w = flag.Int("w", 25, "Duration of a working session")
var s = flag.Int("s", 5, "Duration of a short break")
var l = flag.Int("l", 15, "Duration of a long break")
var p = flag.String("p", "wswswl", "Pattern to apply (for example wswswl)")
var e = flag.String("e", "", "The command to execute when a session is done")

func main() {

	flag.Parse()

	// validate pattern
	pattern := strings.ToLower(*p)
	if pattern == "" {
		fmt.Println("Pattern cannot be empty")
	}
	for _, c := range pattern {
		if c != 'w' && c != 'l' && c != 's' {
			fmt.Printf("Pattern %s contains invalid letter (i.e different from [w,s,l])\n", pattern)
			os.Exit(1)
		}
	}

	writer := uilive.New()
	writer.RefreshInterval = time.Millisecond * 500
	for {
		for _, c := range pattern {
			switch c {
			case 'w':
				confirm(writer, "Start a work session (%d min)? [y/n]", *w)
				runtimer(writer, *w, "Work session : %s")
			case 's':
				confirm(writer, "Take a short break (%d min)? [y/n]", *s)
				runtimer(writer, *s, "Short break : %s")
			case 'l':
				confirm(writer, "Take a long break (%d min)? [y/n]", *l)
				runtimer(writer, *l, "Long break : %s")
			}
		}
		writer.Stop()
	}
}

// Display a confirm message, wait for user input
// If user answer y do nothing else quit
func confirm(writer *uilive.Writer, message string, duration int) {
	fmt.Fprintf(writer, message+"\n", duration)
	writer.Flush()
	c, _, err := getChar()
	if err != nil {
		fmt.Println("Fail to read the response")
		os.Exit(1)
	}
	if c != 'y' {
		os.Exit(1)
	}
}

// runtimer launch a timer and write the counter using the writer
func runtimer(writer *uilive.Writer, duration int, message string) {
	// start listening for updates and render
	d := duration * 60
	writer.Start()
	ticker := time.NewTicker(time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _ = range ticker.C {
			d--
			fmt.Fprintf(writer, message+"\n", formatDuration(d))
			if d == 0 {
				return
			}
		}
	}()
	wg.Wait()
	//execute a command
	if *e != "" {
		v := strings.Split(*e, " ")
		cmd := exec.Command(v[0], v[1:]...)
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Fail to execute command : %s - %v\n", *e, err)
		}
	}
	ticker.Stop()
}

func formatDuration(duration int) string {
	seconds := duration % 60
	minutes := duration / 60
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
