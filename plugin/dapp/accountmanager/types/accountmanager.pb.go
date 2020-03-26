// Code generated by protoc-gen-go. DO NOT EDIT.
// source: accountmanager.proto

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	accountmanager.proto

It has these top-level messages:
	Accountmanager
	AccountmanagerAction
	Register
	ResetKey
	Apply
	Asset
	Transfer
	Supervise
	Account
	AccountReceipt
	ReplyAccountList
	TransferReceipt
	SuperviseReceipt
	QueryExpiredAccounts
	QueryAccountsByStatus
	QueryAccountByID
	QueryAccountByAddr
	QueryBalanceByID
	Balance
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Accountmanager struct {
}

func (m *Accountmanager) Reset()                    { *m = Accountmanager{} }
func (m *Accountmanager) String() string            { return proto.CompactTextString(m) }
func (*Accountmanager) ProtoMessage()               {}
func (*Accountmanager) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type AccountmanagerAction struct {
	// Types that are valid to be assigned to Value:
	//	*AccountmanagerAction_Register
	//	*AccountmanagerAction_ResetKey
	//	*AccountmanagerAction_Transfer
	//	*AccountmanagerAction_Supervise
	//	*AccountmanagerAction_Apply
	Value isAccountmanagerAction_Value `protobuf_oneof:"value"`
	Ty    int32                        `protobuf:"varint,6,opt,name=ty" json:"ty,omitempty"`
}

func (m *AccountmanagerAction) Reset()                    { *m = AccountmanagerAction{} }
func (m *AccountmanagerAction) String() string            { return proto.CompactTextString(m) }
func (*AccountmanagerAction) ProtoMessage()               {}
func (*AccountmanagerAction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isAccountmanagerAction_Value interface {
	isAccountmanagerAction_Value()
}

type AccountmanagerAction_Register struct {
	Register *Register `protobuf:"bytes,1,opt,name=register,oneof"`
}
type AccountmanagerAction_ResetKey struct {
	ResetKey *ResetKey `protobuf:"bytes,2,opt,name=resetKey,oneof"`
}
type AccountmanagerAction_Transfer struct {
	Transfer *Transfer `protobuf:"bytes,3,opt,name=transfer,oneof"`
}
type AccountmanagerAction_Supervise struct {
	Supervise *Supervise `protobuf:"bytes,4,opt,name=supervise,oneof"`
}
type AccountmanagerAction_Apply struct {
	Apply *Apply `protobuf:"bytes,5,opt,name=apply,oneof"`
}

func (*AccountmanagerAction_Register) isAccountmanagerAction_Value()  {}
func (*AccountmanagerAction_ResetKey) isAccountmanagerAction_Value()  {}
func (*AccountmanagerAction_Transfer) isAccountmanagerAction_Value()  {}
func (*AccountmanagerAction_Supervise) isAccountmanagerAction_Value() {}
func (*AccountmanagerAction_Apply) isAccountmanagerAction_Value()     {}

func (m *AccountmanagerAction) GetValue() isAccountmanagerAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AccountmanagerAction) GetRegister() *Register {
	if x, ok := m.GetValue().(*AccountmanagerAction_Register); ok {
		return x.Register
	}
	return nil
}

func (m *AccountmanagerAction) GetResetKey() *ResetKey {
	if x, ok := m.GetValue().(*AccountmanagerAction_ResetKey); ok {
		return x.ResetKey
	}
	return nil
}

