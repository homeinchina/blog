package helper

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/cjyzwg/forestblog/models"
)

func HtmlTemplate(fileName string) (*template.Template, error) {

	return template.ParseFiles(
		"resources/views/"+fileName+".html",
		"resources/views/layouts/head.html",
		"resources/views/layouts/footer.html")
}

func ErrorHtml(errorInfo string) []byte {
	errorHtml := `
			<div style='width: 100%;height: 100vh;display: flex;justify-content: center;align-items: center;'>
				<p style='padding: 10px 20px;background-color: #d9534f;color:#fff;border-radius: 4px;text-align: center;'
				onmouseover="this.style.backgroundColor='#f0ad4e';"
				>` + errorInfo + " :(</p></div>"

	return []byte(errorHtml)
}

func WriteErrorHtml(w http.ResponseWriter, err string) {
	_, newErr := w.Write(ErrorHtml(err))
	if newErr != nil {
		panic(newErr)
	}
}

func SedResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write([]byte(`{"msg": "` + msg + `"}`))
	if err != nil {
		panic(err)
	}
}

func BuildArrByInt(num int) []int {
	var arr []int

	for i := 1; i <= num; i++ {
		arr = append(arr, i)
	}
	return arr
}

func UpdateArticle() {

	deleteCacheErr := os.RemoveAll("cache")
	if deleteCacheErr != nil {
		fmt.Println(deleteCacheErr)
	}

	//生成缓存
	_, err := models.GetMarkdownListByCache("/")

	if err != nil {
		log.Fatalf("生成缓存失败： %s\n", err)
	}
	return
}
