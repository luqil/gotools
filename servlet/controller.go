package servlet

import (
	"errors"
	"html/template"
	"net/http"
	"os"
	"strings"
)

//静态资源通用控制器
type ControllerStatic struct {
	Controller
	//Http请求时的前缀如：/static
	HttpPathPrefix string
	//对应存储文件的相对目录或绝对目录 如：static 或/home/xxx/static
	FilePathPrefix string
}

func (cs *ControllerStatic) HandAction(res http.ResponseWriter, req *http.Request) error {
	httpPath := req.URL.Path
	if !strings.HasPrefix(httpPath, cs.HttpPathPrefix) {
		return errors.New("The Path Is Wrong :" + httpPath)
	}
	filePath := cs.FilePathPrefix + httpPath[len(cs.HttpPathPrefix):]
	f, err := os.Open(filePath)
	if err != nil {
		return errors.New("Not Find File:" + httpPath)
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for {
		i, _ := f.Read(buf)
		if i == 0 {
			break
		}
		res.Write(buf[:i])
	}
	return nil
}

//模板资源通用控制器
type ControllerTemplate struct {
	Controller
	FilePathPrefix string             //对应存储文件的相对目录或绝对目录 如：view 或/home/xxx/view
	ViewPath       string             //视图地址
	Tpl            *template.Template //模板资源
	Method         string             //为空时不进行匹配 如：(GET, POST, PUT, etc.)

	DoAction func(res http.ResponseWriter, req *http.Request, ct *ControllerTemplate) (map[string]interface{}, error) //执行方法,返回前台数据
}

//执行通用模板
func (ct *ControllerTemplate) HandAction(res http.ResponseWriter, req *http.Request) error {
	if ct.Method != "" {
		if req.Method != ct.Method {
			return errors.New("Request Method Is Wrong:" + req.URL.Path)
		}
	}
	if ct.Tpl == nil {
		ct.Tpl = template.New(ct.ViewPath)
		ct.Tpl.ParseFiles(ct.FilePathPrefix + ct.ViewPath)
	}
	data, err := ct.DoAction(res, req, ct)
	if err != nil {
		err = ct.Tpl.Execute(res, data)
	}
	return err
}
