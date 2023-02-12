package ui


type Column struct {
	Properties *Properties
	Children   []UIElement
}

func (column Column) Draw(screen []byte) {
	
	column.Properties.Draw(screen)

	for child := range column.Children {
		
		column.Children[child].SetProperties(
			Size{
				Scale:  column.Properties.Size.Scale,
				Width:  column.Properties.Size.Width,
				Height: column.Properties.Size.Height / len(column.Children),
			},
			Point{
				X: column.Properties.Center.X,
				Y: column.Properties.Center.Y - column.Properties.MaxSize.Height/2 + (2*(len(column.Children) - child - 1)+1)*column.Properties.MaxSize.Height/(len(column.Children)*2),
			},
		)
		println("duh", column.Properties.Center.Y - column.Properties.MaxSize.Height/2 + (child+1)*column.Properties.MaxSize.Height/(len(column.Children)+1))
		column.Children[child].Draw(screen)
	}
}


func (column Column) SetProperties(size Size, center Point) {
	column.Properties.MaxSize = size
	column.Properties.Center = center
}

func (column Column) Debug() {
	println(column.Properties.Center.Y)
}
