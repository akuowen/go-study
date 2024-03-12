package web

// Handler 处理器
type Handler interface {
	Router
	// ServerHttp 真正处理http请求的地方  进行路由处理
	ServerHttp(ctx *Context)
}

var _ Handler = &TreeNodeHandler{}

// TreeNodeHandler 通过树来进行url路由
type TreeNodeHandler struct {
}

func (t *TreeNodeHandler) ServerHttp(ctx *Context) {
	//TODO implement me
	panic("implement me")
}

func NewTreeNodeHandler() Handler {
	return &TreeNodeHandler{}
}

// Route 将url添加到树上
func (t *TreeNodeHandler) Route(method string, pattern string, handlerFunc handlerFunc) {
	//TODO implement me
	panic("implement me")
}
