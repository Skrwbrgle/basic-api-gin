package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Data struct{
  Name string ;
  Age int ;
}

func newData(name string, age int) Data {
  return Data{
    Name: name,
    Age: age,
  }
}

/**
its is comment
*/
func main() {
    
  r := gin.Default()

  var listData []Data;

  listData = append(listData, newData("redo", 22));
  listData = append(listData, newData("yoga", 25));

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "data": listData,
    })
  })        
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
