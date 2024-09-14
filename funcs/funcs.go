package funcs

import (
	"html/template"
	"net/http"
)

func ExctTmple(w http.ResponseWriter, name string, data any) error {
	var tpl *template.Template

	tpl, err := template.ParseGlob("static/templates/*.html")
	if err != nil {
		return err
	}

	err = tpl.ExecuteTemplate(w, name, data)
	if err != nil {
		return err
	}
	return nil
}

func GetArtist(idTarget int, structTarget interface{}) interface{} {
	MyStruct, ok := structTarget.([]map[string]interface{})

	if ok {
		for _, v := range MyStruct {
			myint, ok := v["id"].(float64)
			if ok {
				if int(myint) == idTarget {
					return v
				}
			}
		}
	}

	return nil
}

func ReValueDates(dates []interface{}) []interface{} {
	var res []interface{}
	
	for _, v := range dates {
		if v.(string)[0] == '*' {
			// fmt.Println("test")
			v = v.(string)[1:]
		}
		res = append(res, v)
	}
	
	return res
}

func AddSliceToMap(MySlice []interface{}) []string {
	var res []string

	for _, v := range MySlice {
		var curStr string
		for i, ch := range v.(string) {
			if ch == '*' {
				continue
			}

			if i == len(v.(string))-1 {
				curStr += string(ch)
				res = append(res, curStr)
				curStr = ""
				break
			}

			curStr += string(ch)
		}
	}

	return res
}

func CheckErr(MySlice []error) bool {

	for _, v := range MySlice {
		if v != nil {
			return true
		}
	}

	return false
}
