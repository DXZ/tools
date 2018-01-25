package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type SendOcrData struct {
	Reqid   string     `json:"reqid"`
	Reqtime string     `json:"reqtime"`
	Images  []OcrImage `json:"images"`
}

// func (img *ExpressionImage) Encode() (buf *bytes.Buffer, err error) {
//  buf = new(bytes.Buffer)
//  if img.Format == "png" {
//      err = png.Encode(buf, img.Image)
//      // util.CheckErr(err)
//  } else if img.Format == "jpg" || img.Format == "jpeg" {
//      err = jpeg.Encode(buf, img.Image, nil)
//  } else {
//      err = errors.New("bad format express image")
//  }
//  return
// }

type OcrImage struct {
	Id        int    `json:"id"`
	Imgbase64 string `json:"imgbase64"`
}

func RequestOcrV2(senddata *SendOcrData) (string, error) {
	body_type := "application/json;charset=utf-8"
	tmp, err := json.Marshal(senddata)
	fmt.Println(string(tmp))
	if err != nil {
		return "", err
	}
	req := bytes.NewReader(tmp)

	c := &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
	resp, err := c.Post("http://ocr.snowcloud.ai:18080/serving/api/v1.0/images", body_type, req)
	if err != nil {
		return "", err
	}
	resp_body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	jsonstr := bytes.NewBuffer(resp_body).String()
	return jsonstr, nil
}

func main() {
	senddata := new(SendOcrData)
	senddata.Reqid = "test_seven"
	layout := "20060102150405"
	senddata.Reqtime = time.Now().Format(layout)
	var image_tmp OcrImage
	for i := 0; i < 2; i++ {
		ff, _ := os.Open("a.jpg")
		defer ff.Close()

		img, format, err := image.Decode(ff)
		buf := new(bytes.Buffer)
		err = jpeg.Encode(buf, img, nil)

		fmt.Println(format, err)
		// sourcebuffer := make([]byte, 50000000)
		// n, _ := ff.Read(sourcebuffer)
		// dist := make([]byte, 2*len(buf.Bytes()))
		// base64.StdEncoding.Encode(dist, buf.Bytes())
		//base64压缩
		// sourcestring := string(dist[:len(buf.Bytes())])
		// sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
		sourcestring := base64.StdEncoding.EncodeToString(buf.Bytes()[:])
		// dist := make([]byte, 2*len(buf.Bytes()))
		// base64.StdEncoding.Encode(dist, buf.Bytes())
		// sourcestring := base64.StdEncoding.EncodeToString(buf.Bytes())
		image_tmp.Id = i
		// epimgs_pos_dict[epimg.ImgNo] = make([]int, 4)
		image_tmp.Imgbase64 = sourcestring
		senddata.Images = append(senddata.Images, image_tmp)
	}

	jsonstr, err := RequestOcrV2(senddata)
	fmt.Println(jsonstr, err)
}
