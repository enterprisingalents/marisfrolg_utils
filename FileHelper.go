package Marisfrolg_utils

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/NetEase-Object-Storage/nos-golang-sdk/config"
	"github.com/NetEase-Object-Storage/nos-golang-sdk/logger"
	"github.com/NetEase-Object-Storage/nos-golang-sdk/model"
	"github.com/NetEase-Object-Storage/nos-golang-sdk/nosclient"
	"github.com/NetEase-Object-Storage/nos-golang-sdk/nosconst"
)

//判断目录是否存在
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}

//是否创建目录
func IsCreateDir(path string) (err error) {
	if !IsDirExists(path) {
		err = os.MkdirAll(path, 0766)
	}
	return err
}

/*
创建人：李奇峰
功能：上传文件到网易Nos对象云存储
nosContentType:上传的文件类型，留空则不指定
*/
func UploadFileToNos(fileName string, filePath string, bucketName string, nosContentType string) (ossPath string, err error) {
	if len(bucketName) == 0 {
		bucketName = "ODSAPP"
	}
	conf := &config.Config{
		Endpoint:  "nos-eastchina1.126.net",
		AccessKey: "73e280baf79f401abce7e37f9638e67a",
		SecretKey: "20eb53b14bf54549ba560048903a1457",

		NosServiceConnectTimeout:    500,
		NosServiceReadWriteTimeout:  500,
		NosServiceMaxIdleConnection: 500,

		LogLevel: logger.LogLevel(logger.DEBUG),
		Logger:   logger.NewDefaultLogger(),
	}
	metadata := &model.ObjectMetadata{
		Metadata: map[string]string{
			nosconst.CONTENT_TYPE: nosContentType,
			//  nosconst.CONTENT_MD5:      "OBJECTMD5(文件md5值)",
		},
	}

	putObjectRequest := &model.PutObjectRequest{
		Bucket:   "marisfrolg",
		Object:   bucketName + "/" + fileName,
		FilePath: filePath,
		Metadata: metadata,
	}

	nosClient, err := nosclient.New(conf)

	_, err = nosClient.PutObjectByFile(putObjectRequest)
	if err != nil {
		log.Println(err.Error())
		return "", err
	} else {
		//return  "https://marisfrolg.nos-eastchina1.126.net/ODSAPP/"+fileName+"?imageView&thumbnail=100y100",err
		return "https://marisfrolg.nos-eastchina1.126.net/" + bucketName + "/" + fileName, err
	}

}

func SaveFileToTempDirectory(isNeedPrefix bool, file *multipart.FileHeader) (fileName, filePath string, err error) {
	var (
		dir       string
		existsDir bool
		out       *os.File
		reader    multipart.File
	)
	dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	filePath = strings.Replace(dir, "\\", "/", -1) + "/UploadFile"
	existsDir, _ = PathExists(filePath)
	if !existsDir {
		os.Mkdir(filePath, os.ModePerm)
	}
	if isNeedPrefix {
		fileName = time.Now().Format("20060102150405") + "_" + file.Filename
	} else {
		fileName = file.Filename
	}

	filePath = filePath + "/" + fileName

	if out, err = os.Create(filePath); err != nil {
		goto ERR
	}
	defer out.Close()
	if reader, err = file.Open(); err != nil {
		goto ERR
	}
	defer reader.Close()
	_, err = io.Copy(out, reader)
	return
ERR:
	return
}

/*
  创建人：李奇峰
  功能：读取文件反序列化
*/
func LoadFile(filename string, v interface{}) (err error) {
	var data []byte
	if data, err = ioutil.ReadFile(filename); err != nil {
		return
	}

	if err = json.Unmarshal(data, v); err != nil {
		return
	}
	return
}
