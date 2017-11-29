package servlet

import (
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

//一个简易的Controller
type Controller interface {
	HandAction(res http.ResponseWriter, req *http.Request, hand *Handler) error
}

//一个简易的处理状态Controller
type ControllerStatus interface {
	HandStatus(res http.ResponseWriter, req *http.Request, status int, msg string, hand *Handler) error
}

//简易过滤器 error不为nil时，匹配路径不能访问
type ControllerFilter interface {
	HandFilter(req *http.Request, hand *Handler) error
}
type Handler struct {
	http.Handler
	ViewPrefix       string                      //默认的视图地址前缀
	StaticPrefix     string                      //默认的静态资源地址前缀
	controllers      map[string]*Controller      //注册的控制器
	controllerStatus map[int]*ControllerStatus    //注册的状态控制器
	controllerFilter map[string]*ControllerFilter //注册过滤器
}

func NewHandler(viewPrefix, staticPrefix string) *Handler {
	hand := new(Handler)
	hand.StaticPrefix = staticPrefix
	hand.ViewPrefix = viewPrefix
	return hand
}

//注册一个控制器
//path 以*结尾时，模糊匹配*部分，执行优先级较低
func (hand *Handler) RegController(path string, controller Controller) {
	if hand.controllers == nil {
		hand.controllers = make(map[string]*Controller)
	}
	if hand.controllers[path] != nil {
		log.Fatal("注册Controller路径重复,Path：", path)
	} else {
		hand.controllers[path] = &controller
	}
}

//注册一个状态控制器
//主要为404、500
func (hand *Handler) RegControllerStatus(status int, controller ControllerStatus) {
	hand.controllerStatus[status] = &controller
}

//注册一个过滤器
func (hand *Handler) RegControllerFilter(path string, filter ControllerFilter) {
	if hand.controllerFilter[path] != nil {
		log.Fatal("注册ControllerFilter路径重复,Path：", path)
	} else {
		hand.controllerFilter[path] = &filter
	}
}

//路由执行
func (hand *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	isFind := false //是否已匹配执行
	var err error
	//执行过滤器（可实现登陆拦截，权限拦截等）
	for k, v := range hand.controllerFilter {
		kt := strings.TrimRight(k, "*")
		if path == k || strings.HasPrefix(path, kt) {
			err = (*v).HandFilter(req, hand)
		}
	}

	//过滤器无错误，继续执行
	if err == nil {
		//路径完全匹配时执行
		for k, v := range hand.controllers {
			if path == k {
				err = (*v).HandAction(res, req, hand)
				isFind = true
			}
		}
		//路径模糊匹配时执行
		if !isFind {
			for k, v := range hand.controllers {
				kt := strings.TrimRight(k, "*")
				if strings.HasPrefix(path, kt) {
					err = (*v).HandAction(res, req, hand)
					isFind = true
				}
			}
		}
	}
	//处理404情况
	if err == nil && !isFind {
		if hand.controllerStatus[http.StatusNotFound] != nil {
			err = (*hand.controllerStatus[http.StatusNotFound]).HandStatus(res, req, http.StatusNotFound, "", hand)
		} else {
			http.NotFound(res, req)
		}
	}
	//处理错误的情况
	if err != nil {
		if hand.controllerStatus[http.StatusInternalServerError] != nil {
			err = (*hand.controllerStatus[http.StatusInternalServerError]).HandStatus(res, req, http.StatusInternalServerError, err.Error(), hand)
		} else {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
	}
}

//默认处理器
var DefaultHandler = new(Handler)
