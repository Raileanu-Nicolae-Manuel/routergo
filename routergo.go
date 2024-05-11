package routergo

import (
	"log"
	"os"
)

func routergo() {
	path := "./api"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0777)
		log.Fatalln(err)
	}
}
