package process

import (
	"math"
	"math/rand"

	"regexp"
	"strings"

	"github.com/asaskevich/govalidator"
	"github.com/spf13/cast"
)

//https://blog.csdn.net/u010517268/article/details/112864274

type Manage struct {
	//级别还是算法
	Default string            //默认处理方法
	Dept    map[string]string //字段对应处理方法
}

func (m *Manage) Invalidation(bodys []string, category int, count int, repl string) []string {
	tmp := []string{}
	template := []string{}
	for i := 0; i < count; i++ {
		template = append(template, repl)
	}
	repstr := strings.Join(template, "")
	for _, body := range bodys {
		if len(body) > 0 {
			if len(body) < count {
				tmp = append(tmp, strings.Join(template, ""))
			} else {
				if category == 0 {
					tmp = append(tmp, strings.Replace(body, body[0:count], repstr, -1))
				} else {
					start := len(body) - count
					end := len(body)
					tmp = append(tmp, strings.Replace(body, body[start:end], repstr, -1))
				}
			}
		}
	}
	return tmp

}

func (m *Manage) Random(bodys []string) []string {
	results := []string{}
	for _, body := range bodys {
		tmp := []string{}
		if len(body) > 0 {
			for _, word := range body {
				row := string(word)
				if regexp.MustCompile("^[\u4e00-\u9fa5]{3,8}$").MatchString(row) {
					r := rand.Intn(500)
					row = CHINA[r]
				} else if govalidator.IsUTFLetter(row) {
					r := rand.Intn(26)
					row = ENGLISH[r]
				} else if govalidator.IsInt(row) {
					r := rand.Intn(10)
					row = NUMBER[r]
				}
				tmp = append(tmp, row)
			}

			results = append(results, strings.Join(tmp, ""))
		}
	}
	return results
}

func (m *Manage) DataReplacement(bodys []string) []string {
	count := 0
	category := "zh"
	for _, body := range bodys {
		row := string(body)
		if regexp.MustCompile("^[\u4e00-\u9fa5]{3,8}$").MatchString(row) {
			count += len(body)
			category = "zh"
		} else if govalidator.IsUTFLetter(row) {
			count += len(body)
			category = "en"
		} else if govalidator.IsInt(row) {
			count += len(body)
			category = "int"
		} else if govalidator.IsFloat(row) {
			count += len(body)
			category = "float"
		} else if govalidator.IsNumeric(row) {
			count += len(body)
			category = "number"
		} else {
			count += len(body)
		}
	}
	results := []string{}
	loop := count / len(bodys)
	for i := 0; i < len(bodys); i++ {
		switch category {
		case "zh":
			tmp := []string{}
			for i := 0; i < loop; i++ {
				r := rand.Intn(500)
				row := CHINA[r]
				tmp = append(tmp, row)
			}
			results = append(results, strings.Join(tmp, ""))
		case "en":
			tmp := []string{}
			for i := 0; i < loop; i++ {
				r := rand.Intn(26)
				row := ENGLISH[r]
				tmp = append(tmp, row)
			}
			results = append(results, strings.Join(tmp, ""))
		case "int":
			tmp := []string{}
			for i := 0; i < loop; i++ {
				r := rand.Intn(10)
				row := NUMBER[r]
				tmp = append(tmp, row)
			}
			results = append(results, strings.Join(tmp, ""))
		default:
			count = 10
			if count > loop {
				count = loop
			}
			results = m.Invalidation(bodys, 0, count, "*")

		}
	}
	return results
}

func (m *Manage) SymmetricEncryption(bodys []string) []string {
	results := []string{}
	for _, body := range bodys {
		if content, err := AesEncrypt(body); err == nil {
			results = append(results, string(content))
		} else {
			results = append(results, "")
		}
	}

	return results
}

func (m *Manage) AverageValue(bodys []string) []string {
	sum := 0
	for _, word := range bodys {
		if govalidator.IsFloat(word) || govalidator.IsIn(word) || govalidator.IsNumeric(word) {
			sum += cast.ToInt(word)
		}

	}
	if sum > 0 {
		result := []string{}
		for i := 0; i < len(bodys); i++ {
			result = append(result, cast.ToString(sum/len(bodys)))
		}
		return result
	}
	return bodys
}

func (m *Manage) OffsetRounding(bodys []string) []string {
	results := []string{}
	for _, body := range bodys {
		if govalidator.IsFloat(body) || govalidator.IsNumeric(body) {
			results = append(results, cast.ToString(math.Ceil(cast.ToFloat64(body))))
		}
	}
	return results
}

func (m *Manage) Swith(col string, body []string) []string {
	process := m.Default
	if column, ok := m.Dept[col]; ok {
		process = column
	}
	results := []string{}
	switch process {
	case "1":
		results = m.Invalidation(body, 0, 5, "*")
	case "2":
		results = m.Random(body)
	case "3":
		results = m.DataReplacement(body)
	case "4":
		results = m.SymmetricEncryption(body)
	case "5":
		results = m.AverageValue(body)
	case "6":
		results = m.OffsetRounding(body)
	default:
		results = m.Invalidation(body, 0, 5, "*")
	}
	return results
}
