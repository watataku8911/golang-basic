package main

import (
	"./Pkg/dao"
	"./Pkg/entitys"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)


func deptList(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()
    // テンプレートをパース
	t := template.Must(template.ParseFiles(path + "/static/view/deptList.html.tpl"))


	//DB処理
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
	}

	defer db.Close()
	 deptList := dao.FindAll(db)
	 if deptList == nil {
	 	fmt.Println("から")
	 }

    // テンプレートを描画
    if err := t.ExecuteTemplate(w, "deptList.html.tpl", deptList); err != nil {
      log.Fatal(err)
    }

}

func add(w http.ResponseWriter, r *http.Request)  {
	path, _ := os.Getwd()
	// テンプレートをパース
	t := template.Must(template.ParseFiles(path + "/static/add/goDeptAdd.html.tpl"))
	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "goDeptAdd.html.tpl", nil); err != nil {
		log.Fatal(err)
	}
}

func addNow (w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w,"<h1>400 Bad Request</h1>" +
			"<p><a href='/'>TOPへ</a></p>")
		return
	}

	deptno := r.FormValue("addDeptDeptno")
	dname := r.FormValue("addDeptDname")
	loc := r.FormValue("addDeptLoc")
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
	}
	defer db.Close()

	var validation  entitys.Error
	if len(deptno) == 0 {

		validation.ValidationMsg = append(validation.ValidationMsg, "部門番号は必須です")
		validation.Flg = true
	}

	if len(dname) == 0 {
		validation.ValidationMsg = append(validation.ValidationMsg, "部門名は必須です")
		validation.Flg = true
	}

	count := dao.FindByPkCount(db, deptno)

	if count >= 1 {
		validation.ValidationMsg = append(validation.ValidationMsg, "部門番号が重複してます")
		validation.Flg = true
	}

	if validation.Flg {
		http.Redirect(w, r, "/add", http.StatusMovedPermanently)
		return
	} else {
			dao.Insert(db, deptno, dname, loc)

			http.Redirect(w, r, "/view", http.StatusMovedPermanently)
	}


}

func edit(w http.ResponseWriter, r *http.Request) {
	var deptno string
	deptno = r.FormValue("editDeptDeptno")
	if deptno == "" {
		fmt.Fprintln(w, "不正なアクセス")
	}
	fmt.Println(deptno)

	path, _ := os.Getwd()
	// テンプレートをパース
	t := template.Must(template.ParseFiles(path + "/static/edit/deptEdit.html.tpl"))
	//DB処理
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
	}

	defer db.Close()
	deptList := dao.FindByPk(db, deptno)
	if deptList == nil {
		fmt.Println("空")
	}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "deptEdit.html.tpl", deptList); err != nil {
		log.Fatal(err)
	}
}

func editNow(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w,"<h1>400 Bad Request</h1>" +
			"<p><a href='/'>TOPへ</a></p>")
		return
	}


	dname := r.FormValue("editDeptDname")
	loc := r.FormValue("editDeptLoc")
	deptno := r.FormValue("editDeptDeptno")
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
	}
	defer db.Close()


	dao.Update(db, dname, loc, deptno)


	http.Redirect(w, r, "/view", http.StatusMovedPermanently)

}


func delete(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w,"<h1>400 Bad Request</h1>" +
			"<p><a href='/'>TOPへ</a></p>")
		return
	}

	var deptno string
	deptno = r.FormValue("deleteDeptDeptno")
	path, _ := os.Getwd()
	// テンプレートをパース
	t := template.Must(template.ParseFiles(path + "/static/delete/confirmDeptDelete.html.tpl"))
	//DB処理
	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
	}

	defer db.Close()
	deptList := dao.FindByPk(db, deptno)
	if deptList == nil {
		fmt.Println("空")
	}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "confirmDeptDelete.html.tpl", deptList); err != nil {
		log.Fatal(err)
	}
}

func deleteNow(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w,"<h1>400 Bad Request</h1>" +
			"<p><a href='/'>TOPへ</a></p>")
		return
	}
	deptno := r.FormValue("deleteDeptDeptno")

	db, err := sql.Open("mysql", "scott:tiger@tcp(127.0.0.1:8889)/wp32scott")//通常：ポート番号３３０６、＊manp:8889
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("データベース接続失敗")
	}
	defer db.Close()

	dao.Delete(db,deptno)



	http.Redirect(w, r, "/view", http.StatusMovedPermanently)

}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))

	/**
	 *部門リスト
	 */
	http.HandleFunc("/view/", deptList)
	http.HandleFunc("/add/", add)
	http.HandleFunc("/add/now", addNow)
	http.HandleFunc("/edit", edit)
	http.HandleFunc("/edit/now", editNow)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/delete/now", deleteNow)

	/**
	 *従業員リスト
	 */

	http.ListenAndServe(":8000", nil)
}

