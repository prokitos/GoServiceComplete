package tasks

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

var strConnect = "http://localhost:8888"

func MainRest() {

	var checker bool = false

	for checker == false {
		clearScreen()
		fmt.Println("Select number of operation")
		fmt.Println("1. insert new person")
		fmt.Println("2. delete person by id")
		fmt.Println("3. show all person")
		fmt.Println("4. show person by parameters")
		fmt.Println("5. update person")
		fmt.Println("9. exit")

		var temp string
		fmt.Scanln(&temp)

		switch temp {
		case "1":
			operationInsert()
		case "2":
			operationDelete()
		case "3":
			operationShowAll()
		case "4":
			operationShowByParameters()
		case "5":
			operationUpdate()
		case "9":
			checker = true
			fmt.Println("Okey. Goodbye")
		default:
			fmt.Println("wrong! select correct number")
		}

		fmt.Scanln()
	}

}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func operationUpdate() {
	clearScreen()

	var id string
	var name string
	var surname string
	var patronymic string
	var age string
	var sex string
	var nationality string

	fmt.Print("Please enter person id to update: ")
	fmt.Scanln(&id)
	fmt.Print("Please enter new person name: ")
	fmt.Scanln(&name)
	fmt.Print("Please enter new person surname: ")
	fmt.Scanln(&surname)
	fmt.Print("Please enter new person patronymic: ")
	fmt.Scanln(&patronymic)
	fmt.Print("Please enter new person age: ")
	fmt.Scanln(&age)
	fmt.Print("Please enter new person sex: ")
	fmt.Scanln(&sex)
	fmt.Print("Please enter new person nationality: ")
	fmt.Scanln(&nationality)

	sendRequestUpdate("/update", id, name, surname, patronymic, age, sex, nationality)
	println("Request to update has sended")
}

func sendRequestUpdate(additionalConnect string, id string, name string, surname string, patronymic string, age string, sex string, nationality string) {

	baseURL, _ := url.Parse(strConnect + additionalConnect)
	params := url.Values{}
	params.Add("id", id)
	params.Add("name", name)
	params.Add("surname", surname)
	params.Add("patronymic", patronymic)
	params.Add("age", age)
	params.Add("sex", sex)
	params.Add("nationality", nationality)
	baseURL.RawQuery = params.Encode()

	resp, _ := http.Get(baseURL.String())
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func operationShowByParameters() {
	clearScreen()

	println("Please write a parameters: ")
	println("id")
	println("name")
	println("surname")
	println("age")
	println("sex")
	println("nationality")
	println("limit")
	println("offset")
	println("sort")
	println("exit to end a writing")
	println()

	sendRequestShowByParameters("/show")
}

func sendRequestShowByParameters(additionalConnect string) {

	baseURL, _ := url.Parse(strConnect + additionalConnect)
	params := url.Values{}

	for {

		var first string
		var second string

		fmt.Print("write a paramater: ")
		fmt.Scanln(&first)
		if first == "exit" {
			break
		}

		fmt.Print("write a value: ")
		fmt.Scanln(&second)
		if second == "exit" {
			break
		}

		params.Add(first, second)
	}

	baseURL.RawQuery = params.Encode()

	resp, _ := http.Get(baseURL.String())
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func operationShowAll() {
	clearScreen()

	sendRequestShowAll("/showall")
}

func sendRequestShowAll(additionalConnect string) {

	baseURL, _ := url.Parse(strConnect + additionalConnect)
	params := url.Values{}
	baseURL.RawQuery = params.Encode()

	resp, _ := http.Get(baseURL.String())
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func operationDelete() {
	clearScreen()

	var id string

	fmt.Print("Please enter person id to delete: ")
	fmt.Scanln(&id)

	sendRequestDelete("/delete", id)
}

func sendRequestDelete(additionalConnect string, id string) {

	baseURL, _ := url.Parse(strConnect + additionalConnect)
	params := url.Values{}
	params.Add("id", id)
	baseURL.RawQuery = params.Encode()

	resp, _ := http.Get(baseURL.String())
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}

func operationInsert() {
	clearScreen()

	var name string
	var surname string
	var patronymic string

	fmt.Print("Please enter person name: ")
	fmt.Scanln(&name)
	fmt.Print("Please enter person surname: ")
	fmt.Scanln(&surname)
	fmt.Print("Please enter person patronymic: ")
	fmt.Scanln(&patronymic)

	sendRequestInsert("/insert", name, surname, patronymic)

}

func sendRequestInsert(additionalConnect string, p_name string, p_surname string, p_patron string) {

	baseURL, _ := url.Parse(strConnect + additionalConnect)
	params := url.Values{}
	params.Add("name", p_name)
	params.Add("surname", p_surname)
	params.Add("patronymic", p_patron)
	baseURL.RawQuery = params.Encode()

	resp, _ := http.Get(baseURL.String())
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
