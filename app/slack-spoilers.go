package slackspoilers;

import (
    "encoding/json"
    "encoding/base64"
    "net/http"
)

func init() {
   http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
     //Read the Request Parameter "command"
    command := r.FormValue("command")
    text := r.FormValue("text")
    var response string
     //Ideally do other checks for tokens/username/etc
    if command == "/encode" || command =="/spoiler" {
        response = string(base64Encode([]byte(text)))  
    } else if command == "/decode" || command == "/unspoiler" {
        enbyte, err := base64Decode(text)
        if err != nil {
            response = (err.Error())
        } else {
            response = (string(enbyte))
        }
    }else {
        response = "I do not understand your command:" + "test"
    }

    slackResponse := map[string]string{"response_type": "ephemeral", "text": response}
    js, err := json.Marshal(slackResponse)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}

func base64Encode(src []byte) string {
    return base64.StdEncoding.EncodeToString(src)
}

func base64Decode(src string) ([]byte, error) {
    return base64.StdEncoding.DecodeString(src)
}
