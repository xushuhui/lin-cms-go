package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

type Categories struct {
	ID       int
	Pid      int
	Name     string
	Children []Categories
}

func getNode(list []Categories, pid int) (l []Categories) {
	for _, val := range list {
		if val.Pid == pid {
			if pid == 0 {
				// 顶层
				l = append(l, val)
			} else {
				var children []Categories
				child := val
				children = append(children, child)
			}
		}
	}
	return
}

func Tree(list []Categories, pid int) (newList []Categories) {
	for _, v := range list {
		if v.Pid == pid {
			child := getNode(list, v.ID)
			node := Categories{
				ID:   v.ID,
				Name: v.Name,
				Pid:  v.Pid,
			}
			node.Children = child
			newList = append(newList, node)
		}
	}
	return
}

func TestAddUsers(t *testing.T) {
	categories := []Categories{
		{ID: 1, Pid: 0, Name: "电脑"},
		{ID: 2, Pid: 0, Name: "手机"},
		{ID: 3, Pid: 1, Name: "笔记本"},
		{ID: 4, Pid: 1, Name: "台式机"},
		{ID: 5, Pid: 2, Name: "智能机"},
		{ID: 6, Pid: 2, Name: "功能机"},
		{ID: 7, Pid: 3, Name: "超级本"},
		{ID: 8, Pid: 3, Name: "游戏本"},
	}
	m := Tree(categories, 0)

	fmt.Println(m)
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
