# prototypeパターン概要
最初にprototypeとなるインスタンスを生成し、そのprototypeをコピーし新たなインスタンスを作成するデザインパターン。  
主に、少しだけ違う形のインスタンスを大量に生成する必要がある場合に使用されるパターン？  

# ソースコード説明
今回はtext出力で接頭と語尾に適当な文字列を追加し、出力するという簡単な仕様で実装を行った。  
outputText.goではsingleラインで出力を行い、doubleText.goでは2ラインで出力するようになっている。  

## 参考文献
デザインパターン「Prototype」  
http://qiita.com/shoheiyokoyama/items/61826e158b3c4a259065
