# 算法与数据结构

[![Build Status](https://travis-ci.org/Thrimbda/dune.svg?branch=master)](https://travis-ci.org/Thrimbda/dune)

一个用 `Golang`实现的算法与数据结构的.

目前项目正在专注于基础的数据结构实现。

## 数据结构

基础的元数据均为实现了以下接口的任意数据结构。

```go
type Elem interface {
  Key() int
  String() string
}
```

### 基本数据结构

#### 线性表

线性表接口如下

```go
type List interface {
  setup(size int)
  clear()
  insert(item Elem)
  append(item Elem)
  remove() Elem
  SetFirst()
  next()
  prev()
  length() int
  SetPos(pos int)
  SetValue(value Elem)
  currValue() Elem
  isEmpty() bool
  isInList() bool
  print()
}
```

拥有链表与顺序表两种实现方式。

#### 栈

```go
type Stack interface {
  setup(size int)
  clear()
  push(item Elem)
  pop() Elem
  topValue() Elem
  isEmpty() bool
}
```

#### 队列

```go
type Queue interface {
  setup(size int)
  clear()
  enqueue(item Elem)
  dequeue() Elem
  firstValue() Elem
  isEmpty() bool
}
```

### 二叉树

二叉树节点接口

```go
type BinNode interface {
  Element() Elem
  SetElement(element Elem)
  Left() BinNode
  SetLeft(node BinNode)
  Right() BinNode
  SetRight(node BinNode)
  SetParent(node BinNode)
  Parent() BinNode
  IsLeaf() bool
}
```

#### 二叉搜索树 (BST)

```go
type BST interface {
  Insert(value Elem)
  Search(key int) BST
  Delete(key int)
  Predecessor() BST //寻找前驱节点
  Successor() BST //寻找后继
  Minimum() BST
  Maximum() BST
  InorderWalk() //中序遍历
}
```

二叉搜索树除基本实现以外，还有许多平衡树实现。

目前实现了红黑树。

#### 堆

由于堆为满二叉树，因此使用数组实现，故并未使用上述二叉树结点。

```go
type Heap interface {
  setup(size int)
  heapSize() int
  isLeaf(pos int) bool
  buildHeap()
  leftChild(pos int) int
  rightChile(pos int) int
  parent(pos int) int
  shiftDown(pos int)
  insert(val Elem)
  removeMax() Elem
  remove(pos int) Elem
}
```

