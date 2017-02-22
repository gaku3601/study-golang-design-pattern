# Flyweightパターン概要
Flyweightとは、英語で「フライ級」を意味する単語です。  
フライ級とは、ボクシングなどの体重階級の中でもっとも軽い部類に分類できる階級です。  
もっと軽い階級もありますが。Flyweight パターンとは、インスタンスを共有することでリソースを無駄なく使うことに焦点を当てたパターンです。  

# ソースコード解説
Factoryをsingletonで作成し、そのfactoryを通して各インスタンスを生成することで、  
一意なインスタンスを生成することが可能となるように実装した。  

## 参考文献
TECHSCORE - 20．Flyweight パターン  
https://www.techscore.com/tech/DesignPattern/Flyweight.html/
