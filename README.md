Gocss
==========

Gocss is a simple css parser written in GO, that provides a way to parse css selectors and properties.

Installation
============

```cmd
    $>go get github.com/m2omou/parser-css/src
```

```go
    import "github.com/m2omou/parser-css/src"
```
Code examples
============

####Type:####

The different type of selector, 'All' stand for the '*' css selector

```go
const ( 
	Class = 1
	Id
	All
	Element
	Attribute
)
```

####Property:####

e.g: "color: white;"

```go
type property struct {
        Name string 		// color
        Value string 		// white
}
```

####Selector:####

e.g ".container { color: white, float: left }"

```go
type selector struct {
        Name string             // '.container'
        Properties []property   // [ { Name : 'color', Value : 'white' }, { Name : 'float', Value : 'left'} ]
        Type int                // Class
}
```

####
```go
fd, err := os.Open("./style.css")
    
if err != nil {
        panic(err)
}
        
css := parserCSS(fd)
for s := range selectors {
        fmt.Println(selectors[s].Name)
        for p := range selectors[s].Properties {
                fmt.Println(selectors[s].Properties[p].Name + selectors[s].Properties[p].Value)
        }
}
```

For example this css :

```css
body {
    width: 100%;
    height: 100%;
    Background: red;
}

.container {
    margin-bottom: 20px;
    height: 220px;
    width:300px;
    font-size: 14px;
}
```
Will return the following slice :

```json
[
   {
      "Name":"body",
      "Properties":[
         {
            "Name":"width",
            "Value":"100%"
         },
         {
            "Name":"height",
            "Value":"100%"
         },
         {
            "Name":"Background",
            "Value":"red"
         }
      ],
      "Type":1
   },
   {
      "Name":".container",
      "Properties":[
         {
            "Name":"margin-bottom",
            "Value":"20px"
         },
         {
            "Name":"height",
            "Value":"220px"
         },
         {
            "Name":"width",
            "Value":"300px"
         },
         {
            "Name":"font-size",
            "Value":"14px"
         }
      ],
      "Type":1
   }
]
```	

	
	
	
	
