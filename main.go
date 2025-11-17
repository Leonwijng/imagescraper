package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Login(email, password string) (authToken string, err error) {
	f := url.Values{}
	f.Add("email", email)
	f.Add("password", password)
	f.Add("rememberMe", "on")
	// f.Add("__superform_id", "1r9k94j")

	req, err := http.NewRequest(http.MethodPost, "https://app.meijertheoriecursus.nl/auth/inloggen", strings.NewReader(f.Encode()))
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Sec-CH-UA", `"Chromium";v="142", "Google Chrome";v="142", "Not_A Brand";v="99"`)
	req.Header.Set("Sec-CH-UA-Mobile", "?0")
	req.Header.Set("Sec-CH-UA-Platform", `"Windows"`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("X-SvelteKit-Action", "true")
	req.Header.Set("User-Agent", getUserAgent())
	req.Header.Set("referrer", "https://app.meijertheoriecursus.nl/auth/inloggen")
	req.Header.Set("mode", "cors")
	req.Header.Set("credentials", "include")

	c := http.Client{}

	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed: api returned a non 2xx code")
	}

	var cVal string
	for _, c := range resp.Cookies() {
		if c.Name == "pb_auth" {
			cVal = c.Value
			break
		}
	}

	u, err := url.QueryUnescape(cVal)
	if err != nil {
		return
	}

	var info AuthResponse
	err = json.Unmarshal([]byte(u), &info)
	if err != nil {
		return
	}

	if len(info.Token) <= 0 {
		return "", fmt.Errorf("login failed: missing auth token in cookie")

	}

	return info.Token, nil
}



const (
	ExamPageUrl = "https://app.meijertheoriecursus.nl/app/examens"
	ExamenButtonLink = ".//div[@class='grid gap-2 sm:grid-cols-[auto_120px]']/a"
	ExamStartButton = ".//button[text()='Starten']"
	ExamAsset = ".//img[contains(@class, 'mx-auto object-contain')] | .//video[contains(@class, 'mx-auto')]"
	ExamQuestion = ".//h5[@class='font-heading font-bold text-lg md:text-xl']"
	ExamAnswerOptions = ".//div[contains(@class,'grid gap-2')]//span[@class='relative top-[1px]'] | .//div[contains(@class,'grid gap-2')]//span[contains(@class, 'ml-2')]"
	ExamNextQuestion = ".//div[span[text()='Volgende vraag']]/button"
)

const (
	VideoCoursePageUrl = "https://app.meijertheoriecursus.nl/app/videocursus"
	VideoCourseButton = ".//div/a[span]"
)

func main() {
	godotenv.Load(".env")
	authToken, err := Login(os.Getenv("MEIJER_ACCOUNT_EMAIL"), os.Getenv("MEIJER_ACCOUNT_PASSWORD"))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(authToken)

}
