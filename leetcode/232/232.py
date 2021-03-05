class MyQueue:
        

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.S1 = []
        self.S2 = []

    def push(self, x: int) -> None:
        """
        Push element x to the back of queue.
        """
        self.S1.append(x)

    def pop(self) -> int:
        """
        Removes the element from in front of queue and returns that element.
        """
        res = self.peek()
        self.S2 = self.S2[:-1]
        return res

    def peek(self) -> int:
        """
        Get the front element.
        """
        if len(self.S2) == 0:
            while len(self.S1) != 0:
                self.S2.append(self.S1[-1])
                self.S1 = self.S1[:-1]
        res = self.S2[-1]
        return res

    def empty(self) -> bool:
        """
        Returns whether the queue is empty.
        """
        return len(self.S1) == 0 and len(self.S2) == 0


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