func (m *AccountmanagerAction) GetTransfer() *Transfer {
	if x, ok := m.GetValue().(*AccountmanagerAction_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (m *AccountmanagerAction) GetSupervise() *Supervise {
	if x, ok := m.GetValue().(*AccountmanagerAction_Supervise); ok {
		return x.Supervise
	}
	return nil
}

func (m *AccountmanagerAction) GetApply() *Apply {
	if x, ok := m.GetValue().(*AccountmanagerAction_Apply); ok {
		return x.Apply
	}
	return nil
}

func (m *AccountmanagerAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AccountmanagerAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AccountmanagerAction_OneofMarshaler, _AccountmanagerAction_OneofUnmarshaler, _AccountmanagerAction_OneofSizer, []interface{}{
		(*AccountmanagerAction_Register)(nil),
		(*AccountmanagerAction_ResetKey)(nil),
		(*AccountmanagerAction_Transfer)(nil),
		(*AccountmanagerAction_Supervise)(nil),
		(*AccountmanagerAction_Apply)(nil),
	}
}

func _AccountmanagerAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AccountmanagerAction)
	// value
	switch x := m.Value.(type) {
	case *AccountmanagerAction_Register:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Register); err != nil {
			return err
		}
	case *AccountmanagerAction_ResetKey:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ResetKey); err != nil {
			return err
		}
	case *AccountmanagerAction_Transfer:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Transfer); err != nil {
			return err
		}
	case *AccountmanagerAction_Supervise:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Supervise); err != nil {
			return err
		}
	case *AccountmanagerAction_Apply:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Apply); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("AccountmanagerAction.Value has unexpected type %T", x)
	}
	return nil
}

func _AccountmanagerAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AccountmanagerAction)
	switch tag {
	case 1: // value.register
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Register)
		err := b.DecodeMessage(msg)
		m.Value = &AccountmanagerAction_Register{msg}
		return true, err
	case 2: // value.resetKey
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ResetKey)
		err := b.DecodeMessage(msg)
		m.Value = &AccountmanagerAction_ResetKey{msg}
		return true, err
	case 3: // value.transfer
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Transfer)
		err := b.DecodeMessage(msg)
		m.Value = &AccountmanagerAction_Transfer{msg}
		return true, err
	case 4: // value.supervise
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Supervise)
		err := b.DecodeMessage(msg)
		m.Value = &AccountmanagerAction_Supervise{msg}
		return true, err
	case 5: // value.apply
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Apply)
		err := b.DecodeMessage(msg)
		m.Value = &AccountmanagerAction_Apply{msg}
		return true, err
	default:
		return false, nil
	}
}

func _AccountmanagerAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AccountmanagerAction)
	// value
	switch x := m.Value.(type) {
	case *AccountmanagerAction_Register:
		s := proto.Size(x.Register)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AccountmanagerAction_ResetKey:
		s := proto.Size(x.ResetKey)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AccountmanagerAction_Transfer:
		s := proto.Size(x.Transfer)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AccountmanagerAction_Supervise:
		s := proto.Size(x.Supervise)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AccountmanagerAction_Apply:
		s := proto.Size(x.Apply)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// 注册
type Register struct {
	AccountID string `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
}

func (m *Register) Reset()                    { *m = Register{} }
func (m *Register) String() string            { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()               {}
func (*Register) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Register) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

// 重置公钥
type ResetKey struct {
	AccountID string `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
	Addr      string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
}

func (m *ResetKey) Reset()                    { *m = ResetKey{} }
func (m *ResetKey) String() string            { return proto.CompactTextString(m) }
func (*ResetKey) ProtoMessage()               {}
func (*ResetKey) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ResetKey) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *ResetKey) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

