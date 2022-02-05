package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	//	"io/ioutil"
	"net/http"
	"os"
	//"path"
)

type FileInfo struct {
	Name string `json:name"`
	Wc   string `json:wc`
}

var fileinfo []FileInfo

func Upload(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var update bool
	log.Println("Entered to upload Handler")
	if req.Method == "PATCH" {

		fmt.Fprintf(w, "Method received in Upload handler : %v\n", req.Method)
		log.Println("Method received in Upload handler :", req.Method)
		update = true
	} else if req.Method != "POST" {

		fmt.Fprintf(w, "Wrong Header %v\n", req.Method)
		log.Println("Entered to upload Handler with Wrong Method")
		return
	} else {
		fmt.Fprintf(w, "Method received in Upload handler : %v\n", req.Method)
		log.Println("Method received in Upload handler :", req.Method)
	}

	// Write straight to disk
	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	// We always want to remove the multipart file as we're copying
	// the contents to another file anyway
	defer func() {
		if remErr := req.MultipartForm.RemoveAll(); remErr != nil {
			log.Println("Failed to remove multiform")
		}
	}()

	/*-------------------------
	Muitiple file upload
	--------------------------*/
	files := req.MultipartForm.File["uploadfile"]
	// initiate counter

	for i, _ := range files {
		file, err := files[i].Open()
		if err != nil {
			fmt.Println(err)
			log.Println("Error in file uploading")
			fmt.Fprintf(w, "Error in file uploading\n")
			return
		}
		defer file.Close()

		store := req.FormValue("store")
		if store == "" {
			log.Println("store is null")
			fmt.Fprintf(w, "store is null")
			return
		}
		log.Printf("store is %s\n", store)
		//fmt.Fprintf(w, "%v\n", handler.Header)
		vnfFilePath := "./" + store + "/"
		_, err = os.Stat(vnfFilePath)
		if os.IsNotExist(err) {
			errDir := os.MkdirAll(vnfFilePath, 0755)
			if errDir != nil {
				log.Println(vnfFilePath, " creation is failed with error:", err)
			}
		}
		if _, err := os.Stat(vnfFilePath + files[i].Filename); errors.Is(err, os.ErrNotExist) {
			log.Println("Filename not exists")
		} else {
			log.Println("Filename exists")
			if !(update) {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		destFile, err := os.OpenFile(vnfFilePath+files[i].Filename, os.O_WRONLY|os.O_CREATE, 0666)
		defer destFile.Close()
		log.Println("File to copy : ", files[i].Filename)
		// Write contents of uploaded file to destFile
		if _, err = io.Copy(destFile, file); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("File copy failed with err: ", err)
			return
		}
	}

}

func Remove(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	if req.Method != "DELETE" {

		fmt.Fprintf(w, "Wrong Header %v\n", req.Method)
		log.Println("Entered to upload Handler with Wrong Method")
		return
	} else {
		log.Println("Entered to upload Handler with Delete Method")
	}
	//var err error

	// Write straight to disk
	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Println("error in parsing", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// We always want to remove the multipart file as we're copying
	// the contents to another file anyway
	defer func() {
		if remErr := req.MultipartForm.RemoveAll(); remErr != nil {
			log.Println("Failed to remove multiform")
		}
	}()

	/*-------------------------
	Muitiple file upload
	--------------------------*/
	//files := req.MultipartForm.Value["deletefiles"]
	//for _, filename := range files {
	filename := req.FormValue("deletefiles")
	//fmt.Fprintf(w, "%v\n", handler.Header)
	fmt.Println("Filename for delete :", filename)
	vnfFilePath := "./" + "store" + "/"
	if _, err := os.Stat(vnfFilePath + filename); errors.Is(err, os.ErrNotExist) {
		log.Println("Filename not exists")
		http.Error(w, "File Already Not Exists"+filename, http.StatusInternalServerError)
		return
	} else {
		log.Println("Filename exists")

	}
	err = os.Remove(vnfFilePath + filename)
	log.Println("File to delete : ", filename)
	// Write contents of uploaded file to destFile
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("File delete failed with err: ", err)
		return
	}

	//}

}

func List(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	var count int
	if req.Method != "GET" {

		fmt.Fprintf(w, "Wrong Header %v\n", req.Method)
		log.Println("Entered to upload Handler with Wrong Method")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	files, err := ioutil.ReadDir("store/")
	if err != nil {
		log.Println("File open error ", err)
	}
	fileinfo = make([]FileInfo, len(files))
	for i, file := range files {

		vnfFilePath := "./" + "store" + "/"
		destFile, _ := os.Open(vnfFilePath + file.Name())
		defer destFile.Close()

		// initiate scanner from file handle
		fileScanner := bufio.NewScanner(destFile)

		// tell the scanner to split by words
		fileScanner.Split(bufio.ScanWords)

		count = 0

		// for looping through results
		for fileScanner.Scan() {
			fmt.Printf("word: '%s' - position: '%d'\n", fileScanner.Text(), count)
			count++
		}
		fmt.Println(count)
		fileinfo[i].Name = file.Name()
		fileinfo[i].Wc = strconv.Itoa(count)
	}
	fmt.Println(fileinfo)
	json.NewEncoder(w).Encode(fileinfo)
	return
}

func main() {
	//root, _ := os.Getwd()

	fmt.Printf("\nStarting server in 4000 port\n\n")

	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	file := root + path.Clean(r.URL.String())
	//	http.ServeFile(w, r, file)
	//})

	http.HandleFunc("/upload", Upload)
	http.HandleFunc("/remove", Remove)
	http.HandleFunc("/list", List)
	//	http.HandleFunc("/mutifileupload", MutliUpload)
	http.ListenAndServe(":4000", nil)
}
