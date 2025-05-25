package id

import (
	"fmt"
	"math/rand"
	"powergame.com/common/pkg/enum"
	"strconv"
	"time"
)

func idLast4(id int64) string {
	return fmt.Sprintf("%04d", id%10000)
}

// 左侧填充函数
func leftPad(input string, length int, padChar rune) string {
	for len(input) < length {
		input = string(padChar) + input
	}
	return input
}

// 生成固定长度的随机数
func generatePaddedRandomNumber(length int) string {
	// 使用当前时间戳作为种子
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomNumber := 0
	if length == 5 {
		randomNumber = r.Intn(1e5)
	} else if length == 6 {
		randomNumber = r.Intn(1e6)
	} else if length == 8 {
		randomNumber = r.Intn(1e8)
	} else {
		panic("Invalid length")
	}
	randomStr := strconv.Itoa(randomNumber)
	return leftPad(randomStr, length, '0')
}

// GetMerchantId 商户ID（1+商户类型1位+8位随机数）
func GetMerchantId(merchantType enum.MerchantType) int64 {
	idStr := fmt.Sprintf("%d%d%s", enum.MerchantIdPrefix, merchantType.Index(), generatePaddedRandomNumber(8))
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}

// GetAgentId 代理ID（2+商户ID后4位+代理类型1位+6位随机数）num % 10000 的作用是提取 num 的后 4 位数字
func GetAgentId(merchantId int64, agentType enum.AgentType) int64 {
	idStr := fmt.Sprintf("%d%s%d%s", enum.AgentIdPrefix, idLast4(merchantId), agentType.Index(), generatePaddedRandomNumber(6))
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}

// GetMemberId 会员ID（3+商家ID后4位+会员类型1位+8位随机数）
func GetMemberId(merchantId int64, memberType enum.MemberType) int64 {
	idStr := fmt.Sprintf("%d%s%d%s", enum.MemberIdPrefix, idLast4(merchantId), memberType.Index(), generatePaddedRandomNumber(8))
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}

// GetWalletId 钱包ID（8+借贷标识1位+业务标识2位+用户类型1位+用户ID后4位+5位随机数）
func GetWalletId(dc enum.DcType, bizType int, ownerType enum.OwnerType, ownerId int64) int64 {
	bizTypeStr := fmt.Sprintf("%02d", bizType)
	idStr := fmt.Sprintf("%d%d%s%d%s%s", enum.WalletIdPrefix, dc.Index(), bizTypeStr, ownerType.Index(), idLast4(ownerId), generatePaddedRandomNumber(5))
	id, _ := strconv.ParseInt(idStr, 10, 64)
	return id
}

// GetOrderId 订单ID（系统标识2位+业务标识2位+YYDDMM+5位时间秒数+5位随机数）
func GetOrderId(systemType enum.SystemType, bizType enum.BizType) string {
	now := time.Now()
	date := now.Format("060102")
	secondOfDay := now.Hour()*60*60 + now.Minute()*60 + now.Second()
	idStr := fmt.Sprintf("%d%d%s%d%s", systemType.Index(), bizType.Index(), date, secondOfDay, generatePaddedRandomNumber(5))
	return idStr
}
