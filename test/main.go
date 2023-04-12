package main

import (
	metacli "github.com/FogMeta/go-mc-sdk/client"
	"github.com/filswan/go-swan-lib/logs"
	"os"
)

func main() {
	// Swan API key. Acquire from [Swan Platform](https://console.filswan.com/#/dashboard) -> "My Profile"->"Developer Settings". It can be ignored if `[sender].offline_swan=true`.
	key := "V0schjjl_bxCtSNwBYXXXX"
	// Swan API access token. Acquire from [Swan Platform](https://console.filswan.com/#/dashboard) -> "My Profile"->"Developer Settings". It can be ignored if `[sender].offline_swan=true`.
	token := "fca72014744019a949248874610fXXXX"
	metaUrl := "http://{ip}:8099/rpc/v0"
	metaClient := metacli.NewAPIClient(key, token, metaUrl)

	apiUrl := "http://127.0.0.1:5001"
	inputPath := "./testdata"
	ipfsCid, err := metaClient.UploadFile(apiUrl, inputPath)
	if err != nil {
		logs.GetLogger().Error("upload failed:", err)
		return
	}
	logs.GetLogger().Infoln("upload success, and data cid is: ", ipfsCid)

	datasetName := "dataset-name"
	ipfsGateway := "http://127.0.0.1:8080"
	sourceName := "./testdata"
	ipfsCid = "QmQgM2tGEduvYmgYy54jZaZ9D7qtsNETcog8EHR8XoeyEp"
	info, err := os.Stat(sourceName)
	if err != nil {
		logs.GetLogger().Error("get source file stat information error:", err)
		return
	}
	oneItem := metacli.IpfsData{}
	oneItem.SourceName = sourceName
	oneItem.IpfsCid = ipfsCid
	oneItem.DataSize = info.Size()
	oneItem.IsDirectory = info.IsDir()
	oneItem.DownloadUrl = metacli.PathJoin(ipfsGateway, "ipfs/", ipfsCid)
	ipfsData := []metacli.IpfsData{oneItem}
	err = metaClient.ReportMetaClientServer(datasetName, ipfsData)

	if err != nil {
		logs.GetLogger().Error("report meta client server  failed:", err)
		return
	}
	logs.GetLogger().Infoln("report meta client server success")

	ipfsCid = "QmQgM2tGEduvYmgYy54jZaZ9D7qtsNETcog8EHR8XoeyEp"
	outPath := "./output"
	downloadUrl := "http://127.0.0.1:8080/ipfs/QmQgM2tGEduvYmgYy54jZaZ9D7qtsNETcog8EHR8XoeyEp"
	host := "127.0.0.1"
	port := 6800
	secret := "my_aria2_secret"
	conf := &metacli.Aria2Conf{Host: host, Port: port, Secret: secret}
	err = metaClient.DownloadFile(ipfsCid, outPath, downloadUrl, conf)
	if err != nil {
		logs.GetLogger().Error("download failed:", err)
		return
	}
	logs.GetLogger().Infoln("download success")

	fileName := "testdata"
	ipfsCids, err := metaClient.GetIpfsCidByName(fileName)
	if err != nil {
		logs.GetLogger().Error("get ipfs cid failed:", err)
		return
	}
	logs.GetLogger().Infof("get ipfs cid success: %+v", ipfsCids)

	pageNum := 0
	limit := 10
	// sourceFileList, err := metaClient.GetFileLists(pageNum, limit, metacli.WithShowCar(true))
	sourceFileList, err := metaClient.GetFileLists(pageNum, limit) //default show CAR option is false
	if err != nil {
		logs.GetLogger().Error("get file list failed:", err)
		return
	}
	logs.GetLogger().Infof("get file list success: %+v", sourceFileList)

	ipfsCid = "QmQgM2tGEduvYmgYy54jZaZ9D7qtsNETcog8EHR8XoeyEp"
	sourceFileInfo, err := metaClient.GetFileInfoByIpfsCid(ipfsCid)
	if err != nil {
		logs.GetLogger().Error("get source file info failed:", err)
		return
	}
	logs.GetLogger().Infof("get source file info success: %+v", sourceFileInfo)

	return
}
