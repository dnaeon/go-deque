// Copyright (c) 2022 Marin Atanasov Nikolov <dnaeon@gmail.com>
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions
// are met:
// 1. Redistributions of source code must retain the above copyright
//    notice, this list of conditions and the following disclaimer
//    in this position and unchanged.
// 2. Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE AUTHOR(S) ``AS IS'' AND ANY EXPRESS OR
// IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
// OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE AUTHOR(S) BE LIABLE FOR ANY DIRECT, INDIRECT,
// INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
// THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package deque_test

import (
	"fmt"
	"testing"

	"gopkg.in/dnaeon/go-deque.v1"
)

func TestDeque(t *testing.T) {
	deque := deque.New[int]()
	if deque.Length() != 0 {
		t.Fatal("deque length should be 0")
	}

	// Insert a few items. Resulting deque is: 5 4 1 2 3
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3) // Back item
	deque.PushFront(4)
	deque.PushFront(5) // Front item

	if deque.IsEmpty() {
		t.Fatal("deque should not be empty")
	}

	if deque.Length() != 5 {
		t.Fatal("deque length should be 5")
	}

	backValue := 3
	back, err := deque.PeekBack()
	if err != nil {
		t.Fatal(err)
	}
	if back != backValue {
		t.Fatalf("item at the back should be %d", backValue)
	}

	frontValue := 5
	front, err := deque.PeekFront()
	if err != nil {
		t.Fatal(err)
	}
	if front != frontValue {
		t.Fatalf("item at the front should be %d", frontValue)
	}

	// Consume all items - front to back
	dequeItems := []int{5, 4, 1, 2, 3}
	i := 0
	for !deque.IsEmpty() {
		want := dequeItems[i]
		got, err := deque.PopFront()
		if err != nil {
			t.Fatal(err)
		}

		if want != got {
			t.Fatalf("want %d, got %d", want, got)
		}
		i++
	}

	// Deque should be empty at this point
	if !deque.IsEmpty() {
		t.Fatal("deque should be empty")
	}

	if deque.Length() != 0 {
		t.Fatal("deque length should be 0")
	}
}

func TestPopBack(t *testing.T) {
	deque := deque.New[int]()

	// Deque[int]: 1, 2, 3
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PushBack(3)

	// Consume all items - back to front
	dequeItems := []int{3, 2, 1}
	i := 0
	for !deque.IsEmpty() {
		want := dequeItems[i]
		got, err := deque.PopBack()
		if err != nil {
			t.Fatal(err)
		}

		if want != got {
			t.Fatalf("want %d, got %d", want, got)
		}
		i++
	}
}

func TestEmpty(t *testing.T) {
	deque := deque.New[int]()

	got, err := deque.PopBack()

	if err == nil || got != 0 {
		t.Fatal("expected error, but got nil")
	}

	got, err = deque.PopFront()

	if err == nil || got != 0 {
		t.Fatal("expected error, but got nil")
	}

	got, err = deque.PeekBack()

	if err == nil || got != 0 {
		t.Fatal("expected error, but got nil")
	}

	got, err = deque.PeekFront()

	if err == nil || got != 0 {
		t.Fatal("expected error, but got nil")
	}
}

func TestRotate(t *testing.T) {
	d := deque.New[int]()

	// Deque is: []int{1, 2, 3, 4, 5}
	for i := range 5 {
		d.PushBack(i + 1)
	}

	testCases := []struct {
		n    int
		want []int
	}{
		{n: 1, want: []int{2, 3, 4, 5, 1}},
		{n: -1, want: []int{1, 2, 3, 4, 5}},
		{n: 3, want: []int{3, 4, 5, 1, 2}},
		{n: -2, want: []int{5, 1, 2, 3, 4}},
	}

	if d.IsEmpty() {
		t.Fatal("deque should not be empty")
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("rotate with %d positions", tc.n), func(t *testing.T) {
			if err := d.Rotate(tc.n); err != nil {
				t.Error(err)
			}
		})
	}
}
