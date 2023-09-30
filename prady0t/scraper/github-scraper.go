package scraper

import (
	//"fmt"

	"github.com/fatih/color"
	"github.com/gocolly/colly"
)

type data map[string]string


func (d data)github_scraper(s string){

	d["github_username"] = s
	
	c := colly.NewCollector()
	c.Visit("github.com/"+s)
}

func(d data) printer(){
	black := color.New(color.FgHiWhite, color.BgBlack)
	for key,value := range d{
		black.Printf("%s -> %s",key, value)
	}
}

func Finalized_github_output(){
	d := make(data)
	d.github_scraper("prady0t")
	d.printer()
}