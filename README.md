# iranian-validation-go
Using this package you can validate below things.
#
## Features

### - Mobile Features
- validate iranian mobile
- get mobile operator's name
- get mobile's operator's persian name
- normalize number


## Installation

```sh
go get github.com/aminzohrabifar/iranian-validation-go
```

### Example in your projects

```go
package main

import (
	"fmt"
	"github.com/aminzohrabifar/iranian-validation-go"
	"github.com/aminzohrabifar/iranian-validation-go/src/mobile"
)

func main() {
	operatorname, err := mobile.GetMobileNumberOprator("  sww 916ddw 123dw 4567  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(operatorname)
	// print => mci

	IsMobileValid := mobile.IsMobileNumberValid("  sww 916ddw 123dw 4567  ")
	// IsMobileValid => true

	normalizedNumber := mobile.CleanMobileNumber("  sww 916ddw 123dw 4567  ")
	// print => 916123456
}
```

## Function Reference
#### To get mobile operator name, you need to use the following function:


```go
  operatorname,err := mobile.GetMobileNumberOprator(mobileNumber)
```

| Parameter | Type                  | Description                                     |
| :-------- |:----------------------|:------------------------------------------------|
| `operatorname` | `string` | return operator name                            |
| `err`      | `error`               | returns an error if passed values are not found |
| `mobileNumber` | `string`              | mobilenumber that you want to know              |



#### To check mobile validation, you need to use the following function:


```go
  isMobileNumberValid := mobile.IsMobileNumberValid(mobileNumber)
```

| Parameter | Type     | Description                        |
| :-------- |:---------|:-----------------------------------|
| `isMobileNumberValid` | `bool`   | check your mobile is valid or not  |
| `mobileNumber` | `string` | mobilenumber that you want to know |



#### To normalize and clean mobile number, you need to use the following function:


```go
  normalizedNumber := mobile.CleanMobileNumber(mobileNumber)
```

| Parameter | Type     | Description                        |
| :-------- |:---------|:-----------------------------------|
| `normalizedNumber` | `string` | cleaned mobile number              |
| `mobileNumber` | `string` | mobilenumber that you want to know |


## Authors

- [@Amin Zohrabi Far](https://www.linkedin.com/in/aminzohrabifar)