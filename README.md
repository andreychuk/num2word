
num2word
========

[![GoCard][1]][2]

[1]: https://goreportcard.com/badge/github.com/andreychuk/num2word
[2]: https://goreportcard.com/report/github.com/andreychuk/num2word

[num2word](https://github.com/andreychuk/num2word) - Числа прописью

golang package для написания сумм прописью.

Features
--------

В текущей версии реализовано:

### UaAmount - гроші прописом українською
### RuAmount - гроші прописом російською

* Оригинальное название: number2word
* Источник на SQL: http://oraclub.trecom.tomsk.su/db/web.page?pid=461
* В конференцию relcom.comp.dbms.oracle поместил "Igor Volkov" (volkov@rdtex.msk.ru)
* Fork from [LeKovr/num2word](https://github.com/LeKovr/num2word) 

Usage
-----

```
import (
  "github.com/andreychuk/num2word"
)

...

text := num2word.UaAmount(total, true)
```

License
-------

The MIT License (MIT), see [LICENSE](LICENSE).
