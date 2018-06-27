package main

import (
	"fmt"
	"github.com/gogap/aop"
)

type Auth struct {}

func (p *Auth) Login(userName, password string) bool {
	fmt.Println("@Login", userName, password)
	if userName == "mhc" && password == "123" {
		return true
	}
	return false
}

func (p *Auth) Before(jp aop.JoinPointer) {
	username := ""
	jp.Args().MapTo(func(u, p string) {
		username = u
	})

	fmt.Printf("Before Login: %s\n", username)
}


func (p *Auth) After(username, password string) {
	fmt.Printf("After Login: %s %s\n", username, password)
}


func (p *Auth) Around(pjp aop.ProceedingJoinPointer) {
	fmt.Println("@Begin Around")

	ret := pjp.Proceed("fakeName", "fakePassword")
	ret.MapTo(func(loginResult bool) {
		fmt.Println("@Proceed Result is", loginResult)
	})

	fmt.Println("@End Around")
}

func main() {
	beanFactory := aop.NewClassicBeanFactory()
	beanFactory.RegisterBean("auth", new(Auth))

	aspect := aop.NewAspect("aspect_1", "auth")
	aspect.SetBeanFactory(beanFactory)

	pointcut := aop.NewPointcut("pointcut_1").Execution(`Login()`)
	aspect.AddPointcut(pointcut)

	aspect.AddAdvice(&aop.Advice{Ordering:aop.Before, Method:"Before", PointcutRefID:"pointcut_1"})
	aspect.AddAdvice(&aop.Advice{Ordering:aop.After, Method:"After", PointcutRefID:"pointcut_1"})
	aspect.AddAdvice(&aop.Advice{Ordering:aop.Around, Method:"Around", PointcutRefID:"pointcut_1"})

	gogaAop := aop.NewAOP()
	gogaAop.SetBeanFactory(beanFactory)
	gogaAop.AddAspect(aspect)

	var err error
	var proxy *aop.Proxy

	if proxy, err = gogaAop.GetProxy("auth"); err != nil {
		fmt.Println("get proxy failed", err)
		return
	}

	login := proxy.Method(new(Auth).Login).(func(string, string) bool)("zeal", "gogap")
	fmt.Println(login)
	//if err = proxy.Invoke(new(Auth).Login, "zeal", "errorpassword").End(
	//	func(result bool) {
	//		login = result
	//	}); err != nil {
	//	fmt.Println("invoke proxy func error", err)
	//} else {
	//	fmt.Println("Login result:", login)
	//}
}
