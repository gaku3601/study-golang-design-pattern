# visitorパターン概要
Visitor とは、英語で「訪問者」を意味します。  
Visitor パターンでは、「処理」を訪問者であるVisitorオブジェクトに記述することで、処理の追加を簡単にします。  
処理対象となる、Acceptorオブジェクトは、Visitor オブジェクトを受け入れるaccept(Visitor visitor)メソッドを実装している必要があります。  

# ソースコード解説
今回は受け入れ側となる鈴木さん家へ新米先生と電気工事士が訪問するパターンを考えてみた。  
鈴木さん家はVisitorインターフェースを満たしている訪問者なら誰でも受け入れることができ、  
受けるサービスは各訪問者側で実装している。  
こうすることで、鈴木さん家は訪問者を受け入れるだけで、様々なサービスを受けることができるようになっている。  

## 参考文献
TECHSCORE - 13．Visitor パターン  
https://www.techscore.com/tech/DesignPattern/Visitor.html/
