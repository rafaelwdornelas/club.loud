package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

var abertos []string

type Creators []struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	AvatarURL     string `json:"avatar_url"`
	BackgroundURL string `json:"background_url"`
	ChannelURL    string `json:"channel_url"`
	Status        string `json:"status"`
	Creator       struct {
		Name string `json:"name"`
	} `json:"creator"`
}

func get_content() {
	url := "https://community-api.dropull.gg/api/v1/partners/993174a8-6f2b-43d9-8b95-578724075a8a/creators"

	res, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	} else {

		body, err := ioutil.ReadAll(res.Body)

		if err != nil {
			fmt.Println(err)
		} else {
			var data Creators
			json.Unmarshal(body, &data)
			abertostmp := abertos
			abertos = nil
			fmt.Println(len(data), "criadores encontrados")
			//loop
			for i := 0; i < len(data); i++ {
				//verifica se já foi aberto antes
				if contains(abertostmp, data[i].Username) {
					log.Println(data[i].Username + " Já foi aberto")
					abertos = append(abertos, data[i].Username)
				} else {

					abertos = append(abertos, data[i].Username)
					open(data[i].ChannelURL)
					log.Println(data[i].Username + " aberto")
				}
			}
		}
	}

	//aguardar 1 minutos
	time.Sleep(1 * time.Minute)
	//chamar a função novamente
	get_content()
}

// verifica se o valor já existe no array
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// abre o link no navegador padrão
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	get_content()
}
