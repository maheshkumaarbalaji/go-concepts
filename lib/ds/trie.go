package ds

import (
	"strings"
	"fmt"
)

type RouteNode struct {
	RoutePart string
	Children []*RouteNode
}

func CreateRouteNode(routePart string) *RouteNode {
	newNode := new(RouteNode)
	newNode.RoutePart = strings.TrimSpace(routePart)
	newNode.Children = make([]*RouteNode, 0)
	return newNode
}

func (rn *RouteNode) Insert(RouteParts []string) {
	if len(rn.Children) == 0 {
		// If the root node does not have any child nodes i.e., it is an empty tree
		newNode := CreateRouteNode(RouteParts[0])
		rn.Children = append(rn.Children, newNode)
		if len(RouteParts) > 1 {
			newNode.Insert(RouteParts[1:])
		}
	} else {
		// If the root node has one or more child nodes
		cnFound := false
		var cnNode *RouteNode
		for _, cn := range rn.Children {
			if strings.EqualFold(cn.RoutePart, RouteParts[0]) {
				cnFound = true
				cnNode = cn
				break
			}
		}

		if !cnFound {
			// If none of the child nodes of the root node had the first route part of the given route.
			cnNode = CreateRouteNode(RouteParts[0])
			rn.Children = append(rn.Children, cnNode)
			if len(RouteParts) > 1 {
				cnNode.Insert(RouteParts[1:])
			}
		} else {
			// If one of the child nodes of the root node contained the first route part of the given route.
			if len(RouteParts) > 1 {
				cnNode.Insert(RouteParts[1:])
			}
		}
	}
}

func (rn *RouteNode) Print() []string {
	routeParts := make([]string, 0)
	if len(rn.Children) != 0 {
		for _, cn := range rn.Children {
			childParts := cn.Print()
			for _, cp := range childParts {
				if rn.RoutePart != "" {
					routeParts = append(routeParts, fmt.Sprintf("%s/%s", rn.RoutePart, cp))
				} else {
					routeParts = append(routeParts, cp)
				}
			}
		}
	} else {
		routeParts = append(routeParts, rn.RoutePart)
	}

	return routeParts
}

func GetRouteParts(RoutePath string) []string {
	RoutePath = strings.ToLower(RoutePath)
	RoutePath = strings.TrimSpace(RoutePath)
	RoutePath = strings.TrimSuffix(RoutePath, "/")
	RoutePath = strings.TrimPrefix(RoutePath, "/")
	RouteParts := strings.Split(RoutePath, "/")
	NewRouteParts := make([]string, 0)
	for _, routePart := range RouteParts {
		routePart = strings.TrimSpace(routePart)
		if routePart != "" {
			NewRouteParts = append(NewRouteParts, routePart)
		}
	}

	return NewRouteParts
}

func AddRouteToTree(root *RouteNode, RoutePath string) {
	RouteParts := GetRouteParts(RoutePath)
	fmt.Println("The route parts for ", RoutePath, " are: ", RouteParts)
	root.Insert(RouteParts)
}

func CreateTree() *RouteNode {
	return CreateRouteNode("")
}