// 用户申请服务
type Apply struct {
	AccountID string `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
	// 操作， 1 撤销账户公钥重置, 2 锁定期结束后，执行重置公钥操作
	Op int32 `protobuf:"varint,2,opt,name=op" json:"op,omitempty"`
}

func (m *Apply) Reset()                    { *m = Apply{} }
func (m *Apply) String() string            { return proto.CompactTextString(m) }
func (*Apply) ProtoMessage()               {}
func (*Apply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Apply) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *Apply) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

// 资产类型
type Asset struct {
	Execer string `protobuf:"bytes,1,opt,name=execer" json:"execer,omitempty"`
	Symbol string `protobuf:"bytes,2,opt,name=symbol" json:"symbol,omitempty"`
}

func (m *Asset) Reset()                    { *m = Asset{} }
func (m *Asset) String() string            { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()               {}
func (*Asset) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Asset) GetExecer() string {
	if m != nil {
		return m.Execer
	}
	return ""
}

func (m *Asset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

// 合约内部账户之间转账
type Transfer struct {
	// 资产类型
	Asset *Asset `protobuf:"bytes,1,opt,name=Asset" json:"Asset,omitempty"`
	// from账户
	FromAccountID string `protobuf:"bytes,2,opt,name=fromAccountID" json:"fromAccountID,omitempty"`
	// to账户
	ToAccountID string `protobuf:"bytes,3,opt,name=toAccountID" json:"toAccountID,omitempty"`
	// 转账金额
	Amount int64 `protobuf:"varint,4,opt,name=amount" json:"amount,omitempty"`
}

func (m *Transfer) Reset()                    { *m = Transfer{} }
func (m *Transfer) String() string            { return proto.CompactTextString(m) }
func (*Transfer) ProtoMessage()               {}
func (*Transfer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Transfer) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

func (m *Transfer) GetFromAccountID() string {
	if m != nil {
		return m.FromAccountID
	}
	return ""
}

func (m *Transfer) GetToAccountID() string {
	if m != nil {
		return m.ToAccountID
	}
	return ""
}

func (m *Transfer) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// 管理员监管操作
type Supervise struct {
	// 账户名单
	AccountIDs []string `protobuf:"bytes,1,rep,name=accountIDs" json:"accountIDs,omitempty"`
	// 操作， 1为冻结，2为解冻，3增加有效期,4为授权
	Op int32 `protobuf:"varint,2,opt,name=op" json:"op,omitempty"`
	// 0普通,后面根据业务需要可以自定义，有管理员授予不同的权限
	Level int32 `protobuf:"varint,3,opt,name=level" json:"level,omitempty"`
}

func (m *Supervise) Reset()                    { *m = Supervise{} }
func (m *Supervise) String() string            { return proto.CompactTextString(m) }
func (*Supervise) ProtoMessage()               {}
func (*Supervise) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Supervise) GetAccountIDs() []string {
	if m != nil {
		return m.AccountIDs
	}
	return nil
}

func (m *Supervise) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *Supervise) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type Account struct {
	// 账户名称
	AccountID string `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
	// 地址
	Addr string `protobuf:"bytes,2,opt,name=addr" json:"addr,omitempty"`
	// 上一次公钥地址
	PrevAddr string `protobuf:"bytes,3,opt,name=prevAddr" json:"prevAddr,omitempty"`
	// 账户状态 0 正常， 1表示冻结, 2表示锁定 3,过期注销
	Status int32 `protobuf:"varint,4,opt,name=status" json:"status,omitempty"`
	// 等级权限 0普通,后面根据业务需要可以自定义，有管理员授予不同的权限
	Level int32 `protobuf:"varint,5,opt,name=level" json:"level,omitempty"`
	// 注册时间
	CreateTime int64 `protobuf:"varint,6,opt,name=createTime" json:"createTime,omitempty"`
	// 失效时间
	ExpireTime int64 `protobuf:"varint,7,opt,name=expireTime" json:"expireTime,omitempty"`
	// 锁定时间
	LockTime int64 `protobuf:"varint,8,opt,name=lockTime" json:"lockTime,omitempty"`
	// 主键索引
	Index int64 `protobuf:"varint,9,opt,name=index" json:"index,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Account) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *Account) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Account) GetPrevAddr() string {
	if m != nil {
		return m.PrevAddr
	}
	return ""
}

func (m *Account) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Account) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Account) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Account) GetExpireTime() int64 {
	if m != nil {
		return m.ExpireTime
	}
	return 0
}

func (m *Account) GetLockTime() int64 {
	if m != nil {
		return m.LockTime
	}
	return 0
}

func (m *Account) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type AccountReceipt struct {
	Account *Account `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *AccountReceipt) Reset()                    { *m = AccountReceipt{} }
