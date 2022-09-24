package deque_test

import (
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
