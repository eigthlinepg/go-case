package main

import (
	"sync"
)

type tv struct {
	tvName string
	count  int64
}

var once sync.Once
var t *tv

/*
* Function: 单例模式
* Description: function_description
* Parameters:
  - : param1_name - param1_description

* Returns: return_value
*/

/*func initTv() {*/
//t = tv{
//tvName: "kuangbiao",
//count:  39,
//}
//}

func GetInstance() *tv {

	if t == nil {
		once.Do(func() {

			t = &tv{
				tvName: "santi",
				count:  19,
			}
		})
	}
	return t
}
