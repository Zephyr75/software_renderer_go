package mesh

import (
	"bufio"
	"fmt"
	"image/color"
	"log"
	"os"
	"overdrive/geometry"
)

//read file suzanne.obj in folder obj
func ReadObjFile() Mesh {
	file, err := os.Open("obj/suzanne.obj")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()



	scanner := bufio.NewScanner(file)

	var vertices []geometry.Vector3
	
	var triangles []geometry.Triangle


	for scanner.Scan() {
		if scanner.Text()[0] == 'v' && scanner.Text()[1] == ' ' {
			var vertex geometry.Vector3
			fmt.Sscanf(scanner.Text(), "v %f %f %f", &vertex.X, &vertex.Y, &vertex.Z)
			vertex.LightAmount = color.Black
			vertices = append(vertices, vertex)
		}
		if scanner.Text()[0] == 'f' {
			var face geometry.Triangle
			var v1, v2, v3 int
			var t1, t2, t3, t4, t5, t6 int
			fmt.Sscanf(scanner.Text(), "f %d/%d/%d %d/%d/%d %d/%d/%d", &v1, &t1, &t2, &v2, &t3, &t4, &v3, &t5, &t6)
			face.A = vertices[v1-1]
			face.B = vertices[v2-1]
			face.C = vertices[v3-1]
			triangles = append(triangles, face)
			fmt.Println(face)
		}

		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

	return NewMesh(triangles, geometry.ZeroVector(), geometry.ZeroVector())
	
}