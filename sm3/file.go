package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

type withFile struct {
	f   *os.File
	len int
}

const FILENAME = "links.json"

func (w *withFile) Push(link string, opts ...PushOption) error {
	i := Item{
		Date: time.Now(),
		Link: link,
	}

	for _, opt := range opts {
		opt(&i)
	}

	var (
		stateJSON []byte
		err       error
	)
	if w.len != 0 {
		stateJSON, err = os.ReadFile(FILENAME)
		if err != nil {
			return fmt.Errorf("failed to read from file: %w", err)
		}
	}

	file, err := os.OpenFile(FILENAME, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println("failed to close file: " + err.Error())
		}
	}()

	var state []Item
	if len(stateJSON) != 0 {
		if err = json.Unmarshal(stateJSON, &state); err != nil {
			return fmt.Errorf("failed to unmarshal json file: %w", err)
		}
	}

	state = append(state, i)

	b, err := json.Marshal(state)
	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	_, err = file.Write(b)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	w.len++
	return nil
}

func (w *withFile) Delete(link string, opts ...Option) (int, error) {
	prms := Params{}

	for _, prm := range opts {
		prm(&prms)
	}

	stateJSON, err := os.ReadFile(FILENAME)
	if err != nil {
		return 0, fmt.Errorf("failed to read from file: %w", err)
	}

	var state []Item
	if len(stateJSON) != 0 {
		if err = json.Unmarshal(stateJSON, &state); err != nil {
			return 0, fmt.Errorf("failed to unmarshal json file: %w", err)
		}
	}

	for i := 0; i < len(state); i++ {
		if state[i].Link == link || prms.Tag != nil && slices.Contains(state[i].Tags, *prms.Tag) {
			if prms.Limit != nil {
				if *prms.Limit <= prms.currentLimit {
					break
				}
			}
			if prms.Offset != nil {
				if *prms.Offset > prms.currentOffset {
					prms.currentOffset++
					continue
				}
			}
			state = append(state[:i], state[i+1:]...)
			i--
			prms.currentLimit++
		}
	}

	b, err := json.Marshal(state)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal value: %w", err)
	}

	file, err := os.OpenFile(FILENAME, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return 0, fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println("failed to close file: " + err.Error())
		}
	}()

	_, err = file.Write(b)
	if err != nil {
		return 0, fmt.Errorf("failed to write to file: %w", err)
	}

	w.len--
	return prms.currentLimit, nil
}

func (w *withFile) Get(link string, opts ...Option) ([]Item, error) {
	prms := Params{}

	for _, prm := range opts {
		prm(&prms)
	}

	var (
		stateJSON []byte
		err       error
	)
	if w.len != 0 {
		stateJSON, err = os.ReadFile(FILENAME)
		if err != nil {
			return nil, fmt.Errorf("failed to read from file: %w", err)
		}
	}

	var state []Item
	if len(stateJSON) != 0 {
		if err = json.Unmarshal(stateJSON, &state); err != nil {
			return nil, fmt.Errorf("failed to unmarshal json file: %w", err)
		}
	}

	var res []Item
	for _, i := range state {
		if i.Link == link || prms.Tag != nil && slices.Contains(i.Tags, *prms.Tag) {
			if prms.Limit != nil {
				if *prms.Limit <= prms.currentLimit {
					break
				}
			}
			if prms.Offset != nil {
				if *prms.Offset > prms.currentOffset {
					prms.currentOffset++
					continue
				}
			}
			res = append(res, i)
			prms.currentLimit++
		}
	}

	return res, nil
}

func (w *withFile) Length() int {
	return w.len
}

func (w *withFile) All() ([]Item, error) {
	var (
		stateJSON []byte
		err       error
	)
	if w.len != 0 {
		stateJSON, err = os.ReadFile(FILENAME)
		if err != nil {
			return nil, fmt.Errorf("failed to read from file: %w", err)
		}
	}

	var state []Item
	if len(stateJSON) != 0 {
		if err = json.Unmarshal(stateJSON, &state); err != nil {
			return nil, fmt.Errorf("failed to unmarshal json file: %w", err)
		}
	}

	return state, nil
}

func (w *withFile) Clear() error {
	file, err := os.OpenFile(FILENAME, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	if err = file.Close(); err != nil {
		log.Println("failed to close file: " + err.Error())
	}

	w.len = 0

	return nil
}
