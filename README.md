# plt

グラフを生成するCLI。
***未完成***

## 使い方

```bash
plt testdata/dummy.tsv out.png
cat testdata/dummy.tsv | plt out.png

plt testdata/dummy.tsv testdata/dummy2.tsv testdata/dummy3.tsv out.png

plt testdata/dummy.tsv -t bar out.png
plt testdata/dummy.tsv -t line out.png

plt --separator , testdata/dummy.csv out.png
plt --width 1280 --height 760 testdata/dummy.tsv out.png

plt --d horizontal testdata/dummy.tsv out.png
plt --direction h testdata/dummy.tsv out.png
plt --direction vertical testdata/dummy.tsv out.png
plt --direction v testdata/dummy.tsv out.png

plt --noheader testdata/dummy.tsv out.png
plt --norowheader testdata/dummy.tsv out.png
```
