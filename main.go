package main

import (
	"github.com/mayankpathak10299/appointy_assignment/api"
	"encoding/json"
	"image"
	"net/http"
	"time"
	"user"
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	api.StartServer()
}

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
	PostedTimestamp time.Duration
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
