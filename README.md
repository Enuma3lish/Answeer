# Answeer
SQL is part of answear 1 and 2.
# API Folder 
In API folder is demand for api part written by Golang http and use docker to package my API.
test:http://localhost:8080/rate?source=USD&target=JPY&amount=1,525
result=> {"source":"USD","target":"JPY","amount":"170,496.52","msg":"success"}
# How to build the image
=> docker build -t currency-exchange-service .
# Experience share.
below is my answer.
# 1
=>目前沒有特別注意過有沒有使用過但是基本的精神coding時有符合.
SOLID 是OOP programing和design的五個基本原則的首字母縮寫，旨在提高軟體設計的靈活性、可維護性和可擴展性。
S:Single Responsibility Principle, (SRP）：一個class應該只有一個引起變化的原因，或者說一個class應該僅負責一個功能．
Open/Closed Principle, OCP）：軟體實體（class、module、function等）應該對擴展開放，對修改封閉。這意味著在不修改現有code的情況下可以對系統進行擴展。
Liskov Substitution Principle, LSP）：子類型必須能夠替換其基類型而不影響程序的正確性。這確保了繼承體系中的一致性，避免子類對父類行為的意外改變。
Interface Segregation Principle, ISP）：client端不應該被迫依賴於它們不使用的介面。換句話說，使用多個專門的介面，而不是一個通用的介面。
Dependency Inversion Principle, DIP）：high-level module 不應該依賴於low-level module，二者都應該依賴於抽象；抽象不應該依賴於細節，細節應該依賴於抽象。
The probelm ISP solve:如果一個介面定義了太多的方法，client端可能會被迫實現這些它們不需要的方法，導致不必要的代碼複雜度和耦合。
ISP=> 通過定義多個細粒度的介面，每個介面包含一組相關的方法，這樣client端只需要依賴它們真正需要的介面，從而減少不必要的依賴和耦合。
The probelm DIP soleve: high-level-module 直接依賴於low-level-module，這導致high-level-module的變化會影響low-level-module，反之亦然。這種緊密耦合使得系統難以擴展和維護。
DIP=> 通過引入抽象層（介面或抽象類），使高層模塊依賴於抽象而不是具體實現。低層模塊實現這些抽象，從而實現依賴注入。這樣，高層模塊和低層模塊之間的耦合減少，系統的靈活性和可維護性提高。
conclusion:ISP 通過精細化介面定義減少了不必要的依賴，而 DIP 通過依賴抽象而非具體實現來降低模塊之間的耦合。
# 2
因為避免進低程式的可讀性所以通常只有在比較簡單的函數會使用FP去取代攏長的程式碼。
Class 是一個template，用於定義對象的數據結構和行為。class包含屬性（數據字段）和方法（函數）。

實體（Instance)是類別的具體實現，也就是一個object。實體是根據class創建的，擁有類別中定義的屬性和方法。
# 3
定義介面（Interface）的意義在於它提供了一種抽象的方式來描述一組相關的方法，而不涉及這些方法的具體實現。
多型性（Polymorphism）是物件導向程式設計（OOP）的一個核心概念，允許一個介面可以有多種不同的實現。多型性使得一個方法可以針對不同的對象類型進行操作，而無需知道對象的具體類型．
# 4
Git rebase 重新整合分支歷史，使提交歷史更加線性和整潔。
Git merge fast-forward 模式簡單地移動分支指針，不創建新的合併提交。
Git merge non-fast-forward 模式在兩個分支合併時創建一個新的合併提交，保留了分支的完整歷史。