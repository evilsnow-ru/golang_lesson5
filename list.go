//Package contains non thread-safe LinkedList implementation
package otus_lesson5

const (
	linksNone uint8 = iota
	linksLeft
	linksBoth
	linksRight
)

type appendStrategy uint8

const (
	front appendStrategy = iota
	back
)

type Item struct {
	nextItem *Item
	prevItem *Item
	val interface{}
	container *List
}

func (item *Item) Value() interface{} {
	return item.val
}

func (item *Item) Next() *Item {
	return item.nextItem
}

func (item *Item) Prev() *Item {
	return item.prevItem
}

//Setter methods to avoid constructions like "item.prevItem.nextItem = ..."
func (item *Item) setNextItem(nextItem *Item) {
	item.nextItem = nextItem
}

func (item *Item) setPreviousItem(prevItem *Item) {
	item.prevItem = prevItem
}

func (item *Item) Remove() {
	if item.container != nil {
		//Check next & right Item reference state
		switch item.references() {
		case linksBoth:
			nextItem := item.nextItem
			previousItem := item.prevItem

			previousItem.setNextItem(nextItem)
			nextItem.setPreviousItem(previousItem)

		//This item is only one on the container
		case linksNone:
			item.container.lastItem = nil
			item.container.firstItem = nil

		case linksLeft:
			item.prevItem.setNextItem(nil)
			item.container.lastItem = item.prevItem

		case linksRight:
			item.nextItem.setPreviousItem(nil)
			item.container.firstItem = item.nextItem
		}

		//Decrement container list items count & remove reference to them to avoid multiple Remove() calls
		item.container.items--
		item.container = nil
		item.prevItem = nil
		item.nextItem = nil
	}
}

func (item *Item) references() uint8 {
	if item.prevItem == nil && item.nextItem == nil {
		return linksNone
	}

	if item.nextItem == nil {
		return linksLeft
	}

	if item.prevItem == nil {
		return linksRight
	}

	return linksBoth
}

type List struct {
	items int
	firstItem *Item
	lastItem *Item
}

func (list *List) First() *Item {
	return list.firstItem
}

func (list *List) Last() *Item {
	return list.lastItem
}

func (list *List) Len() int {
	return list.items
}

func (list *List) PushFront(v interface{}) {
	appendItem(front, list, v)
}

func (list *List) PushBack(v interface{}) {
	appendItem(back, list, v)
}

func appendItem(strategy appendStrategy, list *List, v interface{}) {
	newItem := &Item{val: v, container: list}

	if list.items == 0 {
		list.firstItem = newItem
		list.lastItem = newItem
	} else {
		switch strategy {
		case front:
			newItem.setNextItem(list.firstItem)
			list.firstItem.setPreviousItem(newItem)
			list.firstItem = newItem

		case back:
			newItem.setPreviousItem(list.lastItem)
			list.lastItem.setNextItem(newItem)
			list.lastItem = newItem
		}
	}

	list.items++
}

//Named as NewList because package name is "otus_lesson5"
func NewList() *List {
	return &List{}
}