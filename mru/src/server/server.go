package server

import (
	//	"command"
	"context"
	"controller"
	"encoding/json"
	"github.com/gorilla/websocket"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mcase"
	"net/http"
	"result"
	"rut"
	"script"
	"sync"
	"task"
	"time"
	"util"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	filePeriod = 10 * time.Second
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

var Engine *controller.Controller

func Product(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]
	log.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.New("product.html").Delims("|||", "|||").ParseFiles("asset/web/template/product.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println("Cannot parse form: ", err.Error())
			return
		}

		log.Println(r.Form)

		device := &http.Cookie{
			Name:  "Device",
			Value: r.FormValue("device"),
			Path:  "runscript",
		}

		username := &http.Cookie{
			Name:  "Username",
			Value: r.FormValue("username"),
			Path:  "runscript",
		}

		password := &http.Cookie{
			Name:  "Password",
			Value: r.FormValue("password"),
			Path:  "runscript",
		}

		ip := &http.Cookie{
			Name:  "ip",
			Value: r.FormValue("ip"),
			Path:  "runscript",
		}

		http.SetCookie(w, device)
		http.SetCookie(w, username)
		http.SetCookie(w, password)
		http.SetCookie(w, ip)

		dev, err := rut.New(&rut.RUT{
			Device:   r.FormValue("device"),
			IP:       r.FormValue("ip"),
			Port:     "23",
			Username: r.FormValue("username"),
			Password: r.FormValue("passowrd"),
		})
		dev.Init()
		if err != nil {
			io.WriteString(w, err.Error())
		} else {
			sess.RUT = dev
			t, err := template.New("script.html").Delims("|||", "|||").ParseFiles("asset/web/template/script.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html")
			if err != nil {
				log.Println(err)
				io.WriteString(w, err.Error())
				return
			}

			err = t.Execute(w, nil)
			if err != nil {
				log.Println(err.Error())
			}
		}
	} else {
		http.Redirect(w, r, "/invalid", http.StatusTemporaryRedirect)
	}
}

func VUETree(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("vuetree.html").Delims("|||", "|||").ParseFiles("asset/web/template/vuetree.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		values := r.Form
		for k, v := range values {
			log.Println(k, v[0])
			io.WriteString(w, k+":"+v[0])
		}
	}
}

func VUEtest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("vuetest.html").Delims("|||", "|||").ParseFiles("asset/web/template/vuetest.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		values := r.Form
		for k, v := range values {
			log.Println(k, v[0])
			io.WriteString(w, k+":"+v[0])
		}
	}
}

func JSTree(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("jstree.html").Delims("|||", "|||").ParseFiles("asset/web/template/jstree.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		values := r.Form
		for k, v := range values {
			log.Println(k, v[0])
			io.WriteString(w, k+":"+v[0])
		}
	}
}

func JSONTree(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("jsontree.html").Delims("|||", "|||").ParseFiles("asset/web/template/jsontree.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		values := r.Form
		for k, v := range values {
			log.Println(k, v[0])
			io.WriteString(w, k+":"+v[0])
		}
	}
}

func LaunchTree(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("launchtree.html").Delims("|||", "|||").ParseFiles("asset/web/template/launchtree.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		values := r.Form
		for k, v := range values {
			log.Println(k, v[0])
			io.WriteString(w, k+":"+v[0])
		}
	}
}

