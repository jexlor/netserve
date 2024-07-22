## Netserve 🪄

Getting started with Netserve.
----------------------------
```bash
git clone https://github.com/jexlor/netserve.git
```

First step.
----------------------------
```bash
cd netserve
```
Run:
----------------------------
```bash
go run main.go
```
Enter destination url:
---------------------------
![Screenshot from 2024-07-21 18-08-17](https://github.com/user-attachments/assets/5c857cce-51b8-48bf-b842-e5b85f451eaa)

Serve content.
---------------------------
```bash
bash serve_content.sh
```
---------------------------
<strong>Netserve</strong> is mostly used for E-books and such static sites which can work without internet. You can download some books and then access it anytime in your localhost even if you are offline.
You don't have to worry about site structure as <strong>Netserve</strong> will do it for you.

<strong>Note!</strong> Script will only run if you have Go 1.22.0 installed (or newer version).
The script itself is only supported for unix like platforms for now.
