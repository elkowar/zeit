package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"flag"

	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"

  "github.com/atotto/clipboard"
)

func main() {
	toClipboard := flag.Bool("c", false, "Write to clipboard")
	asDiscord := flag.Bool("d", false, "Format as discord timestamp")
	flag.Parse()
	text := flag.Args()


	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	r, err := w.Parse(strings.Join(text, " "), time.Now())
	if err != nil {
		log.Fatal(err)
	}
	if r == nil {
		log.Fatal("No valid time description found")
	}

	var output string
	if *asDiscord {
		output = fmt.Sprintf("<t:%d>\n", r.Time.Unix())
	} else {
		output = fmt.Sprintf("%d", r.Time.Unix())
	}
	if *toClipboard {
		clipboard.WriteAll(output)
	}
	fmt.Println(output)
}
