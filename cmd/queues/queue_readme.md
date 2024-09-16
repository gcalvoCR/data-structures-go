# Queues

A queue is a linear data structure that follows the First In, First Out (FIFO) principle. This means that the first element added to the queue is the first one to be removed. It's similar to a line of people waiting for service; the first person in line is the first to be served.

## Key Operations:
- Push: Adds an element to the top of the stack.
- Pop: Removes the element from the top of the stack.
- Peek/Top: Returns the top element of the stack without removing it.
- IsEmpty: Checks if the stack is empty.

## Characteristics:
- LIFO Order: The last element added is the first to be removed.
- Dynamic Size: The stack can grow and shrink as elements are added and removed.
- Access Time: Access to elements is restricted to the top of the stack.

## Applications:
- Function Call Management: Used to keep track of function calls in programming.
- Undo Mechanisms: Used in applications like text editors to implement undo features.
- Expression Evaluation: Used in parsing expressions, like in postfix or prefix notation evaluations.



Key Operations:
Enqueue: Adds an element to the end of the queue.
Dequeue: Removes the element from the front of the queue.
Front/Peek: Returns the element at the front of the queue without removing it.
IsEmpty: Checks if the queue is empty.
Characteristics:
FIFO Order: The first element added is the first to be removed.
Dynamic Size: The queue can grow and shrink as elements are added and removed.
Access Time: Access to elements is restricted to the front and end of the queue.
Applications:
Task Scheduling: Used in operating systems to manage tasks in scheduling algorithms.
Breadth-First Search (BFS): Used in graph traversal algorithms.
Print Spooling: Used to manage print jobs in the order they are received.