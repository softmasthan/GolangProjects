package main

import (
	"log"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	//Referred site:- https://github.com/jlaffaye/ftp
	//Connecting Ftp Details with username and password
	c, err := ftp.Dial("ipAdress:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("username", "password")
	if err != nil {
		log.Fatal(err)
	}

	//Give local path drive where ftp.zip fiel is present
	sourceFile, err := os.Open("F:\\ftp.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	// Do something with the FTP conn
	//Remote ftp.zip file has going to move in server
	err = c.Stor("ftp.zip", sourceFile)
	if err != nil {
		panic(err)
	}

	if err := c.Quit(); err != nil {
		log.Fatal(err)
	}
}
