package pkg
import (
	"fmt"
	"github.com/golang/protobuf/proto"
)
import proo "github.com/velann21/todo_list_activity_manager/pkg/proto"
import "github.com/velann21/todo_list_activity_manager/pkg/helpers"

func UserReq(){
	user := &proo.User{}
	user.FirstName = "Velan"
	user.LastName = "Nadhakumar"
	user.Age = 19


	data, err := proto.Marshal(user)
	if err != nil{

	}

	resp, err := helpers.HttpRequest("POST", "http://localhost:7000/api/v1/users/printuser", data)
	if err != nil{

	}
	fmt.Println(resp)
}




