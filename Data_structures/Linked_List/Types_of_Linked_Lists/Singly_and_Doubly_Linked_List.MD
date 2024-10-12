# Singly & Doubly Linked List
## Singly Linked List
- Node structure:
  - Data
  - Pointer to the next node.
  
  ```go
  type Node struct {
    data int //store data
    next *Node // pointer to the next node
  }
  ```