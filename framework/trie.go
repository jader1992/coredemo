package framework

import (
	"errors"
	"strings"
)

/**
 * 主要实现了字典树：
 * 1. 定义树和节点的数据结构
 * 2. 增加编写函数： "增加路由规则"
 * 3. 编写函数："查找路由"
 * 4. 将这些路由添加到框架中
 */

// 代表节点
type node struct {
	isLast bool // 该节点是否能成为一个独立的uri, 是否自身就是一个终极节点
	segment string // uri中的字符串
	handlers []ControllerHandler // 中间件 + 控制器
	childes []*node // 子节点
}

// Tree 代表树结构
type Tree struct {
	root *node
}

// 初始化节点
func newCode() *node  {
	return &node{
		isLast: false,
		segment: "",
		childes: []*node{},
	}
}

// NewTree 初始化树
func NewTree() *Tree {
	root := newCode()
	return &Tree{
		root: root,
	}
}

// 增加路由节点，路由节点有先后顺序
/*
/book/list
/book/:id (冲突)
/book/:id/name
/book/:student/age
/:user/name
/:user/name/:age (冲突)
*/

func (tree *Tree) AddRouter(uri string, handlers []ControllerHandler) error {
	n := tree.root

	// 检测这个路由是否已经在路由树上
	if n.matchNode(uri) != nil {
		return errors.New("route exist: " + uri)
	}

	segments := strings.Split(uri, "/")

	// 对每个segment
	for index, segment := range segments {
		// 最终进入Node segment的字段
		if !isWildSegment(segment) {
			segment = strings.ToUpper(segment)
		}
		isLast := index == len(segments)-1

		var objNode *node // 标记是否有合适的子节点

		childNodes := n.filterChildNodes(segment)
		// 如果有匹配的节点
		if len(childNodes) > 0 {
			// 如果有segment相同的子节点，则选择这个子节点
			for _, cNode := range childNodes {
				if cNode.segment == segment {
					objNode = cNode
					break
				}
			}
		}

		if objNode == nil {
			// 创建一个当前的node节点
			cNode := newCode()
			cNode.segment = segment
			if isLast {
				cNode.isLast = true
				cNode.handlers = handlers
			}
			n.childes = append(n.childes, cNode)
			objNode = cNode
		}

		n = objNode
	}

	return nil
}

// 判断路由是否已经在节点的所有子节点树中存在了
func (n *node) matchNode(uri string) *node {
	// 使用分隔符将uri切割为两个部分
	segments := strings.SplitN(uri, "/", 2)
	// 如果只能分割成一个段，说明 URI 中没有分隔符了，这时候再检查下一级节点中是否有匹配这个段的节点就行。
	// 如果分割成了两个段，我们用第一个段来检查下一个级节点中是否有匹配这个段的节点。

	// 第一个部分用于匹配下一层子节点
	segment := segments[0]
	if !isWildSegment(segment) {
		// 统一换成大写
		segment = strings.ToUpper(segment)
	}

	// 匹配符合的下一层节点
	cNodes := n.filterChildNodes(segment)
	// 如果当前子节点没有一个符合，那么说明这个uri一定是之前不存在, 直接返回nil
	if cNodes == nil || len(cNodes) == 0 {
		return nil
	}

	// 如果只有一个segment，则是最后一个标记
	if len(segments) == 1 {
		// 如果segment已经是最后一个节点，判断这些cNodes是否有isLast标志
		for _, tn := range cNodes {
			if tn.isLast {
				return tn
			}
		}

		// 都不是最后一个节点
		return nil
	}

	// 如果有2个segment
	for _, tn := range cNodes {
		// 递归每个子节点继续进行查找
		tnMatch := tn.matchNode(segments[1])
		if tnMatch != nil {
			return tnMatch
		}
	}
	return nil
}

// 过滤下一层满足segment规则的子节点
func (n *node) filterChildNodes(segment string) []*node {
	if len(n.childes) == 0 {
		return nil
	}

	// 如果segment是通配符，则所有下一层子节点都满足需求
	if isWildSegment(segment) {
		return n.childes
	}

	nodes := make([]*node, 0, len(n.childes))
	// 过滤所有下一层的子节点
	for _, cNode := range n.childes {
		if isWildSegment(cNode.segment) {
			// 如果下一层节点有通配符，则满足要求
			nodes = append(nodes, cNode)
		} else if cNode.segment == segment {
			// 如果下一层节点没有通配符，但是文本完全匹配，则满足要求
			nodes = append(nodes, cNode)
		}
	}

	return nodes
}

// 判断一个segment是否是通用segment，即以:开头
func isWildSegment(segment string) bool {
	return strings.HasPrefix(segment, ":")
}

// FindHandler 匹配uri
func (tree *Tree) FindHandler(uri string) []ControllerHandler {
	matchNode := tree.root.matchNode(uri)
	if matchNode == nil {
		return nil
	}
	return matchNode.handlers
}
