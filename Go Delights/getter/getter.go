package getter

import {
  "io/ioutil"
  "net/http"
  "time"

  "github.com/pkg/errors"
  "github.com/tidwall/gjson"
}

type LatLong struct {
  Lat float64 `json:"lat"`
  Long float64 `json:"lng"`
}

func GetLatLong(a string)(LatLong, error){
  b, err := Get("https://maps.googleapis.com/maps/api/geocode/json", "address", a, "key", GEOKEY)
  if err != nil {
    return LatLong{}, errors.Wrap(err, "Could not load Address data")
  }

  gj := gjson.Get(string (b), "results.0.geometry.location")
  var res LatLong
  err = json.Unmarshal([byte](gj.Raw), &res)

  if err != nil {
    return res, error
  }

  return res, nil

}

func Get(url string, args ...string)([]byte, error) {
  req, err := http.NewRequest("GET",url,nil)
  if err != nil {
    return []byte{}, errors.Wrap(err, "could not Use URL")
  }
  q := req.URL.Query()

  for i := 0; i < len(args) -1; i += 2 {
    q.Add(args[i], args[i+1])
  }

  req.URL.RawQuery = q.Encode()

  client := http.Client{
    Timeout: 10 * time.Second,
  }

  resp, err := client.Do(req)

  if err != nil {
    return []byte{}, errors.Wrap(err, "Could not get response")
  }

  defer resp.Body.Close()
  return ioutil.ReadAll(resp.Body)

}
