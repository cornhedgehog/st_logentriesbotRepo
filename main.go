package main

import (
  "fmt"
  "net/http"
)

func main() {
    // resp_p, err_p := http.Post("login.microsoft.online.com/common/oauth2/v2.0/token", " HTTP/1.1 Host: login.microsoftonline.com Content-Type: application/x-www-form-urlencoded", &buf)
    // resp, err := http.Get("https://api.skype.net")
    // fmt.Println(resp)
    // fmt.Println(err)


    url := "login.microsoftonline.com/common/oauth2/v2.0/token"
        fmt.Println("URL:>", url)

        var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
        req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
        req.Header.Set("type", "message/text")
        req.Header.Set("text", "test")

        req.Body.Set("Grand Type", "client_credentials")
        req.Body.Set("scope", "https://graph.microsoft.com/.default")
        //req.Body.Set("scope", "https://graph.microsoft.com/.default")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            panic(err)
        }
        defer resp.Body.Close()

        fmt.Println("response Status:", resp.Status)
        fmt.Println("response Headers:", resp.Header)
        body, _ := ioutil.ReadAll(resp.Body)
        fmt.Println("response Body:", string(body))


}

  // func handler(w http.ResponseWriter, r *http.Request) {
  //     fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
  // }
  //
  // func main() {
  //     http.HandleFunc("/", handler)
  //     http.ListenAndServe(":8080", nil)
  // }
