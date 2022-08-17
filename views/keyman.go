package views

import (
	"math"
)

// GetBugOrigin
// @Description: 根据崩溃堆栈解析出来的objects切片进行计算，返回出错率前五高的函数及每个函数前五名责任人
// @param  objects 解析出来的堆栈函数切片
// @return  bugOrigin 本次错误的主要责任人
// @author  Halokk
func GetBugOrigin(objects []ObjectInfo) (bugOringin []bugOriginInfo) {
	var methods []ObjectInfo
	var frameNumber []int
	var relevanceDistance []int

	//  筛选出堆栈中每个函数并统计出现在堆栈中的次数
	for dist, object := range objects {
		flag, index := false, 0
		for i, method := range methods {
			if object.objectId == method.objectId {
				flag = true
				index = i
				break
			}
		}

		if flag {
			frameNumber[index]++
		} else {
			methods = append(methods, object)
			frameNumber = append(frameNumber, 1)
			relevanceDistance = append(relevanceDistance, dist+1)
		}
	}

	//  计算出错率
	for i, method := range methods {
		var bugMethod bugOriginInfo
		bugMethod.object = method
		bugMethod.wrongRate = CalculateComtribution(method.confidence, frameNumber[i], len(objects), relevanceDistance[i])
		bugMethod.owners = CalculateOwnerWeight(method.objectId)
		bugOringin = append(bugOringin, bugMethod)
		if len(bugOringin) == 5 {
			break
		}
	}

	return bugOringin
}

//	@Description: innerValue的计算模型
//	@param addLines 代码新增行数
//	@return	f(addLines) = (zoomY∗(−1∗arctan((𝑎𝑑𝑑−translation)/zoomX)+π/2)+adjust) 模型的结果
//	@author Halokk 2022-08-12 14:25:46
func CalculateInnerModel(addLines int) float64 {
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
func CalculateInnerValue(oldConfidence float64, add, new, delete, old int) (innerValue float64) {
	oldPart := (1.0 - float64(add)/float64(new)) * (1.0 - float64(delete)/float64(old)) * oldConfidence
	newPart := float64(add) / float64(new) * CalculateInnerModel(add)
	innerValue = oldPart + newPart
	return
}

//	@Description: 根据定义链计算信息熵
//	@param	objectId
//	@return comentropy 信息熵
//	@author Halokk 2022-08-12 16:09:29
func CalculateComentropy(objectId string) (comentropy float64) {
	node := getChain(objectId)
	comentropy = math.Log(math.E + float64(len(node.childs)))
	if len(node.childs) != 0 {
		average := 0.0
		for _, object := range node.childs {
			average += CalculateComentropy(object.objectId)
		}
		average /= float64(len(node.childs))
		comentropy *= average
	}
	return
}

//	@Description: 当函数发生变更时，根据innerValue和comentropy计算置信度
//  @param innerValue
//  @param comentropy 信息熵
//	@return	confidence 置信度
//	@author Halokk 2022-08-12 14:42:25
func CalculateConfidence(object UncalculateObjectInfo, oldConfidence float64) float64 {
	innerValue := CalculateInnerValue(oldConfidence, object.addedLineCount, object.newlineCount,
		object.deletedlineCount, object.oldlineCount)
	comentropy := CalculateComentropy(object.objectId)
	return math.Pow(innerValue, comentropy)
}

//	@Description: 当函数没有发生变更时，增加其置信度
//  @param oldConfidence 旧的置信度
//  @return confidence 置信度
//	@author Halokk 2022-08-12 16:24:15
func HightenConfidence(oldConfidence float64) float64 {
	return 1.2349 - math.Pow(0.2, oldConfidence-0.1)
}

//	@Description: 根据置信度、出现在堆栈中的频率、与直接错误函数的距离计算对本次错误的贡献
//  @param confidence 置信度
//  @param frameNumbers 出现在堆栈中的次数
//  @param totalNumbers 堆栈中总数
//  @param relevanceDistance 与直接错误函数的距离
//  @param comtribution 对本次错误的贡献
//	@author Halokk 2022-08-12 16:31:46
func CalculateComtribution(confidence float64, frameNumbers, totalNumbers, relevanceDistance int) float64 {
	return (1.0 / confidence) * (float64(frameNumbers) / float64(totalNumbers)) * (1.0 / float64(relevanceDistance))
}

//	@Description: 根据每次commit函数改动的比例以及迭代次序赋予责任人权重
//  @param objectId  函数ID
//  @return	[author]weight
//	@author Halokk 2022-08-12 17:37:36
func CalculateOwnerWeight(objectId string) map[string]float64 {
	bugOwners := make(map[string]float64, 0)
	historys := getHistory(objectId)
	for _, history := range historys {
		commit, _ := history.commitHistory, history.objectHistory
		owner := commit.commitAuthor
		weight := 1.0
		for _, notChangedPart := range historys {
			if notChangedPart.commitHistory.commitHash != history.commitHistory.commitHash {
				weight *= float64((notChangedPart.objectHistory.addedLineCount)+notChangedPart.objectHistory.deletedlineCount) /
					float64(notChangedPart.objectHistory.newlineCount+notChangedPart.objectHistory.deletedlineCount)
			} else {
				weight *= float64(notChangedPart.objectHistory.addedLineCount) /
					float64(notChangedPart.objectHistory.newlineCount)
				break
			}
		}
		if _, ok := bugOwners[owner]; !ok {
			bugOwners[owner] = 0
		}
		bugOwners[owner] += weight
		if len(bugOwners) == 5 {
			break
		}
	}
	return bugOwners
}
