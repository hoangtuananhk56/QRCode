// go-qrcode
// Copyright 2014 Tom Harwood

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	qrcode "QRcodeLib/qrCodeLib"
	"strconv"
	"time"
)

var size = 256

// func random -> namefile

var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ0123456789"
var countAlpha = 62

// generates a random string of fixed size
func randNameFile(size int) string {
	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}

func createNameFileImg(typeQr string) string {
	var namefile string
	switch typeQr {
	case "URL":
		namefile = "MQR_URL_" + randNameFile(countAlpha)
	case "TEXT":
		namefile = "MQR_TEXT_" + randNameFile(countAlpha)
		//fb, pdf, vcard, youtube, sms, mms, ....
	}
	return namefile + ".png"
}
func createNewFolderImg() string {
	var err error
	t := time.Now().Local()
	x := t.Month()
	month := int(x)
	sYear := strconv.Itoa(t.Year())
	sMonth := strconv.Itoa(month)
	sDay := strconv.Itoa(t.Day())
	sHour := strconv.Itoa(t.Hour())
	sMinute := strconv.Itoa(t.Minute())
	// s_second := strconv.Itoa(t.Second())

	folder := []string{sYear, sMonth, sDay, sHour, sMinute /*, s_second*/}
	pathFolder := "../static/img/"
	var countFolder = 5
	for i := 0; i < countFolder; i++ {
		pathFolder += folder[i] + "/"
		os.Mkdir(pathFolder, os.ModeSticky|0755)
		if err != nil {
			log.Println("Error creating directory")
			log.Println(err)
		}
	}
	return pathFolder
}

func createImage(text string, typeQr string) {
	var png []byte
	var err error
	var qr *qrcode.QRCode
	qr, err = qrcode.New(text, qrcode.Highest)
	png, err = qr.PNG(size)
	createNewFolderImg()
	createNameFileImg(typeQr)
	fh, err := os.Create(createNewFolderImg() + createNameFileImg(typeQr))
	checkError(err)
	fh.Write(png)
	defer fh.Close()

}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(randNameFile(10))
	createImage("http://24h.com.vn", "URL")
	p := fmt.Println
	p(createNameFileImg("URL"))
}
