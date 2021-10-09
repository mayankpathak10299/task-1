package main

import (
	"encoding/json"
	"image"
	"net/http"
	"time"
	"user"
)

/*Create an User*/
type User struct {
    Id string `json:"Id"`
    Name string `json:"Name"`
    Email string `json:"Email"`
    Password String  `json:"Password"`
}

/*Get a user using id*/

func main() {

    user, err := user.Current()
    if err != nil {
        panic(err)
    }

    // Current User
    fmt.Println("Hi " + user.Name + " (id: " + user.Uid + ")")
    fmt.Println("Username: " + user.Username)
    fmt.Println("Home Dir: " + user.HomeDir)

    fmt.Println("Real User: ")
}

/*Create a Post*/
type Post struct {
	Id       string `json:"id"`
	Caption  string `json:"caption"`
	Image_url string `json:"Image_url"`
	time     string `json:"time"`
}
/*Password*/
type Password struct{
	pass char
	}
func main() {
    fmt.Println("Encryption Program v0.01")

    text := []byte("My Super Secret Code Stuff")
    key := []byte("passphrasewhichneedstobe32bytes!")

    c, err := aes.NewCipher(key)
    if err != nil {
        fmt.Println(err)
    }

    gcm, err := cipher.NewGCM(c)
  
    if err != nil {
        fmt.Println(err)
 }

/*Get a post using id*/
package ig-post

import (
	"fmt"
	"os"
	"testing"
)

func GetAllPostCode(t *testing.T) {
	codes, err := GetAllPostCode(
		os.Getenv("IG_TEST_USERNAME"),
		os.Getenv("IG_DS_USER_ID"),
		os.Getenv("IG_SESSIONID"),
		os.Getenv("IG_CSRFTOKEN"))
	if err != nil {
		t.Error(err)
		return
	}
	for _, code := range codes {
		fmt.Printf("URL: https://www.instagram.com/p/%s/\n", code)
	}
}