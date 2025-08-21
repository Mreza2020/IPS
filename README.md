# IPS (Image Processing Service)

This package analyzes the image and makes changes based on your wishes.
Before a detailed explanation, I must mention the licenses of this project:

The names of the packages used along with the licenses for each are as follows:

1-mysql    License  ==  [License](https://github.com/go-sql-driver/mysql?tab=MPL-2.0-1-ov-file "License mysql")
Mozilla Public License Version 2.0

---------------------------------------

2-gin      License  ==  [License](https://github.com/gin-gonic/gin?tab=MIT-1-ov-file "License gin") 
Copyright (c) 2014 Manuel Martínez-Almeida

---------------------------------------

3-uuid      License  ==  [License](https://github.com/google/uuid?tab=License-1-ov-file "License uuid") 
Copyright (c) 2009,2014 Google Inc. All rights reserved.

---------------------------------------

4-Imaging      License  ==  [License](https://github.com/disintegration/imaging?tab=MIT-1-ov-file "License Imaging")
Copyright (c) 2012 Grigory Dryapak


---------------------------------------

5-jwt-go      License  ==  [License](https://github.com/golang-jwt/jwt?tab=MIT-1-ov-file "jwt-go")
Copyright (c) 2012 Dave Grijalva
Copyright (c) 2021 golang-jwt maintainers

---------------------------------------

Well, after fully understanding this package, now it's time to learn how to use it:


Package features:

  * [Login](#Login)  
  * [Sign](#Sign)
  * [Upload](#Upload)
  * [Compress](#Compress)
  * [Crop](#Crop)
  * [ApplyFilter](#ApplyFilter)
  * [Flip](#Flip)
  * [ChangeFormat](#ChangeFormat)
  * [Resize](#Resize)
  * [Rotate](#Rotate)
  * [Watermark](#Watermark)


  
---------------------------------------
## Login
> [!IMPORTANT]
>The database of this project is built in Docker, so the code is local.

This information needs to be entered on the domain http://localhost:8080/IPS/login.
```json
{
  "username" : "admin",
  "password" : "admin"
}
```
Gives token code
The token time is 1 hour.

## Sign
This information needs to be entered on the domain http://localhost:8080/IPS/sign.

```json
{
  "username" : "admin",
  "password" : "admin"
}
```

## Upload
This information needs to be entered on the domain http://localhost:8080/IPS/Upload

```
    headers = {
    "Authorization": "Bearer {Token}",
    "Content-Type": "application/json"  
    }
```

```json
{
  "image" : "address"
}
```
Saves the file name in the database.


## Compress

This information needs to be entered on the domain http://localhost:8080/IPS/Compress

```
    headers = {
    "Authorization": "Bearer {Token}",
    "Content-Type": "application/json"  
    }
```

```json
{
  "file": "Build/10.png",
  "quality": "10"
}
```

## Crop

This information needs to be entered on the domain http://localhost:8080/IPS/Crop

```
    headers = {
    "Authorization": "Bearer {Token}",
    "Content-Type": "application/json"  
    }
```

```json
{
  "file": "Build/10.png",
  "width": "800",
  "height": "100",
  "x": "200",
  "y": "500"
}
```

## ApplyFilter

This information needs to be entered on the domain http://localhost:8080/IPS/ApplyFilter

```
    headers = {
    "Authorization": "Bearer {Token}",
    "Content-Type": "application/json"  
    }
```

```json
{
  "file": "Build/10.png",
  "filter": "grayscale"

}
```
It has 3 modes:

1- grayscale
2- sepia
3- invert

## Flip

This information needs to be entered on the domain http://localhost:8080/IPS/Flip

```
    headers = {
    "Authorization": "Bearer {Token}",
    "Content-Type": "application/json"  
    }
```

```json
{
  "file": "Build/10.png",
  "mode": "horizontal"

}
```

## ChangeFormat

This information needs to be entered on the domain http://localhost:8080/IPS/ChangeFormat

```
    headers = {
    "Authorization": "Bearer {Token}",
    "Content-Type": "application/json"  
    }
```

```json
{
  "file": "Build/10.png",
  "format": "jpg"

}
```




## Resize

This information needs to be entered on the domain http://localhost:8080/IPS/Resize

```
    headers = {
    "Authorization": "Bearer {Token}",
    "Content-Type": "application/json"  
    }
```

```json
{
  "file": "Build/10.png",
  "width": "500",
  "height": "500"
}
```


## Rotate

This information needs to be entered on the domain http://localhost:8080/IPS/Rotate

```
    headers = {
    "Authorization": "Bearer {Token}",
    "Content-Type": "application/json"  
    }
```

```json
{
  "file": "Build/Createananimeimag.png",
  "rotate": "90"
}
```



## Watermark

This information needs to be entered on the domain http://localhost:8080/IPS/Watermark

```
    headers = {
    "Authorization": "Bearer {Token}",
    "Content-Type": "application/json"  
    }
```

```json
{
  "file": "Build/10.png",
  "watermark": "Build/11.png",
  "opacity": "50"
}
```
