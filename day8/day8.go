package day8

import (
	"fmt"
	"strings"
)

type Node struct {
	name  string
	left  string
	right string
}

func Solve1(input string) int {
	instructions := strings.Split(input, "\n\n")[0]
	nodeBlock := strings.Split(input, "\n\n")[1]
	nodes := strings.Split(nodeBlock, "\n")
	nodeTree := map[string]Node{}
	for _, node := range nodes {
		if node == "" {
			continue
		}
		mapNodes := strings.Split(strings.Split(node, " = ")[1], ", ")
		nodeTree[strings.Split(node, " = ")[0]] = Node{strings.Split(node, " = ")[0], mapNodes[0][1:], mapNodes[1][:3]}
	}

	currentNode := nodeTree["AAA"]
	steps := 0
	for currentNode != nodeTree["ZZZ"] {
		for _, instruction := range instructions {
			if currentNode == nodeTree["ZZZ"] {
				break
			}
			steps += 1
			if instruction == 'L' {
				currentNode = nodeTree[currentNode.left]
				continue
			}
			if instruction == 'R' {
				currentNode = nodeTree[currentNode.right]
				continue
			}
			fmt.Println("Shouldn't get here")
		}
	}
	return steps
}

func Solve2(input string) int {
	instructions := strings.Split(input, "\n\n")[0]
	nodeBlock := strings.Split(input, "\n\n")[1]
	nodes := strings.Split(nodeBlock, "\n")
	nodeTree := map[string]Node{}
	for _, node := range nodes {
		if node == "" {
			continue
		}
		mapNodes := strings.Split(strings.Split(node, " = ")[1], ", ")
		nodeTree[strings.Split(node, " = ")[0]] = Node{strings.Split(node, " = ")[0], mapNodes[0][1:], mapNodes[1][:3]}
	}

	currentNodes := FindNodesEndingWith(nodeTree, "A")
	steps := 0
	for !AreNodesEndingWith(currentNodes, "Z") {
		for _, instruction := range instructions {
			// fmt.Println(string(instruction))
			// fmt.Println(currentNodes)
			if AreNodesEndingWith(currentNodes, "Z") {
				break
			}
			steps += 1
			currentNodes = MapNodes(currentNodes, instruction, nodeTree)
			// fmt.Println(currentNodes)
		}
		fmt.Println(steps)
	}
	return steps
}

func FindNodesEndingWith(nodes map[string]Node, ending string) []Node {
	validNodes := make([]Node, 0)
	for _, node := range nodes {
		if string(node.name[len(node.name)-1]) == ending {
			validNodes = append(validNodes, node)
		}
	}
	return validNodes
}

func AreNodesEndingWith(nodes []Node, ending string) bool {
	for _, node := range nodes {
		if string(node.name[len(node.name)-1]) != ending {
			return false
		}
	}
	return true
}

func MapNodes(nodes []Node, instruction rune, nodeTree map[string]Node) []Node {
	mappedNodes := make([]Node, 0)
	if instruction == 'L' {
		for _, node := range nodes {
			mappedNodes = append(mappedNodes, nodeTree[node.left])
		}
		return mappedNodes
	}
	if instruction == 'R' {
		for _, node := range nodes {
			mappedNodes = append(mappedNodes, nodeTree[node.right])
		}
		return mappedNodes
	}
	fmt.Println("Should not get here.")
	return mappedNodes
}

func FindNode(nodeTree []Node, name string) Node {
	for _, node := range nodeTree {
		if node.name == name {
			return node
		}
	}
	fmt.Println("Should have found a node.")
	return nodeTree[0]
}
