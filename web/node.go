package web

const (
	// ROOT 根节点
	ROOT = iota
	// STATIC 	静态节点
	STATIC
	// PARAM 	带参数节点
	PARAM
	// REG 正则
	REG
	// STAR * 匹配
	STAR
)

type match = func(string2 string, ctx *Context) bool

type node struct {
	child       []*node
	match       match
	nodeType    int
	handlerFunc handlerFunc
}

func newRootNode() *node {
	return &node{
		child: make([]*node, 2),
		match: func(string2 string, ctx *Context) bool {
			return true
		},
		nodeType: ROOT,
	}
}

func newStaticNode() *node {
	return &node{
		nodeType: STATIC,
	}
}

func newParamNode() *node {
	return &node{
		nodeType: PARAM,
	}
}

func newRegNode() *node {
	return &node{
		nodeType: REG,
	}
}

func newStarNode() *node {
	return &node{
		nodeType: STAR,
	}
}
