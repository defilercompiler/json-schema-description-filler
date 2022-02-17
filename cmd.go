package main

import (
    "encoding/json"
    "fmt"
)
  
// declaring a struct
type Country struct {
  
    // defining struct variables
    Name      string
    Capital   string
    Continent string
}
  
// main function
func main() {
  
    // defining a struct instance
    var country []Country
  
    // JSON array to be decoded
    // to an array in golang
    Data := []byte(`
    [
        {"Name": "Japan", "Capital": "Tokyo", "Continent": "Asia"},
        {"Name": "Germany", "Capital": "Berlin", "Continent": "Europe"},
        {"Name": "Greece", "Capital": "Athens", "Continent": "Europe"},
        {"Name": "Israel", "Capital": "Jerusalem", "Continent": "Asia"}
    ]`)
  
    // decoding JSON array to
    // the country array
    err := json.Unmarshal(Data, &country)
  
    if err != nil {
  
        // if error is not nil
        // print error
        fmt.Println(err)
    }
  
    // printing decoded array
    // values one by one
    for i := range country {
        fmt.Println(country[i].Name + " - " + country[i].Capital + 
                                     " - " + country[i].Continent)
    }
}