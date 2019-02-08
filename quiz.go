package main

import (
    "encoding/csv"
    "flag"
    "fmt"
    "log"
    "os"
    "time"
)

type problem struct {
    question string
    answer   string
}

func getCSVData(fileName string) []problem {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    fmt.Printf("Opened file %q\n", file.Name())
    reader := csv.NewReader(file)
    // fmt.Printf("Created reader %v\n", reader)
    csvData, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("CSV Data has %v records\n", len(csvData))
    var ret = make([]problem, len(csvData))
    for i, line := range csvData {
        ret[i] = problem{
            question: line[0],
            answer:   line[1],
        }
    }
    return ret
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
    // parse arguments using flag
    var fileName = flag.String("quizCSV", "quiz.csv", "The path to the CSV file of questions/answers")
    var timeout = flag.Int("timeout", 30, "The timeout in seconds to wait for")
    flag.Parse()
    var problems = getCSVData(*fileName)
    // fmt.Printf("Hit enter when you're ready to start!")
    // TODO: scan for enter hit and start the counter
    ticker := time.NewTicker(1 * time.Second)
    answeredCorrectly := 0
    go func() {
        seconds := 0
        for {
            // <-time.After(1 * time.Second)
            select {
            case <-ticker.C:
                seconds++
                if seconds >= *timeout {
                    fmt.Println("You're out of time! :(")
                    printAnswersRight(answeredCorrectly, len(problems))
                    os.Exit(0)
                }
            }
        }
    }()
    for i, problem := range problems {
        fmt.Printf("\nQuestion %d: %s\n", i, problem.question)
        var answer string
        if _, err := fmt.Scan(&answer); err == nil && problem.answer == answer {
            answeredCorrectly++
            fmt.Printf("Correct!\n")
        } else {
            fmt.Printf("Incorrect! Correct answer is %v\n", problem.answer)
        }
    }
    printAnswersRight(answeredCorrectly, len(problems))
}
