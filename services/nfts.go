package services

import (
	db "github.com/Touch/datasource"
	"github.com/Touch/models"
)

type NFTInfo struct {
	Personality string `json:"personality"` // 个性
	Level       string `json:"level"`       // level
	Name        string `json:"name"`        // name
	URL         string `json:"url"`         // url
}

func GetNftInfo(ctx *db.Context) ([]NFTInfo, error) {
	ctx.Logger().Info("获取NFT数据信息")
	// 查找指定用户的nft

	nfts, err := models.ListNFT(ctx, &models.NFTCondition{})
	if err != nil {
		return nil, err
	}

	var nftInfos []NFTInfo
	for _, v := range nfts {
		nftInfo := NFTInfo{
			Personality: v.Personality,
			Level:       v.Level,
			Name:        v.Name,
			URL:         v.URL,
		}

		nftInfos = append(nftInfos, nftInfo)
	}

	return nftInfos, nil
}

//func GetNftByPersonality(ctx *db.Context) (*NFTnfo, error) {
//	ctx.Logger().Info("获取Personality信息")
//	// 根据用户当前信息，查找到用户自己的nft信息
//	// 用户表
//	// nft表
//	// 查找当前personality的nft
//
//}
