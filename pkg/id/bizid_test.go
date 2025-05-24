package id

import (
	"fmt"
	"powergame.com/powergame/pkg/enum"
	"testing"
)

func TestGetMerchantId(t *testing.T) {
	type args struct {
		merchantType enum.MerchantType
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "case01", args: args{
			merchantType: enum.MerchantSelf,
		}, want: "10"},
		{name: "case02", args: args{
			merchantType: enum.MerchantThird,
		}, want: "11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMerchantId(tt.args.merchantType)
			fmt.Println(got)
		})
	}
}

func TestGetAgentId(t *testing.T) {
	type args struct {
		merchantId int64
		agentType  enum.AgentType
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "case01", args: args{
			merchantId: 10000001,
			agentType:  enum.AgentDirect,
		}},
		{name: "case02", args: args{
			merchantId: 10000002,
			agentType:  enum.AgentSub,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAgentId(tt.args.merchantId, tt.args.agentType)
			fmt.Println(got)
		})
	}
}

func TestGetMemberId(t *testing.T) {
	type args struct {
		merchantId int64
		memberType enum.MemberType
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "case01", args: args{
			merchantId: 10000001,
			memberType: enum.MemberDirect,
		}},
		{name: "case02", args: args{
			merchantId: 10000001,
			memberType: enum.MemberSub,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetMemberId(tt.args.merchantId, tt.args.memberType)
			fmt.Println(got)
		})
	}
}

func TestGetWalletId(t *testing.T) {
	type args struct {
		dc        enum.DcType
		bizType   int
		ownerType enum.OwnerType
		ownerId   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "case01", args: args{
			dc: enum.Credit, bizType: 1,
			ownerType: enum.OwnerMerchant, ownerId: 10000001},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetWalletId(tt.args.dc, tt.args.bizType, tt.args.ownerType, tt.args.ownerId)
			fmt.Println(got)
		})
	}
}
func TestGetOrderId(t *testing.T) {
	type args struct {
		systemType enum.SystemType
		bizType    enum.BizType
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "case01",
			args: args{
				systemType: enum.SystemOrder,
				bizType:    enum.BizPayOrder,
			},
		},
		{
			name: "case02",
			args: args{
				systemType: enum.SystemPay,
				bizType:    enum.BizPayOrder,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetOrderId(tt.args.systemType, tt.args.bizType)
			fmt.Println(got)
		})
	}
}
