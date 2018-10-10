package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Enqueue_Empty(t *testing.T) {
	aq := animalQueue{}
	aq.enqueue(nil)

	require.Equal(t, 0, aq.dogs.Length())
	require.Equal(t, 0, aq.cats.Length())
}

func Test_Enqueue_InvalidType(t *testing.T) {
	aq := animalQueue{}
	aq.enqueue(animal{name: "test"})

	require.Equal(t, 0, aq.dogs.Length())
	require.Equal(t, 0, aq.cats.Length())
}

func Test_Enqueue_Normal(t *testing.T) {
	aq := animalQueue{}
	aq.enqueue(dog{animal: animal{name: "test"}})
	aq.enqueue(cat{animal: animal{name: "test2"}})

	require.Equal(t, 1, aq.dogs.Length())
	require.Equal(t, 1, aq.cats.Length())
}

func Test_DequeueAny_Normal(t *testing.T) {
	aq := animalQueue{}
	aq.enqueue(dog{animal: animal{name: "test"}})
	aq.enqueue(cat{animal: animal{name: "test2"}})

	iDog := aq.dequeueAny()
	dog, _ := iDog.(dog)

	require.Equal(t, "test", dog.name)
	require.Equal(t, 1, dog.order)

	iCat := aq.dequeueAny()
	cat, _ := iCat.(cat)
	require.Equal(t, "test2", cat.name)
	require.Equal(t, 2, cat.order)
}

func Test_DequeueAny_DequeueItemsBefore(t *testing.T) {
	aq := animalQueue{}
	aq.enqueue(dog{animal: animal{name: "test"}})
	aq.enqueue(cat{animal: animal{name: "test2"}})
	aq.enqueue(dog{animal: animal{name: "test3"}})

	dog1 := aq.dequeueDog()
	require.Equal(t, "test", dog1.name)
	require.Equal(t, 1, dog1.order)

	iCat := aq.dequeueAny()
	cat, _ := iCat.(cat)

	require.Equal(t, "test2", cat.name)
	require.Equal(t, 2, cat.order)

	iDog2 := aq.dequeueAny()
	dog2, _ := iDog2.(dog)
	require.Equal(t, "test3", dog2.name)
	require.Equal(t, 3, dog2.order)
}

func Test_DequeueDog_Normal(t *testing.T) {
	aq := animalQueue{}
	aq.enqueue(dog{animal: animal{name: "test"}})
	aq.enqueue(cat{animal: animal{name: "test2"}})

	dog := aq.dequeueDog()
	require.Equal(t, "test", dog.name)
	require.Equal(t, 1, dog.order)

	dog2 := aq.dequeueDog()

	require.Equal(t, "", dog2.name)
}

func Test_DequeueCat_Normal(t *testing.T) {
	aq := animalQueue{}
	aq.enqueue(dog{animal: animal{name: "test"}})
	aq.enqueue(cat{animal: animal{name: "test2"}})

	cat := aq.dequeueCat()
	require.Equal(t, "test2", cat.name)
	require.Equal(t, 2, cat.order)

	cat2 := aq.dequeueCat()

	require.Equal(t, "", cat2.name)
}

func Test_Dequeue_Empty(t *testing.T) {
	aq := animalQueue{}

	animal := aq.dequeueAny()

	require.Nil(t, animal)
}
