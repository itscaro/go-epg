package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	"github.com/kataras/iris"
)

func main() {
	//iris.Favicon("./favicon.ico")
	iris.Get("/", func(ctx *iris.Context) {
		ctx.Render("index.html", nil)
	})

	iris.Get("/xml", func(ctx *iris.Context) {
		var result Tv
		data, _ := ioutil.ReadFile("./xmltvutf8.xml")
		xml.Unmarshal(data, &result)

		var str string

		for _, programme := range result.Programme {
			channel := findChannel(programme.Attr_channel, result.Channel)
			str += fmt.Sprintf("%+v-%+v %+v on %+v\n", programme.Attr_start, programme.Attr_stop, programme.Title.Text, channel.Display_name[len(channel.Display_name)-1].Text)

		}
		ctx.Text(iris.StatusOK, str)
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