func (m *AccountReceipt) String() string            { return proto.CompactTextString(m) }
func (*AccountReceipt) ProtoMessage()               {}
func (*AccountReceipt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *AccountReceipt) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type ReplyAccountList struct {
	Accounts   []*Account `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
	PrimaryKey string     `protobuf:"bytes,2,opt,name=primaryKey" json:"primaryKey,omitempty"`
}

func (m *ReplyAccountList) Reset()                    { *m = ReplyAccountList{} }
func (m *ReplyAccountList) String() string            { return proto.CompactTextString(m) }
func (*ReplyAccountList) ProtoMessage()               {}
func (*ReplyAccountList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ReplyAccountList) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *ReplyAccountList) GetPrimaryKey() string {
	if m != nil {
		return m.PrimaryKey
	}
	return ""
}

type TransferReceipt struct {
	FromAccount *Account `protobuf:"bytes,1,opt,name=FromAccount" json:"FromAccount,omitempty"`
	ToAccount   *Account `protobuf:"bytes,2,opt,name=ToAccount" json:"ToAccount,omitempty"`
	Index       int64    `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
}

func (m *TransferReceipt) Reset()                    { *m = TransferReceipt{} }
func (m *TransferReceipt) String() string            { return proto.CompactTextString(m) }
func (*TransferReceipt) ProtoMessage()               {}
func (*TransferReceipt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *TransferReceipt) GetFromAccount() *Account {
	if m != nil {
		return m.FromAccount
	}
	return nil
}

func (m *TransferReceipt) GetToAccount() *Account {
	if m != nil {
		return m.ToAccount
	}
	return nil
}

func (m *TransferReceipt) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

// 回执日志
type SuperviseReceipt struct {
	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts" json:"accounts,omitempty"`
	Op       int32      `protobuf:"varint,2,opt,name=op" json:"op,omitempty"`
	Index    int64      `protobuf:"varint,3,opt,name=index" json:"index,omitempty"`
}

func (m *SuperviseReceipt) Reset()                    { *m = SuperviseReceipt{} }
func (m *SuperviseReceipt) String() string            { return proto.CompactTextString(m) }
func (*SuperviseReceipt) ProtoMessage()               {}
func (*SuperviseReceipt) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *SuperviseReceipt) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *SuperviseReceipt) GetOp() int32 {
	if m != nil {
		return m.Op
	}
	return 0
}

func (m *SuperviseReceipt) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type QueryExpiredAccounts struct {
	PrimaryKey string `protobuf:"bytes,1,opt,name=primaryKey" json:"primaryKey,omitempty"`
	// 第一次需要传入逾期时间，时间戳
	ExpiredTime int64 `protobuf:"varint,2,opt,name=expiredTime" json:"expiredTime,omitempty"`
	// 单页返回多少条记录，默认返回10条
	// 0降序，1升序，默认降序
	Direction int32 `protobuf:"varint,3,opt,name=direction" json:"direction,omitempty"`
}

func (m *QueryExpiredAccounts) Reset()                    { *m = QueryExpiredAccounts{} }
func (m *QueryExpiredAccounts) String() string            { return proto.CompactTextString(m) }
func (*QueryExpiredAccounts) ProtoMessage()               {}
func (*QueryExpiredAccounts) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *QueryExpiredAccounts) GetPrimaryKey() string {
	if m != nil {
		return m.PrimaryKey
	}
	return ""
}

func (m *QueryExpiredAccounts) GetExpiredTime() int64 {
	if m != nil {
		return m.ExpiredTime
	}
	return 0
}

