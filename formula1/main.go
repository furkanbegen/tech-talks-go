package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"text/tabwriter"
)

//Rankins represents team ranking response
type Rankins struct {
	Response []TeamRanking `json:"response"`
}

// Team represents a Formula 1 team
type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TeamRanking represens a ranking of a team
type TeamRanking struct {
	Position int  `json:"position"`
	Team     Team `json:"team"`
	Points   int  `json:"points"`
	Season   int  `json:"season"`
}

func main() {

	//api key almak için https://rapidapi.com/api-sports/api/api-formula-1 adresini ziyaret edebilirsiniz
	apiKey := flag.String("apikey", "", "API key for team rankings")
	season := flag.String("season", "2019", "Season information for team rankings")
	flag.Parse()

	url := "https://api-formula-1.p.rapidapi.com/rankings/teams?season=" + *season

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "api-formula-1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", string(*apiKey))

	if err != nil {
		panic(err.Error())
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var rankings Rankins
	json.Unmarshal(body, &rankings)

	w := tabwriter.NewWriter(os.Stdout, 35, 0, 2, ' ', tabwriter.TabIndent)

	for _, ranking := range rankings.Response {
		fmt.Fprintf(w, "%s\t\t%d\n", ranking.Team.Name, ranking.Points)
	}

	w.Flush()

}
