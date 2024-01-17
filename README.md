<!-- PROJECT LOGO -->

<div align="center">
  <a href="https://github.com/kaandesu/banana-converter-go">
    <img src="./logo.webp" alt="Banana Converter Logo" width="140">
  </a>

  <br>

  <!-- Title -->
  <h3 align="center">Banana Converter</h3>

  <!-- DESCRIPTION -->
  <p align="center">
    Convert lengths to and from bananas with this simple Golang library. Supports various units for convenient conversions. 
    <br />
    Simple, blazingly fast and free! 🍌
    <br />        
    <br />
    <a href="https://github.com/kaandesu/banana-converter-go/issues">Report Bug</a>
    <a href="https://github.com/kaandesu/banana-converter-go/issues">Request Feature</a>
  </p>
</div>

### Install

```bash
go get github.com/kaandesu/banana-converter-go
```

### Usage

```go
package main

import (
	"fmt"
	"github.com/kaandesu/banana-converter-go/converter"
)

func main() {
	result, err := converter.Convert(1.5, "banana", "cm")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %.2f\n", result)
	}
}
```

## Contributing

Contributions to the project is highly appreciated. If you have any suggestions/questions/requests please consider [opening an issue](https://github.com/kaandesu/banana-converter-go/issues/new). If you want to contribute to the project, fixing an open issue is greatly recommended and appreciated.

### Roadmap

- More banana options for **i18n** support:
  - [ ] Cavendish Banana.
  - [ ] Pisang Raja.
  - [ ] Red Banana.
  - [ ] Lady Finger Banana.
  - [ ] Blue Java Banana.
  - [ ] Plantain.
  - [ ] Manzano Banana.
  - [ ] Burro Banana.
  - [ ] Çikita

* [ ] Weight converter.
* [ ] Detailed documentation.

### Contact

| Maintainer                              |
| --------------------------------------- |
| [kaandesu](https://github.com/kaandesu) |
