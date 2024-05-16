package models

import (
	"encoding/json"
	db "github.com/Touch/datasource"
	"os"
)

// nft表
type NFT struct {
	Personality string `json:"personality"` // 个性
	Level       string `json:"level"`       // level
	Name        string `json:"name"`        // name
	URL         string `json:"url"`         // url
}

type NFTCondition struct {
	Personality string `json:"personality"` // 个性
}

func ListNFT(ctx *db.Context, condition *NFTCondition) ([]NFT, error) {
	ctx.Logger().Info("通过NFT获取nft")

	// 读取文件中的数据
	dir, _ := os.Getwd()
	data, err := os.ReadFile(dir + "/conf/data.json")
	if err != nil {
		return nil, err
	}

	// 文件unmarshal
	var nfts []NFT
	err = json.Unmarshal(data, &nfts)
	if err != nil {
		return nil, err
	}

	return nfts, nil
}
