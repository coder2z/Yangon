package tools

import "log"

func MustCheck(err error) {
	if err != nil {
		log.Panic(err)
	}
}
