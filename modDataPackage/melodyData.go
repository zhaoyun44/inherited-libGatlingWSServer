package modDataPackage

import (
	"fmt"

	"github.com/gatlinglab/libGatlingWSServer/internal/honorMelody"
	"github.com/gatlinglab/libGatlingWSServer/modProtocol"
)

/*type IWJWSSocket interface {
	Write(msg []byte) error
	WriteBinary(msg []byte) error
	Close() error
	IsClosed() bool
	LocalAddr()
	RemoteAddr()
}*/

type CWJMelodyDataHandler struct {
	handlerClose   modProtocol.CBWJClosedHandler
	handlerConnect modProtocol.CBWJConnectedHandler
	//handlerMessage modProtocol.CBWJMessageHandler
	adapter *modProtocol.ProtocolAdapter
}

func newMelodyDataHandler() *CWJMelodyDataHandler {
	return &CWJMelodyDataHandler{adapter: modProtocol.NewProtocolAdapter()}
}

func (pInst *CWJMelodyDataHandler) OnClose(session *honorMelody.Session, code int, reason string) error {
	fmt.Println("connect: ", session)
	wjSession := session.DataAdapter.(modProtocol.IWJSocket)
	pInst.handlerClose(wjSession)
	return nil
}

func (pInst *CWJMelodyDataHandler) OnConnect(session *honorMelody.Session) {
	wjSession := modProtocol.NewCWJSessionServer(session)
	session.DataAdapter = wjSession
	pInst.handlerConnect(wjSession)
}

func (pInst *CWJMelodyDataHandler) OnMessage(session *honorMelody.Session, msg []byte) {
	wjSession := session.DataAdapter.(modProtocol.IWJSocket)
	fmt.Println("melody data: ", string(msg), msg)
	pInst.adapter.OnMessage(wjSession, msg)
}

func (pInst *CWJMelodyDataHandler) WsHandlerClose(fn modProtocol.CBWJClosedHandler) {
	pInst.handlerClose = fn
}

func (pInst *CWJMelodyDataHandler) WsHandlerConnect(fn modProtocol.CBWJConnectedHandler) {
	pInst.handlerConnect = fn
}

func (pInst *CWJMelodyDataHandler) WsHandlerMessage(fn modProtocol.CBWJMessageHandler) {
	//pInst.handlerMessage = fn
	pInst.adapter.WsHandlerMessage(fn)
}
