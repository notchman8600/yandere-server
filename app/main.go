package main

import "d-meeting.com/d-server/infra"

func main() {
	router := infra.NewRouter()
	router.Run()
}