func Script(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("Device")
	if err != nil {
		io.WriteString(w, "You must connect to a device first: ")
		return
	}

	if r.Method == "GET" {
		t, err := template.New("script.html").Delims("|||", "|||").ParseFiles("asset/web/template/script.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		values := r.Form
		for k, v := range values {
			log.Println(k, v[0])
			io.WriteString(w, k+":"+v[0])
		}
	}
}

func JSNotify(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("jsnotify.html").Delims("|||", "|||").ParseFiles("asset/web/template/jsnotify.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func TreeView(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	log.Printf("%q", sess.NewCache)
	if r.Method == "GET" {
		log.Println("Dump tree")
		encoder := json.NewEncoder(w)
		err := encoder.Encode(sess.NewCache.TreeView().Children)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func NewCase(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "GET" {
		t, err := template.New("newcase.html").Delims("|||", "|||").ParseFiles("asset/web/template/newcase.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		js, _ := json.Marshal(sess.NewCache.TreeView().Children)
		//log.Println(string(js))
		err = t.Execute(w, string(js))
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		log.Printf("%#q\n", r.Form)
		var newcase mcase.Case
		for k, v := range r.Form {
			log.Println(k, ":", v)
		}

		err := json.Unmarshal([]byte(r.FormValue("newcase")), &newcase)
		if err != nil {
			log.Println(newcase, err)
			io.WriteString(w, err.Error())
			return
		}
		sess.NewCache.AddCase(&newcase)
	}
}

func NewTask(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "GET" {
		cookies := r.Cookies()
		for _, cookie := range cookies {
			log.Println(cookie)
		}

		t, err := template.New("newtask.html").Delims("|||", "|||").ParseFiles("asset/web/template/newtask.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html", "asset/web/template/caseheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		js, _ := json.Marshal(sess.NewCache.TreeView().Children)
		err = t.Execute(w, string(js))
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		caseid, err := r.Cookie("CASEID")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		r.ParseForm()
		log.Println(r.Form)
		for k, v := range r.Form {
			log.Println(k, v)
		}
		var newtask task.Task
		err = json.Unmarshal([]byte(r.FormValue("newtask")), &newtask)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		err = sess.NewCache.AddTask(caseid.Value, &newtask)
		if err != nil {
			log.Println(err.Error())
			io.WriteString(w, err.Error())
		}
	}
}

func DumpSubGroup(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "GET" {
		r.ParseForm()
		cookies := r.Cookies()
		for _, cookie := range cookies {
			log.Println(cookie)
		}
		encoder := json.NewEncoder(w)
		c, err := sess.NewCache.GetSubGroupByID(r.FormValue("id"))
		if err != nil {
			io.WriteString(w, err.Error())
		}
		err = encoder.Encode(c)
		if err != nil {
			log.Println(err.Error())
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}
func DumpFeature(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "GET" {
		r.ParseForm()
		cookies := r.Cookies()
		for _, cookie := range cookies {
			log.Println(cookie)
		}
		encoder := json.NewEncoder(w)
		c, err := sess.NewCache.GetFeatureByID(r.FormValue("id"))
		if err != nil {
			io.WriteString(w, err.Error())
		}
		err = encoder.Encode(c)
		if err != nil {
			log.Println(err.Error())
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func DumpGroup(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "GET" {
		r.ParseForm()
		cookies := r.Cookies()
		for _, cookie := range cookies {
			log.Println(cookie)
		}

		log.Println("DumpGroup: ", r.FormValue("id"))
		encoder := json.NewEncoder(w)
		c, err := sess.NewCache.GetGroupByID(r.FormValue("id"))
		if err != nil {
			io.WriteString(w, err.Error())
		}
		err = encoder.Encode(c)
		if err != nil {
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func DumpCase(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "GET" {
		r.ParseForm()
		cookies := r.Cookies()
		for _, cookie := range cookies {
			log.Println(cookie)
		}
		encoder := json.NewEncoder(w)
		c, err := sess.NewCache.GetCaseByID(r.FormValue("id"))
		if err != nil {
			io.WriteString(w, err.Error())
		}
		err = encoder.Encode(c)
		if err != nil {
			log.Println(err.Error())
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func DumpTask(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "GET" {
		r.ParseForm()
		cookies := r.Cookies()
		for _, cookie := range cookies {
			log.Println(cookie)
		}

		caseid, err := r.Cookie("CASEID")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		if r.FormValue("id") == "" {
			io.WriteString(w, "Task ID must be set")
			return
		}
		encoder := json.NewEncoder(w)
		c, err := sess.NewCache.GetTaskByID(caseid.Value, r.FormValue("id"))
		if err != nil {
			io.WriteString(w, err.Error())
		}
		err = encoder.Encode(c)
		if err != nil {
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func TaskInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		log.Println(r.FormValue("id"))
		t, err := template.New("taskinfo.html").Delims("|||", "|||").ParseFiles("asset/web/template/taskinfo.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html", "asset/web/template/caseheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		if r.FormValue("id") == "" {
			io.WriteString(w, "Case ID is not set!")
			return
		}

		cookie := &http.Cookie{
			Name:  "TASKID",
			Value: r.FormValue("id"),
		}

		http.SetCookie(w, cookie)
		err = t.Execute(w, nil)
		if err != nil {
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func CaseInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		log.Println(r.FormValue("id"))
		t, err := template.New("caseinfo.html").Delims("|||", "|||").ParseFiles("asset/web/template/caseinfo.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html", "asset/web/template/caseheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		if r.FormValue("id") == "" {
			io.WriteString(w, "Case ID is not set!")
			return
		}

		featureid := &http.Cookie{
			Name:  "CASEID",
			Value: r.FormValue("id"),
			Path:  "/",
		}

		id := &http.Cookie{
			Name:  "UNIQUE",
			Value: r.FormValue("id"),
			Path:  "/",
		}

		http.SetCookie(w, featureid)
		http.SetCookie(w, id)
		err = t.Execute(w, nil)
		if err != nil {
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func GroupInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		log.Println(r.FormValue("id"))
		t, err := template.New("groupinfo.html").Delims("|||", "|||").ParseFiles("asset/web/template/groupinfo.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html", "asset/web/template/groupheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		if r.FormValue("id") == "" {
			io.WriteString(w, "GROUP ID is not set!")
			return
		}

		cookie := &http.Cookie{
			Name:  "GROUPID",
			Value: r.FormValue("id"),
			Path:  "dumpgroup",
		}

		http.SetCookie(w, cookie)
		err = t.Execute(w, nil)
		if err != nil {
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func SubGroupInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		log.Println(r.FormValue("id"))
		t, err := template.New("subgroupinfo.html").Delims("|||", "|||").ParseFiles("asset/web/template/subgroupinfo.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html", "asset/web/template/subgroupheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		if r.FormValue("id") == "" {
			io.WriteString(w, "SUBGROUP ID is not set!")
			return
		}

		cookie := &http.Cookie{
			Name:  "SGID",
			Value: r.FormValue("id"),
		}

		http.SetCookie(w, cookie)
		err = t.Execute(w, nil)
		if err != nil {
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func FeatureInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		log.Println(r.FormValue("id"))
		t, err := template.New("featureinfo.html").Delims("|||", "|||").ParseFiles("asset/web/template/featureinfo.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html", "asset/web/template/featureheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		if r.FormValue("id") == "" {
			io.WriteString(w, "Feature ID is not set!")
			return
		}

		featureid := &http.Cookie{
			Name:  "FEATUREID",
			Value: r.FormValue("id"),
			Path:  "featureinfo",
		}

		id := &http.Cookie{
			Name:  "UNIQUUNIQUEE",
			Value: r.FormValue("id"),
			Path:  "featureinfo",
		}

		http.SetCookie(w, featureid)
		http.SetCookie(w, id)
		err = t.Execute(w, nil)
		if err != nil {
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func DeleteNode(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "POST" {
		r.ParseForm()

		if _, ok := sess.CaseResult[sessionid]; ok {
			io.WriteString(w, "You are runing another cases")
			return
		}

		if r.FormValue("type") == "GROUP" {
			sess.NewCache.DelGroupByID(r.FormValue("id"))
		} else if r.FormValue("type") == "SUBGROUP" {
			sess.NewCache.DelSubGroupByID(r.FormValue("id"))
		} else if r.FormValue("type") == "FEATURE" {
			sess.NewCache.DelFeatureByID(r.FormValue("id"))
		} else if r.FormValue("type") == "CASE" {
			sess.NewCache.DelCaseByID(r.FormValue("id"))
		} else if r.FormValue("type") == "TASK" {
			caseid, err := r.Cookie("CASEID")
			if err != nil {
				io.WriteString(w, "Case ID is not set when delete task")
				return
			}
			sess.NewCache.DelTaskByID(caseid.Value, r.FormValue("id"))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Invalid request") //A proper status code in more usefull.
		}
	}
}

func GetDUTCountByID(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "GET" {
		r.ParseForm()
		log.Println(r.FormValue("id"))

		if r.FormValue("id") == "" {
			io.WriteString(w, "ID is not set!")
			return
		}

		count, err := sess.NewCache.GetDUTCountByID(r.FormValue("id"))
		if err != nil {
			log.Println(err.Error(), "+++++++")
			w.WriteHeader(http.StatusForbidden)
		}

		encoder := json.NewEncoder(w)
		err = encoder.Encode(struct{ DUTCount int }{DUTCount: count})
		if err != nil {
			io.WriteString(w, "adfasdkfjaskdfjaskdfjaksdjfkasjdffuckyou")
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func CheckIsReadyForRunByID(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "GET" {
		r.ParseForm()
		log.Println(r.FormValue("id"))

		if r.FormValue("id") == "" {
			io.WriteString(w, "ID is not set!")
			return
		}

		type Res struct {
			IsError bool
			Ready   bool
			Message string
		}

		var res Res
		ready, err := sess.NewCache.CheckIsReadyForRunByID(r.FormValue("id"))
		if err != nil {
			res.IsError = true
			res.Message = err.Error()
		}

		res.Ready = ready
		encoder := json.NewEncoder(w)
		err = encoder.Encode(res)
		if err != nil {
			io.WriteString(w, "adfasdkfjaskdfjaskdfjaksdjfkasjdffuckyou")
		}
	} else if r.Method == "POST" {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func SetDUTsByID(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "GET" {
		r.ParseForm()
		log.Println(r.FormValue("id"))
		t, err := template.New("setduts.html").Delims("|||", "|||").ParseFiles("asset/web/template/setduts.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html", "asset/web/template/caseheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		log.Println(r.FormValue("id"))

		if r.FormValue("id") == "" {
			log.Println("ID is not set")
			io.WriteString(w, "ID is not set!")
			return
		}

		if r.FormValue("duts") == "" {
			io.WriteString(w, "Invalid Input")
			return
		}

		var con []*rut.Config
		err := json.Unmarshal([]byte(r.FormValue("duts")), &con)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		log.Printf("%#v", con)
		log.Printf("%#q", con)

		var duts []*rut.RUT
		for _, cn := range con {
			d, err := rut.GetRUTByConfig(cn)
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
			duts = append(duts, d)
		}

		err = sess.NewCache.SetDUTsByID(r.FormValue("id"), duts)
		if err != nil {
			io.WriteString(w, err.Error())
		}
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
}

func RunCases(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	if r.Method == "POST" {
		r.ParseForm()

		if _, ok := sess.CaseResult[sessionid]; ok {
			io.WriteString(w, "You are runing another cases")
			return
		}

		sess.CaseResult[sessionid] = make(chan *result.Result, 10)

		log.Println(sess.CaseResult[sessionid])
		if r.FormValue("type") == "GROUP" {
			go func() {
				wg := sync.WaitGroup{}
				for res := range sess.NewCache.RunCasesByGroupID(r.FormValue("id")) {
					wg.Add(1)
					log.Printf("%#v", res)
					go func(r *result.Result) {
						sess.CaseResult[sessionid] <- res
						wg.Done()
					}(res)
				}
				wg.Wait()
				close(sess.CaseResult[sessionid])

			}()
		} else if r.FormValue("type") == "SUBGROUP" {
			go func() {
				wg := sync.WaitGroup{}
				for res := range sess.NewCache.RunCasesBySubGroupID(r.FormValue("id")) {
					wg.Add(1)
					go func(r *result.Result) {
						sess.CaseResult[sessionid] <- r
						wg.Done()
					}(res)
				}
				wg.Wait()
				close(sess.CaseResult[sessionid])
			}()
		} else if r.FormValue("type") == "FEATURE" {
			go func() {
				wg := sync.WaitGroup{}
				for res := range sess.NewCache.RunCasesByFeatureID(r.FormValue("id")) {
					wg.Add(1)
					go func(r *result.Result) {
						sess.CaseResult[sessionid] <- r
						wg.Done()
					}(res)
				}
				wg.Wait()
				close(sess.CaseResult[sessionid])

			}()
		} else if r.FormValue("type") == "CASE" {
			go func() {
				wg := sync.WaitGroup{}
				for res := range sess.NewCache.RunCaseByID(r.FormValue("id")) {
					wg.Add(1)
					go func(r *result.Result) {
						sess.CaseResult[sessionid] <- r
						wg.Done()
					}(res)
				}
				wg.Wait()
				close(sess.CaseResult[sessionid])
			}()
		} else if r.FormValue("type") == "TASK" {
			caseid, err := r.Cookie("CASEID")
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
			go func() {
				wg := sync.WaitGroup{}
				for res := range sess.NewCache.RunTaskByID(caseid.Value, r.FormValue("id")) {
					wg.Add(1)
					go func(r *result.Result) {
						sess.CaseResult[sessionid] <- r
						wg.Done()
					}(res)
				}
				wg.Wait()
				close(sess.CaseResult[sessionid])
			}()
		} else {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Invalid request") //A proper status code in more usefull.
		}
	}
	io.WriteString(w, r.Host)
}

func NewRunScript(w http.ResponseWriter, r *http.Request) {
	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	sess, _ := Engine.Sessions[sessionid]

	_, err := r.Cookie("Device")
	if err != nil {
		io.WriteString(w, "You must connect to a device first: ")
		return
	}

	r.ParseForm()

	log.Println(r.Method)
	for k, v := range r.Form {
		log.Println(k, v)
	}

	cookies := r.Cookies()
	for _, c := range cookies {
		log.Println(c)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	log.Println(string(body))

	var sc script.Script
	err = json.Unmarshal([]byte(r.FormValue("Script")), &sc)
	if err != nil {
		log.Println(err.Error())
		io.WriteString(w, err.Error())
		return
	}

	sess.Result = sess.RUT.RunScript(&sc)
	log.Printf("+++++++++++++%q++++++++++++++\n", sc)
	log.Println(sc)
	io.WriteString(w, r.Host)
}

func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}

func writer(ws *websocket.Conn, sessionid string) {
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		pingTicker.Stop()
		ws.Close()
	}()

	sess, _ := Engine.Sessions[sessionid]
	for {
		select {
		case res, ok := <-sess.Result:
			if !ok {
				break
			}
			util.AppendToFile("script.log", []byte(res.Result))
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.TextMessage, []byte(res.Result)); err != nil {
				return
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func WS(w http.ResponseWriter, r *http.Request) {
	log.Println("Websocket Openend")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)
	go writer(ws, sessionid)
	reader(ws)
}

func RunCaseResultWS(w http.ResponseWriter, r *http.Request) {
	log.Println("Websocket Openend")

	sid := r.Context().Value("SESSIONID")
	sessionid, _ := sid.(string)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	go TestCaseResultWriter(ws, sessionid)
	TestCaseResultWSKeepAlive(ws)
}

func TestCaseResultWSKeepAlive(ws *websocket.Conn) {
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	go func() {
		for {
			_, _, err := ws.ReadMessage()
			if err != nil {
				break
			}
		}
	}()

	go func() {
		pingTicker := time.NewTicker(pingPeriod)
		defer func() {
			pingTicker.Stop()
		}()

		for _ = range pingTicker.C {
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}()

}
func TestCaseResultWriter(ws *websocket.Conn, sessionid string) {
	defer func() {
		ws.Close()
	}()

	sess, _ := Engine.Sessions[sessionid]
	_, ok := sess.CaseResult[sessionid]
	if !ok {
		panic("Result channel has beend remove")
	}

	for res := range sess.CaseResult[sessionid] {
		log.Printf("GetResult: %v", res)
		ws.SetWriteDeadline(time.Now().Add(writeWait))
		js, err := json.Marshal(res)
		if err != nil {
			log.Println("error happend when encoding result")
			continue
		}
		if err := ws.WriteMessage(websocket.TextMessage, js); err != nil {
			log.Println(err.Error())
		}
	}
	delete(sess.CaseResult, sessionid)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("login.html").Delims("|||", "|||").ParseFiles("asset/web/template/login.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		//log.Println(string(js))
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	} else if r.Method == "POST" {
		r.ParseForm()
		log.Printf("%#q\n", r.Form)
		for k, v := range r.Form {
			log.Println(k, ":", v)
		}

		var err error = nil
		newsess := Engine.GetSessionByUsernameAndPassword(r.FormValue("username"), r.FormValue("password"))
		if newsess == nil {
			log.Println("1212")
			newsess, err = Engine.AddSessionByUsernameAndPassword(r.FormValue("username"), r.FormValue("password"))
			if err != nil {
				log.Printf("Cannot add user: %s with: %s\n", r.FormValue("username"), err.Error())
				//w.WriteHeader(http.StatusNotAcceptable)
				return
			}
		}
		cookie := &http.Cookie{
			Name:  "SESSIONID",
			Value: newsess.ID,
		}

		log.Println(newsess.ID)

		http.SetCookie(w, cookie)
		w.WriteHeader(http.StatusOK)
		//	io.WriteString(w, util.GenerateSessionIDByUserNameAndPassword(r.FormValue("username"), r.FormValue("password")))
		t, err := template.New("main.html").Delims("|||", "|||").ParseFiles("asset/web/template/main.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		//log.Println(string(js))
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func MainPage(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method)
	log.Println("+++++++++")
	if r.Method == "GET" {
		log.Println("+++++++++")
		t, err := template.New("main.html").Delims("|||", "|||").ParseFiles("asset/web/template/main.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		//log.Println(string(js))
		err = t.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().AddDate(0, 0, -1)
	cookie := http.Cookie{Name: "SESSIONID", Value: "alice_cooper@gmail.com", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func Start() {
	mux := http.NewServeMux()
	//@liwei: This need more analysis.
	mux.HandleFunc("/vuetree", VUETree)
	mux.HandleFunc("/vuetest", VUEtest)
	mux.HandleFunc("/jstree", JSTree)
	mux.HandleFunc("/jsontree", JSONTree)
	mux.HandleFunc("/launchtree", LaunchTree)
	mux.HandleFunc("/runscript", NewRunScript)
	mux.HandleFunc("/script", Script)
	mux.HandleFunc("/product", Product)
	mux.HandleFunc("/jsnotify", JSNotify)
	mux.HandleFunc("/treeview", TreeView)
	mux.HandleFunc("/newcase", NewCase)
	mux.HandleFunc("/newtask", NewTask)
	mux.HandleFunc("/dumpcase", DumpCase)
	mux.HandleFunc("/dumptask", DumpTask)
	mux.HandleFunc("/dumpgroup", DumpGroup)
	mux.HandleFunc("/dumpsubgroup", DumpSubGroup)
	mux.HandleFunc("/dumpfeature", DumpFeature)
	mux.HandleFunc("/caseinfo", CaseInfo)
	mux.HandleFunc("/taskinfo", TaskInfo)
	mux.HandleFunc("/groupinfo", GroupInfo)
	mux.HandleFunc("/subgroupinfo", SubGroupInfo)
	mux.HandleFunc("/featureinfo", FeatureInfo)
	mux.HandleFunc("/runcases", RunCases)
	mux.HandleFunc("/delete", DeleteNode)
	mux.HandleFunc("/getdutcountbyid", GetDUTCountByID)
	mux.HandleFunc("/setdutsbyid", SetDUTsByID)
	mux.HandleFunc("/isinitialized", CheckIsReadyForRunByID)
	mux.HandleFunc("/ws", WS)
	mux.HandleFunc("/runcaseresultws", RunCaseResultWS)
	mux.HandleFunc("login", Login)
	mux.HandleFunc("/mainpage", MainPage)
	mux.HandleFunc("/", Login)
	mux.HandleFunc("/test", Test)
	mux.Handle("/asset/web/", http.FileServer(http.Dir(".")))

	//Add context support
	contextedMux := AddContextSupport(mux)
	log.Fatal(http.ListenAndServe(":8080", contextedMux))
}

func Test(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Context().Value("SESSIONID"))
}

func AddContextSupport(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, "-", r.RequestURI)
		cookie, _ := r.Cookie("SESSIONID")
		if cookie != nil {
			ctx := context.WithValue(r.Context(), "SESSIONID", cookie.Value)
			// WithContext returns a shallow copy of r with its context changed
			// to ctx. The provided ctx must be non-nil.
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func init() {
	Engine = controller.New()
}
