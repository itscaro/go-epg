package main

///////////////////////////
/// structs
///////////////////////////

type Actor struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Category struct {
	Attr_lang string `xml:" lang,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Channel struct {
	Attr_id      string          `xml:" id,attr"  json:",omitempty"`
	Display_name []*Display_name `xml:" display-name,omitempty" json:"display-name,omitempty"`
}

type Credits struct {
	Actor     []*Actor     `xml:" actor,omitempty" json:"actor,omitempty"`
	Director  []*Director  `xml:" director,omitempty" json:"director,omitempty"`
	Presenter []*Presenter `xml:" presenter,omitempty" json:"presenter,omitempty"`
}

type Date struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Desc struct {
	Attr_lang string `xml:" lang,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Director struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Display_name struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Episode_num struct {
	Attr_system string `xml:" system,attr"  json:",omitempty"`
	Text        string `xml:",chardata" json:",omitempty"`
}

type Length struct {
	Attr_units string `xml:" units,attr"  json:",omitempty"`
	Text       string `xml:",chardata" json:",omitempty"`
}

type Presenter struct {
	Text string `xml:",chardata" json:",omitempty"`
}

type Programme struct {
	Attr_channel string         `xml:" channel,attr"  json:",omitempty"`
	Attr_start   string         `xml:" start,attr"  json:",omitempty"`
	Attr_stop    string         `xml:" stop,attr"  json:",omitempty"`
	Category     []*Category    `xml:" category,omitempty" json:"category,omitempty"`
	Credits      *Credits       `xml:" credits,omitempty" json:"credits,omitempty"`
	Date         *Date          `xml:" date,omitempty" json:"date,omitempty"`
	Desc         *Desc          `xml:" desc,omitempty" json:"desc,omitempty"`
	Episode_num  []*Episode_num `xml:" episode-num,omitempty" json:"episode-num,omitempty"`
	Length       *Length        `xml:" length,omitempty" json:"length,omitempty"`
	Rating       *Rating        `xml:" rating,omitempty" json:"rating,omitempty"`
	Star_rating  *Star_rating   `xml:" star-rating,omitempty" json:"star-rating,omitempty"`
	Sub_title    *Sub_title     `xml:" sub-title,omitempty" json:"sub-title,omitempty"`
	Title        *Title         `xml:" title,omitempty" json:"title,omitempty"`
}

type Rating struct {
	Attr_system string `xml:" system,attr"  json:",omitempty"`
	Value       *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Root struct {
	Tv *Tv `xml:" tv,omitempty" json:"tv,omitempty"`
}

type Star_rating struct {
	Value *Value `xml:" value,omitempty" json:"value,omitempty"`
}

type Sub_title struct {
	Attr_lang string `xml:" lang,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Title struct {
	Attr_lang string `xml:" lang,attr"  json:",omitempty"`
	Text      string `xml:",chardata" json:",omitempty"`
}

type Tv struct {
	Attr_generator_info_name string       `xml:" generator-info-name,attr"  json:",omitempty"`
	Attr_generator_info_url  string       `xml:" generator-info-url,attr"  json:",omitempty"`
	Attr_source_info_name    string       `xml:" source-info-name,attr"  json:",omitempty"`
	Channel                  []*Channel   `xml:" channel,omitempty" json:"channel,omitempty"`
	Programme                []*Programme `xml:" programme,omitempty" json:"programme,omitempty"`
}

type Value struct {
	Text string `xml:",chardata" json:",omitempty"`
}

///////////////////////////
