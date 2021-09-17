package main

import (
	"fmt"
	"io"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"syscall"
	"unsafe"
)

// https://gist.github.com/NaniteFactory/0bd94e84bbe939cda7201374a0c261fd
func MessageBox(hwnd uintptr, caption, title string, flags uint) int {
	ret, _, _ := syscall.NewLazyDLL("user32.dll").NewProc("MessageBoxW").Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(caption))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(flags))
	return int(ret)
}

func MessageBoxPlain(title, caption string) int {
	const (
		NULL  = 0
		MB_OK = 0
	)
	return MessageBox(NULL, caption, title, MB_OK)
}

func main() {
	userprofile := os.Getenv("userprofile")
	buffer := fmt.Sprintf("[InternetShortcut]\nURL=fivem://connect/4rdm.cf\nIconIndex=0\nHotKey=0\nIDList=\nIconFile=%s\\4RDM\\4rdmlogo.ico\n[{000214A0-0000-0000-C000-000000000046}]\nProp3=19,0\n", userprofile)

	response, err := http.Get("https://4rdm.pl/assets/logo.ico")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	err = os.Mkdir(fmt.Sprintf("%s\\4RDM", userprofile), 0755)
	if err != nil {}

	out, err := os.Create(fmt.Sprintf("%s\\4RDM\\4rdmlogo.ico", userprofile))
	if err != nil {
		panic(err)
	}
	defer out.Close()

	io.Copy(out, response.Body)

	ioutil.WriteFile(fmt.Sprintf("%s\\Desktop\\4RDM.url", userprofile), []byte(buffer), fs.ModeAppend)

	MessageBoxPlain("Ukończono", "Skrót utworzony pomyślnie!")
}
