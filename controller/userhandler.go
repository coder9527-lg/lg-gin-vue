package controller

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shuwenhe/shuwen-shop/model"
	"github.com/shuwenhe/shuwen-shop/service"
	"github.com/shuwenhe/shuwen-shop/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	flag, _ := service.IsLogin(r) // Determine whether you have logged in
	if flag {
		GetPageBooksByPrice(w, r) // Go to homepage
	} else {
		phone := r.PostFormValue("phone")       // Get username
		password := r.PostFormValue("password") // Get password
		user, _ := service.CheckUserNameAndPassword(phone, string(password))
		if user.ID > 0 { // Username and password are correct
			uuid := utils.CreateUUID() // Generate UUID
			session := &model.Session{ // Create a session
				SessionID: uuid,
				Phone:     user.Phone,
				UserID:    user.ID,
			}
			service.AddSession(session) // Write the session to the database, there is an identification in the database
			cookie := http.Cookie{      // Create a cookie associated with the session
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie) // Send cookie to browser
			t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
			t.Execute(w, user)
		} else {
			t := template.Must(template.ParseFiles("views/pages/user/login.html"))
			t.Execute(w, "Incorrect username or password")
		}
	}
}

// Logout Logout
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value // Get the value of the cookie
		service.DeleteSession(cookieValue)
		cookie.MaxAge = -1        // Set cookie invalidation
		http.SetCookie(w, cookie) // Send the modified cookie to the browser
	}
	GetPageBooksByPrice(w, r) // Go to homepage
}

// Regist Regist user
func Regist(w http.ResponseWriter, r *http.Request) {
	phone := r.PostFormValue("phone")
	password := r.PostFormValue("password")
	user, _ := service.CheckUserName(phone)
	if user.ID > 0 {
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "Username already exists")
	} else {
		service.SaveUser(phone, password)
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "Incorrect username or password")
	}
}

// CheckUserName Verify that the username exists by sending an Ajax request
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	phone := r.PostFormValue("phone")
	user, _ := service.CheckUserName(phone)
	if user.ID > 0 {
		w.Write([]byte("用户名已经存在！"))
	} else {
		w.Write([]byte("<font style = 'color:#04be02'>用户名不存在！</font>"))
	}
}

func Register(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	fmt.Println("name =***=", name)
	fmt.Println("password =***=", password)
	fmt.Println("phone=***=", phone)
	c.JSON(200, gin.H{
		"msg": "Register success",
	})

	if len(phone) != 11 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "The phone num must be 11 digits!",
		})
	}

	if len(password) < 6 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code": 422,
			"msg":  "password num cannt be less than 6 digits",
		})
	}
	if len(name) == 0 {
		name = RandomString(10)
		return
	}

	log.Println(name, password, phone)

}

func RandomString(n int) string {
	var letters = []byte("*dueihduieuidhuiehuideui")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for k, _ := range result {
		result[k] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
