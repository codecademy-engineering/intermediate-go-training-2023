# Creating a path with a parameter in Gin
```go
r.GET("/somepath/:myparam", myHandlerMethod)
```

# Reading a path parameter in Gin

```go
func (h *SomeStruct) GetUsingParam(c *gin.Context) {
	myParam := c.Param("myparam")
	// some logic to use the param
	c.JSON(200, myResult)
}
```
