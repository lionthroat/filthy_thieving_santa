package main

import (
  "fmt"
  "math/rand"
  "os"
  "text/tabwriter"
  "time"
)

// to-do: since golang is fully garbage-collected, I need to eliminate memory
// management functions
func main() {
  rand.Seed(time.Now().UnixNano()) // seed random number generator

  // make a new tabwriter table
  // minwidth, tabwidth, padding, padchar, flags
  t := new(tabwriter.Writer)
  t.Init(os.Stdout, 8, 8, 0, '\t', 0)
  defer t.Flush()

  fmt.Printf("\n\n\n\tFilthy Thieving Santa:\n\tHe's taking back your toys.\n\n\n\n")

  fmt.Printf("\tOk, kid, Santa's going on a heist.\n")
  fmt.Printf("\tWhich bag is he taking?\n\n")
  fmt.Printf("\t1. Dainty Fanny Pack: 10 lb. capacity\n")
  fmt.Printf("\t2. Leather Knapsack:  25 lb. capacity\n")
  fmt.Printf("\t3. SANTA'S TOY BAG:   75 lb. capacity\n")
  fmt.Printf("\tEnter choice: ")
  var choice string
  fmt.Scanln(&choice)
  W:= 75

  min := 3
  max := 20
  var n = rand.Intn(max - min + 1) + min    // number of items available

  var Weights[20]int  // weights of available items
  var Values[20]int   // values of available items

  var Stash[20]string    // everything you could steal
  var Knapsack[20]string // everything you DID steal
  _ = Knapsack

  // nobody asked for this...
  var items = [60]string{
    "guitar", "necklace", "rare frog", "hammer", "coat rack",
    "laptop", "milkshake", "broken printer", "flannel shirt", "toy car",
    "Fabergé egg", "adequate Rembrandt forgery", "Haida flute", "Left air pod",
    "mackerel", "fanny pack", "technical manual", "candle", "undented hydroflask",
    "dented hydroflask", "phone charger", "earrings", "cursed stone", "vase",
    "fossil", "meteorite", "jar of teeth", "L.L. Bean half-zip", "treasure map",
    "skeleton key", "safety deposit box key", "car key", "dubloons", "old coins",
    "stamps", "Beanie baby", "leaf", "pebble", "Russian-English dictionary",
    "crochet scarf", "manila envelope of blackmail material", "fuzz", "NFT", "lipstick",
    "mix tape", "Chaka Khan record", "walkman without batteries", "batteries", "EMF jammer",
    "Rosemary's Baby poster", "ratchet", "can of soup", "aquarium filter", "mermaid soap",
    "Rolex", "diamonds", "forbidden flour", "lottery ticket", "Declaration of Independence"}
  _ = items

  for i := 0; i < n; i++ {
		Values[i] = rand.Intn(220 - 1 + 1) + 1
    Weights[i] = rand.Intn(60 - 1 + 1) + 1
    Stash[i] = items[rand.Intn(60)] // need to update to not pick repeats
	}

  fmt.Fprintf(t, "\n %s\t%s\t%s\t", "Val", "Wt", "Item")
	fmt.Fprintf(t, "\n %s\t%s\t%s\t", "----", "----", "----")
  for i := 0; i < n; i++ {
    fmt.Fprintf(t, "\n $%d\t %d lbs \t%s \t", Values[i], Weights[i], Stash[i])
  }
  fmt.Fprintf(t, "\n\n")

  var maxvalue = OptimalThievery(W, Weights, Values, Stash, Knapsack, n);
  fmt.Printf("N = %d W = %d, Max = $%d\n", n, W, maxvalue);

}

func Free(Weights [20]int, Values [20]int) {
  // needed to learn diff between nil slice and empty slice here
  // https://yourbasic.org/golang/clear-slice/
  // Weights = nil
  // Values = nil
  return
}

// A utility function that returns
// maximum of two integers
func Max(a int, b int) int {
  // had to replace my ternary statement because that's not a thing in Go
  if a > b {
    return a
  } else {
    return b
  }
}

// Returns the maximum value that
// can be put in a knapsack of capacity W
func OptimalThievery(W int, Weights [20]int, Values [20]int, Stash [20]string, Knapsack [20]string,  n int) int {

  s:= 0
  _ = s
  // set up a 2D array (well, a slice of slices...) for my solution grid!
  K := make([][]int, n + 1)
  for i := range K {
    K[i] = make([]int, W + 1)
  }

  for i := 0; i <= n; i++ {
    for w:= 0; w <= W; w++ {
      if (i == 0 || w == 0) {
          K[i][w] = 0
      } else if Weights[i - 1] <= w {
          K[i][w] = Max(Values[i-1] + K[i - 1][w - Weights[i - 1]], K[i - 1][w]);
      } else {
          K[i][w] = K[i-1][w]
      }
    }
  }

  return K[n][W];
}
