package common

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Page struct {
	PageNum    int         `json:"pageNum"`
	PageSize   int         `json:"pageSize"`
	TotalPage  int         `json:"totalPage"`
	TotalCount int64       `json:"totalCount"`
	List       interface{} `json:"list"`
}

func PageUtil(pageNum int, pageSize int, totalCount int64, list interface{}) Page {
	i := int64(pageSize)
	totalPage := totalCount / i
	if totalCount%i > 0 {
		totalPage = totalPage + 1
	}
	return Page{PageNum: pageNum, PageSize: pageSize, TotalPage: int(totalPage), TotalCount: totalCount, List: list}
}

func Succes(data interface{}) (result Result) {
	result.Code = 200
	result.Msg = "ok"
	result.Data = data
	return result
}

func Failed(code int, msg string) (result Result) {
	result.Code = code
	result.Msg = msg
	return result
}
