package main

import (
  "engo.io/engo"
  "github.com/coderconvoy/frog2/play"
)

func main() {
  opts := engo.RunOptions {
    Width: 500,
    Height: 500,
    Title: "Froggy2",
  }
  
  engo.Run(opts, &play.PlayScene)
  
}