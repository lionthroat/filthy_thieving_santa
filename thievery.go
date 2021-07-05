package main

import (
  "fmt"
  "math/rand"
  "time"
)

// to-do: since golang is fully garbage-collected, I need to eliminate memory
// management functions
func main() {
  rand.Seed(time.Now().UnixNano()) // seed random number generator

  fmt.Printf("Filthy Thieving Knapsack\n")

  // want three different sized knapsacks, but largest by default for now
  // fanny pack:                 10 lbs
  // artisanal leather knapsack: 25 lbs
  // santa's toy bag:            75 lbs
  W:= [3]int{10, 25, 75}
  _ = W

  var n = 0      // number of items available
  _ = n
  // used to initialize pointers to arrays here, but that needs to be ported
  // bc the array size is *part* of the type in Go...

  var Weights[20]int  // weights of available items
  var Values[20]int   // values of available items
  var Stash[20]string //everything you stole
  _ = Weights
  _ = Values
  _ = Stash

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
    "Rolex", "polished charoite", "forbidden flour", "lottery ticket", "Declaration of Independence"}
  _ = items

  // for (int i = 0; i < 5; i++) {
  //   n = rand() % 20 + 1;
  //   Values = GenerateValues(n);
  //   Weights = GenerateWeights(n);
  //   int MaxValue_Dynamic = Dynamic_Knapsack(W, Weights, Values, n);
  //   printf("N = %d W = %d, Max Recursive = $%d, Max Dynamic = $%d\n", n, W, MaxValue_Recursive, MaxValue_Dynamic);
  // }
  // return(0);
}

// Fill an array with pseudo-random weights,
// with values from 1 to 80. Return pointer.
func GenerateWeights(n int) int {
  var x = 5
  _ = x
  // int* tmparr = (int*)malloc(n * sizeof(int));
  //
  // for (int i = 0; i < n; i++) {
  //   int tmpwt = rand() % 80 + 1;
  //   tmparr[i] = tmpwt;
  // }
  // return tmparr;
  return 1
}

// Fill an array with pseudo-random dollar values
// from 1 to 120. Return pointer.
func GenerateValues(n int) int {
  var x = 5
  _ = x
  // int* tmparr = (int*)malloc(n * sizeof(int));
  //
  // for (int i = 0; i < n; i++) {
  //   int tmpvl = rand() % 120 + 1;
  //   tmparr[i] = tmpvl;
  // }
  // return tmparr;
  return 1
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
func Knapsack(W int, Weights [20]int, Values [20]int, n int) int {
  var i, w int
  _ = i
  _ = w
    // int K[n + 1][W + 1];
    //
    // // Build table K[][] in bottom up manner
    // for (i = 0; i <= n; i++)
    // {
    //     for (w = 0; w <= W; w++)
    //     {
    //         if (i == 0 || w == 0)
    //             K[i][w] = 0;
    //         else if (Weights[i - 1] <= w)
    //             K[i][w] = Max(Values[i - 1]
    //                       + K[i - 1][w - Weights[i - 1]],
    //                       K[i - 1][w]);
    //         else
    //             K[i][w] = K[i - 1][w];
    //     }
    // }
    //
    // return K[n][W];
    return 1
}
