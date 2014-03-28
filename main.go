package main

import (
    "code.google.com/p/google-api-go-client/calendar/v3"
)

func main() {
    args := os.Args


    var config = &oauth.Config{
        ClientId:     "", // from https://code.google.com/apis/console/
        ClientSecret: "", // from https://code.google.com/apis/console/
        Scope:        calendar.CalendarScope,
        AuthURL:      "https://accounts.google.com/o/oauth2/auth",
        TokenURL:     "https://accounts.google.com/o/oauth2/token",
    }
}
