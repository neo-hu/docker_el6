package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/neo-hu/docker_el6/docker/pkg/namesgenerator"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(namesgenerator.GetRandomName(0))
}
