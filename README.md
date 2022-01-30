File Upload server and client
-----------------------------

How to test

Run Server in a 1 tab

go run fileuploadserver.go

Run Client in another tab

Different commands used.

 go run fileuploadclient.go add file2.txt 
 go run fileuploadclient.go add file2.txt 
 go run fileuploadclient.go update file2.txt 
 go run fileuploadclient.go wc
 go run fileuploadclient.go update file2.txt file2.txt file3.txt file4.txt
 go run fileuploadclient.go wc
 go run workingfileuploadclient.go rm file2.txt
 go run workingfileuploadclient.go ls


Need to complete :

Docker file unit testcases,creation and pod creation.

Tested logs are there in Tested_logs
