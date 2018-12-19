package dns

import (
	"designPattern/ABE_oberver_responsibilitychain/b_observer_dns/observer"
	"math/rand"
	"time"
	"fmt"
	"strings"
)

type IServer interface {
	observer.Observer
	IsLocal(recorder *Recorder) bool
	SetUpperServer(server IServer)
	ResponsFromUpperServer(recorder *Recorder)
	Sign(recorder *Recorder)
}

type DnsServer struct {
	*observer.Observable
	specific IServer
}

func NewDnsServer(specific IServer) *DnsServer {
	return &DnsServer{
		observer.NewObservable(),
		specific,
	}
}

func (l *DnsServer) Update(o observer.IObservable, arg interface{}) {
	recoder := arg.(*Recorder)
	if l.specific.IsLocal(recoder) {
		recoder.Ip = l.genIpAddress();
	} else {
		l.specific.ResponsFromUpperServer(recoder)
	}
	l.specific.Sign(recoder)
}

func (l *DnsServer) IsLocal(recorder *Recorder) bool {
	panic("not implement")
	return true
}

func (l *DnsServer)SetUpperServer(server IServer) {
	l.DeleteObservers()
	l.AddObserver(server)
}

func (l *DnsServer) ResponsFromUpperServer(recorder *Recorder) {
	l.SetChanged()
	l.NotifyObservers(recorder)
}

func (l *DnsServer) Sign(recorder *Recorder){
	panic("not implement")
}

func (l *DnsServer) genIpAddress() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%d.%d.%d.%d", r.Intn(255), r.Intn(255), r.Intn(255), r.Intn(255))
}

///////////////////////////////////////////////////

type SHDnsServer struct {
	*DnsServer
}

func NewSHDnsServer() *SHDnsServer {
	shd := new(SHDnsServer)
	ds := NewDnsServer(shd)
	shd.DnsServer = ds
	return shd
}

func (l *SHDnsServer) Sign(recorder *Recorder) {
	recorder.Owner = "shang hai"
}

func (l *SHDnsServer) IsLocal(recorder *Recorder) bool {
	return strings.HasSuffix(recorder.Domain, ".sh.cn")
}

///////////////////////////////////////////////////

type ChinaDnsServer struct {
	*DnsServer
}

func NewChinaDnsServer() *ChinaDnsServer {
	shd := new(ChinaDnsServer)
	ds := NewDnsServer(shd)
	shd.DnsServer = ds
	return shd
}

func (l *ChinaDnsServer) Sign(recorder *Recorder) {
	recorder.Owner = "china"
}

func (l *ChinaDnsServer) IsLocal(recorder *Recorder) bool {
	return strings.HasSuffix(recorder.Domain, ".cn")
}

///////////////////////////////////////////////////

type GlobalDnsServer struct {
	*DnsServer
}

func NewGlobalDnsServer() *GlobalDnsServer {
	shd := new(GlobalDnsServer)
	ds := NewDnsServer(shd)
	shd.DnsServer = ds
	return shd
}

func (l *GlobalDnsServer) Sign(recorder *Recorder) {
	recorder.Owner = "global"
}

func (l *GlobalDnsServer) IsLocal(recorder *Recorder) bool {
	return true
}
