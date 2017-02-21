package main

func main() {
	//ファイルの作成
	file1 := NewFile("file1")
	file2 := NewFile("file2")
	file3 := NewFile("file3")
	file4 := NewFile("file4")
	file5 := NewFile("file5")
	file6 := NewFile("file6")
	//ディレクトの作成
	dir1 := NewDirectory("dir1")
	dir1.Add(file1)
	dir1.Add(file2)
	dir1.Add(file3)
	dir2 := NewDirectory("dir2")
	dir2.Add(file4)
	dir2.Add(file5)
	dir2.Add(file6)

	dir1.Add(dir2)
	dir1.Remove()
}
