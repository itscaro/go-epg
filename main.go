package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"strconv"
	"time"

	"golang.org/x/net/html/charset"

	"github.com/iris-contrib/middleware/logger"
	"github.com/kataras/go-template/html"
	"github.com/kataras/iris"
)

const (
	timeXmlLayout = "20060102150405 -0700"
	timeLayout    = "15:04"
)

type MyProgramme struct {
	Title        string
	Channel      string
	Start        time.Time
	End          time.Time
	DivStyle     string
	RawProgramme *Programme
	RawChannel   *Channel
}

func main() {
	iris.Config.IsDevelopment = true

	iris.UseTemplate(html.New(html.Config{
		Funcs: template.FuncMap{
			"attr": func(s string) template.HTMLAttr {
				return template.HTMLAttr(s)
			},
			"safe": func(s string) template.HTML {
				return template.HTML(s)
			},
			"css": func(s string) template.CSS {
				return template.CSS(s)
			},
			"formatTime": func(t time.Time) string {
				return t.Format(timeLayout)
			},
		},
	}))

	iris.Use(logger.New())

	iris.Get("/prime", func(ctx *iris.Context) {
		_, epgByDate := processData()

		for _, pByChannel := range epgByDate[time.Now().Format("20060102")] {
			for _, p := range pByChannel {
				length := 0
				if p.RawProgramme.Length != nil {
					length, _ = strconv.Atoi(p.RawProgramme.Length.Text)
				}
				if p.Start.Hour() > 18 && length > 45 {
					ctx.Write("%s %s %s: %s\n", p.Title, p.Channel, p.Start.Format(timeLayout), p.RawProgramme.Desc.Text)
				}
			}
		}
	})

	//iris.Favicon("./favicon.ico")
	iris.Get("/", func(ctx *iris.Context) {
		epgByChannel, epgByDate := processData()

		by := ctx.URLParam("by")
		switch by {
		case "channel":
			ctx.MustRender("epgByChannel.html", epgByChannel)
		case "date":
			ctx.MustRender("epgByDate.html", epgByDate)
		default:
			ctx.MustRender("epg.html", epgByDate[time.Now().Format("20060102")])
		}
	})

	iris.Listen(":8080")
}

func findChannel(id string, channels []*Channel) *Channel {
	for _, channel := range channels {
		if channel.Attr_id == id {
			return channel
		}
	}

	return nil
}

func processData() (epgByChannel map[string]map[string][]MyProgramme, epgByDate map[string]map[string][]MyProgramme) {
	var result Tv

	data, _ := ioutil.ReadFile("./tnt.xml")
	decoder := xml.NewDecoder(bytes.NewReader(data))
	decoder.CharsetReader = charset.NewReaderLabel
	if err := decoder.Decode(&result); err != nil {
		iris.Logger.Panicln(err)
	}

	epgByChannel = make(map[string]map[string][]MyProgramme)
	epgByDate = make(map[string]map[string][]MyProgramme)

	for _, programme := range result.Programme {
		channel := findChannel(programme.Attr_channel, result.Channel)
		channelName := channel.Display_name[len(channel.Display_name)-1].Text
		start, _ := time.Parse(timeXmlLayout, programme.Attr_start)
		end, _ := time.Parse(timeXmlLayout, programme.Attr_stop)
		str := fmt.Sprintf(
			"%+v-%+v %s on %s\n",
			start.Format(timeLayout),
			end.Format(timeLayout),
			programme.Title.Text,
			channel.Display_name[len(channel.Display_name)-1].Text,
		)
		if programme.Credits != nil {
			if len(programme.Credits.Director) > 0 {
				str += fmt.Sprint("Director\n")
				for _, director := range programme.Credits.Director {
					str += fmt.Sprintf("- %s\n", director.Text)
				}
			}

			if len(programme.Credits.Actor) > 0 {
				str += fmt.Sprint("Actors\n")
				for _, actor := range programme.Credits.Actor {
					str += fmt.Sprintf("- %s\n", actor.Text)
				}
			}

			if len(programme.Credits.Presenter) > 0 {
				str += fmt.Sprint("Presenters\n")
				for _, presenter := range programme.Credits.Presenter {
					str += fmt.Sprintf("- %s\n", presenter.Text)
				}
			}
		}
		str += fmt.Sprintf("\n")

		styleWidth := end.Sub(start).Minutes() / 60 / 24 * 100
		p := MyProgramme{
			Title:        programme.Title.Text,
			Channel:      channel.Display_name[len(channel.Display_name)-1].Text,
			Start:        start,
			End:          end,
			DivStyle:     "width: " + strconv.FormatFloat(styleWidth, 'f', 2, 64) + "%;",
			RawProgramme: programme,
			RawChannel:   channel,
		}
		//fmt.Printf("%+v", p)
		if epgByChannel[channelName] == nil {
			epgByChannel[channelName] = make(map[string][]MyProgramme)
		}
		if epgByDate[start.Format("20060102")] == nil {
			epgByDate[start.Format("20060102")] = make(map[string][]MyProgramme)
		}

		epgByChannel[channelName][start.Format("20060102")] = append(epgByChannel[channelName][start.Format("20060102")], p)
		epgByDate[start.Format("20060102")][channelName] = append(epgByDate[start.Format("20060102")][channelName], p)
	}

	return
}
