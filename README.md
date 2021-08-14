# custom vet


```
$ cd cmd/nosetter
$ go install
$ go vet -vettool=$(which nosetter)
```


# References
- https://agniva.me/vet/2019/01/21/vet-analyzer.html
- https://medium.com/a-journey-with-go/go-how-to-build-your-own-analyzer-f6d83315586f
- [How to add suggested fixers](https://arslan.io/2020/07/07/using-go-analysis-to-fix-your-source-code/)
