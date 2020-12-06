package main

import (
    "bufio"
    "os"
    "fmt"
    "encoding/json"
    "github.com/jeremywohl/flatten"
    // "github.com/fatih/color"
)

var result map[string]interface{}

func main() {
    // c := color.New(color.BgBlack)
    colorCyan := "\033[36m"
    colorReset := "\033[0m"

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := scanner.Text()
        if line == "." {
            break
        }
        flattenJson, _ := flatten.FlattenString(line, "", flatten.DotStyle)
        json.Unmarshal([]byte(flattenJson), &result)
        for k,v := range result {
          fmt.Print(string(colorCyan),"[", k, "] ", string(colorReset), v, " ")
        }
        fmt.Println("\n")
    }
}
