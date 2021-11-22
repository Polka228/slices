package main

func snake(row, acol int)[][]int{
	a:=make([][]int,row)
	for i:=range{
		a[i]=make([]int,col)
	}
	for i:=range a{
		for j :=range a[0]{
			x++
			a[i][j]=x
		}

	}
	return a
}

func main() {
	const row, col = 5, 7
	for _, v := range snake(row, col) {
		for _, x := range v {
			fmt.Printf("%3d", x)
		}
		fmt.Println()
	}
}
