package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/Superredstone/Random-Good-Hanime/Lib"
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

	Log = (Lib.Credits())

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Nothing to see here..")
	})
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Currently working")
	})
	http.HandleFunc("/", updateLog)

	fmt.Println(Lib.Credits())
	fmt.Println("Started server on localhost:" + cfg.Port)

	helpArray := []string{"/neko", "/lewdneko", "/sfwfoxes", "/wallpapers", "/mobileWallpapers", "/hentai", "/ass", "/bdsm", "/cum", "/doujin", "/femdom", "/maid", "/orgy", "/panties", "/nsfwwallpapers", "/nsfwmobilewallpapers", "/netorare", "/git", "/blowjob", "/feet", "/pussy", "/uglybastard", "/uniform", "/gangbang", "/foxgirl", "/cumslut", "/glasses", "/thighs", "/tentacles", "/masturbation", "/school", "/yuri", "/zettairyouiki"}

	//Bot options
	b, err := tb.NewBot(tb.Settings{
		Token:  cfg.Token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	//Handle commands
	for i := 0; i < len(helpArray); i++ {
		handleCommands(helpArray, b, i)
	}

	b.Handle("/start", func(m *tb.Message) {
		b.Send(m.Sender, "Use /help to start watching good animes")
		cron("/start", m.Chat.Username, m.Chat.FirstName)
	})
	b.Handle("/help", func(m *tb.Message) {
		b.Send(m.Sender, Lib.HelpMessage())
		cron("/help", m.Chat.Username, m.Chat.FirstName)
	})
	b.Handle("/besthentai", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveRandomGoodHanimeAPI("random"))}

		b.Send(m.Sender, photo)
		cron("/besthentai", m.Chat.Username, m.Chat.FirstName)
	})
	b.Handle("/cat", func(m *tb.Message) {
		photo := &tb.Photo{File: tb.FromURL(retrieveCat())}

		b.Send(m.Sender, photo)
		cron("/cat", m.Chat.Username, m.Chat.FirstName)
	})

	//Start bot
	fmt.Printf("Bot started\n\n")

	b.Start()
}

func handleCommands(helpArray []string, b *tb.Bot, i int) {
	b.Handle(helpArray[i], func(m *tb.Message) {
		helpWithoutSlash := strings.Replace(helpArray[i], "/", "", 1)

		photo := &tb.Photo{File: tb.FromURL(retrieveHentai(helpWithoutSlash))}

		b.Send(m.Sender, photo)
		cron(helpArray[i], m.Chat.Username, m.Chat.FirstName)
	})
}

func cron(commandToPrint, msgUsername, msgFirstName string) {
	dt := time.Now()
	dtFormatted := dt.Format("01-02-2006 15:04:05")

	goodFormat := "[" + dtFormatted + "] @" + msgUsername + " " + msgFirstName + " ==> " + commandToPrint + "\n"

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
func retrieveRandomGoodHanimeAPI(parameter string) string {
	api := "https://random-good-hanime-api.herokuapp.com/api/v1/hentai/"

	response, err := http.Get(api + parameter)
	if err != nil {
		fmt.Println(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var data HentaiAPI
	json.Unmarshal(responseData, &data)

	return data.File
}

func retrieveCat() string {
	api := "https://aws.random.cat/meow"

	response, err := http.Get(api)
	if err != nil {
		fmt.Println(err)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var data catWebsite
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
	//Update localhost:{port}
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

type HentaiAPI struct {
	File string `json:"file"`
}

type catWebsite struct {
	Url string `json:"file"`
}

type Config struct {
	Port  string `yaml:"port"`
	Token string `yaml:"token"`
}
