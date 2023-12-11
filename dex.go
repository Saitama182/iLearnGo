package main

import (
	"fmt"
	"math"
)

// Graph представляет собой структуру данных для представления графа.
type Graph struct {
	nodes map[string]map[string]int // Карта смежности для хранения ребер и их весов.
}

// Dijkstra реализует алгоритм Дейкстры для нахождения кратчайших путей в графе.
func (g *Graph) Dijkstra(startNode string) map[string]int {
	// Инициализация карты для хранения расстояний от начальной вершины до остальных вершин.
	distances := make(map[string]int)

	// Инициализация карты для отслеживания посещенных вершин.
	visited := make(map[string]bool)

	// Инициализация очереди для обработки вершин.
	queue := make(map[string]int)

	// Инициализация расстояния от начальной вершины до самой себя равным 0.
	distances[startNode] = 0

	// Инициализация расстояний до всех остальных вершин как бесконечность.
	for node := range g.nodes {
		if node != startNode {
			distances[node] = math.MaxInt64
		}
	}

	// Добавление начальной вершины в очередь.
	queue[startNode] = 0

	// Основной цикл алгоритма Дейкстры.
	for len(queue) > 0 {
		// Находим вершину с минимальным расстоянием из очереди.
		var currentNode string
		minDistance := math.MaxInt64
		for node, dist := range queue {
			if dist < minDistance {
				minDistance = dist
				currentNode = node
			}
		}

		// Удаляем текущую вершину из очереди.
		delete(queue, currentNode)

		// Помечаем текущую вершину как посещенную.
		visited[currentNode] = true

		// Обновляем расстояния до соседних вершин.
		for neighbor, weight := range g.nodes[currentNode] {
			if !visited[neighbor] {
				// Рассчитываем новое расстояние.
				newDistance := distances[currentNode] + weight

				// Если новое расстояние меньше текущего, обновляем его.
				if newDistance < distances[neighbor] {
					distances[neighbor] = newDistance

					// Добавляем соседнюю вершину в очередь для дальнейшей обработки.
					queue[neighbor] = newDistance
				}
			}
		}
	}

	// Возвращаем карту с кратчайшими расстояниями до каждой вершины.
	return distances
}

func main() {
	// Пример использования алгоритма Дейкстры.

	// Создаем новый граф.
	graph := Graph{
		nodes: map[string]map[string]int{
			"A": {"B": 1, "C": 4},
			"B": {"A": 1, "C": 2, "D": 5},
			"C": {"A": 4, "B": 2, "D": 1},
			"D": {"B": 5, "C": 1},
		},
	}

	// Задаем начальную вершину для алгоритма.
	startNode := "A"

	// Вызываем алгоритм Дейкстры.
	shortestPaths := graph.Dijkstra(startNode)

	// Выводим результаты.
	for node, distance := range shortestPaths {
		fmt.Printf("Shortest path from %s to %s is %d\n", startNode, node, distance)
	}
}
