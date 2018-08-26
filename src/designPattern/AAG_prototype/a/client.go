package a

import (
	"fmt"
	"math/rand"
	"bytes"
	"time"
	"designPattern/AAG_prototype/a/adv"
	"sync"
)

var MAXCOUNT = 6

func SendMail(m *adv.Mail) {
	fmt.Println("title: ", m.GetSubject(), "receiver: ", m.GetReceiver(), "...send successfully.")
}

func GetRandString(length int) string {
	s := "abcdefghhigklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	sb := []rune(s)
	l := len(s)
	var res bytes.Buffer
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i:=0; i<length; i++ {
		res.WriteRune(sb[r.Intn(l)])
	}
	return res.String()
}

func main() {
	mail := adv.NewMail(adv.NewAdvTemplate())
	var wg sync.WaitGroup
	wg.Add(MAXCOUNT)
	for i:= 0; i<MAXCOUNT; i++ {
		go func(m *adv.Mail) {
			cloneMail := m.Clone()
			cloneMail.SetAppellation(GetRandString(5))
			cloneMail.SetReceiver(GetRandString(5) + "@" + GetRandString(8) + ".com")
			SendMail(cloneMail)
			wg.Done()
		}(mail)
	}
	wg.Wait()
}
