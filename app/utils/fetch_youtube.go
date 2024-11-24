package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
)

var baseUrl = "https://music.youtube.com/"

func GetBanner(artistName string) (string, error) {
	path, err := search(url.QueryEscape(artistName))
	if err != nil {
		return "", err
	}

	url, err := getBanner(string(path))
	if err != nil {
		return "", err
	}

	return url, err
}

func search(artistName string) ([]byte, error) {
	request, err := http.NewRequest("GET", baseUrl+"search?q="+artistName, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:132.0) Gecko/20100101 Firefox/132.0")
	request.Header.Add("Accept-Language", "en-US,en;q=0.5")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	pattern := `text\\x22:\\x22([\w ]+)\S+browseId\\x22:\\x22(\S*)\\x22,\\x22browseEndpointContextSupportedConfigs\\x22:\\x7b\\x22browseEndpointContextMusicConfig\\x22:\\x7b\\x22pageType\\x22:\\x22MUSIC_PAGE_TYPE_ARTIST`
	artistPathRgx, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	matches := artistPathRgx.FindSubmatch(body)
	if len(matches) < 3 {
		fmt.Println(matches)
		fmt.Fprintln(os.Stderr, string(body))
		return nil, errors.New("artist not found")
	}

	return matches[2], nil
}

func getBanner(path string) (string, error) {
	request, err := http.NewRequest("GET", baseUrl+"channel/"+path, nil)
	if err != nil {
		return "", err
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:132.0) Gecko/20100101 Firefox/132.0")
	request.Header.Add("Accept-Language", "en-US,en;q=0.5")

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	pattern := `Action menu\\x22\\x7d\\x7d\\x7d\\x7d,\\x22thumbnail\\x22:\\x7b\\x22musicThumbnailRenderer\\x22:\\x7b\\x22thumbnail\\x22:(\\x7b\\x22thumbnails\\x22:\\x5b(?:\\x7b\\x22url\\x22:\\x22https:\\\/\\\/lh3.googleusercontent.com\\\S+\\x3dw\d+-h\d+\S+\\x22,\\x22width\\x22:\d+,\\x22height\\x22:\d+\\x7d,)+\\x7b\\x22url\\x22:\\x22https:\\\/\\\/lh3.googleusercontent.com\\\S*\\x3dw\d+-h\d+\S*\\x22,\\x22width\\x22:\d+,\\x22height\\x22:\d+\\x7d\\x5d\\x7d)`
	artistPathRgx, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	matches := artistPathRgx.FindSubmatch(body)
	rawJson := htoa(string(matches[1]))
	jsonData, err := Decode(rawJson)

	if err != nil {
		return "", err
	}

	var thumbnails []Object
	err = jsonData.Get(&thumbnails, ".thumbnails")
	if err != nil {
		return "", err
	}

	url, ok := thumbnails[len(thumbnails)-1]["url"].(string)
	if !ok {
		return "", errors.New("url isn't a string")
	}

	return url, nil
}

func htoa(txt string) string {
	result := ""
	rText := []rune(txt)
	length := len(rText)

	for i := 0; i < length; i++ {
		char := rText[i]

		if char == '\\' && i+1 < length && rText[i+1] == 'x' && i+2 < length {
			switch rText[i+2] {
			case '7':
				if i+3 < length {
					if rText[i+3] == 'b' {
						result += "{"
					} else if rText[i+3] == 'd' {
						result += "}"
					} else {
						result += string(char)
					}
				} else {
					result += string(char)
				}
				i += 3 // Skip the processed escape sequence
			case '5':
				if i+3 < length {
					if rText[i+3] == 'b' {
						result += "["
					} else if rText[i+3] == 'd' {
						result += "]"
					} else {
						result += string(char)
					}
				} else {
					result += string(char)
				}
				i += 3 // Skip the processed escape sequence
			case '2':
				if i+3 < length && rText[i+3] == '2' {
					result += "\""
				} else {
					result += string(char)
				}
				i += 3 // Skip the processed escape sequence
			case '3':
				if i+3 < length && rText[i+3] == 'd' {
					result += "="
					i += 3 // Skip the processed escape sequence
				} else {
					result += string(char)
				}
			default:
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}
	return result
}
