package gomonkey_demo

type PrivateMethodStruct struct {
}

func (item *PrivateMethodStruct) ok() bool {
	return item != nil
}

func (item *PrivateMethodStruct) Happy() string {
	if item.ok() {
		return "happy"
	}
	return "unhappy"
}

func (item PrivateMethodStruct) haveEaten() bool {
	return item != PrivateMethodStruct{}
}

func (item PrivateMethodStruct) AreYouHungry() string {
	if item.haveEaten() {
		return "I am full"
	}

	return "I am hungry"
}
