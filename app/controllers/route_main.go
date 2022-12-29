package controllers

import (
	"log"
	"net/http"

	"sample_todo/app/models"
)

func top(w http.ResponseWriter, r *http.Request){
	_,err:= session(w,r)
	if err != nil{
		generateHTML(w,"Hello","layout","public_navbar","top")
	}else {
		http.Redirect(w, r, "/todos", redirectCode)

	}
}

func index(w http.ResponseWriter, r *http.Request){
	sess,err:= session(w,r)
	if err != nil{
		http.Redirect(w, r, "/", redirectCode)
	}else{
		user,err := sess.GetUserBySession()
		if err != nil{
			log.Println(err)
		}
		todos,_ := user.GetTodosByUser()
		user.Todos = todos
		generateHTML(w,user,"layout","private_navbar","index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request){
	// セッション確認
	_,err:= session(w,r)
	if err != nil{
		http.Redirect(w, r, "/login", redirectCode)
	}else{
		generateHTML(w,nil,"layout","private_navbar","todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request){
	// セッション確認
	sess,err:= session(w,r)
	if err != nil{
		http.Redirect(w, r, "/login", redirectCode)
	}else{
		err := r.ParseForm()
		if err != nil{
			log.Println(err)
		}
		user,err := sess.GetUserBySession()
		if err != nil{
			log.Println(err)
		}
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content);err !=nil {
			log.Println(err)
		}
		http.Redirect(w, r,"/todos",redirectCode)
	}

}

func todoEdit(w http.ResponseWriter, r *http.Request,id int){
	sess,err:= session(w,r)
	if err != nil{
		http.Redirect(w, r, "/login", redirectCode)
	}else{
		_,err := sess.GetUserBySession()
		if err != nil{
			log.Println(err)
		}
		t,err := models.GetTodo(id)
		if err != nil{
			log.Println(err)
		}
		generateHTML(w,t,"layout","private_navbar","todo_edit")

	}
}

func todoUpdate(w http.ResponseWriter, r *http.Request,id int){
	sess,err:= session(w,r)
	if err != nil{
		http.Redirect(w, r, "/login", redirectCode)
	}else{
		err := r.ParseForm()
		if err != nil{
			log.Println(err)
		}
		user,err := sess.GetUserBySession()
		if err != nil{
			log.Println(err)
		}
		content := r.PostFormValue("content")
		t:= &models.Todo{ID:id,Content:content,UserID:user.ID}
		if err := t.UpdateTodo(); err != nil{
			log.Println(err)
		}
		http.Redirect(w, r,"/todos",redirectCode)
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request,id int){
	sess,err:= session(w,r)
	if err != nil{
		http.Redirect(w, r, "/login", redirectCode)
	}else{
		_,err := sess.GetUserBySession()
		if err != nil{
			log.Println(err)
		}
		t,err:= models.GetTodo(id)
		if err != nil{
			log.Println(err)
		}
		if err := t.DeleteTodo(); err != nil{
			log.Println(err)
		}
		http.Redirect(w, r,"/todos",redirectCode)


	}
}