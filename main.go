package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//////////////////////////// User ///////////////////////////////////
	urlUser := "http://localhost:8000/api/user"

	respUser, err := http.Get(urlUser)
	if err != nil {
		fmt.Println("Error fetching User:", err)
		return
	}
	defer respUser.Body.Close()

	bodyUser, err := ioutil.ReadAll(respUser.Body)
	if err != nil {
		fmt.Println("Error reading User response:", err)
		return
	}

	fmt.Println("User response:", string(bodyUser))

	//////////////////////////// Avatar ///////////////////////////////////
	urlAvatar := "http://localhost:8000/api/avatar"

	respAvatar, err := http.Get(urlAvatar)
	if err != nil {
		fmt.Println("Error fetching Avatar:", err)
		return
	}
	defer respAvatar.Body.Close()

	bodyAvatar, err := ioutil.ReadAll(respAvatar.Body)
	if err != nil {
		fmt.Println("Error reading Avatar response:", err)
		return
	}

	fmt.Println("Avatar response:", string(bodyAvatar))

	//////////////////////////// Diamond ///////////////////////////////////
	urlDiamond := "http://localhost:8000/api/diamond"

	respDiaomnd, err := http.Get(urlDiamond)
	if err != nil {
		fmt.Println("Error fetching Avatar:", err)
		return
	}
	defer respDiaomnd.Body.Close()

	bodyDiamond, err := ioutil.ReadAll(respDiaomnd.Body)
	if err != nil {
		fmt.Println("Error reading Avatar response:", err)
		return
	}

	fmt.Println("Diamond response:", string(bodyDiamond))

	//////////////////////////// Quiz ///////////////////////////////////
	urlQuiz := "http://localhost:8000/api/diamond"

	respQuiz, err := http.Get(urlQuiz)
	if err != nil {
		fmt.Println("Error fetching Avatar:", err)
		return
	}
	defer respDiaomnd.Body.Close()

	bodyQuiz, err := ioutil.ReadAll(respQuiz.Body)
	if err != nil {
		fmt.Println("Error reading Avatar response:", err)
		return
	}

	fmt.Println("Diamond response:", string(bodyQuiz))

}
