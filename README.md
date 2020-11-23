Hello,

You can find on this repository my test.
I disrespect the subject, but I try my best to show you my skill and my motivation. I have created a backend server to try to follow the good practice for the architecture and code. 

To run this project, you need to install 

Three dependencies:

Gin is the backend Framework, and gorm the ORM.

go get -u github.com/gin-gonic/gin

go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite


You can find three end points operational and one another not finish.



POST http://localhost:8080/article

With this body in raw

{
    "author":"hugo",
    "title": "hugo le test",
    "content": "les montagne sont belle5"
}


GET http://localhost:8080/articles
Whit no parameter


GET http://localhost:8080/articles/keyWord

With this body in raw

{
    "word":"plage"
}

Thank you for your time,
Looking forward to hearing from you 

Kind Regards 

Hugo