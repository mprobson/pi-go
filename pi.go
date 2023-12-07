package main

// https://github.com/TheRemote/PiBenchmarks/issues/26
// jq, fx, python -m json.tool

import (
  "os/exec"
  "strconv"
  "log"
  "fmt"
  "encoding/json"
  //"net/url"
  "slices"
  "os"
  "cmp"
)

type Board struct {
  Manufacturer string
  Model string
  Countboard int
  Popularity int
  Brand string
  Ranking int
  Offer_url string
  Amz_us string
  Amz_ca string
  Amz_uk string
  Amz_jp string
  Amz_au string
  Amz_de string
  Amz_es string
  Amz_fr string
  Amz_it string
  Amz_nl string
  Amz_pl string
  Amz_se string
  Amz_sg string
  Ali string
}

type Dump struct {
  Result []Board
  Response_code int
  Response_desc string
  Previouspage int
  Nextpage int
}

func printBoard(input Board){
  fmt.Printf("Mfr. : %s\n", input.Manufacturer)
  fmt.Printf("Model: %s\n", input.Model)
  fmt.Printf("Count: %d\n", input.Countboard)
  fmt.Printf("Pop. : %d\n", input.Popularity)
  fmt.Printf("Brand: %s\n", input.Brand)
  fmt.Printf("Rank : %d\n", input.Ranking)
}

func printBoardSpeed(input Board){
  fmt.Printf("Brand: %s\n", input.Brand)
  fmt.Printf("Rank : %d\n", input.Ranking)
}

func cmpBoardSpeed(a Board, b Board) int {
  return cmp.Compare(a.Ranking, b.Ranking)
}

func cmpBoardPop(a Board, b Board) int {
  return cmp.Compare(a.Popularity, b.Popularity)
}

// unfinished - needs to store ouput, maybe -o option
// could also use ranges [1-37] instead of a loop
// could also store data directly as json object
func download() {
  const baseStr = "https://pibenchmarks.com/api/boards/"

  // 37
  for i := 1; i <= 1; i++ {
    // curl -k -L https://pibenchmarks.com/api/benchmark/62444
    curlCmd := exec.Command("curl", "-k", "-L", baseStr + strconv.Itoa(i) + "/")
    //fmt.Println(baseStr + strconv.Itoa(i) + "/")
    //err := curlCmd.Run()
    out, err := curlCmd.Output()
    if err != nil {
      log.Fatal(err) // panic in other tutorial
    }
    fmt.Println(string(out))
  }
}

func main() {
  dat, err := os.ReadFile("./data/dump1.json")
  if err != nil {
    log.Fatal(err) // log.Panic "better"? (both better than normal panic?)
  }
  //fmt.Print(string(dat))

  var data Dump // name
  err = json.Unmarshal(dat, &data) 
  if err != nil { // fn like in by example
    log.Fatal(err)
  }

  boards := data.Result
  //fmt.Printf("%+v", boards)
  //fmt.Println(boards[0])
  printBoard(boards[0]) 
  fmt.Println()

  slices.SortFunc(boards, cmpBoardPop)

  printBoard(boards[0])
  fmt.Println()

  slices.SortFunc(boards, cmpBoardSpeed)

  printBoardSpeed(boards[0])
  fmt.Println()
  for i := 1; i <= 10; i++ {
    printBoardSpeed(boards[i])
    fmt.Println()
  } 
}
