package main

func main() {
	bubble := NewBubbleSorter()
	quick := NewQuickSorter()
	timeSorter := NewTimeSorter(bubble)
	timeSorter.timeSort()

	timeSorter2 := NewTimeSorter(quick)
	timeSorter2.timeSort()

	/*
		出力結果
			計測を開始します。
			バブルソートを実行します。
			計測を終了します。
			計測を開始します。
			クイックソートを実行します。
			計測を終了します。
	*/
}
