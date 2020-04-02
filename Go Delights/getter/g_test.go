package getter

var ADDRS []string = []string{
  "BA5 8NP", "BS1 5SP", "BA3 4PQ", "BA1 3TJ",
  "BA7 1PQ", "BA3 4LJ", "BA2 4SD", "BA4 5D0",
  "BS7 1PQ", "BS3 4LJ", "BS2 4SD", "BS4 5D0",
  "CR7 1PQ", "CR3 4LJ", "CR2 4SD", "CR4 5D0",
  "YA7 1PQ", "YA3 4LJ", "YA2 4SD", "YA4 5D0",
}

func Test_Get(t *testing.T){

  wg := workers.New(10)
  locs := (make[]LatLong, len(ADDRS))

  doget := func(a string, i int) {
    return func ()  {
      l, err := GetLatLong(a)
      if err != nil {
        t.Log(err)
      }
      locs[i] = l
    }
  }

  for i, a := range ADDRS {
    wg.Add(doget(a, i))
    }

    wg.Wait()

  fmt.Println(locs)
}
