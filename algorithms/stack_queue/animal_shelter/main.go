package main

import (
	"../queue"
)

type animal struct {
	order int
	name  string
}

type dog struct {
	animal
	bark string
}

type cat struct {
	animal
	cry string
}

type animalQueue struct {
	dogs  queue.Queue
	cats  queue.Queue
	order int
}

func (aq *animalQueue) enqueue(a interface{}) {
	aq.order++
	dog, ok := a.(dog)
	if ok {
		dog.order = aq.order
		aq.dogs.Add(dog)
		return
	}
	cat, ok := a.(cat)
	if ok {
		cat.order = aq.order
		aq.cats.Add(cat)
		return
	}
}
func (aq *animalQueue) dequeueAny() animal {
	if aq.dogs.Length() == 0 && aq.cats.Length() == 0 {
		return animal{}
	}

	if aq.dogs.Length() == 0 {
		return aq.dequeueCat().animal
	}
	if aq.cats.Length() == 0 {
		return aq.dequeueDog().animal
	}
	iDog, _ := aq.dogs.Peek()
	iCat, _ := aq.cats.Peek()
	dog, _ := iDog.(dog)
	cat, _ := iCat.(cat)
	if dog.order < cat.order {
		aq.dogs.Remove()
		return dog.animal
	}
	aq.cats.Remove()
	return cat.animal
}

func (aq *animalQueue) dequeueDog() *dog {
	result, err := aq.dogs.Remove()
	if err == nil {
		dog, ok := result.(dog)
		if ok {
			return &dog
		}
	}
	return nil
}

func (aq *animalQueue) dequeueCat() *cat {
	result, err := aq.cats.Remove()
	if err == nil {
		cat, ok := result.(cat)
		if ok {
			return &cat
		}
	}
	return nil
}

func main() {

}
