# Email Generator Module for Golang

This is a simple Go module that allows you to generate random email addresses. It's useful for various testing and development purposes where you need placeholder email addresses.

![Golang Logo](https://cdn-images-1.medium.com/max/892/1*skL4jh12fS0W9TZ2Hru-1g.jpeg)

## How to Use

1. Install the module by running:
   ```shell
   go get github.com/tamathecxder/randomail
   ```
2. Import the module in your Go code:
    ```go
    import (
        "fmt"
        "github.com/tamathecxder/randomail"
    )
    ```
3.  Generate random email addresses using the 
<code>GenerateRandomEmails</code> function:

    ```go
    emails := randomail.GenerateRandomEmails(10)

	for i, email := range emails {
		fmt.Printf("Random email #%d: %s\n", i+1, email)
	}
    ```

## Customization

You can customize the list of email domains by modifying the domains variable in the emailgenerator/emailgenerator.go file.

```go
var domains = []string{"example.com", "test.com", "yourdomain.com"}
```

## License

This module is distributed under the MIT License. You are free to use and modify it for your projects. If you find it helpful or have suggestions, feel free to contribute!

Happy email generation!
