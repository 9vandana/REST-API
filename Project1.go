package main

import (
    "log"
    "net/http"
    "database/sql"
    "github.com/gorilla/handlers"
    _ "github.com/go-sql-driver/mysql"
    // "rest-api-sites/rest-api-sites/route"
    "io"
    "io/ioutil"
    "fmt"
    "github.com/gorilla/mux"
    "encoding/json"
)

var (
    name string
    role string 
    uri string
    label string 
    address string
)

func main() {
    
    router := NewRouter()
	router.HandleFunc("/sites", GetSites).Methods("GET")
    router.HandleFunc("/sites/{name}", GetSite).Methods("GET")
    router.HandleFunc("/sites/{name}", DeleteSite).Methods("DELETE")
    router.HandleFunc("/sites/{name}", CreateSite).Methods("POST")
    router.HandleFunc("/sites/{name}", EditSite).Methods("PUT")
     
    allowedOrigins := handlers.AllowedOrigins([]string{"*"})
    allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	
    log.Fatal(http.ListenAndServe(":8000", handlers.CORS(allowedOrigins, allowedMethods)(router)))
}

type Site struct {
    Name string   `json:"name,omitempty"`
    Role string   `json:"role,omitempty"`
    URI  string   `json:"uri,omitempty"`
    AccessPoints   *AccessPoint `json:"accesspoint,omitempty"`
}

type AccessPoint struct {
    Label string `json:"label,omitempty"`
    URL string `json:"url,omitempty"`
}

type Message struct {
    Msg string   `json:"msg,omitempty"`
}

var sites []Site
var access_pts []AccessPoint
var msg1 Message

func NewRouter() *mux.Router {
    router := mux.NewRouter()
    return router
}

func GetSites(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:test123@/restapiDB")
    if err != nil {
       panic(err)
    }
    defer db.Close()
    rows, err := db.Query("select name, role, uri from Site")
    if err != nil {
       panic(err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&name, &role, &uri)
        if err != nil {
           panic(err)
        }
        rowp, err := db.Query("select label, address from accesspoint where name = ?",&name)
        if err != nil {
            panic(err)
        }
        for rowp.Next() {
            err := rowp.Scan(&label, &address)
            if err != nil {
                panic(err)
            }
            sites = append(sites, Site{Name: name, Role: role, URI: uri , AccessPoints: &AccessPoint{Label: label, URL: address} })
        }
    }
    err = rows.Err()
    if err != nil {
       panic(err)
    }
    json.NewEncoder(w).Encode(sites)    
}

func GetSite(w http.ResponseWriter, r *http.Request) {
    name = " "
    params := mux.Vars(r)
    name = params["name"]
    db, err := sql.Open("mysql", "root:test123@/restapiDB")
    if err != nil {
       panic(err)
    }
    defer db.Close()
    rows, err := db.Query("select name, role, uri from Site where name = ?", &name)
    if err != nil {
       panic(err)
    }
    defer rows.Close()
    for rows.Next() {
        err := rows.Scan(&name, &role, &uri)
        if err != nil {
           panic(err)
        }
        rowp, err := db.Query("select label, address from accesspoint where name = ?",&name)
        if err != nil {
            panic(err)
        }
        for rowp.Next() {
            err := rowp.Scan(&label, &address)
            if err != nil {
                panic(err)
            }
            sites = append(sites, Site{Name: name, Role: role, URI: uri , AccessPoints: &AccessPoint{Label: label, URL: address} })
        }
    }
    err = rows.Err()
    if err != nil {
       panic(err)
    }
    json.NewEncoder(w).Encode(sites)    
}

func DeleteSite(w http.ResponseWriter, r *http.Request) {
    name = " "
    params := mux.Vars(r)
    name = params["name"]
    db, err := sql.Open("mysql", "root:test123@/restapiDB")
    if err != nil {
       panic(err)
    }
    defer db.Close()
    delAP, err := db.Query("delete from accesspoint where name = ?", &name)
    if err != nil {
       panic(err)
    }
    defer delAP.Close()
    delSite, err := db.Query("delete from Site where name = ?", &name)
    if err != nil {
       panic(err)
    }
    defer delSite.Close()
    msg1.Msg = "Deleted site and its access points"
    json.NewEncoder(w).Encode(msg1)
}

func CreateSite(w http.ResponseWriter, r *http.Request) {
    //params := mux.Vars(r)
    var site Site
    //_ = json.NewDecoder(r.Body).Decode(&site)
    //var site Site 
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    fmt.Println(string(body))
    err = json.Unmarshal(body, &site)
    db, err := sql.Open("mysql", "root:test123@/restapiDB")
    if err != nil {
       panic(err)
    }
    defer db.Close()
    //rows, err := db.Query("insert into Site values(?,?,?)", &site.Name,&site.Role,&site.URI)
    rows, err := db.Query("insert into Site values(?,?,?)",&site.Name,&site.Role,&site.URI)
    if err != nil {
       panic(err)
    }
    defer rows.Close()
    msg1.Msg = "Created site and its access points"
    json.NewEncoder(w).Encode(msg1)
}

func EditSite(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var site Site
    _ = json.NewDecoder(r.Body).Decode(&site)
    site.Name = params["name"]
    for index, item := range sites {
         if item.Name == params["name"] {            
             sites = append(sites[:index], sites[index+1:]...)
             sites = append(sites, site)
         }
    for _, item := range sites {
        if item.Name == params["name"] {     
            item.URI = params["uri"]
            item.Role = params["role"]
            //sites = append(sites[:index], sites[index+1:]...)
            //sites = append(sites, site)
        }

    }
    json.NewEncoder(w).Encode(&Site{})
}
}