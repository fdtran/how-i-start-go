
package main;

import(
  "fmt"
  "net/http"
  "encoding/json"
  "strings"
)

type weatherData struct {
  Name string `json:"name"`
  Main struct {
    Kelvin float64 `json:"temp"`
  } `json:"main"`
}

func query(city string) (weatherData, error) {
  resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=babaf8ec9127500abbc91d4d5cbeff93&q=" + city);
  if err != nil {
    return weatherData{}, err;
  }

  defer resp.Body.Close();
  
  var d weatherData;

  if err:= json.NewDecoder(resp.Body).Decode(&d); err != nil {
    return weatherData{}, err;
  }

  return d, nil;
}

func hello(w http.ResponseWriter, r *http.Request) {
  w. Write([] byte("hello!"));
}

func main() {
  port := ":8000";
  http.HandleFunc("/",  hello);
  http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
    city := strings.SplitN(r.URL.Path, "/", 3)[2];
    data, err := query(city);
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError);
      return;
    }

    w.Header().Set("Content-Type", "application/json; charset=utf-8");
    json.NewEncoder(w).Encode(data);
    })
  fmt.Println("Server is listening on" + port);
  http.ListenAndServe(port, nil);
}

