package test

import (
	"fmt"
	"testing"
)
type I interface{
   F()
}

type T struct {
}

func (t *T) F() {
}

func makeI() I {
   var r *T
   if r == nil {
       fmt.Println("I am nil at makeI")
   }
   return r
}


func TestAddUsers(t *testing.T) {
	i := makeI()
	
	if i != nil {
	   fmt.Println("I am not nil at main")
	}


}
func TestUserAll(t *testing.T) {
	
	//t.Run("TestEditUsers", testEditUsers)
	//t.Run("TestDeleteUsers", testDeleteUsers)
}
