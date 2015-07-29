package main

import (
	"fmt"
	"bytes"
	"time"
	"io"	
	"io/ioutil"
	"os/exec"
	"github.com/stacktic/dropbox"
	"github.com/stianeikeland/go-rpio"	
)

func NopReadCloser(x io.Reader) io.ReadCloser {
	return &nopReadCloser{x}
}

type nopReadCloser struct {
	io.Reader
}

func (x *nopReadCloser) Close() error {
	return nil
}

func toggleLight(pinNr int) {
	pin := rpio.Pin(pinNr)
	pin.Output()
	pin.High()
	time.Sleep(500 * time.Millisecond)
	pin.Low()
}

func shoot(db *dropbox.Dropbox) {

        cmd := exec.Command("raspistill", "-n", "-o", "-")
        out, err := cmd.StdoutPipe()
        if(err != nil) {
		fmt.Println("cmd.StdoutPipe()")
                fmt.Println(err)
        }       
        err = cmd.Start()
        if(err != nil) {
		fmt.Println("cmd.Start()")
                fmt.Println(err)
        }
	data, err := ioutil.ReadAll(out)
        if(err != nil) {
		fmt.Println("ioutil.ReadAll(out)")
		fmt.Println(err)
	}
	
	cmd.Wait()
	
	now := time.Now().Format("2006-01-02 150405")
	len := int64(len(data))
	bytes := NopReadCloser(bytes.NewReader(data))
	
	_, err = db.FilesPut(bytes, len, fmt.Sprintf("%s.jpg", now) , true, "");
	if(err != nil) {
		fmt.Println("db.FilesPut(...)")
		fmt.Println(err)
	}

}

func main() {

	err := rpio.Open()
	if(err != nil) {
		fmt.Println("rpio.Open()")
                fmt.Println(err)
        }      
	// Unmap gpio memory when done
	defer rpio.Close()

	toggleLight(9)

	var db *dropbox.Dropbox

	db = dropbox.NewDropbox()

	db.SetAppInfo(clientid, clientsecret)
	db.SetAccessToken(token)

	toggleLight(10)

	shoot(db)

	toggleLight(11)
}
