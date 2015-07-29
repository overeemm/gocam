package main

import (
	"fmt"
	"bytes"
	"time"
	"io"	
	"io/ioutil"
	"os/exec"
	"github.com/stacktic/dropbox"
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

	var db *dropbox.Dropbox

	db = dropbox.NewDropbox()

	db.SetAppInfo(clientid, clientsecret)
	db.SetAccessToken(token)

	//for i := 0; i < 2500; i++ {
		shoot(db)
	//	time.Sleep(5 * time.Minute)
	//}
}
