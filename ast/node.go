package ast

import "strings"

type Node interface {
	String() string
}

func MapString(nodes []Node) []string {
	stringslice := make([]string, len(nodes))
	for i, n := range nodes {
		stringslice[i] = n.String()
	}
	return stringslice
}

func SliceToString[T Node](nodes []T, sep string) string {
	stringslice := make([]string, len(nodes))
	for i, n := range nodes {
		stringslice[i] = n.String()
	}
	return strings.Join(stringslice, sep)
}
