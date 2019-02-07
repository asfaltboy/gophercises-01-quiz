package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "time"
)

func getCSVData() [][]string {
    file, err := os.Open("quiz.csv")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Opened file %q\n", file.Name())
    reader := csv.NewReader(file)
    // fmt.Printf("Created reader %v\n", reader)
    csvData, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("CSV Data has %v records\n", len(csvData))
    return csvData
}

func checkIfTimeExpired(t time.Ticker) {
    for {
        // <-time.After(1 * time.Second)
        select {
        case <-t.C:
            fmt.Println("1 second tick")
        }
        // go checkIfTimeExpired()
    }
}

func printAnswersRight(correctCount int, total int) {
    fmt.Printf("\nYou answered %v questions correctly (out of %v)", correctCount, total)
}

func main() {
    csvData := getCSVData()
    // fmt.Printf("Hit enter when you're ready to start!")
    // TODO: scan for enter hit and start the counter
    ticker := time.NewTicker(1 * time.Second)
    const timeout = 30
    answeredCorrectly := 0
    go func() {
        seconds := 0
        for {
            // <-time.After(1 * time.Second)
            select {
            case <-ticker.C:
                seconds++
                if seconds >= timeout {
                    fmt.Println("You're out of time! :(")
                    printAnswersRight(answeredCorrectly, len(csvData))
                    os.Exit(0)
                }
            }
        }
    }()
    for i := 0; i < len(csvData); i++ {
        // question, answer := csvData[i]
        question := csvData[i][0]
        correctString := csvData[i][1]
        var correct int
        _, err := fmt.Sscan(correctString, &correct)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("\nQuestion % d: %s\n", i, question)
        var answer int
        if _, err := fmt.Scan(&answer); err == nil && correct == answer {
            answeredCorrectly++
            fmt.Printf("Correct!\n")
        } else {
            fmt.Printf("Incorrect! Correct answer is %d\n", correct)
        }
    }
    printAnswersRight(answeredCorrectly, len(csvData))
}
