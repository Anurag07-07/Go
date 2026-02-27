package main

import (
	// "fmt"
	// "bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("a.txt")
	// if err!=nil {
	// 	//Log the Error
	// 	panic(err)
	// }
	// //File Info
	// fileInfo,err := f.Stat()

	// if err!=nil {
	// 	//Log the Error
	// 	panic(err)
	// }

	// fmt.Println("file name: ",fileInfo.Name())
	// fmt.Println("file is dir or not: ",fileInfo.IsDir())
	// fmt.Println("file name: ",fileInfo.Size())

	//Read the File
	// f,err := os.Open("a.txt")
	// if err!=nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte,10)
	
	// d,err:=f.Read(buf)

	// if err!=nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	fmt.Println("data",d,string(buf[i]))
	// }


	// data,err := os.ReadFile("a.txt")
	// if err!=nil {
	// 	panic(err)
	// }

	// fmt.Println(string(data))

	// dir,err:=os.Open("../")
	// if err!=nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo,err:=dir.ReadDir(5)
	// // fileInfo,err:=dir.ReadDir(-1) //for multiple file

	// for _,fi := range fileInfo{
	// 	fmt.Println(fi.Name())
	// }

	// f,err:=os.Create("a2.txt")
	
	// if err!=nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("hi go")

	// bytes:=[]byte("Hello Golang")
	// f.Write(bytes)


	// sourceFile,err:=os.OpenFile("example.txt")
	// if err!=nil {
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile,err :=os.Create("example.txt")
	// if err!=nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader:=bufio.NewReader(sourceFile)
	// writer:=bufio.NewWriter(destFile)


	// for{
	// 	b,err:=reader.ReadByte()
	// 	if err!=nil {
	// 		if err.Error()!="EOF"{
	// 			panic(err)
	// 		}
	// 		break
	// 	}

	// 	e:=writer.WriteByte(b)
	// 	if e!=nil{
	// 		panic(e)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("Written to new file sucessfully")

	err:=os.Remove("a2.txt")
	if err!=nil {
		panic(err)
	}

	fmt.Println("file deleted successfully")
}