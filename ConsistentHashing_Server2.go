package main
import (
    "github.com/julienschmidt/httprouter"
    "encoding/json"
    "net/http"
    "strconv"
)

type JsonObjReq struct{
    Name string
}

type JsonObjRes struct{
    Key int `json:"key"`
    Value string `json:"value"`
}


var jsonObj JsonObjReq

var store[] string 

var lent int
func getKey(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    id,_ := strconv.Atoi(p.ByName("id"))
    jsonObjRes := JsonObjRes{id,store[id]}
    
    js, err := json.Marshal(jsonObjRes)
        if err != nil {
        panic(err)
        }
        rw.Write(js)
}
func getAllKeys(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
    jsonObjArr := make([]JsonObjRes,lent)
    var i int
    i=0
   
    for ind,ele :=range store{
        
        if ele!=""{
        jsonObjArr[i] = JsonObjRes{ind,ele}
        i++
        }
    }
    
    js, err := json.Marshal(jsonObjArr)
        if err != nil {
        panic(err)
        }
        rw.Write(js)
}
func putKey(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
        
    id,_ := strconv.Atoi(p.ByName("id"))
    value := p.ByName("value")
    store[id] = value
    lent++
    rw.WriteHeader(http.StatusCreated)
}

func main() {
    lent =0
    store = make([]string,15)
    mux := httprouter.New()
    mux.GET("/keys/:id", getKey)
    mux.PUT("/keys/:id/:value",putKey)
    mux.GET("/keys",getAllKeys)
    server := http.Server{
            Addr:        "127.0.0.1:3002",
            Handler: mux,
    }
    server.ListenAndServe()
}