package main

import (
	// "io"
	"log"
	"net/http"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"

)
//Get in webpage
func getLinks(url string) ([]string){

	// Get the body of url
	resp, err := http.Get(url)
	if err != nil {
	   log.Fatalln(err)
	}
	defer resp.Body.Close()	

	// get the html from body 
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var numberOfHyperlinks int = 0 
	var hyperLinks []string

	//function for filtering out just links
	filterhtml := func(i int, s *goquery.Selection) bool{
        link, _ := s.Attr("href")
		stringLink := strings.HasPrefix(link, "https")
        return stringLink
    }

	//find all href tags
	document.Find("body a").FilterFunction(filterhtml).Each(func(_ int, tag *goquery.Selection) {
        link, _ := tag.Attr("href")
		//filter out just one domain
		if strings.HasPrefix(link, url){
			numberOfHyperlinks = numberOfHyperlinks + 1 
			// works the first time around 
			hyperLinks = append(hyperLinks, link)
			fmt.Println("hyperlinks in document find")
			fmt.Println(hyperLinks)

			}	
    })
	return hyperLinks
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func getAllLinks(url string) []string{

	var links []string 
	// add links on first page to link
	links = getLinks(url)
	var allLinks []string
	// add links to allLinks 	
	allLinks = append(allLinks, links...)
	var linksOnPage []string
	// for all of allLinks, find all of their links
	for i := 0; i < len(allLinks); i++{
		fmt.Println("printing all link index")
		fmt.Println(allLinks[i])
		fmt.Println(i)
		//PROBLEM HERE 
		linksOnPage = getLinks(allLinks[i])
		// only returning 0 or 1 link from getlinks
		fmt.Println("printing lenth of linksOnPage")
		fmt.Println(len(linksOnPage))

		// go through all links on on links on a page and if not in allLinks then add them 
		// get the lenth of the links being added 
		for i := 0; i < len(linksOnPage); i++{
			// fmt.Println("code getting to here")
			if contains(allLinks, linksOnPage[i]) == false{
			allLinks = append(allLinks, linksOnPage[i])
			//NOT GETTING HERE
			// fmt.Println("code getting to beyond if statement")
			// fmt.Println(linksOnPage[i])
			}
		}
	}

	return allLinks
}

func main(){
	// var url string;

	//print statement 
	fmt.Println("I'm a crawler")
	// fmt.Printf("Please enter the domain you would like to search:")

	//get domain to crawl
	// fmt.Scanf("%s", &url)
	// fmt.Println(url)

	allLinks := getAllLinks("https://www.vodafone.com")
	fmt.Println("all the links in the domain")
	fmt.Println(allLinks)
	
}