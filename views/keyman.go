package views

import (
	"fmt"
	"math"
)

func Test() {
	fmt.Println(calculateInnerModel(200))
}

//	@Description: innerValue的计算模型
//	@param addLines 代码新增行数
//	@return	f(addLines) = (zoomY∗(−1∗arctan((𝑎𝑑𝑑−translation)/zoomX)+π/2)+adjust) 模型的结果
//	@author Halokk 2022-08-12 14:25:46
func calculateInnerModel(addLines int) float64 {
	//  translation 横向平移单位数 ,zoomX 横向缩放比例 ,zoomY 纵向缩放比例 ,adjust 调节变量
	translation, zoomX, zoomY, adjust := 200, 100, 0.3, 0.1966
	return (math.Atan(float64((addLines-translation)/zoomX))*(-1)+math.Pi/2)*zoomY + adjust
}

//	@Description: 根据代码行数变更和旧的置信度来计算innerValue
//	@param oldConfidence 旧的置信度
//	@return add 新增的行数
//  @param new 当前的行数
//  @param delete 删除的行数
//  @param old 原本的行数
//	@author Halokk 2022-08-12 15:08:03
func calculateInnerValue(oldConfidence float64, add, new, delete, old int) (innerValue float64) {
	oldPart := (1.0 - float64(add)/float64(new)) * (1.0 - float64(delete)/float64(old)) * oldConfidence
	newPart := float64(add) / float64(new) * calculateInnerModel(add)
	innerValue = oldPart + newPart
	return
}

//	@Description: 根据定义链计算信息熵
//	@param
//	@return comentropy 信息熵
//	@author Halokk 2022-08-12 16:09:29
func calculateComentropy() (comentropy float64) {

	return
}

//	@Description: 当函数发生变更时，根据innerValue和comentropy计算置信度
//  @param innerValue
//  @param comentropy 信息熵
//	@return	confidence 置信度
//	@author Halokk 2022-08-12 14:42:25
func calculateConfidence(innerValue, comentropy float64) float64 {
	return math.Pow(innerValue, comentropy)
}

//	@Description: 当函数没有发生变更时，置信度应增加
//  @param oldConfidence 旧的置信度
//  @return confidence 置信度
//	@author Halokk 2022-08-12 16:24:15
func hightenConfidence(oldConfidence float64) float64 {
	return (1.2349 - math.Pow(0.2, oldConfidence-0.1))
}

//	@Description: 根据置信度、出现在堆栈中的频率、与直接错误函数的距离计算对本次错误的贡献
//  @param confidence 置信度
//  @param frameNumbers 出现在堆栈中的次数
//  @param totalNumbers 堆栈中总数
//  @param relevanceDistance 与直接错误函数的距离
//  @param comtribution 对本次错误的贡献
//	@author Halokk 2022-08-12 16:31:46
func calculateComtribution(confidence float64, frameNumbers, totalNumbers, relevanceDistance int) float64 {
	return (1.0 / confidence) * (float64(frameNumbers) / float64(totalNumbers)) * (1.0 / float64(relevanceDistance))
}

//	@Description: 根据每次commit函数改动的比例以及迭代次序赋予责任人权重
//  @param objectInfo 函数的结构体
//  @param relevanceDistance 与直接错误函数的距离
//  @return	[author-commitTime]weight
//	@author Halokk 2022-08-12 17:37:36
func getBugOwner(objectId string) (bugOwners map[string]float64) {
	historys := getHistory(objectId)
	for _, history := range historys {
		commit, _ := history.commitHistory, history.objectHistory
		owner := commit.commitAuthor + "-" + commit.commitTime
		weight := 0.0
		for _, notChangedPart := range historys {
			if notChangedPart.commitHistory.commitHash != history.commitHistory.commitHash {
				weight *= (1.0 - float64(notChangedPart.objectHistory.addedLineCount)/float64(notChangedPart.objectHistory.newlineCount))
			} else {
				weight *= float64(notChangedPart.objectHistory.addedLineCount) / float64(notChangedPart.objectHistory.newlineCount)
				break
			}
		}
		bugOwners[owner] = weight
	}

	return
}
