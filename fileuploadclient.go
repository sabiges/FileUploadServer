package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	//"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

type FileInfo struct {
	Name string `json:name"`
	Wc   string `json:wc`
}

var fileinfo []FileInfo

func listFile(targetUrl string, flag string) error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", targetUrl, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}

	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(&fileinfo)
	if flag == "listfile" {
		for _, filelist := range fileinfo {
			fmt.Printf("API Response For list file names are   %+v\n", filelist.Name)
		}
	} else {
		for _, filelist := range fileinfo {
			fmt.Printf("API Response For WordCount of file names (%v) is  %+v\n", filelist.Name, filelist.Wc)
		}
	}

	return nil
}

func postFile(filenames []string, targetUrl string, flag string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	valueWriter, err1 := bodyWriter.CreateFormField("store")
	if err1 != nil {
		fmt.Println("error writing to store")
		return err1
	}
	_, err2 := io.Copy(valueWriter, strings.NewReader("store"))
	if err2 != nil {
		return err2
	}

	for _, filename := range filenames {
		// this step is very important
		fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
		if err != nil {
			fmt.Println("error writing to buffer")
			return err
		}

		// open file handle
		fh, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file")
			return err
		}
		defer fh.Close()

		//iocopy
		_, err = io.Copy(fileWriter, fh)
		if err != nil {
			return err
		}
	}
	//contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	var err3 error
	var req *http.Request
	if flag == "update" {
		req, err3 = http.NewRequest("PATCH", targetUrl, bytes.NewReader(bodyBuf.Bytes()))
	} else {
		req, err3 = http.NewRequest("POST", targetUrl, bytes.NewReader(bodyBuf.Bytes()))
	}
	if err3 != nil {
		fmt.Println("Error return from file upload server :", err3)
		return err3
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	rsp, errclient := client.Do(req)
	if errclient != nil {
		fmt.Println("Request failed with response code: ", errclient)
		return errclient
	}
	if rsp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with response code: %d", rsp.StatusCode)
	} else {
		fmt.Println("Api request %s successfully finished ", flag)
	}
	return nil
}

func removeFile(filenames []string, targetUrl string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	valueWriter, err1 := bodyWriter.CreateFormField("store")
	if err1 != nil {
		fmt.Println("error writing to store")
		return err1
	}
	_, err2 := io.Copy(valueWriter, strings.NewReader("store"))
	if err2 != nil {
		return err2
	}

	/*
		for _, filename := range filenames {
			// this step is very important
			y1, _ := bodyWriter.CreateFormField("deletefiles")
			_, err2 = io.Copy(y1, strings.NewReader(filenames[0]))
			if err2 != nil {
				return err2
			}
		}
	*/
	valueWriter1, err2 := bodyWriter.CreateFormField("deletefiles")
	if err2 != nil {
		fmt.Println("error writing to deletefile")
		return err2
	}
	_, err2 = io.Copy(valueWriter1, strings.NewReader(filenames[0]))
	if err2 != nil {
		return err2
	}
	bodyWriter.Close()
	var err3 error
	var req *http.Request
	req, err3 = http.NewRequest("DELETE", targetUrl, bytes.NewReader(bodyBuf.Bytes()))
	if err3 != nil {
		fmt.Println("Error return from file upload server :", err3)
		return err3
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	rsp, errclient := client.Do(req)
	if errclient != nil {
		fmt.Println("Request failed with response code: ", errclient)
		return errclient
	}
	if rsp.StatusCode != http.StatusOK {
		fmt.Printf("Request failed with response code: %d", rsp.StatusCode)
	} else {
		fmt.Printf("Api request  successfully finished \n")
	}
	return nil
}

// sample usage
func main() {
	target_url := "http://localhost:4000/upload"
	target_url_rm := "http://localhost:4000/remove"
	target_url_ls := "http://localhost:4000/list"
	target_url_wc := "http://localhost:4000/list"
	//filename := "./test.json"

	if len(os.Args) < 2 {
		fmt.Println("\nWrong input :", os.Args)
		help()
		os.Exit(1)
	}

	switch os.Args[1] {

	case "add":
		fmt.Println("Add is called with files :", os.Args[2:])
		//postFile(filename, target_url)
		postFile(os.Args[2:], target_url, "add")

	case "ls":
		fmt.Println("List is called with files :")
		listFile(target_url_ls, "listfile")

	case "rm":
		fmt.Println("Rm is called with files :", os.Args[2:])
		removeFile(os.Args[2:], target_url_rm)

	case "update":
		fmt.Println("Update is called with files :", os.Args[2:])
		postFile(os.Args[2:], target_url, "update")

	case "wc":
		fmt.Println("Wc is called with files :", os.Args[2:])
		listFile(target_url_wc, "wordcount")

	default:
		fmt.Println("No option is given")
		help()
	}

}

func help() {

	printdata := `
Tool Excecution commands supported

<client binary> add <file1> <file2> ...
<client binary> ls
<client binary> rm <file1> <file2> ...
<client binary> update <file1> <file2> ...
<client binary> wc
 `

	fmt.Println(printdata)
	os.Exit(1)

}
