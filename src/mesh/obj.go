package mesh

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"overdrive/src/geometry"
	"overdrive/src/material"
)

// Creates a mesh from a .obj file with the given name and assigns it a given material
func ReadObjFile(name string, mtl material.Material) Mesh {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var vertices []geometry.Vector3
	var coordinatesX []float32
	var coordinatesY []float32
	var triangles []geometry.Triangle

	for scanner.Scan() {
		// Add vertices to the list
		if scanner.Text()[0] == 'v' && scanner.Text()[1] == ' ' {
			var vertex geometry.Vector3
			fmt.Sscanf(scanner.Text(), "v %f %f %f", &vertex.X, &vertex.Y, &vertex.Z)
			vertex.LightAmount = color.Black
			vertices = append(vertices, vertex)
		}
		// Add texture coordinates to the list
		if scanner.Text()[0] == 'v' && scanner.Text()[1] == 't' {
			var x, y float32
			fmt.Sscanf(scanner.Text(), "vt %f %f", &x, &y)
			coordinatesX = append(coordinatesX, x)
			coordinatesY = append(coordinatesY, y)
		}
		// Create triangles from the vertices and texture coordinates
		if scanner.Text()[0] == 'f' {
			var face geometry.Triangle
			var v1, v2, v3 int
			var t1, t2, t3, t4, t5, t6 int
			fmt.Sscanf(scanner.Text(), "f %d/%d/%d %d/%d/%d %d/%d/%d", &v1, &t1, &t2, &v2, &t3, &t4, &v3, &t5, &t6)
			vertices[v1-1].U = coordinatesX[t1-1]
			vertices[v1-1].V = coordinatesY[t1-1]
			vertices[v2-1].U = coordinatesX[t3-1]
			vertices[v2-1].V = coordinatesY[t3-1]
			vertices[v3-1].U = coordinatesX[t5-1]
			vertices[v3-1].V = coordinatesY[t5-1]

			face.A = vertices[v1-1]
			face.B = vertices[v2-1]
			face.C = vertices[v3-1]
			triangles = append(triangles, face)
			// fmt.Println(coordinatesX[t1-1], coordinatesY[t1-1])
			// fmt.Println(face.A.U, face.A.V)
			// fmt.Println("----------------")
		}

		// fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return NewMesh(triangles, geometry.ZeroVector(), geometry.ZeroVector(), mtl)
}
