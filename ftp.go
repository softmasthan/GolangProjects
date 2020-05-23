package main

import (
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	//Referred site:- https://github.com/jlaffaye/ftp
	
	//Initializing IP Address with port 21 Ex:- 127.0.0.1:21
	c, err := ftp.Dial("ipAdress:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	
	//Connecting Ftp Details with username and password
	err = c.Login("username", "password")
	if err != nil {
		log.Fatal(err)
	}

	//Give full local path of a drive, where your file is present  Ex: F:\\ftp.zip
	sourceFile, err := os.Open("F:\\ftp.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Do something with the FTP connection
	//Moving a ftp.zip file from local PC to Remote Server or some other FTP connected system
	err = c.Stor("ftp.zip", sourceFile)
	if err != nil {
		panic(err)
	}

	//Closing the connection of a ftp
	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
