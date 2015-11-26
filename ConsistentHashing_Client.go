
package main
import (
    "fmt"
    "net/http"
   "io/ioutil"
    "hash/crc32"
    "sort"
    "strconv"
)

func main(){
	client := &http.Client{}
	url1 :="http://localhost:3000"
	url2 :="http://localhost:3001"
	url3 :="http://localhost:3002"
    
	url1Hash := []byte(url1)
	url2Hash := []byte(url2)
	url3Hash := []byte(url3)
	hashServ :=make(map[uint32]string)
	

	hashServ[crc32.ChecksumIEEE(url1Hash)] =url1
	hashServ[crc32.ChecksumIEEE(url2Hash)] =url2
	hashServ[crc32.ChecksumIEEE(url3Hash)] =url3

	var keys []int
	for k := range hashServ{
		keys = append(keys,int(k))
	}
	sort.Ints(keys)

	for _,k := range keys{
		fmt.Println("The sorted keys are ",k)
	}
	
	keyMap :=make(map[uint32]string)
	keyMap[1]="a"
	keyMap[2]="b"
	keyMap[3]="c"
	keyMap[4]="d"
	keyMap[5]="e"
	keyMap[6]="f"
	keyMap[7]="g"
	keyMap[8]="h"
	keyMap[9]="i"
	keyMap[10]="j"

	keyToServUrlMap :=make(map[uint32]uint32)
	var tmpVal uint32
	for i:=1;i<=10;i++{
		tmpVal = crc32.ChecksumIEEE([]byte(strconv.Itoa(i)))
	
		for ind,k := range keys{
			if(tmpVal< uint32(k)){
				keyToServUrlMap[tmpVal] = uint32(k)
			}
			if ind==0 && tmpVal >uint32(k) {
				keyToServUrlMap[tmpVal] = uint32(k)
			}
		}
	}
	
	for k := range keyToServUrlMap{
		for i:=1;i<=10;i++{
			if crc32.ChecksumIEEE([]byte(strconv.Itoa(i))) ==k{
				req1, _ := http.NewRequest("PUT",hashServ[keyToServUrlMap[k]]+"/keys/"+strconv.Itoa(i)+"/"+keyMap[uint32(i)],nil)
				response, err:= client.Do(req1)
				fmt.Println("final url is",hashServ[keyToServUrlMap[k]]+"/keys/"+strconv.Itoa(i)+"/"+keyMap[uint32(i)])
				fmt.Println("Res",response,"err",err)
			}
		}
	}
	req, _ := http.NewRequest("GET",url1+"/keys",nil)
	response, err:= client.Do(req)
	if err!=nil{
		fmt.Println(err)
	}
	contents,_ := ioutil.ReadAll(response.Body)
	fmt.Println("The GET ALL KEYS for url:",url1,"is:",string(contents))

	req, _ = http.NewRequest("GET",url2+"/keys",nil)
	response, err= client.Do(req)
	if err!=nil{
		fmt.Println(err)
	}
	contents,_ = ioutil.ReadAll(response.Body)
	fmt.Println("The GET ALL KEYS for url:",url2,"is:",string(contents))

	req, _ = http.NewRequest("GET",url3+"/keys",nil)
	response, err= client.Do(req)
	if err!=nil{
		fmt.Println(err)
	}
	contents,_ = ioutil.ReadAll(response.Body)
	fmt.Println("The GET ALL KEYS for url:",url3,"is:",string(contents))

}
