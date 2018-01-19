package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const(
	UPLOAD_DIR="./uploads"
)
func uploadHandler(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		io.WriteString(w,"<html><form method=\"POST\" action=\"/upload\" " +
			"enctype=\"multipart/form-data\">" +
			"Chose an image to upload:<input name=\"image\" type=\"file\" />" +
			"<br><input type=\"submit\" value=\"Upload\" />" +
			"</form></html>")
		return
	}
	if r.Method=="POST"{
		f,h,err:=r.FormFile("image")
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		filename:=h.Filename
		defer f.Close()
		t,err:=os.Create(getCurrentPath()+"/"+filename)
		if err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		defer t.Close()
		if _,err:=io.Copy(t,f);err!=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}
		http.Redirect(w,r,"/view?id="+filename,http.StatusFound)
	}
}
func viewHandler(w http.ResponseWriter,r *http.Request){
	imageId:=r.FormValue("id")
	imagePath:=getCurrentPath()+"/"+imageId
	w.Header().Set("Content-Type","image")
	http.ServeFile(w,r,imagePath)
}
func uploadHandler2(w http.ResponseWriter,r *http.Request){
	if r.Method=="GET"{
		io.WriteString(w,"<html><form method=\"POST\" action=\"/upload\" "+
			" enctype=\"multipart/form-data\">"+
			"Choose an image to upload:<input name=\"image\" type=\"file\" /><br>"+
			"<input type=\"submit\" value=\"Upload\" />"+
			"</form></html>")
		return
	}
}
func getCurrentPath() string {
	s, err := exec.LookPath(os.Args[0])
	checkErr(err)
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	http.HandleFunc("/view",viewHandler)
	http.HandleFunc("/upload",uploadHandler)
	err:=http.ListenAndServe(":8087",nil)
	if err!=nil{
		log.Fatal("ListenAndServe:",err.Error())
	}
}
