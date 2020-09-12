package test

import (
	"testing"
)

func testAddUsers(t *testing.T) {

	testcase := []TestCase{
		{
			code:     40001,
			param:    `name=baibai&created_by=admin`,
			errMsg:   `该名称已存在`,
			method:   "POST",
			desc:     "验证插入名称已存在",
			showBody: true,
			url:      "/api/v1/users",
			ext1:     1,
		},
		{
			code:     0,
			param:    `name=5678999999&created_by=admin`,
			errMsg:   `创建成功`,
			method:   "POST",
			desc:     "验证创建成功",
			showBody: true,
			url:      "/api/v1/users",
			ext1:     1,
		},
	}
	call(t, testcase)

}
func TestUserAll(t *testing.T) {
	t.Run("TestAddUsers", testAddUsers)
	//t.Run("TestEditUsers", testEditUsers)
	//t.Run("TestDeleteUsers", testDeleteUsers)
}
