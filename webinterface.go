package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "path/filepath"
    "sync"
    "runtime"
    "time"
    "os/exec"
    "os"
    "io/ioutil"
)

func deliverFile(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        var data []byte
        var err error
        switch {
          case r.URL.Path == "/info.json":
             data, err = ioutil.ReadFile("/root/storage/info.json")
          case r.URL.Path == "/work.sh":
             data, err = ioutil.ReadFile("/root/work.sh")
          case r.URL.Path == "/db-plugin.code":
             data, err = ioutil.ReadFile("/root/storage/db-plugin.code")
        }
        if err != nil {
            p(fmt.Sprintf("Error: could not find file specific to this bucket %s", r.URL.Path))
            fmt.Fprintf(w, "Error: file %s does not exist on this system", r.URL.Path[1:])
        } else {
            fmt.Fprintf(w, "%s", string(data))
        }
    }
}

func indexpage(w http.ResponseWriter, r *http.Request) {
    //fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        data, err := Asset("views/index.html") // remove the leading '/'
        if err != nil {
            p(fmt.Sprintf("Error: could not find file views/index.html"))
            panic(err)
        }
        t, _ := template.New("dummy").Parse(string(data))
        t.Execute(w, nil)
    }
}

func returnViews(w http.ResponseWriter, r *http.Request) {
    //fmt.Printf("path: %s %s\n", r.URL.Path, r.Method) // get request method
    if r.Method == "GET" {
        data, err := Asset(r.URL.Path[1:]) // remove the leading '/'
        if err != nil {
            p(fmt.Sprintf("Error: could not find file %s", r.URL.Path))
            panic(err)
        }
        // if we have a html file we should use the templates, if we have any other file don't use templates
        if filepath.Ext(r.URL.Path) == ".html" {         
            t, _ := template.New("dummy").Parse(string(data))
            t.Execute(w, nil)
        } else {
            //fmt.Printf("export: %s\n%s\n", r.URL.Path, string(data))
            fmt.Fprintf(w, "%s\n", string(data))
        }
    }
}

func update(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    if r.Method == "POST" {
           file := r.FormValue("file")
           content := r.FormValue("content")
           if file == "info.json" {
               // save content if file exists
               fn := "/root/storage/info.json"
               if _, err := os.Stat(fn); os.IsNotExist(err) {
                  p(fmt.Sprintf("no such file or directory: %s", fn))
                  return
               }
               f, err := os.Create(fn)
               if err != nil {
                   p(fmt.Sprintf("Error: could not open file for writing %s", fn))
               }
               defer f.Close()
               f.WriteString(content)
           }
           if file == "work.sh" {
               // save content if file exists
               fn := "/root/work.sh"
               if _, err := os.Stat(fn); os.IsNotExist(err) {
                  p(fmt.Sprintf("no such file or directory: %s", fn))
                  return
               }
               f, err := os.Create(fn)
               if err != nil {
                   p(fmt.Sprintf("Error: could not open file for writing %s", fn))
               }
               defer f.Close()
               f.WriteString(content)
           }
           if file == "db-plugin.code" {
               // save content if file exists
               fn := "/root/storage/db-plugin.code"
               if _, err := os.Stat(fn); os.IsNotExist(err) {
                  p(fmt.Sprintf("no such file or directory: %s", fn))
                  return
               }
               f, err := os.Create(fn)
               if err != nil {
                   p(fmt.Sprintf("Error: could not open file for writing %s", fn))
               }
               defer f.Close()
               f.WriteString(content)
           }
    }
}

func waitServer(url string) bool {
	tries := 20
	for tries > 0 {
		resp, err := http.Get(url)
		if err == nil {
			resp.Body.Close()
			return true
		}
		time.Sleep(200 * time.Millisecond)
		tries--
	}
	return false
}

func startBrowser(url string) bool {
	// try to start the browser
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

func webinterface( verbose bool, init bool, wg sync.WaitGroup ) {
    defer wg.Done()
    http.HandleFunc("/", indexpage) // setting router rule
    http.HandleFunc("/views/", returnViews)
    http.HandleFunc("/info.json", deliverFile)
    http.HandleFunc("/work.sh", deliverFile)
    http.HandleFunc("/db-plugin.code", deliverFile)
    http.HandleFunc("/update", update)
    // try to open the web-browser
    go func() {
        if waitServer(":9090/") && startBrowser(":9090/") {
            if verbose { 
              fmt.Printf("tried to open localhost:9090 in local webbrowser")
            }
        } else {
            p(fmt.Sprintf("listen to localhost:9090"))
        }
    }()
    
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