func (m *QueryExpiredAccounts) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type QueryAccountsByStatus struct {
	// 账户状态 1 正常， 2表示冻结, 3表示锁定
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	// 主键索引
	PrimaryKey string `protobuf:"bytes,3,opt,name=primaryKey" json:"primaryKey,omitempty"`
	// 0降序，1升序，默认降序
	Direction int32 `protobuf:"varint,5,opt,name=direction" json:"direction,omitempty"`
}

func (m *QueryAccountsByStatus) Reset()                    { *m = QueryAccountsByStatus{} }
func (m *QueryAccountsByStatus) String() string            { return proto.CompactTextString(m) }
func (*QueryAccountsByStatus) ProtoMessage()               {}
func (*QueryAccountsByStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *QueryAccountsByStatus) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *QueryAccountsByStatus) GetPrimaryKey() string {
	if m != nil {
		return m.PrimaryKey
	}
	return ""
}

func (m *QueryAccountsByStatus) GetDirection() int32 {
	if m != nil {
		return m.Direction
	}
	return 0
}

type QueryAccountByID struct {
	AccountID string `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
}

func (m *QueryAccountByID) Reset()                    { *m = QueryAccountByID{} }
func (m *QueryAccountByID) String() string            { return proto.CompactTextString(m) }
func (*QueryAccountByID) ProtoMessage()               {}
func (*QueryAccountByID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *QueryAccountByID) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

type QueryAccountByAddr struct {
	Addr string `protobuf:"bytes,1,opt,name=addr" json:"addr,omitempty"`
}

func (m *QueryAccountByAddr) Reset()                    { *m = QueryAccountByAddr{} }
func (m *QueryAccountByAddr) String() string            { return proto.CompactTextString(m) }
func (*QueryAccountByAddr) ProtoMessage()               {}
func (*QueryAccountByAddr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *QueryAccountByAddr) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

type QueryBalanceByID struct {
	AccountID string `protobuf:"bytes,1,opt,name=accountID" json:"accountID,omitempty"`
	Asset     *Asset `protobuf:"bytes,2,opt,name=asset" json:"asset,omitempty"`
}

func (m *QueryBalanceByID) Reset()                    { *m = QueryBalanceByID{} }
func (m *QueryBalanceByID) String() string            { return proto.CompactTextString(m) }
func (*QueryBalanceByID) ProtoMessage()               {}
func (*QueryBalanceByID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *QueryBalanceByID) GetAccountID() string {
	if m != nil {
		return m.AccountID
	}
	return ""
}

func (m *QueryBalanceByID) GetAsset() *Asset {
	if m != nil {
		return m.Asset
	}
	return nil
}

type Balance struct {
	Balance int64 `protobuf:"varint,1,opt,name=balance" json:"balance,omitempty"`
	Frozen  int64 `protobuf:"varint,2,opt,name=frozen" json:"frozen,omitempty"`
}

func (m *Balance) Reset()                    { *m = Balance{} }
func (m *Balance) String() string            { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()               {}
func (*Balance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *Balance) GetBalance() int64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Balance) GetFrozen() int64 {
	if m != nil {
		return m.Frozen
	}
	return 0
}

func init() {
	proto.RegisterType((*Accountmanager)(nil), "types.Accountmanager")
	proto.RegisterType((*AccountmanagerAction)(nil), "types.AccountmanagerAction")
	proto.RegisterType((*Register)(nil), "types.Register")
	proto.RegisterType((*ResetKey)(nil), "types.ResetKey")
	proto.RegisterType((*Apply)(nil), "types.Apply")
	proto.RegisterType((*Asset)(nil), "types.asset")
	proto.RegisterType((*Transfer)(nil), "types.Transfer")
	proto.RegisterType((*Supervise)(nil), "types.Supervise")
	proto.RegisterType((*Account)(nil), "types.account")
	proto.RegisterType((*AccountReceipt)(nil), "types.AccountReceipt")
	proto.RegisterType((*ReplyAccountList)(nil), "types.ReplyAccountList")
	proto.RegisterType((*TransferReceipt)(nil), "types.TransferReceipt")
	proto.RegisterType((*SuperviseReceipt)(nil), "types.SuperviseReceipt")
	proto.RegisterType((*QueryExpiredAccounts)(nil), "types.QueryExpiredAccounts")
	proto.RegisterType((*QueryAccountsByStatus)(nil), "types.QueryAccountsByStatus")
	proto.RegisterType((*QueryAccountByID)(nil), "types.QueryAccountByID")
	proto.RegisterType((*QueryAccountByAddr)(nil), "types.QueryAccountByAddr")
	proto.RegisterType((*QueryBalanceByID)(nil), "types.QueryBalanceByID")
	proto.RegisterType((*Balance)(nil), "types.balance")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Accountmanager service

type AccountmanagerClient interface {
}

type accountmanagerClient struct {
	cc *grpc.ClientConn
}

func NewAccountmanagerClient(cc *grpc.ClientConn) AccountmanagerClient {
	return &accountmanagerClient{cc}
}

// Server API for Accountmanager service

type AccountmanagerServer interface {
}

func RegisterAccountmanagerServer(s *grpc.Server, srv AccountmanagerServer) {
	s.RegisterService(&_Accountmanager_serviceDesc, srv)
}

var _Accountmanager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.accountmanager",
	HandlerType: (*AccountmanagerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "accountmanager.proto",
}

func init() { proto.RegisterFile("accountmanager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 731 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4d, 0x6b, 0xdb, 0x4c,
	0x10, 0x8e, 0xac, 0xc8, 0xb6, 0x26, 0xef, 0xeb, 0x98, 0xc5, 0x0d, 0xa2, 0x94, 0x62, 0x96, 0x1c,
	0x4c, 0x69, 0x43, 0x48, 0x29, 0x85, 0xb6, 0x17, 0x9b, 0xb4, 0x24, 0xb4, 0x97, 0x6c, 0x7c, 0x2e,
	0x28, 0xf2, 0x24, 0x88, 0xca, 0x96, 0x58, 0xad, 0x4d, 0xd4, 0x3f, 0xd0, 0x5b, 0x4f, 0xfd, 0xaf,
	0xbd, 0x96, 0x9d, 0x5d, 0x7d, 0x39, 0x69, 0x43, 0x4e, 0xda, 0x79, 0xe6, 0xd9, 0xf9, 0xde, 0x11,
	0x8c, 0xc2, 0x28, 0x4a, 0xd7, 0x2b, 0xb5, 0x0c, 0x57, 0xe1, 0x0d, 0xca, 0xa3, 0x4c, 0xa6, 0x2a,
	0x65, 0x9e, 0x2a, 0x32, 0xcc, 0xf9, 0x10, 0x06, 0xd3, 0x96, 0x9a, 0xff, 0xea, 0xc0, 0xa8, 0x0d,
	0x4d, 0x23, 0x15, 0xa7, 0x2b, 0xf6, 0x0a, 0xfa, 0x12, 0x6f, 0xe2, 0x5c, 0xa1, 0x0c, 0x9c, 0xb1,
	0x33, 0xd9, 0x3b, 0xd9, 0x3f, 0x22, 0x23, 0x47, 0xc2, 0xc2, 0x67, 0x3b, 0xa2, 0xa2, 0x18, 0x7a,
	0x8e, 0xea, 0x33, 0x16, 0x41, 0x67, 0x8b, 0x6e, 0x60, 0x43, 0x37, 0x67, 0x4d, 0x57, 0x32, 0x5c,
	0xe5, 0xd7, 0x28, 0x03, 0xb7, 0x45, 0x9f, 0x5b, 0x58, 0xd3, 0x4b, 0x0a, 0x3b, 0x06, 0x3f, 0x5f,
	0x67, 0x28, 0x37, 0x71, 0x8e, 0xc1, 0x2e, 0xf1, 0x87, 0x96, 0x7f, 0x59, 0xe2, 0x67, 0x3b, 0xa2,
	0x26, 0xb1, 0x43, 0xf0, 0xc2, 0x2c, 0x4b, 0x8a, 0xc0, 0x23, 0xf6, 0x7f, 0x96, 0x3d, 0xd5, 0xd8,
	0xd9, 0x8e, 0x30, 0x4a, 0x36, 0x80, 0x8e, 0x2a, 0x82, 0xee, 0xd8, 0x99, 0x78, 0xa2, 0xa3, 0x8a,
	0x59, 0x0f, 0xbc, 0x4d, 0x98, 0xac, 0x91, 0x4f, 0xa0, 0x5f, 0xa6, 0xc9, 0x9e, 0x81, 0x6f, 0x6b,
	0x7a, 0x7e, 0x4a, 0xa5, 0xf0, 0x45, 0x0d, 0xf0, 0x0f, 0x9a, 0x69, 0xb3, 0xfa, 0x27, 0x93, 0x31,
	0xd8, 0x0d, 0x17, 0x0b, 0x49, 0xe5, 0xf1, 0x05, 0x9d, 0xf9, 0x1b, 0xf0, 0x28, 0xa4, 0x07, 0xae,
	0x0e, 0xa0, 0x93, 0x66, 0x74, 0xd1, 0x13, 0x9d, 0x34, 0xe3, 0x6f, 0xc1, 0x0b, 0xf3, 0x1c, 0x15,
	0x3b, 0x80, 0x2e, 0xde, 0x62, 0x64, 0x7b, 0xe4, 0x0b, 0x2b, 0x69, 0x3c, 0x2f, 0x96, 0x57, 0x69,
	0x62, 0xbd, 0x59, 0x89, 0xff, 0x74, 0xa0, 0x5f, 0x56, 0x98, 0x71, 0xf0, 0xa6, 0xda, 0x8a, 0xed,
	0x6f, 0x59, 0x23, 0xb2, 0x2c, 0x8c, 0x8a, 0x1d, 0xc2, 0xff, 0xd7, 0x32, 0x5d, 0x4e, 0xab, 0xd8,
	0x8c, 0xbd, 0x36, 0xc8, 0xc6, 0xb0, 0xa7, 0xd2, 0x9a, 0xe3, 0x12, 0xa7, 0x09, 0xe9, 0x80, 0xc2,
	0xa5, 0x3e, 0x53, 0xfb, 0x5c, 0x61, 0x25, 0x7e, 0x01, 0x7e, 0xd5, 0x41, 0xf6, 0x1c, 0xa0, 0xca,
	0x39, 0x0f, 0x9c, 0xb1, 0x3b, 0xf1, 0x45, 0x03, 0xd9, 0x2e, 0x03, 0x1b, 0x81, 0x97, 0xe0, 0x06,
	0x13, 0x72, 0xe8, 0x09, 0x23, 0xf0, 0xdf, 0x0e, 0xf4, 0xec, 0xa5, 0xc7, 0x77, 0x84, 0x3d, 0x85,
	0x7e, 0x26, 0x71, 0x33, 0xd5, 0xb8, 0xc9, 0xa3, 0x92, 0xa9, 0xaa, 0x2a, 0x54, 0xeb, 0x9c, 0x92,
	0xf0, 0x84, 0x95, 0xea, 0x38, 0xbc, 0x46, 0x1c, 0x3a, 0x9b, 0x48, 0x62, 0xa8, 0x70, 0x1e, 0x2f,
	0x91, 0x86, 0xcc, 0x15, 0x0d, 0x44, 0xeb, 0xf1, 0x36, 0x8b, 0xa5, 0xd1, 0xf7, 0x8c, 0xbe, 0x46,
	0x74, 0x24, 0x49, 0x1a, 0x7d, 0x23, 0x6d, 0x9f, 0xb4, 0x95, 0xac, 0x3d, 0xc6, 0xab, 0x05, 0xde,
	0x06, 0x3e, 0x29, 0x8c, 0xc0, 0xdf, 0x55, 0xcf, 0x5b, 0x60, 0x84, 0x71, 0xa6, 0xd8, 0xa4, 0x2a,
	0x85, 0x6d, 0xf2, 0xa0, 0x6c, 0xb2, 0xe5, 0x95, 0x6a, 0xfe, 0x15, 0x86, 0x02, 0xb3, 0xa4, 0xb0,
	0x06, 0xbe, 0xc4, 0xb9, 0x62, 0x2f, 0xa0, 0x6f, 0xd5, 0xa6, 0x1b, 0x77, 0xaf, 0x57, 0x7a, 0x9d,
	0x4d, 0x26, 0xe3, 0x65, 0x28, 0x8b, 0x72, 0x05, 0xf8, 0xa2, 0x81, 0xf0, 0x1f, 0x0e, 0xec, 0x97,
	0x93, 0x57, 0x46, 0x77, 0x0c, 0x7b, 0x9f, 0xea, 0x39, 0xfa, 0x4b, 0x84, 0x4d, 0x0a, 0x7b, 0x09,
	0xfe, 0xbc, 0x9c, 0x2a, 0xbb, 0x67, 0xb6, 0xf9, 0x35, 0xa1, 0xae, 0x92, 0xdb, 0xac, 0xd2, 0x02,
	0x86, 0xd5, 0xc8, 0x95, 0x91, 0x3c, 0x26, 0xd3, 0x7b, 0xa6, 0xf0, 0x1e, 0x2f, 0x1b, 0x18, 0x5d,
	0xac, 0x51, 0x16, 0x1f, 0xa9, 0xa1, 0x8b, 0xe9, 0xfd, 0x75, 0x72, 0xb6, 0xeb, 0xa4, 0x9f, 0x92,
	0x99, 0x81, 0x05, 0x35, 0xbe, 0x43, 0x36, 0x9b, 0x90, 0x9e, 0xe9, 0x45, 0x2c, 0x91, 0xd6, 0xb4,
	0x9d, 0xfc, 0x1a, 0xe0, 0x4b, 0x78, 0x42, 0x7e, 0x4b, 0x87, 0xb3, 0xe2, 0xd2, 0x0c, 0x69, 0x3d,
	0xbc, 0x4e, 0x6b, 0x78, 0xdb, 0x01, 0xb9, 0x77, 0x02, 0x6a, 0xb9, 0xf3, 0xb6, 0xdd, 0x1d, 0xc3,
	0xb0, 0xe9, 0x6e, 0x56, 0x9c, 0x9f, 0x3e, 0xb0, 0x30, 0x27, 0xc0, 0xda, 0x37, 0xe8, 0x69, 0x95,
	0x4f, 0xd1, 0x69, 0x2c, 0xc7, 0xb9, 0xb5, 0x3d, 0x0b, 0x93, 0x70, 0x15, 0xe1, 0xc3, 0xb6, 0xf5,
	0x46, 0xa3, 0xed, 0x65, 0x47, 0x63, 0x6b, 0xa3, 0xd1, 0x87, 0xbf, 0x87, 0xde, 0x95, 0x31, 0xc8,
	0x82, 0xea, 0x48, 0xa6, 0x5c, 0x51, 0x69, 0x0e, 0xa0, 0x7b, 0x2d, 0xd3, 0xef, 0xb8, 0xb2, 0x0d,
	0xb0, 0xd2, 0xc9, 0x10, 0x06, 0xed, 0xff, 0xeb, 0x55, 0x97, 0x7e, 0xb0, 0xaf, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x84, 0x94, 0xab, 0xf2, 0x78, 0x07, 0x00, 0x00,
}