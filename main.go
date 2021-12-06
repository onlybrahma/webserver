package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/del", delHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Starting server at port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func delHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error Parsing the form : %v", err)
		return
	}

	//	n := strings.ToLower(r.FormValue("name"))

	f, err := os.OpenFile("temp.csv", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//	userExists := false
	var m [][]string
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		field := strings.Split(line, ",")

		//		if strings.ToLower(field[0]) == n {
		//	var row []string
		//		row = append(row, field[0], field[1])
		m = append(m, []string{field[0], field[1]})
		//		fmt.Fprintf(w, "Field values %s & %s then array %s\n", field[0], field[1], m)
		//		return

	}
	//	fmt.Fprint(w, m)
	//	if !userExists {
	//		fmt.Fprintf(w, "%s user does not exists... \n", n)
	//	}
	fmt.Fprintf(w, "the array %s and a value %s, %s\n", m, m[0][0], m[0][1])
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error Parsing the form : %v", err)
		return
	}

	n := strings.ToLower(r.FormValue("name"))

	f, err := os.OpenFile("temp.csv", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	userExists := false

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		field := strings.Split(line, ",")
		if strings.ToLower(field[0]) == n {
			fmt.Fprintf(w, "%s user exists and is %v years old!\n", field[0], field[1])
			userExists = true
			return
		}
	}
	if !userExists {
		fmt.Fprintf(w, "%s user does not exists... \n", n)
	}

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Parse form error: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful!\n")
	n := r.FormValue("name")
	a := r.FormValue("age")
	fmt.Fprintf(w, "Name : %s\n", n)
	fmt.Fprintf(w, "Age : %s\n", a)

	fmt.Fprintf(w, "\n*** File Content ***\n")
	fmt.Fprintf(w, "________________________")
	fmt.Fprintf(w, "\nName\tAge\n")

	file, err := os.OpenFile("temp.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	if _, err := file.WriteString(n + "," + a + "\n"); err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile("temp.csv", os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	/*		i := 0
			t := 0

			for scanner.Scan() {
				fmt.Fprint(w, t, "\t", i)
				i++
			}
	*/
	for scanner.Scan() {
		line := scanner.Text()
		field := strings.Split(line, ",")
		textToPrint := field[0] + "\t" + field[1] + "\n"
		fmt.Fprint(w, textToPrint)
	}

}
