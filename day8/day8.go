package day8

import (
	"fmt"
	"slices"
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

	initialNodes := FindNodesEndingWith(nodeTree, "A")
	currentNodes := FindNodesEndingWith(nodeTree, "A")
	steps := 0
	lastCycles := map[int]int{}
	for !AreNodesEndingWith(currentNodes, "Z") {
		for _, instruction := range instructions {
			for i, node := range currentNodes {
				if node.name == initialNodes[i].name {
					_, exists := lastCycles[i]
					if !exists {
						lastCycles[i] = steps
					}
				}
			}
			if len(lastCycles) == 6 {
				break
			}
			steps += 1
			currentNodes = MapNodes(currentNodes, instruction, nodeTree)
		}

		if len(lastCycles) == 6 {
			fmt.Println(lastCycles)
			break
		}
		// 18157
		// 14363
	}
	mcm := make([][]int, 0)
	for _, value := range lastCycles {
		list := make([]int, 0)
		for i := 1; i < 1000; i++ {
			list = append(list, value*i)
		}
		mcm = append(mcm, list)
	}
	for _, value := range mcm[0] {
		total := 0
		for _, minCoMul := range mcm {
			if slices.Contains(minCoMul, value) {
				total += 1
			}
		}
		if total > 5 {
			return value
		}
	}
	solution := 1
	return solution
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
