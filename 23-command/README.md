# Commandパターン概要
あるオブジェクトに対して要求を送るということは、そのオブジェクトのメソッドを呼び出すことと同じです。  
そして、メソッドにどのような引数を渡すか、ということによって要求の内容は表現されます。  
さまざまな要求を送ろうとすると、引数の数や種類を増やさなければなりませんが、それには限界があります。  
そこで要求自体をオブジェクトにしてしまい、そのオブジェクトを引数に渡すようにします。それがCommandパターンです。  

## 参考文献
TECHSCORE - 22．Commandパターン  
https://www.techscore.com/tech/DesignPattern/Command.html/  
まさるblog - デザインパターンを学ぶ～その21：Commandパターン(1)～  
http://blogs.wankuma.com/masaru/archive/2008/11/21/161907.aspx#165750  
