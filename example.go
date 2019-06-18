package main
import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main(){
    r := gin.Default()
    r.LoadHTMLGlob("templates/*");
    r.GET("/ping",func(c *gin.Context){
        c.JSON(200, gin.H{
            "message":"pong",
        });
    })

    r.GET("/array", func(c *gin.Context){
        var values []int
        for i := 0; i < 5; i++ {
            values = append(values, i)
        }
        c.HTML(http.StatusOK, "array.tmpl", gin.H{"values": values})
    });
    
    r.Run()
}
