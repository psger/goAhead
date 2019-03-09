package main

import (
	"fmt"
	"github.com/vladimirvivien/gowfs"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	//path := GetCurPath()
	//fmt.Println(path)
	fs, err := gowfs.NewFileSystem(gowfs.Configuration{Addr: "192.168.0.109:50070", User: "root"})
	if err != nil{
		log.Fatal(err)
	}

	//delete
	//ok, err := fs.Delete(gowfs.Path{Name:"hello.go"}, false)
	//fmt.Println(ok)

	//checksum, err := fs.GetFileChecksum(gowfs.Path{Name: "/test/hello.txt"})
	//	//if err != nil {
	//	//	log.Fatal(err)
	//	//}
	//fmt.Println (checksum)

	// 上传
	var r io.Reader
	r, err = os.Open("/Users/passengerray/Downloads/V2RayX.app")
	if err != nil {
		fmt.Println(err)
	}
	ok, err := fs.Create(
		//bytes.NewBufferString("121313121212"),
		r,
		gowfs.Path{Name:"v2ray.app"},
		false,
		0,
		0,
		0700,
		0,
		"audio/aiff",
	)
	if err != nil || !ok {
			log.Fatal("",  ":", err)
		} else {
		log.Println ("input file created.")
	}



	//createTestDir(fs, "/hello1.txt")

	//下载
	//data, err := fs.Open(gowfs.Path{Name:"hello.go"}, 0, 512, 2048)
	//rcvdData, _ := ioutil.ReadAll(data)
	//fmt.Println(string(rcvdData))
}

func GetCurPath() string {
	file, _ := exec.LookPath(os.Args[0])

	//得到全路径，比如在windows下E:\\golang\\test\\a.exe
	path, _ := filepath.Abs(file)

	rst := filepath.Dir(path)

	return rst
}

func createTestDir(fs *gowfs.FileSystem, hdfsPath string) {
	path := gowfs.Path{Name:hdfsPath}
	ok, err := fs.MkDirs(path, 0744)
	if err != nil || !ok {
		log.Fatal("Unable to create test directory ", hdfsPath, ":", err)
	}
	log.Println ("HDFS Path ", path.Name, " created.")
}
