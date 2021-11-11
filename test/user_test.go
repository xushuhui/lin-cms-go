package test

import (
	"encoding/json"
	"fmt"
	"github.com/xushuhui/goal/utils"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAddUsers(t *testing.T) {
	m := make(map[int64]AlbumResponse)
	for i := 1; i <= 501; i = i + 10 {
		url := `http://localhost:4455/api/albums?offset=` + utils.IntToString(i) + `&limit=10&app_id=8456804462776431`
		fmt.Println(url)
		items := do(url)
		fmt.Println(items)
		for _, v := range items {

			_, ok := m[v.Id]

			if ok {
				fmt.Println("------------------------------------------------------------------------------------------")
				fmt.Println("repeat ", i, m[v.Id])
			} else {
				m[v.Id] = v
			}
		}

	}

}

type AlbumResponse struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	IssueDate string `json:"issue_date"`
	IssueTime string `json:"issue_time"`
}
type Response struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    PageResponse `json:"data"`
}
type PageResponse struct {
	Items []AlbumResponse `json:"items"`
	Total int64           `json:"total"`
}

func do(url string) []AlbumResponse {
	rsps, err := http.Get(url)
	if err != nil {
		fmt.Println("Request failed:", err)
		return nil
	}
	defer rsps.Body.Close()

	body, err := ioutil.ReadAll(rsps.Body)
	if err != nil {
		fmt.Println("Read body failed:", err)
		return nil
	}
	var resp Response
	json.Unmarshal(body, &resp)
	return resp.Data.Items

}
func TestUserAll(t *testing.T) {

}
