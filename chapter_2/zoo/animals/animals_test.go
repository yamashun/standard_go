package animals

import (
	"testing"
)

func TestElephantFood(t *testing.T) {
	expect := "Grass"
	actual := ElephantFood()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestMonkeyFood(t *testing.T) {
	expect := "Banana"
	actual := MonkeyFood()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}

func TestRabbitFood(t *testing.T) {
	expect := "Carrot"
	actual := RabbitFood()

	if expect != actual {
		t.Errorf("%s != %s", expect, actual)
	}
}
