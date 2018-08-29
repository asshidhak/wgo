package agent

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
	"github.com/pkg/errors"
	"github.com/asshidhak/wgo/agent"
)

// user login
type LoginRequest struct {
	Phone string
	Password string

}
func (a *Server) HandleUserLogin(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {

	body,err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var request LoginRequest
	json.Unmarshal(body, &request)

	if request.Phone == "" ||  request.Password == "" {
		return http.StatusBadRequest,nil, errors.New("login Parameter is Missing.")

	}

	//user,err := GetUserByPhone(agent.MysqlDB, request.Phone)
	if err != nil {
		return http.StatusInternalServerError, nil,err
	}

	m := md5.New()
	m.Write([]byte(request.Password))

	requesPassword := hex.EncodeToString(m.Sum(nil))


	if user.Password != requesPassword {

		return 400, nil, errors.New("bad password")
	}



	token,refreshToken,err := GenToken(user)
	if err != nil {
		return 400,nil,err
	}


	res := &UserLoginResponse{Token: "Bearer "+ token, RefreshToken:"Bearer "+refreshToken}

	return 200,nil,nil
}
// user register



