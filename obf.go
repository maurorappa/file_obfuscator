package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

	
var (
	 header bytes.Buffer
	 ext string
	 err error
)

func filechange(ext string) (written int) {
	written = 0
	switch ext {
        case "elf64": 
	//  ELF 64-bit LSB (SYSV)
		header.WriteByte(0x7f)
		header.WriteByte(0x45)
		header.WriteByte(0x4c)
		header.WriteByte(0x46)
		header.WriteByte(0x02)
		header.WriteByte(0x01)
		header.WriteByte(0x01)
		header.WriteByte(0)
	case "png":
	//PNG image data
		header.WriteByte(0x89)
		header.WriteByte(0x50)
		header.WriteByte(0x4e)
		header.WriteByte(0x47)
		header.WriteByte(0x0d)
		header.WriteByte(0x0a)
		header.WriteByte(0x1a)
		header.WriteByte(0x0a)
	case "jpg":
	//JPG image data
		header.WriteByte(0xff)
		header.WriteByte(0xd8)
		header.WriteByte(0xe0)
		header.WriteByte(0)

	case "ico":
	//JPG image data
		header.WriteByte(0)
		header.WriteByte(0)
		header.WriteByte(0x1)
		header.WriteByte(0)
		header.WriteByte(0x2)
		header.WriteByte(0)
		header.WriteByte(0x10)
		header.WriteByte(0x10)
		header.WriteByte(0)
		header.WriteByte(0)
		header.WriteByte(1)
		header.WriteByte(0)
		header.WriteByte(0x20)
		header.WriteByte(0)
		header.WriteByte(0x68)
		header.WriteByte(0x4)
		header.WriteByte(0)
		header.WriteByte(0)
		header.WriteByte(0x26)
		header.WriteByte(0)
		header.WriteByte(0)
		header.WriteByte(0)
		header.WriteByte(0x20)
		header.WriteByte(0x20)

	case "wav":
	//RIFF (little-endian) data, WAVE audio
		header.WriteByte(0x52)
		header.WriteByte(0x49)
		header.WriteByte(0x46)
		header.WriteByte(0x46)
		header.WriteByte(0xb0)
		header.WriteByte(0xca)
		header.WriteByte(0x03)
		header.WriteByte(0)
		header.WriteByte(0x57)
		header.WriteByte(0x41)
		header.WriteByte(0x56)
		header.WriteByte(0x45)
		header.WriteByte(0x66)
	default:
  		fmt.Println("Unrecognized extension, see the help")
		os.Exit(127)
	}
	return header.Len()
}

func main() {
        ext := flag.String("e","jpg","file extension to use as wrapper")
	action := flag.String("a", "c", "action: c=camoufage, r=reveal")
	filename := flag.String("f", "", "file to work on")
	flag.Usage = func() {
		fmt.Printf("This simple program allows you to wrap any file in a more innocuous file like wav, jpg, png or ico (the only 4 implemented)\nThere are two action defined:\n\t'c' means camoufage which will prepend the desired file's headers\n \t'r' means reveal, it strips the fake header and recreate the original file\nTo hide a file:\n\t./obf -a c -f cat -e png\nTo re-create the original file:\n\t./obf -a r -f cat.png\n")
		os.Exit(0)
	}
	flag.Parse()
	if *action == "c" {

		to, err := os.OpenFile(*filename+"."+*ext, os.O_RDWR|os.O_CREATE, 0666)
		if err != nil {
			log.Fatal(err)
		}
		_ = filechange(*ext)
		// write the new file header
		header.WriteTo(to)
		//copy the original file
		data, err := ioutil.ReadFile(*filename)
		_, err = to.Write(data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Done\n")
		to.Close()

	}
	if *action == "r" {
		file_parts := strings.Split(*filename,".")
		fmt.Printf("extension to be stripped: %s\n", file_parts[1])
		fake_header := int64(filechange(file_parts[1]))

		f, errm := os.Open(*filename)
		if errm != nil {
			log.Fatal(errm)
		}
		stat, _ := f.Stat()
		orig_size := stat.Size() - fake_header
		fmt.Printf("file size: %d, fake header: %d\n", stat.Size(), fake_header)
		_, _ = f.Seek(fake_header, 0)
		b3 := make([]byte, orig_size)
		_, _ = io.ReadAtLeast(f, b3, 2)
		orig_file := *filename + "_orig"
		newf, _ := os.OpenFile(orig_file, os.O_RDWR|os.O_CREATE, 0666)
		_, err := newf.Write(b3)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Done")	
		newf.Close()
	}
}
