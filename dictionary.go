package main

import (
    "os"
    "fmt"
    "flag"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "time"
)

const version = "1.0.0"

type Entry struct {
    Fl          string      `json:"fl"`
    Hwi         Hwi         `json:"hwi"`
    ShortDef    []string    `json:"shortdef"`
}

type Hwi struct {
    HW string `json:"hw"`
    Prs []Pr   `json:"prs"`
}

type Pr struct {
    MW string `json:"mw"`
}

type DictionaryTool struct {
    APIKey string
}


func (dt *DictionaryTool) FetchWordData(word string) (*Entry, error) {
    url := fmt.Sprintf("https://www.dictionaryapi.com/api/v3/references/collegiate/json/%s?key=%s", word, dt.APIKey)

    netClient := http.Client{Timeout: time.Second * 10}
    resp, err := netClient.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error fetching data: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %v", err)
    }

    var entries [1]Entry
    if err := json.Unmarshal(body, &entries); err != nil {
        return nil, fmt.Errorf("Requested word does not exists in dictionary. \nPossible sugestions %v", string(body))
    }
    if len(entries) == 0 {
        return nil, fmt.Errorf("no entries found")
    }
    return &entries[0], nil
}

func (dt *DictionaryTool) QueryDictionary(word string) {
    entry, err := dt.FetchWordData(word)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    if len(entry.Hwi.Prs) > 0 {
        fmt.Printf("%s (%s): %s\n", entry.Hwi.Prs[0].MW, entry.Fl, entry.ShortDef[0])
    } else {
        fmt.Printf("%s (%s): %s\n", word, entry.Fl, entry.ShortDef[0])   
    }
}


func main() {
    showVersion := flag.Bool("V", false, "Print the version")
    flag.Parse()
    if *showVersion {
        fmt.Println("Application Version:", version)
        os.Exit(0)
    }

    if len(os.Args) < 2 {
        fmt.Println("Please provide a word to search.")
        os.Exit(1)
    }

    word := os.Args[1]
    apiKey := os.Getenv("DICTIONARY_API_KEY")
    if apiKey == "" {
        fmt.Println("API key not set. Please set the DICTIONARY_API_KEY environment variable.")
        os.Exit(1)
    }

    tool := DictionaryTool{ APIKey: apiKey }
    tool.QueryDictionary(word)
}
