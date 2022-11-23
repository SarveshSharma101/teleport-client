package utility

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type Teleport struct {
	Identity_key string `yaml:"identity_key"`
	Address      string `yaml:"address"`
}

func GetIdentityKey(path string) Teleport {

	config, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	teleportConfig := &map[string]Teleport{}
	err = yaml.Unmarshal(config, teleportConfig)
	if err != nil {
		panic(err)
	}

	return (*teleportConfig)["teleport"]
}

func GetLastUpdateDuration(duration string) int {
	location, err := time.LoadLocation("CET")
	if err != nil {
		fmt.Println(err)
	}
	tNow := time.Now().In(location)
	fmt.Println("------------------------>", tNow)
	if tNow.Hour() >= 7 && tNow.Hour() < 22 {
		hm := strings.Split(duration, ":")
		h, err := strconv.Atoi(hm[0])
		if err != nil {
			fmt.Println("err while converting hour from string to int:", err)
		}
		m, err := strconv.Atoi(strings.ReplaceAll(hm[1], " ", ""))
		if err != nil {
			fmt.Println("err while converting min from string to int:", err)
		}

		tEdge := time.Date(tNow.Year(), tNow.Month(), tNow.Day(), h, m, 0, 0, location)
		// fmt.Println("------------->", tEdge)
		tSub := tNow.Sub(tEdge).Minutes()
		if int(tSub) == 0 {
			return 1
		}
		return int(tSub)
	}
	return -100
}
