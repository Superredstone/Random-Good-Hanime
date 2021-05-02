package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
	"gopkg.in/yaml.v3"
)

var (
	Log string
)

func main() {
	//Config///////////////////////////////////////////////
	if _, err := os.Stat("config.yml"); os.IsNotExist(err) {
		f, err := os.Create("config.yml")
		if err != nil {
			fmt.Println(err)
		}

		f.WriteString(`port : 8080
token : "1771420037:AAHrIvi1RFh5aUcID_rkS7EXOvjcCYX77sc"`)
	}

	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println("Invalid file")
		f.WriteString(`port : 8080
token : "1771420037:AAHrIvi1RFh5aUcID_rkS7EXOvjcCYX77sc"`)
	}
	defer f.Close()

	var cfg Config

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println(err)
	}

	//Server///////////////////////////////////////////////
	go startServer()
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Nothing to see here..")
	})
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Currently working")
	})
	http.HandleFunc("/", updateLog)

	fmt.Println("Started server on localhost:8080")

	const helpMessage = `SFW
	/neko
	/lewdneko
	/sfwfoxes
	/wallpapers
	/mobileWallpapers
	NSFW
	/hentai
	/ass
	/bdsm
	/cum
	/doujin
	/femdom
	/maid
	/orgy
	/panties
	/nsfwwallpapers
	/nsfwmobilewallpapers
	/netorare
	/gif
	/blowjob
	/feet
	/pussy
	/uglybastard
	/uniform
	/gangbang
	/foxgirl
	/cumslut
	/glasses
	/thighs
	/tentacles
	/masturbation
	/school
	/yuri
	/zettairyouiki`

	//Bot options
	b, err := tb.NewBot(tb.Settings{
		Token:  cfg.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(m *tb.Message) {
		cron("/start", m.Chat.FirstName)
	})
	b.Handle("/help", func(m *tb.Message) {
		b.Send(m.Sender, helpMessage)
		cron("/help", m.Chat.FirstName)
	})
	b.Handle("/ass", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("ass"))}

		b.Send(m.Sender, photo)
		cron("/ass", m.Chat.FirstName)
	})
	b.Handle("/bdsm", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("bdsm"))}

		b.Send(m.Sender, photo)
		cron("/bdsm", m.Chat.FirstName)
	})
	b.Handle("/cum", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("cum"))}

		b.Send(m.Sender, photo)
		cron("/cum", m.Chat.FirstName)
	})
	b.Handle("/doujin", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("doujin"))}

		b.Send(m.Sender, photo)
		cron("/doujin", m.Chat.FirstName)
	})
	b.Handle("/femdom", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("femdom"))}

		b.Send(m.Sender, photo)
		cron("/femdom", m.Chat.FirstName)
	})
	b.Handle("/hentai", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("hentai"))}

		b.Send(m.Sender, photo)
		cron("/hentai", m.Chat.FirstName)
	})
	b.Handle("/maid", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("maid"))}

		b.Send(m.Sender, photo)
		cron("/maid", m.Chat.FirstName)
	})
	b.Handle("/orgy", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("orgy"))}

		b.Send(m.Sender, photo)
		cron("/orgy", m.Chat.FirstName)
	})
	b.Handle("/panties", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("panties"))}

		b.Send(m.Sender, photo)
		cron("/panties", m.Chat.FirstName)
	})
	b.Handle("/nsfwwallpapers", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("nsfwwallpapers"))}

		b.Send(m.Sender, photo)
		cron("/nsfwwallpapers", m.Chat.FirstName)
	})
	b.Handle("/nsfwmobilewallpapers", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("nsfwmobilewallpapers"))}

		b.Send(m.Sender, photo)
		cron("/nsfwmobilewallpapers", m.Chat.FirstName)
	})
	b.Handle("/netorare", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("netorare"))}

		b.Send(m.Sender, photo)
		cron("/netorare", m.Chat.FirstName)
	})
	b.Handle("/gif", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("gif"))}

		b.Send(m.Sender, photo)
		cron("/gif", m.Chat.FirstName)
	})
	b.Handle("/blowjob", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("blowjob"))}

		b.Send(m.Sender, photo)
		cron("/blowjob", m.Chat.FirstName)
	})
	b.Handle("/feet", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("feet"))}

		b.Send(m.Sender, photo)
		cron("/feet", m.Chat.FirstName)
	})
	b.Handle("/pussy", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("pussy"))}

		b.Send(m.Sender, photo)
		cron("/pussy", m.Chat.FirstName)
	})
	b.Handle("/uglybastard", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("uglybastard"))}

		b.Send(m.Sender, photo)
		cron("/uglybastard", m.Chat.FirstName)
	})
	b.Handle("/uniform", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("uniform"))}

		b.Send(m.Sender, photo)
		cron("/uniform", m.Chat.FirstName)
	})
	b.Handle("/gangbang", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("gangbang"))}

		b.Send(m.Sender, photo)
		cron("/gangbang", m.Chat.FirstName)
	})
	b.Handle("/foxgirl", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("foxgirl"))}

		b.Send(m.Sender, photo)
		cron("/foxgirl", m.Chat.FirstName)
	})
	b.Handle("/cumslut", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("cumslut"))}

		b.Send(m.Sender, photo)
		cron("/cumslut", m.Chat.FirstName)
	})
	b.Handle("/glasses", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("glasses"))}

		b.Send(m.Sender, photo)
		cron("/glasses", m.Chat.FirstName)
	})
	b.Handle("/thighs", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("thighs"))}

		b.Send(m.Sender, photo)
		cron("/thighs", m.Chat.FirstName)
	})
	b.Handle("/tentacles", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("tentacles"))}

		b.Send(m.Sender, photo)
		cron("/tentacles", m.Chat.FirstName)
	})
	b.Handle("/masturbation", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("masturbation"))}

		b.Send(m.Sender, photo)
		cron("/masturbation", m.Chat.FirstName)
	})
	b.Handle("/school", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("school"))}

		b.Send(m.Sender, photo)
		cron("/school", m.Chat.FirstName)
	})
	b.Handle("/yuri", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("yuri"))}

		b.Send(m.Sender, photo)
		cron("/yuri", m.Chat.FirstName)
	})
	b.Handle("/zettairyouiki", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("zettai-ryouiki"))}

		b.Send(m.Sender, photo)
		cron("/zettairyouiki", m.Chat.FirstName)
	})

	//SFW
	b.Handle("/neko", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("neko"))}

		b.Send(m.Sender, photo)
		cron("/neko", m.Chat.FirstName)
	})
	b.Handle("/lewdneko", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("lewdneko"))}

		b.Send(m.Sender, photo)
		cron("/lewdneko", m.Chat.FirstName)
	})
	b.Handle("/sfwfoxes", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("sfwfoxes"))}

		b.Send(m.Sender, photo)
		cron("/sfwfoxes", m.Chat.FirstName)
	})
	b.Handle("/wallpapers", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("wallpapers"))}

		b.Send(m.Sender, photo)
		cron("/wallpapers", m.Chat.FirstName)
	})
	b.Handle("/mobileWallpapers", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveHentai("mobileWallpapers"))}

		b.Send(m.Sender, photo)
		cron("/mobileWallpapers", m.Chat.FirstName)
	})

	//Start bot
	fmt.Printf("Bot started\n\n")

	b.Start()
}

func cron(commandToPrint, msgSender string) {
	dt := time.Now()
	dtFormatted := dt.Format("01-02-2006 15:04:05")

	goodFormat := "[" + dtFormatted + "] " + msgSender + " ==> " + commandToPrint + "\n"

	Log = Log + "\n" + goodFormat

	fmt.Println(goodFormat)

	writeLogFile()
}

func retrieveHentai(parameter string) string {
	api := "https://akaneko-api.herokuapp.com/api/"

	response, err := http.Get(api + parameter)
	if err != nil {
		fmt.Println(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var data Website
	json.Unmarshal(responseData, &data)

	return data.Url
}

func startServer() {
	f, err := os.Open("config.yml")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	var cfg Config

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println(err)
	}

	log.Fatal(http.ListenAndServe(":"+cfg.Port, nil))
}

func updateLog(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, Log)
}

func writeLogFile() {
	logFile, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}

	logFile.WriteString(Log)
}

/////////////////////////////////////////////////////////////
type Website struct {
	Url string `json:"url"`
}

type Config struct {
	Port  string `yaml:"port"`
	Token string `yaml:"token"`
}
