package core

import (
	"os"
	"regexp"
	"strconv"
)

func ApplyChanges(str string) string {
	// Création de modèles d'expressions régulières à retrouver dans le texte.
	rE := regexp.MustCompile(`(\((?:hex|bin)\)|\((?:cap|up|low)(?:,\s?\d+)?\))|[.|,|?|!|:|;|']{1,}|([\w\d\(\)']{1,})`)
	match := rE.FindAllString(str, -1)
	res := ""
	for idx, item := range match {
		switch {
		case item == "(bin)":
			match[idx-1] = strconv.Itoa(AtoiBase(match[idx-1], "01"))
		case item == "(hex)":
			match[idx-1] = strconv.Itoa(AtoiBase(match[idx-1], "0123456789ABCDEF"))
		case item == "(cap)":
			match[idx-1] = Capitalize(match[idx-1])
		case item == "(up)":
			match[idx-1] = ToUpper(match[idx-1])
		case item == "(low)":
			match[idx-1] = ToLower(match[idx-1])
		case isNTag(item):
			n, tag := nTag(item)
			cpt := 0
			for _, i := range match[idx-n : idx] {
				if tag == "up" {
					match[idx-n+cpt] = ToUpper(i)
				}
				if tag == "cap" {
					match[idx-n+cpt] = Capitalize(i)
				}
				if tag == "low" {
					match[idx-n+cpt] = ToLower(i)
				}
				cpt++
			}
		}
	}
	open := false
	for i, word := range match {
		if i == 0 {
			res += word
		} else if word != "(low)" && word != "(up)" && word != "(cap)" && word != "(hex)" && word != "(bin)" && !isNTag(word) {
			if word == "'" && !open {
				res += " " + word
				open = true
			} else if word == "'" && open {
				res += word + " "
				open = false

			} else if word[0] == '.' || word[0] == ',' || word[0] == ';' || word[0] == '?' || word[0] == '!' || word[0] == ':' {
				res += word
			} else if word == "a" || word == "A" {
				if isvowel(rune(match[i+1][0])) {
					if word == "a" {
						res += " an"
					} else if word == "A" {
						res += " An"
					}
				} else {
					res += " " + word
				}
			} else {
				if match[i-1] == "'" {
					res += word
				} else {
					res += " " + word
				}
			}
		}
	}
	return res
}

func isvowel(r rune) bool {
	if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' || r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U' {
		return true
	}
	return false
}

func isNTag(s string) bool {
	if len(s) > 5 {
		if s[0:5] == "(up, " {
			return true
		}
		if s[0:6] == "(cap, " || s[0:6] == "(low, " {
			return true
		}
	}
	return false
}

func nTag(s string) (int, string) {
	if s[0:5] == "(up, " {
		i, _ := strconv.Atoi(s[5 : len(s)-1])
		return i, "up"
	}
	if s[0:6] == "(cap, " {
		i, _ := strconv.Atoi(s[6 : len(s)-1])
		return i, "cap"
	}
	if s[0:6] == "(low, " {
		i, _ := strconv.Atoi(s[6 : len(s)-1])
		return i, "low"
	}
	return 0, ""
}

func GetFileContent(path string) []byte {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return content
}

func CreateFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return file
}

func ToUpper(s string) (res string) {
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			res += string(char - 32)
		} else {
			res += string(char)
		}
	}
	return res
}

func ToLower(s string) (res string) {
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			res += string(char + 32)
		} else {
			res += string(char)
		}
	}
	return res
}

func Capitalize(s string) string {
	tmp := ""
	res := ""
	for _, val := range s {
		if !(val >= 'A' && val <= 'Z') && !(val >= 'a' && val <= 'z') && !(val >= '0' && val <= '9') {
			res += maj(tmp)
			res += string(val)
			tmp = ""
		} else {
			tmp += string(val)
		}
	}
	res += maj(tmp)
	return res
}

func maj(s string) string {
	res := ""
	tab := []string{}
	for idx, char := range s {
		if idx == 0 && char >= 'a' && char <= 'z' {
			tab = append(tab, string(char-32))
		} else if idx == 0 && char >= 'A' && char <= 'Z' {
			tab = append(tab, string(char))
		} else {
			if char >= 'A' && char <= 'Z' {
				tab = append(tab, string(char+32))
			} else {
				tab = append(tab, string(char))
			}
		}
	}
	for _, val := range tab {
		res += string(val)
	}
	return res
}

func AtoiBase(s string, base string) int {
	res := 0
	cpt := len(s) - 1
	idx := 0
	for _, char := range s {
		for ind, c := range base {
			if char == c {
				idx = ind
			}
		}
		res += idx * RecursivePower(len(base), cpt)
		cpt--
	}
	return res
}

func RecursivePower(nb int, power int) (res int) {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}
