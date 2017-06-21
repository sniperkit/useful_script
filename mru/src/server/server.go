package server

import (
	//	"command"
	"cache"
	"encoding/json"
	"github.com/gorilla/websocket"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mcase"
	"net/http"
	"newcache"
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

type Server struct {
	RUT        *rut.RUT
	Result     <-chan result.Result
	CaseResult map[string]chan *result.Result
	CaseDB     cache.Cache
	NewCache   *newcache.NewCache
}

var DefaultServer Server

func Start() {
	DefaultServer.Start()
}

type Page struct {
	Link        string
	Description string
}

func Product(w http.ResponseWriter, r *http.Request) {
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
			DefaultServer.RUT = dev
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

func ZTreeMenu(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("ztreemenu.html").Delims("|||", "|||").ParseFiles("asset/web/template/ztreemenu.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, DefaultServer.CaseDB)
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

func ResponsiveNav(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("responsive.html").Delims("|||", "|||").ParseFiles("asset/web/template/responsive.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		err = t.Execute(w, DefaultServer.CaseDB)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func TreeView(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Println("Dump tree")
		encoder := json.NewEncoder(w)
		err := encoder.Encode(DefaultServer.NewCache.TreeView().Children)
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func NewCase(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.New("newcase.html").Delims("|||", "|||").ParseFiles("asset/web/template/newcase.html", "asset/web/template/vuefooter.html", "asset/web/template/vueheader.html", "asset/web/template/treenav.html")
		if err != nil {
			log.Println(err)
			io.WriteString(w, err.Error())
			return
		}

		cookie := &http.Cookie{
			Name:  "SESSIONID",
			Value: util.GenerateSessionID(),
		}

		http.SetCookie(w, cookie)
		js, _ := json.Marshal(DefaultServer.NewCache.TreeView().Children)
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
		DefaultServer.NewCache.AddCase(&newcase)
	}
}

func NewTask(w http.ResponseWriter, r *http.Request) {
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

		js, _ := json.Marshal(DefaultServer.NewCache.TreeView().Children)
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

		err = DefaultServer.NewCache.AddTask(caseid.Value, &newtask)
		if err != nil {
			log.Println(err.Error())
			io.WriteString(w, err.Error())
		}
	}
}

func DumpSubGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		cookies := r.Cookies()
		for _, cookie := range cookies {
			log.Println(cookie)
		}
		encoder := json.NewEncoder(w)
		c, err := DefaultServer.NewCache.GetSubGroupByID(r.FormValue("id"))
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
	if r.Method == "GET" {
		r.ParseForm()
		cookies := r.Cookies()
		for _, cookie := range cookies {
			log.Println(cookie)
		}
		encoder := json.NewEncoder(w)
		c, err := DefaultServer.NewCache.GetFeatureByID(r.FormValue("id"))
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
	if r.Method == "GET" {
		r.ParseForm()
		cookies := r.Cookies()
		for _, cookie := range cookies {
			log.Println(cookie)
		}

		log.Println("DumpGroup: ", r.FormValue("id"))
		encoder := json.NewEncoder(w)
		c, err := DefaultServer.NewCache.GetGroupByID(r.FormValue("id"))
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
	if r.Method == "GET" {
		r.ParseForm()
		cookies := r.Cookies()
		for _, cookie := range cookies {
			log.Println(cookie)
		}
		encoder := json.NewEncoder(w)
		c, err := DefaultServer.NewCache.GetCaseByID(r.FormValue("id"))
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
		c, err := DefaultServer.NewCache.GetTaskByID(caseid.Value, r.FormValue("id"))
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
	if r.Method == "POST" {
		r.ParseForm()

		sessionid, err := r.Cookie("SESSIONID")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		log.Println("SESSIONiD: ", sessionid.Value)
		if _, ok := DefaultServer.CaseResult[sessionid.Value]; ok {
			io.WriteString(w, "You are runing another cases")
			return
		}

		if r.FormValue("type") == "GROUP" {
			DefaultServer.NewCache.DelGroupByID(r.FormValue("id"))
		} else if r.FormValue("type") == "SUBGROUP" {
			DefaultServer.NewCache.DelSubGroupByID(r.FormValue("id"))
		} else if r.FormValue("type") == "FEATURE" {
			DefaultServer.NewCache.DelFeatureByID(r.FormValue("id"))
		} else if r.FormValue("type") == "CASE" {
			DefaultServer.NewCache.DelCaseByID(r.FormValue("id"))
		} else if r.FormValue("type") == "TASK" {
			caseid, err := r.Cookie("CASEID")
			if err != nil {
				io.WriteString(w, "Case ID is not set when delete task")
				return
			}
			DefaultServer.NewCache.DelTaskByID(caseid.Value, r.FormValue("id"))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Invalid request") //A proper status code in more usefull.
		}
	}
}

func GetDUTCountByID(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		r.ParseForm()
		log.Println(r.FormValue("id"))

		if r.FormValue("id") == "" {
			io.WriteString(w, "ID is not set!")
			return
		}

		count, err := DefaultServer.NewCache.GetDUTCountByID(r.FormValue("id"))
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
		ready, err := DefaultServer.NewCache.CheckIsReadyForRunByID(r.FormValue("id"))
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

		err = DefaultServer.NewCache.SetDUTsByID(r.FormValue("id"), duts)
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
	if r.Method == "POST" {
		r.ParseForm()

		sessionid, err := r.Cookie("SESSIONID")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		log.Println("SESSIONiD: ", sessionid.Value)
		if _, ok := DefaultServer.CaseResult[sessionid.Value]; ok {
			io.WriteString(w, "You are runing another cases")
			return
		}

		DefaultServer.CaseResult[sessionid.Value] = make(chan *result.Result, 10)

		log.Println(DefaultServer.CaseResult[sessionid.Value])
		if r.FormValue("type") == "GROUP" {
			go func() {
				wg := sync.WaitGroup{}
				for res := range DefaultServer.NewCache.RunCasesByGroupID(r.FormValue("id")) {
					wg.Add(1)
					log.Printf("%#v", res)
					go func(r *result.Result) {
						DefaultServer.CaseResult[sessionid.Value] <- res
						wg.Done()
					}(res)
				}
				wg.Wait()
				close(DefaultServer.CaseResult[sessionid.Value])

			}()
		} else if r.FormValue("type") == "SUBGROUP" {
			go func() {
				wg := sync.WaitGroup{}
				for res := range DefaultServer.NewCache.RunCasesBySubGroupID(r.FormValue("id")) {
					wg.Add(1)
					go func(r *result.Result) {
						DefaultServer.CaseResult[sessionid.Value] <- r
						wg.Done()
					}(res)
				}
				wg.Wait()
				close(DefaultServer.CaseResult[sessionid.Value])
			}()
		} else if r.FormValue("type") == "FEATURE" {
			go func() {
				wg := sync.WaitGroup{}
				for res := range DefaultServer.NewCache.RunCasesByFeatureID(r.FormValue("id")) {
					wg.Add(1)
					go func(r *result.Result) {
						DefaultServer.CaseResult[sessionid.Value] <- r
						wg.Done()
					}(res)
				}
				wg.Wait()
				close(DefaultServer.CaseResult[sessionid.Value])

			}()
		} else if r.FormValue("type") == "CASE" {
			go func() {
				wg := sync.WaitGroup{}
				for res := range DefaultServer.NewCache.RunCaseByID(r.FormValue("id")) {
					wg.Add(1)
					go func(r *result.Result) {
						DefaultServer.CaseResult[sessionid.Value] <- r
						wg.Done()
					}(res)
				}
				wg.Wait()
				close(DefaultServer.CaseResult[sessionid.Value])
			}()
		} else if r.FormValue("type") == "TASK" {
			caseid, err := r.Cookie("CASEID")
			if err != nil {
				io.WriteString(w, err.Error())
				return
			}
			go func() {
				wg := sync.WaitGroup{}
				for res := range DefaultServer.NewCache.RunTaskByID(caseid.Value, r.FormValue("id")) {
					wg.Add(1)
					go func(r *result.Result) {
						DefaultServer.CaseResult[sessionid.Value] <- r
						wg.Done()
					}(res)
				}
				wg.Wait()
				close(DefaultServer.CaseResult[sessionid.Value])
			}()
		} else {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, "Invalid request") //A proper status code in more usefull.
		}
	}
	io.WriteString(w, r.Host)
}

func NewRunScript(w http.ResponseWriter, r *http.Request) {
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
		io.WriteString(w, err.Error())
		return
	}

	DefaultServer.Result = DefaultServer.RUT.RunScript(&sc)
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

func writer(ws *websocket.Conn) {
	pingTicker := time.NewTicker(pingPeriod)
	defer func() {
		pingTicker.Stop()
		ws.Close()
	}()
	for {
		select {
		case res, ok := <-DefaultServer.Result:
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
	go writer(ws)
	reader(ws)
}

func RunCaseResultWS(w http.ResponseWriter, r *http.Request) {
	log.Println("Websocket Openend")
	sessionid, err := r.Cookie("SESSIONID")
	if err != nil {
		io.WriteString(w, "Unknown session: "+err.Error())
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}
	go TestCaseResultWriter(ws, sessionid.Value)
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
	_, ok := DefaultServer.CaseResult[sessionid]
	if !ok {
		panic("Result channel has beend remove")
	}

	for res := range DefaultServer.CaseResult[sessionid] {
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
	delete(DefaultServer.CaseResult, sessionid)
}

func (s *Server) Start() {
	//@liwei: This need more analysis.
	http.HandleFunc("/ztreemenu", ZTreeMenu)
	http.HandleFunc("/vuetree", VUETree)
	http.HandleFunc("/vuetest", VUEtest)
	http.HandleFunc("/jstree", JSTree)
	http.HandleFunc("/jsontree", JSONTree)
	http.HandleFunc("/launchtree", LaunchTree)
	http.HandleFunc("/runscript", NewRunScript)
	http.HandleFunc("/script", Script)
	http.HandleFunc("/product", Product)
	http.HandleFunc("/jsnotify", JSNotify)
	http.HandleFunc("/responsive", ResponsiveNav)
	http.HandleFunc("/treeview", TreeView)
	http.HandleFunc("/newcase", NewCase)
	http.HandleFunc("/newtask", NewTask)
	http.HandleFunc("/dumpcase", DumpCase)
	http.HandleFunc("/dumptask", DumpTask)
	http.HandleFunc("/dumpgroup", DumpGroup)
	http.HandleFunc("/dumpsubgroup", DumpSubGroup)
	http.HandleFunc("/dumpfeature", DumpFeature)
	http.HandleFunc("/caseinfo", CaseInfo)
	http.HandleFunc("/taskinfo", TaskInfo)
	http.HandleFunc("/groupinfo", GroupInfo)
	http.HandleFunc("/subgroupinfo", SubGroupInfo)
	http.HandleFunc("/featureinfo", FeatureInfo)
	http.HandleFunc("/runcases", RunCases)
	http.HandleFunc("/delete", DeleteNode)
	http.HandleFunc("/getdutcountbyid", GetDUTCountByID)
	http.HandleFunc("/setdutsbyid", SetDUTsByID)
	http.HandleFunc("/isinitialized", CheckIsReadyForRunByID)
	http.HandleFunc("/ws", WS)
	http.HandleFunc("/runcaseresultws", RunCaseResultWS)
	http.HandleFunc("/", NewCase)
	http.Handle("/asset/web/", http.FileServer(http.Dir(".")))
	log.Panic(http.ListenAndServe(":8080", nil))
}

func init() {
	DefaultServer.NewCache = newcache.New("V8300")
	DefaultServer.NewCache.Restore()
	DefaultServer.CaseResult = make(map[string]chan *result.Result)
}
