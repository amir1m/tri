package todo

import (
	"encoding/json"
	"io/ioutil"

	"fmt"
)

type Item struct {
	Text string
  Priority int
}

func (i *Item) SetPriority(pri int)  {
  switch pri {
  case 1:
    i.Priority = 1
  case 2:
    i.Priority = 2
  case 3:
    i.Priority = 3
  default:
    i.Priority = 2

  }

}

func SaveItems(filename string, items []Item) error {

	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	err = ioutil.WriteFile(filename, b, 0644)

	if err != nil {
		return nil
	}

	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, nil
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	return items, nil

}
