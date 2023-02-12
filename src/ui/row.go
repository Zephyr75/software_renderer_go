package ui

type Row struct {
	Properties *Properties
	Children   []UIElement
}

func (row Row) Draw(screen []byte) {
	
	row.Properties.Draw(screen)

	for child := range row.Children {
		
		row.Children[child].SetProperties(
			Size{
				Scale:  row.Properties.Size.Scale,
				Width:  row.Properties.Size.Width / len(row.Children),
				Height: row.Properties.Size.Height,
			},
			Point{
				X: row.Properties.Center.X - row.Properties.MaxSize.Width/2 + (2*child+1)*row.Properties.MaxSize.Width/(len(row.Children)*2),
				Y: row.Properties.Center.Y,
			},
		)
		row.Children[child].Draw(screen)
	}
}


func (row Row) SetProperties(size Size, center Point) {
	row.Properties.MaxSize = size
	row.Properties.Center = center
}

func (row Row) Debug() {
	println(row.Properties.Center.Y)
}
