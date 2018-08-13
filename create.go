package main

import (
	"LazyMe/lazy"
)

func main() {
	lc := lazy.NewClient()

	lc.CreateRepo()
}
