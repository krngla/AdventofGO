package AdventofGO/meta

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"fmt"
)

func HTTPwithCookies(url, sessionid string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.AddCookie(&http.Cookie{Name:	"session", Value: sessionid})

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(url + "\nresp.StatusCode: " + strconv.Itoa(resp.StatusCode))
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)

}

func main()	{
	url := "https://adventofcode.com/2022/day/4/input"
	session := ""

	b, err := HTTPwithCookies(url, session)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
