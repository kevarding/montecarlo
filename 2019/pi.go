package main

import (
  "fmt"
  "math"
  "math/rand"
  "os"
)



func main() {
  rand.Seed(42)

  const nr_of_pebbles int = 1e4

  var x [nr_of_pebbles]float64
  var y [nr_of_pebbles]float64

  pebbles_in_circ := 0

  f, err := os.Create("data.dat")
  if err != nil {
    fmt.Println(err)
    return
  }

  for i := 0; i < nr_of_pebbles; i++ {
    x[i] = rand.Float64();
    y[i] = rand.Float64();
    fmt.Fprintf(f, "%f, %f\n", x[i], y[i])
    if r := math.Pow(x[i],2)+math.Pow(y[i],2); r < 1 {
      pebbles_in_circ += 1;
    }
  }

  err = f.Close()
  if err != nil {
    fmt.Println(err)
    return
  }

  approxpi := 4.*float64(pebbles_in_circ)/float64(nr_of_pebbles)
  fmt.Printf("Pi (with %d pebbles) is approximately %f\n", nr_of_pebbles, approxpi)
}
