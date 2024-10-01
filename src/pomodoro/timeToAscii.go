package pomodoro

import (
	"strconv"
	"strings"
)

var ASCIINumbers map[string]string = map[string]string{
	"0": `
  ___   
 / _ \  
| | | | 
| | | | 
| |_| | 
 \___/  
`,
	"1": `
 __ 
/_ |
 | |
 | |
 | |
 |_|
`,
	"2": `
 ___  
|__ \ 
   ) |
  / / 
 / /_ 
|____|
`,
	"3": `
 ____  
|___ \ 
  __) |
 |__ < 
 ___) |
|____/ 
`,
	"4": `
 _  _   
| || |  
| || |_ 
|__   _|
   | |  
   |_|  
`,
	"5": `
 _____ 
| ____|
| |__  
|___ \ 
 ___) |
|____/ 
`,
	"6": `
   __  
  / /  
 / /_  
| '_ \ 
| (_) |
 \___/ 
`,
	"7": `
 ______ 
|____  |
    / / 
   / /  
  / /   
 /_/    
`,
	"8": `
  ___  
 / _ \ 
| (_) |
 > _ < 
| (_) |
 \___/ 
`,
	"9": `
  ___  
 / _ \ 
| (_) |
 \__, |
   / / 
  /_/  
	`,
	"colon": `    
 _ 
(_)
   
 _ 
(_)
   
	`,
	"space": `
 
 
 
 
 
 
 `,
}

type ASCIIArtChar struct {
	Char      string
	ASCIIChar string
}

func (char *ASCIIArtChar) GetNthLine(nthLine int) string {
	return strings.Split(char.ASCIIChar, "\n")[nthLine]
}

func (char *ASCIIArtChar) SetASCIIArt() {
	if char.Char == " " {
		char.Char = "space"
	} else if char.Char == ":" {
		char.Char = "colon"
	}
	char.ASCIIChar = ASCIINumbers[char.Char]
}

func StringTimeToASCII(timeString string) string {
	var ASCIIArtNumbers []ASCIIArtChar
	var timeStringList []string
	var char ASCIIArtChar
	var fullASCIIChar string
	var space ASCIIArtChar = ASCIIArtChar{
		Char: " ",
	}
	space.SetASCIIArt()

	timeStringList = strings.Split(timeString, "")
	for _, str := range timeStringList {
		char = ASCIIArtChar{
			Char: str,
		}
		char.SetASCIIArt()
		ASCIIArtNumbers = append(ASCIIArtNumbers, char)
	}

	for j := 0; j < 6; j++ {
		for i := range ASCIIArtNumbers {
			fullASCIIChar = fullASCIIChar + ASCIIArtNumbers[i].GetNthLine(j+1)
		}
		fullASCIIChar = fullASCIIChar + "\n"
	}

	return fullASCIIChar

}

func FormatTimeAmmounts(timeAmmount uint64) string {
	var strTime string = strconv.FormatUint(timeAmmount, 10)
	if len(strTime) == 1 {
		strTime = "0" + strTime
	}
	return strTime
}

func SecToTimeFormat(sec uint64) string {
	var minutes uint64 = 0
	var hours uint64 = 0
	var timeString string

	minutes = sec / 60
	if minutes >= 1 {
		sec = sec % 60
	} else {
		return FormatTimeAmmounts(sec)
	}
	hours = minutes / 60
	if hours >= 1 {
		minutes = minutes % 60
		timeString = FormatTimeAmmounts(hours) + ":" +
			FormatTimeAmmounts(minutes) +
			":" + FormatTimeAmmounts(sec)
	} else {
		timeString = FormatTimeAmmounts(minutes) +
			":" + FormatTimeAmmounts(sec)
		return timeString
	}

	return timeString
}
