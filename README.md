## Netserve 🪄

Getting started with Netserve.
----------------------------
```bash
git clone https://github.com/jexlor/netserve.git
```

First step.
----------------------------
```bash
cd Godoc
```
```bash
open main.go
```
Scroll down and find 'urlStr' variable 
```go
urlStr := "example.com" //paste url here
```
Run:

```bash
go run main.go
```
Serve content.
---------------------------
```bash
bash serve_content.sh
```

<strong>Netserve</strong> is mostly used for E-books and such static sites which can work without internet. You can download some books and then access it anytime in your localhost even if you are offline.
You don't have to worry about site structure as <strong>Netserve</strong> will do it for you.
