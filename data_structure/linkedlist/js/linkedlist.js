class Node {

  constructor(val, next = null) {
    this.value = val
    this.next = next
  }

}

class Linkedlist {
  constructor(head = null) {
    this.head = head
    this.length = 0
  }

  setNewHead(head) {
    this.head = head
  }

  InsertAtBeggning(val) {
    let node = new Node(val,this.head)
    this.head = node
    this.length++
  }

  Print() {
    if (this.length == 0) {
      console.log("No nodes in list")
    }
    let ptr = this.head
    for (let i = 0; i < this.length; i++) {
      console.log("Node: ", ptr)
      ptr = ptr.next
    }
  }
}

module.exports = {
  Linkedlist
}
