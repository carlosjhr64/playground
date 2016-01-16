// Just a way to play with and test my go libraries.
package main

import (
  "os"
  "fmt"
  "github.com/carlosjhr64/jd"
  "github.com/carlosjhr64/to"
  "github.com/carlosjhr64/log"
  "github.com/carlosjhr64/hopt"
  "github.com/carlosjhr64/semantic"
)

func checks(){
  likes := semantic.Likes
  likes(jd.VERSION,       "jd-0.0.1")
  likes(to.VERSION,       "to-1.0.0")
  likes(log.VERSION,      "log-0.1.0")
  likes(hopt.VERSION,     "hopt-0.0.2")
  likes(semantic.VERSION, "semantic-2.0.0")
}

const VERSION = "0.0.0.alpha"

func initialize() {
  hopt.Help = `The Playground!

Usage:
  %s [options]

Options:
  --h --help     Show this help text.
  -v --version   Show the version.
  --jd=DATE      Give it a valid date, get a valid jd number.
  --validate     Ensure the data is valid.
`
  hopt.Version = VERSION
}

// Aliases
var puts = fmt.Println
var printf = fmt.Printf
var fp = fmt.Fprintf
var sto = os.Stdout
var ste = os.Stderr

// Local Globals
var date string = ""

func configure() {
  hopt.Parse()
  date = hopt.Tos("--jd")
  validate := hopt.Tob("--validate")
  if validate && date != "" { to.Date(&date) }
  hopt.Destroy()
}

func play(){
  if date != "" {
    f := log.New("Date \"%s\" get me \"%v\".\n")
    n, e := jd.ToNumber(date)
    if e != nil {
      f.Out(date, e)
    } else {
      f.Out(date, n)
    }
  }
}

func main() {
  checks()
  initialize()
  configure()
  play()
}
