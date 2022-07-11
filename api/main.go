package main

import (
    "net/http"
    "github.com/goji/param"
)

type SignupForm struct {
    Name string `param:"name"`
    Email string `param:"email_address"`
    // We use a struct tag with "-" to ignore a value.
    Password string `param:"-"`
}

// FormHandler accepts a POST request, and would typically handle a HTML 
// form with a format like this:
// 
//  <form action="/signup/submit" method="POST">
//  <input name="name" type="text">
//  <input name="email_address" type="text">
//  <input name="password" type="password">
//  <input type="submit" value="Signup!">
//  </form>
//
// The 'name' attributes should match up with those of our struct fields. If 
// they don't, we use the aforementioned struct tags to translate them.
func FormHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "No good!", 400)
        return
    }

    var signupForm SignupForm{}
	
	// Parse url.Values (in this case, r.PostForm) and 
    // a pointer to our struct so that param can populate it.
    err := param.Parse(r.PostForm, &signupForm)
    if err != nil {
        http.Error(w, "Real bad.", 500)
        return
    }

    // Now we can:
    // - Perform some validation on our values
    // - Hash the user password with bcrypt or scrypt
    // - Store the results in our database
    // - (the world is our oyster!)
}