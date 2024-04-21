package main

import (
	"derhauck/driving-journal-estimate/cmd"
)

// @Version 1.0.0
// @Title Driving Journal Estimate Server
// @Description API to use CLI features via REST
// @ContactName derhauck
// @Server https://driving.kateops.com Server-1
func main() {
	cmd.Execute()
}
