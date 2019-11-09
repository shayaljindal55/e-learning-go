// package controller

// import (
// 	"base/config/db"
// 	"base/model"
// 	"context"
// 	"encoding/json"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/mongodb/mongo-go-driver/bson"
// 	"golang.org/x/crypto/bcrypt"
// )

// func LoginHandler(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	var user model.User
// 	body, _ := ioutil.ReadAll(r.Body)
// 	err := json.Unmarshal(body, &user)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	collection, err := db.GetDBCollection()

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	var result model.User
// 	var res model.ResponseResult

// 	err = collection.FindOne(context.TODO(), bson.D{{"username", user.Username}}).Decode(&result)

// 	if err != nil {
// 		res.Error = "Invalid username"
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))

// 	if err != nil {
// 		res.Error = "Invalid password"
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"username":  result.Username,
// 		"firstname": result.FirstName,
// 		"lastname":  result.LastName,
// 	})

// 	tokenString, err := token.SignedString([]byte("secret"))

// 	if err != nil {
// 		res.Error = "Error while generating token,Try again"
// 		json.NewEncoder(w).Encode(res)
// 		return
// 	}

// 	result.Token = tokenString
// 	result.Password = ""

// 	json.NewEncoder(w).Encode(result)

// }
package controller

