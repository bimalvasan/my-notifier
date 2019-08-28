package collectors

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func GetUnassignedTickets()  {	
    response, err := http.Get("https://jira")
    if err != nil {
        fmt.Printf("The JIRA HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}