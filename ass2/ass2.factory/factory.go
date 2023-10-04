package main

import "fmt"

type Sneakers interface {
	Name() string
}

type NikeSneakers struct{}

func (n *NikeSneakers) Name() string {
	return "Nike Sneakers"
}

type AdidasSneakers struct{}

func (a *AdidasSneakers) Name() string {
	return "Adidas Sneakers"
}

type AsicsSneakers struct{}

func (a *AsicsSneakers) Name() string {
	return "Asics Sneakers"
}

//all sneakers
type SneakersFactory interface {
	CreateSneakers() Sneakers
}

//only nike
type NikeFactory struct{}

func (n *NikeFactory) CreateSneakers() Sneakers {
	return &NikeSneakers{}
}

//only adidas
type AdidasFactory struct{}

func (a *AdidasFactory) CreateSneakers() Sneakers {
	return &AdidasSneakers{}
}

//only asics
type AsicsFactory struct{}

func (a *AsicsFactory) CreateSneakers() Sneakers {
	return &AsicsSneakers{}
}

func main() {

	nikeFactory := &NikeFactory{}
	adidasFactory := &AdidasFactory{}
	asicsFactory := &AsicsFactory{}

	nikeSneakers := nikeFactory.CreateSneakers()
	adidasSneakers := adidasFactory.CreateSneakers()
	asicsSneakers := asicsFactory.CreateSneakers()

	fmt.Println("Dayana's sneakers collection:")

	fmt.Println(nikeSneakers.Name())
	fmt.Println(adidasSneakers.Name())
	fmt.Println(asicsSneakers.Name())

}
