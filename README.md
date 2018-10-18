# gram

_gram_ is a package for simulating TinyRAM (http://www.scipr-lab.org/specs.html) written in GO.


## build

```bash
go get github.com/LayerXcom/gram
dep ensure
cd $GOPATH/src/github.com/LayerXcom/gram
go build .
```

## usage
```bash
./gram -path example/testForSuccess.asm -r 100 -t 10 -pi "2" -ai "2"
```

## References

[1] http://www.scipr-lab.org/specs.html

[2] https://www.scipr-lab.org/doc/TinyRAM-spec-0.991.pdf

[3] Ben-Sasson, Eli, et al. "SNARKs for C: Verifying program executions succinctly and in zero knowledge." Advances in Cryptology–CRYPTO 2013. Springer, Berlin, Heidelberg, 2013. 90-108.
