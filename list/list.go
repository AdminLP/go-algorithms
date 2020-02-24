package main

type List struct {
	head   *Item
	length int
}

type Item struct {
	value interface{}
	next  *Item
}

func New() *List {
	list := &List{}
	list.length = 0

	return list
}

func (list *List) First() *Item {
	return list.head
}

func (item *Item) Next() *Item {
	return item.next
}

func Insert(value interface{}, list *List) *List {
	item := &Item{
		value: value,
	}

	if list.head == nil {
		list.head = item
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = item
	}
	list.length++

	return list
}

func Has(value interface{}, list *List) bool {
	if list.head == nil {
		return false
	}

	current := list.First()

	for {
		if current.value == value {
			return true
		}

		if current.next != nil {
			current = current.next
		} else {
			return false
		}
	}
}

func Remove(value interface{}, list *List) *List {
	if list.head == nil {
		return list
	}

	current := list.First()

	var prev *Item

	for {
		if current.value == value {
			if prev != nil {
				prev.next = nil
			} else {
				list.head = nil
			}
			list.length--

			return list
		}

		if current.next != nil {
			prev = current
			current = current.next
		} else {
			return list
		}
	}
}

func Length(list *List) int {
	return list.length
}

func main() {}
