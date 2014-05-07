parser-css
==========

A simple CSS parser written in GO, that provides a way to parse CSS selectors and properties.

Installation
============

    go get github.com/mouradsabour/parser-css/src

Example
============

Property:
```go
type property struct {
        Name string
        Value string
}
```

Selector:
```go
type selector struct {
        Name string
        Properties []property
        Type int
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

	
	
	
	
