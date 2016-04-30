package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/jangaraj/dockereventbeat/beater"
)

func main() {
	err := beat.Run("dockereventbeat", "1.0.0-rc1", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
