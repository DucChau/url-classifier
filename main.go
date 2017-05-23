package main

import (
	"fmt"
	textapi "github.com/AYLIEN/aylien_textapi_go"
)

func main() {
	auth := textapi.Auth{"40ca115a", "e2e22cf08f8cf470d51a7fd295939d7b"}
	client, err := textapi.NewClient(auth, true)

	if err != nil {
		panic(err)
	}

	urls := []string{"http://hollywood-481.comfortkeepers.com/home/care-services/respite-care",
		"http://www.hfyny.org/HowYouCanHelp.aspx",
		"http://offbeat.topix.com/slideshow/17194/slide4",
		"https://m.theborgata.com/casino/mbrm/home?login=1",
		"http://www.tomshardware.com/answers/id-2655761/turn-overclocking-asus-suite.html",
		"http://editablecalendar.com/may-2017-printable-calendar-holidays/"}

	fmt.Println("id,label,score,confident")

	for _, url := range urls {
		params := &textapi.ClassifyByTaxonomyParams{URL: url, Taxonomy: "iab-qag"}

		classifications, err := client.ClassifyByTaxonomy(params)

		if err != nil {
			panic(err)
		}

		for _, c := range classifications.Categories {
			fmt.Printf("%s,%s,%f,%t\n", c.Id, c.Label, c.Score, c.Confident)
		}
	}
}

/*
id,label,score,confident
IAB6,Family & Parenting,0.301972,true
IAB6-8,Special Needs Kids,0.297900,true
IAB10-5,Home Repair,0.207285,false
IAB9,Hobbies & Interests,0.069373,false
IAB10,Home & Garden,0.061193,false
IAB9-6,Candle & Soap Making,0.049769,false
IAB16-4,Dogs,0.215377,false
IAB16,Pets,0.096849,true
IAB12-3,Local News,0.579185,false
IAB21-1,Apartments,0.400887,false
IAB25-4,Profane Content,0.204159,false
IAB23-10,Pagan/Wiccan,0.165288,false
IAB13-9,Options,0.114367,false
IAB19-6,Cell Phones,0.070753,true
IAB9,Hobbies & Interests,0.054815,false
IAB19,Technology & Computing,0.054189,false
IAB12,News,0.050953,false
IAB9-28,Screenwriting,0.050697,false
IAB25,Non-Standard Content,0.047173,false
IAB13,Personal Finance,0.046603,false
IAB21,Real Estate,0.044862,false
IAB23,Religion & Spirituality,0.044785,false
IAB17,Sports,0.044729,false
IAB17-24,Paintball,0.043873,false
IAB26-3,Spyware/Malware,0.394163,false
IAB26,Illegal Content,0.085996,true
IAB19-16,Graphics Software,0.085751,true
IAB19,Technology & Computing,0.084258,true
IAB18-6,Accessories,0.355643,false
IAB10-4,Gardening,0.178102,false
IAB19-30,Shareware/Freeware,0.063815,true
IAB18,Style & Fashion,0.060785,false
IAB9,Hobbies & Interests,0.053867,false
IAB19,Technology & Computing,0.051785,false
IAB17,Sports,0.050024,false
IAB10,Home & Garden,0.048766,false
IAB9-4,Birdwatching,0.047254,false
IAB17-42,Walking,0.041589,false
*/
