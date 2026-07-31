package main

import (
"fmt"
"net/http"
"encoding/json"
"os"
"strings"
)

type Eng struct {
	En rune `json:"en"`
	Rd string `json:"rd"`
}
type ret struct {
	Text string `json:"text"`
}
 var codeMethod = []Eng{
 	{'a', "$#"},{'b', "*#"},{'c', "5&"},{'d', "¢€"}, {'e', "@#"},{'f', "%="},
 	{'g', "5_"}, {'h', "&-"},{'i', "+#"},{'j', "×#"}, {'k', "@+"},{'l', "$$"},
 	{'m', "~|~"}, {'n', "?!"},{'o', "&$"},{'p', "`^"}, {'q', "¥√"},{'r', "°•"},
 	{'s', "£¢"}, {'t', "$&"},{'u', "!?"},{'v', "^`"}, {'w', "√¥"},{'x', "=%"},
 	{'y', "#2"}, {'z', "$2."},{'A', "5$"},{'B', "$5"}, {'C', "_?_"},{'D', "~?~"},
 	{'E', "||"}, {'F', "®©"},{'G', "©®"},{'H', "[{"}, {'I', "}]"},{'J', "()"},
 	{'K', "@+"}, {'L', "@-"},{'M', "@÷"},{'N', "@×"}, {'O', "×@"},{'P', "÷@"},
 	{'Q', ":;"}, {'R', ";:"},{'S', "<>"},{'T', "><"}, {'U', "v^"},{'V', "^v"},
 	{'W', "19"}, {'X', "70"},{'Y', "64"},{'Z', "58"}, {'1', "2"},{'2', "02"},
 	{'3', "aZ?"}, {'4', "_Z_"},{'5', "-z-"},{'6', "@12"}, {'7', "™✓"},{'8', "™®"},
 	{'9', "`|`"}, {'0', "&2&"},{'$', "2&2"},{'&', "@1$"}, {'(', "++"},{')', "--"},
 	{'@', "#@$3"}, {';', "bs64"},{':', "or$"},{'.', "&@#"}, {' ', "**_"},{',', "***"},
 		 }

 func toEnglish(write http.ResponseWriter, read *http.Request) {
 	write.Header().Set("Access-Control-Allow-Origin", "*")
write.Header().Set("Access-Control-Allow-Methods", "POST")
write.Header().Set("Access-Control-Allow-Headers", "Content-Type")
 	var rets ret
 	var cipheredWord string
 	if read.Method != "POST" {
 		http.Error(write,"Invalid Method! Only POST methods allowed", http.StatusMethodNotAllowed)
 		return
 	}
 	err := json.NewDecoder(read.Body).Decode(&rets)
 	if err != nil {
 		http.Error(write, "Invalid JSON", http.StatusBadRequest)
 		return
 	}
 	for _, ch := range rets.Text {
 		for it := 0; it < len(codeMethod); it++ {
 			if ch == codeMethod[it].En {
 				cipheredWord += codeMethod[it].Rd +" "
 			} 
 		}
 	}
 	response := []ret {{cipheredWord}}
 	json.NewEncoder(write).Encode(response)
 }

func toRucdaryDecoded(write http.ResponseWriter, read *http.Request) {
	write.Header().Set("Access-Control-Allow-Origin", "*")
write.Header().Set("Access-Control-Allow-Methods", "POST")
write.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	
	if read.Method != "POST" {
		http.Error(write, "Only POST methods are allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var vex []ret
	var txt string
err := json.NewDecoder(read.Body).Decode(&vex)
	if err != nil {
		http.Error(write, "Invalid JSON request", http.StatusBadRequest)
		return
	}
	oth := strings.Split(vex[0].Text, " ")
	for i := 0; i < len(oth); i++ {
		for j := 0; j < len(codeMethod); j++ {
			if oth[i] == codeMethod[j].Rd {
				txt += string(codeMethod[j].En)
			}
		}
	}
	stf := []ret {{txt}}
	json.NewEncoder(write).Encode(stf)
	}

func main() {
	http.HandleFunc("/en", toEnglish)
	http.HandleFunc("/rd", toRucdaryDecoded)
	port := os.Getenv("PORT");
	if port == "" {
		port = "8082"
	}
	fmt.Println("Server running at port 8082")
	http.ListenAndServe(":"+port, nil)
}
