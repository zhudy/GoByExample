package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)
func check(err error){
	if err != nil{
		panic(err)
	}
}
func main() {
	f, err := ioutil.TempFile("", "sample")
	check(err)
	fmt.Println("temp file name: ", f.Name())
	defer os.Remove(f.Name())

	_, err = f.Write([]byte{1,2,3,4})
	check(err)
	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("temp dir name:", dname)
	defer os.RemoveAll(dname)
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1,2}, 0666)
	check(err)
}