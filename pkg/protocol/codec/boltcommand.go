package codec

import "net"

//command defination
type BoltCommand struct {
	cmdCode int16
	version byte
	cmdType byte
	codec   byte
	//protoSwitchStatus byte

	id            int
	classLength   int16
	headerLength  int16
	contentLength int

	class   []byte
	header  []byte
	content []byte

	invokeContext interface{}
}

type BoltRequestCommand struct {
	//rpc command
	BoltCommand

	//request command
	timeout int

	//rpc request command
	requestObject interface{}
	requestClass  string

	//customSerializer CustomSerializer
	requestHeader map[string]string
	arriveTime    int64
}

type BoltResponseCommand struct {
	//rpc command
	BoltCommand

	//response command
	responseStatus     int16
	responseTimeMillis int64
	responseHost       net.Addr
	cause              error

	//rpc response command
	responseObject interface{}
	responseClass  string

	//customSerializer CustomSerializer
	responseHeader map[string]string

	errorMsg string
}

func (b *BoltCommand) GetCmdCode() int16{
	return b.cmdCode
}

func (b *BoltCommand) GetClass() []byte{
	return b.class
}

func (b *BoltCommand) GetHeader() []byte{
	return b.header
}

func (b *BoltCommand) GetContent() []byte{
	return b.content
}

func (b *BoltRequestCommand) SetTimeout(timeout int) {
	b.timeout = timeout
}
func (b *BoltRequestCommand) SetArriveTime(arriveTime int64) {
	b.arriveTime = arriveTime
}
func (b *BoltRequestCommand) SetRequestHeader(headerMap map[string]string) {
	b.requestHeader = headerMap
}
func (b *BoltRequestCommand) GetRequestHeader(headerMap map[string]string) {
	b.requestHeader = headerMap
}

func (b *BoltResponseCommand) SetResponseStatus(status int16) {
	b.responseStatus = status
}
func (b *BoltResponseCommand) SetResponseTimeMillis(responseTime int64) {
	b.responseTimeMillis = responseTime
}
func (b *BoltResponseCommand) SetResponseHost(host net.Addr) {
	b.responseHost = host
}