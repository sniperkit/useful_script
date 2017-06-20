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
	RUT      *rut.RUT
	Result   <-chan result.Result
	CaseDB   cache.Cache
	NewCache *newcache.NewCache
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

		js, _ := json.Marshal(DefaultServer.NewCache.TreeView().Children)
		log.Println(string(js))
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
		json.Unmarshal([]byte(r.FormValue("newcase")), &newcase)
		log.Println(newcase)
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

		cookie := &http.Cookie{
			Name:  "CASEID",
			Value: r.FormValue("id"),
			Path:  "dumpsubgroup",
		}

		http.SetCookie(w, cookie)
		err = t.Execute(w, nil)
		if err != nil {
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
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

		log.Printf("++++++++++++++++%s-------------\n", r.FormValue("id"))
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

		cookie := &http.Cookie{
			Name:  "FEATUREID",
			Value: r.FormValue("id"),
			Path:  "featureinfo",
		}

		http.SetCookie(w, cookie)
		err = t.Execute(w, nil)
		if err != nil {
			io.WriteString(w, err.Error())
		}
	} else if r.Method == "POST" {
		io.WriteString(w, "Invalid request") //A proper status code in more usefull.
	}
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

func (s *Server) Start() {
	//@liwei: This need more analysis.
	http.HandleFunc("/ztreemenu", ZTreeMenu)
	http.HandleFunc("/vuetree", VUETree)
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
	http.HandleFunc("/ws", WS)
	http.HandleFunc("/", NewCase)

	http.Handle("/asset/web/", http.FileServer(http.Dir(".")))
	log.Panic(http.ListenAndServe(":8080", nil))
}

func init() {

	DefaultServer.NewCache = newcache.New("V8300")

	DefaultServer.NewCache.Restore()

}
